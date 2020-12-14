package validator

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
	"sync/atomic"
)

type tagType uint8

const (
	typeDefault tagType = iota
	typeOmitEmpty
	typeNoStructLevel
	typeStructOnly
	typeDive
	typeOr
	typeExists
)

type structCache struct {
	lock sync.Mutex
	m    atomic.Value // map[reflect.Type]*cStruct
}

func (sc *structCache) Get(key reflect.Type) (c *cStruct, found bool) {
	c, found = sc.m.Load().(map[reflect.Type]*cStruct)[key]
	return
}

func (sc *structCache) Set(key reflect.Type, value *cStruct) {

	m := sc.m.Load().(map[reflect.Type]*cStruct)

	nm := make(map[reflect.Type]*cStruct, len(m)+1)
	for k, v := range m {
		nm[k] = v
	}
	nm[key] = value
	sc.m.Store(nm)
}

type tagCache struct {
	lock sync.Mutex
	m    atomic.Value // map[string]*cTag
}

func (tc *tagCache) Get(key string) (c *cTag, found bool) {
	c, found = tc.m.Load().(map[string]*cTag)[key]
	return
}

func (tc *tagCache) Set(key string, value *cTag) {

	m := tc.m.Load().(map[string]*cTag)

	nm := make(map[string]*cTag, len(m)+1)
	for k, v := range m {
		nm[k] = v
	}
	nm[key] = value
	tc.m.Store(nm)
}

type cStruct struct {
	Name   string          //结构体名称
	fields map[int]*cField //结构体对应的字段map
	fn     StructLevelFunc //结构体校验器 （结构体类型->结构体校验器）
}

type cField struct {
	Idx     int    //字段下标
	Name    string //字段名
	AltName string //
	cTags   *cTag  //Field对应的cTag规则,是一个链表（一串的规则）.
}

type cTag struct {
	tag            string  //标签
	aliasTag       string  //
	actualAliasTag string  //
	param          string  //如果是比较类型的标签，这里存放的是比较的值，比如说 min=10，这里存放的是【10】这个值
	hasAlias       bool    //是否有别名校验器标签
	typeof         tagType //对应的tagType
	hasTag         bool    //是否存在tag标签
	fn             Func    //当前cTag对应的【tag标签校验器】
	next           *cTag   //下一个cTag标签
}

func (v *Validate) extractStructCache(current reflect.Value, sName string) *cStruct {

	v.structCache.lock.Lock()
	defer v.structCache.lock.Unlock() // leave as defer! because if inner panics, it will never get unlocked otherwise!

	typ := current.Type()

	// could have been multiple trying to access, but once first is done this ensures struct
	// isn't parsed again.
	//read note 从缓存里面获取对应结构体类型的处理.
	cs, ok := v.structCache.Get(typ)
	if ok {
		return cs
	}

	//read note 如果缓存里面拿不到的话，就要从结构的Field中去解析
	cs = &cStruct{Name: sName, fields: make(map[int]*cField), fn: v.structLevelFuncs[typ]}

	numFields := current.NumField()

	var ctag *cTag
	var fld reflect.StructField
	var tag string
	var customName string

	//read note 进行字段的处理
	for i := 0; i < numFields; i++ {

		fld = typ.Field(i)

		//read note 内嵌类型和小写的处理，直接pass
		// 类似这种的:
		/* type Derive struct {
		    	Base           ---内嵌
		   }
		*/
		if !fld.Anonymous && fld.PkgPath != blank {
			continue
		}

		//read note 获取tag（通过之前Config设置的【tagName】属性去获取对应的tag）
		tag = fld.Tag.Get(v.tagName)

		//read note 如果是忽略【-】的话，跳过
		if tag == skipValidationTag {
			continue
		}

		customName = fld.Name

		//read note config中【fieldNameTag】的处理,会设置到【cField】的【AltName】字段上，应该只是错误输出用的
		if v.fieldNameTag != blank {

			name := strings.SplitN(fld.Tag.Get(v.fieldNameTag), ",", 2)[0]

			// dash check is for json "-" (aka skipValidationTag) means don't output in json
			if name != "" && name != skipValidationTag {
				customName = name
			}
		}

		// NOTE: cannot use shared tag cache, because tags may be equal, but things like alias may be different
		// and so only struct level caching can be used instead of combined with Field tag caching

		//read note struct层的校验规则和Field层的校验规则不能共用（标签可能相同，但是别名会不一样..）

		if len(tag) > 0 {
			ctag, _ = v.parseFieldTagsRecursive(tag, fld.Name, blank, false)
		} else {
			// even if field doesn't have validations need cTag for traversing to potential inner/nested
			// elements of the field.
			ctag = new(cTag)
		}

		cs.fields[i] = &cField{Idx: i, Name: fld.Name, AltName: customName, cTags: ctag}
	}

	v.structCache.Set(typ, cs)

	return cs
}

func (v *Validate) parseFieldTagsRecursive(tag string, fieldName string, alias string, hasAlias bool) (firstCtag *cTag, current *cTag) {

	var t string
	var ok bool
	noAlias := len(alias) == 0
	//read note：这边会把tag根据【,】进行分割处理，得到对应的tag组，所以我们在写的时候可以写入,分割的标签
	tags := strings.Split(tag, tagSeparator)

	for i := 0; i < len(tags); i++ {

		t = tags[i]

		if noAlias {
			alias = t
		}

		//read note 如果有别名校验器，则回去查找别名校验器（第一个和后面的设置不一样.调用next和没有调用next的区别）
		if v.hasAliasValidators {
			// check map for alias and process new tags, otherwise process as usual
			if tagsVal, found := v.aliasValidators[t]; found {

				if i == 0 {
					//read note 对别名校验器进行解析(一个子递归的过程.)除第一个校验器外，后面的校验器是通过一个next指向的链表连接起来
					firstCtag, current = v.parseFieldTagsRecursive(tagsVal, fieldName, t, true)
				} else {
					next, curr := v.parseFieldTagsRecursive(tagsVal, fieldName, t, true)
					current.next, current = next, curr

				}

				continue
			}
		}

		//read note 设置对应的默认标签（第一个和后面的设置不一样.调用next和没有调用next的区别）
		if i == 0 {
			current = &cTag{aliasTag: alias, hasAlias: hasAlias, hasTag: true}
			firstCtag = current
		} else {
			current.next = &cTag{aliasTag: alias, hasAlias: hasAlias, hasTag: true}
			current = current.next
		}

		//read note 判断用 【,】分割后的标签，前面几个是处理特殊标签的.如果不是特殊标签，则往下处理设置tag的值
		switch t {
		case diveTag:
			current.typeof = typeDive
			continue

		case omitempty:
			current.typeof = typeOmitEmpty
			continue

		case structOnlyTag:
			current.typeof = typeStructOnly
			continue

		case noStructLevelTag:
			current.typeof = typeNoStructLevel
			continue

		case existsTag:
			current.typeof = typeExists
			continue

		default:
			// if a pipe character is needed within the param you must use the utf8Pipe representation "0x7C"
			orVals := strings.Split(t, orSeparator)

			//read note 或的条件进行拆分.

			for j := 0; j < len(orVals); j++ {
				//read note 解析【=】标签
				vals := strings.SplitN(orVals[j], tagKeySeparator, 2)

				if noAlias {
					alias = vals[0]
					current.aliasTag = alias
				} else {
					current.actualAliasTag = t
				}

				//read note 如果或的标签成立，也就是说 A|B|C ABC会组装成一个链（非头节点需要往下一个节点去处理，所以这边需要把current指向它的next）
				if j > 0 {
					current.next = &cTag{aliasTag: alias, actualAliasTag: current.actualAliasTag, hasAlias: hasAlias, hasTag: true}
					current = current.next
				}

				current.tag = vals[0]
				if len(current.tag) == 0 {
					panic(strings.TrimSpace(fmt.Sprintf(invalidValidation, fieldName)))
				}

				//read note 查找对应的校验规则
				if current.fn, ok = v.validationFuncs[current.tag]; !ok {
					panic(strings.TrimSpace(fmt.Sprintf(undefinedValidation, fieldName)))
				}

				//read note 设置为typeOr.这个应该在后面处理对应字段的时候会拿到并且通过or的方式进行处理
				if len(orVals) > 1 {
					current.typeof = typeOr
				}

				if len(vals) > 1 {
					current.param = strings.Replace(strings.Replace(vals[1], utf8HexComma, ",", -1), utf8Pipe, "|", -1)
				}
			}
		}
	}

	return
}
