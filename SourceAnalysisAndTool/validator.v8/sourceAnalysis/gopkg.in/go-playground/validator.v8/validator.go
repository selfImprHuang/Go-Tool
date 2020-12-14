/**
 * Package validator
 *
 * MISC:
 * - anonymous structs - they don't have names so expect the Struct name within StructErrors to be blank
 *
 */

package validator

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"
)

const (
	utf8HexComma    = "0x2C" //read note 在tag里面代表【,】;需要这么书写，直接写【,】会被识别成分割符
	utf8Pipe        = "0x7C" //read note 在tag里面代表【|】;需要这么书写，直接写【|】会被识别成条件符
	tagSeparator    = ","    //read note tag标签分割符
	orSeparator     = "|"    //read note tag判断条件【或】
	tagKeySeparator = "="    //read note tag里面需要等值 使用=

	structOnlyTag    = "structonly"    //read note 结构体上有的tagType,因为结构体的特殊性,他在获取cTag的时候是从第二个开始的
	_                                  //  也就是说第二个才是生效的,所以当只书写structonly的时候,他会略过这个tag,那么必须跟着前面的一个标签才行,比如说 `valid="require,structonly"`,这样才会生效
	noStructLevelTag = "nostructlevel" //  同上

	omitempty         = "omitempty" //read note 如果字段未设值则忽略它
	skipValidationTag = "-"         //read note 忽略字段
	diveTag           = "dive"      //read note 深入到slice, array or map 里面去校验里面的字段是否正确，每一层都需要多一个【div】来标识
	_                               //read note 比如说 [][]string "gt=0,dive,dive,required" gt是校验[][]string长度，第一个div是校验[]string,第二个div是校验string
	existsTag         = "exists"    //read note 校验值存在即可（除非是nil则会报错），和require的区别是说，require对默认值也会报错

	//read note 数组结构的名称组成
	arrayIndexFieldName = "%s" + leftBracket + "%d" + rightBracket
	mapIndexFieldName   = "%s" + leftBracket + "%v" + rightBracket

	//read note 字段校验错误信息打印
	fieldErrMsg = "Key: '%s' Error:Field validation for '%s' failed on the '%s' tag"
	//read note 报错信息(非字段校验错误)
	invalidValidation       = "Invalid validation tag on field %s"
	undefinedValidation     = "Undefined validation function on field %s"
	validatorNotInitialized = "Validator instance not initialized"
	fieldNameRequired       = "Field Name Required"
	tagRequired             = "Tag Required"
)

var (
	timeType      = reflect.TypeOf(time.Time{})
	timePtrType   = reflect.TypeOf(&time.Time{})
	defaultCField = new(cField)
)

// StructLevel contains all of the information and helper methods
// for reporting errors during struct level validation
type StructLevel struct {
	TopStruct     reflect.Value
	CurrentStruct reflect.Value
	errPrefix     string
	nsPrefix      string
	errs          ValidationErrors
	v             *Validate
}

// ReportValidationErrors accepts the key relative to the top level struct and validatin errors.
// Example: had a triple nested struct User, ContactInfo, Country and ran errs := validate.Struct(country)
// from within a User struct level validation would call this method like so:
// ReportValidationErrors("ContactInfo.", errs)
// NOTE: relativeKey can contain both the Field Relative and Custom name relative paths
// i.e. ReportValidationErrors("ContactInfo.|cInfo", errs) where cInfo represents say the JSON name of
// the relative path; this will be split into 2 variables in the next valiator version.
func (sl *StructLevel) ReportValidationErrors(relativeKey string, errs ValidationErrors) {
	for _, e := range errs {

		idx := strings.Index(relativeKey, "|")
		var rel string
		var cRel string

		if idx != -1 {
			rel = relativeKey[:idx]
			cRel = relativeKey[idx+1:]
		} else {
			rel = relativeKey
		}

		key := sl.errPrefix + rel + e.Field

		e.FieldNamespace = key
		e.NameNamespace = sl.nsPrefix + cRel + e.Name

		sl.errs[key] = e
	}
}

// ReportError reports an error just by passing the field and tag information
// NOTE: tag can be an existing validation tag or just something you make up
// and precess on the flip side it's up to you.
func (sl *StructLevel) ReportError(field reflect.Value, fieldName string, customName string, tag string) {

	field, kind := sl.v.ExtractType(field)

	if fieldName == blank {
		panic(fieldNameRequired)
	}

	if customName == blank {
		customName = fieldName
	}

	if tag == blank {
		panic(tagRequired)
	}

	ns := sl.errPrefix + fieldName

	switch kind {
	case reflect.Invalid:
		sl.errs[ns] = &FieldError{
			FieldNamespace: ns,
			NameNamespace:  sl.nsPrefix + customName,
			Name:           customName,
			Field:          fieldName,
			Tag:            tag,
			ActualTag:      tag,
			Param:          blank,
			Kind:           kind,
		}
	default:
		sl.errs[ns] = &FieldError{
			FieldNamespace: ns,
			NameNamespace:  sl.nsPrefix + customName,
			Name:           customName,
			Field:          fieldName,
			Tag:            tag,
			ActualTag:      tag,
			Param:          blank,
			Value:          field.Interface(),
			Kind:           kind,
			Type:           field.Type(),
		}
	}
}

// Validate contains the validator settings passed in using the Config struct
type Validate struct {
	tagName             string                           //校验起作用的tag名
	fieldNameTag        string                           //
	validationFuncs     map[string]Func                  //规则类型的校验			【tag标签】->			校验规则
	structLevelFuncs    map[reflect.Type]StructLevelFunc //规则结构体的校验		【结构体类型】->			校验规则
	customTypeFuncs     map[reflect.Type]CustomTypeFunc  //类型校验器			【数据类型】->			校验规则
	aliasValidators     map[string]string                //别名校验器 			【别名匹配规则组合】-> 	校验规则
	hasCustomFuncs      bool                             //是否存在类型校验器
	hasAliasValidators  bool                             //是否有别名校验器
	hasStructLevelFuncs bool                             //是否有结构体校验器
	tagCache            *tagCache                        //tag对应的【校验规则方法】的缓存
	structCache         *structCache                     //结构体对应的【校验规则方法】的缓存
	errsPool            *sync.Pool                       //校验错误奖池
}

func (v *Validate) initCheck() {
	if v == nil {
		panic(validatorNotInitialized)
	}
}

// Config contains the options that a Validator instance will use.
// It is passed to the New() function
type Config struct {
	TagName      string
	FieldNameTag string
}

// CustomTypeFunc allows for overriding or adding custom field type handler functions
// field = field value of the type to return a value to be validated
// example Valuer from sql drive see https://golang.org/src/database/sql/driver/types.go?s=1210:1293#L29
type CustomTypeFunc func(field reflect.Value) interface{}

// Func accepts all values needed for file and cross field validation
// v             = validator instance, needed but some built in functions for it's custom types
// topStruct     = top level struct when validating by struct otherwise nil
// currentStruct = current level struct when validating by struct otherwise optional comparison value
// field         = field value for validation
// param         = parameter used in validation i.e. gt=0 param would be 0
type Func func(v *Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldtype reflect.Type, fieldKind reflect.Kind, param string) bool

// StructLevelFunc accepts all values needed for struct level validation
type StructLevelFunc func(v *Validate, structLevel *StructLevel)

// ValidationErrors is a type of map[string]*FieldError
// it exists to allow for multiple errors to be passed from this library
// and yet still subscribe to the error interface
type ValidationErrors map[string]*FieldError

// Error is intended for use in development + debugging and not intended to be a production error message.
// It allows ValidationErrors to subscribe to the Error interface.
// All information to create an error message specific to your application is contained within
// the FieldError found within the ValidationErrors map
func (ve ValidationErrors) Error() string {

	buff := bytes.NewBufferString(blank)

	for key, err := range ve {
		buff.WriteString(fmt.Sprintf(fieldErrMsg, key, err.Field, err.Tag))
		buff.WriteString("\n")
	}

	return strings.TrimSpace(buff.String())
}

// FieldError contains a single field's validation error along
// with other properties that may be needed for error message creation
type FieldError struct {
	FieldNamespace string
	NameNamespace  string
	Field          string
	Name           string
	Tag            string
	ActualTag      string
	Kind           reflect.Kind
	Type           reflect.Type
	Param          string
	Value          interface{}
}

// New creates a new Validate instance for use.
func New(config *Config) *Validate {

	//read note tag校验器缓存，初始化
	tc := new(tagCache)
	tc.m.Store(make(map[string]*cTag))

	//read note 结构体缓存，初始化一个Map，防止后面添加的时候报错
	sc := new(structCache)
	sc.m.Store(make(map[reflect.Type]*cStruct))

	v := &Validate{
		tagName:      config.TagName,
		fieldNameTag: config.FieldNameTag,
		tagCache:     tc,
		structCache:  sc,
		//read note 自定义一个错误池
		errsPool: &sync.Pool{New: func() interface{} {
			return ValidationErrors{}
		}}}

	//read note 设置别名规则类（一组） 可参考： baked_in.go/bakedInAliasValidators
	if len(v.aliasValidators) == 0 {
		// must copy alias validators for separate validations to be used in each validator instance
		v.aliasValidators = map[string]string{}
		for k, val := range bakedInAliasValidators {
			//read note 默认别名校验器注册
			v.RegisterAliasValidation(k, val)
		}
	}

	//read note 设置默认的校验方法.可参考： 	baked_in.go/bakedInValidators
	if len(v.validationFuncs) == 0 {
		// must copy validators for separate validations to be used in each instance
		v.validationFuncs = map[string]Func{}
		for k, val := range bakedInValidators {
			//read note 默认tag校验器注册
			v.RegisterValidation(k, val)
		}
	}

	return v
}

// RegisterStructValidation registers a StructLevelFunc against a number of types
// NOTE: this method is not thread-safe it is intended that these all be registered prior to any validation
func (v *Validate) RegisterStructValidation(fn StructLevelFunc, types ...interface{}) {
	v.initCheck()

	if v.structLevelFuncs == nil {
		v.structLevelFuncs = map[reflect.Type]StructLevelFunc{}
	}

	for _, t := range types {
		v.structLevelFuncs[reflect.TypeOf(t)] = fn
		fmt.Println(reflect.TypeOf(t))
	}

	v.hasStructLevelFuncs = true
}

// RegisterValidation adds a validation Func to a Validate's map of validators denoted by the key
// NOTE: if the key already exists, the previous validation function will be replaced.
// NOTE: this method is not thread-safe it is intended that these all be registered prior to any validation
func (v *Validate) RegisterValidation(key string, fn Func) error {
	v.initCheck()

	if key == blank {
		return errors.New("Function Key cannot be empty")
	}

	if fn == nil {
		return errors.New("Function cannot be empty")
	}

	_, ok := restrictedTags[key]

	if ok || strings.ContainsAny(key, restrictedTagChars) {
		panic(fmt.Sprintf(restrictedTagErr, key))
	}

	v.validationFuncs[key] = fn

	return nil
}

// RegisterCustomTypeFunc registers a CustomTypeFunc against a number of types
// NOTE: this method is not thread-safe it is intended that these all be registered prior to any validation
//read note 注册类型的校验器，应该在校验之前就注册完成，因为该注册方法不是线程安全的.
func (v *Validate) RegisterCustomTypeFunc(fn CustomTypeFunc, types ...interface{}) {
	v.initCheck()

	if v.customTypeFuncs == nil {
		v.customTypeFuncs = map[reflect.Type]CustomTypeFunc{}
	}

	for _, t := range types {
		v.customTypeFuncs[reflect.TypeOf(t)] = fn
	}

	v.hasCustomFuncs = true
}

// RegisterAliasValidation registers a mapping of a single validationstag that
// defines a common or complex set of validation(s) to simplify adding validation
// to structs. NOTE: when returning an error the tag returned in FieldError will be
// the alias tag unless the dive tag is part of the alias; everything after the
// dive tag is not reported as the alias tag. Also the ActualTag in the before case
// will be the actual tag within the alias that failed.
// NOTE: this method is not thread-safe it is intended that these all be registered prior to any validation
func (v *Validate) RegisterAliasValidation(alias, tags string) {
	v.initCheck()

	_, ok := restrictedTags[alias]

	if ok || strings.ContainsAny(alias, restrictedTagChars) {
		panic(fmt.Sprintf(restrictedAliasErr, alias))
	}

	v.aliasValidators[alias] = tags
	v.hasAliasValidators = true
}

// Field validates a single field using tag style validation and returns nil or ValidationErrors as type error.
// You will need to assert the error if it's not nil i.e. err.(validator.ValidationErrors) to access the map of errors.
// NOTE: it returns ValidationErrors instead of a single FieldError because this can also
// validate Array, Slice and maps fields which may contain more than one error
//read note
func (v *Validate) Field(field interface{}, tag string) error {
	v.initCheck()

	if len(tag) == 0 || tag == skipValidationTag {
		return nil
	}

	errs := v.errsPool.Get().(ValidationErrors)
	fieldVal := reflect.ValueOf(field)

	//read note 从tag缓存中获取 tag标签对应的校验方法
	ctag, ok := v.tagCache.Get(tag)
	if !ok {
		//read note 加锁
		v.tagCache.lock.Lock()
		defer v.tagCache.lock.Unlock()

		// could have been multiple trying to access, but once first is done this ensures tag
		// isn't parsed again.
		//read note 加锁之后在这边再获取一次，因为在上一次判断到加锁的过程中，可能有函数已经把对应方法加载进来了
		ctag, ok = v.tagCache.Get(tag)
		if !ok {
			//read note
			ctag, _ = v.parseFieldTagsRecursive(tag, blank, blank, false)
			//read note 进行标签设值
			v.tagCache.Set(tag, ctag)
		}
	}

	v.traverseField(fieldVal, fieldVal, fieldVal, blank, blank, errs, false, false, nil, nil, defaultCField, ctag)

	//read note 判断校验错误是否存在
	if len(errs) == 0 {
		v.errsPool.Put(errs)
		return nil
	}

	return errs
}

// FieldWithValue validates a single field, against another fields value using tag style validation and returns nil or ValidationErrors.
// You will need to assert the error if it's not nil i.e. err.(validator.ValidationErrors) to access the map of errors.
// NOTE: it returns ValidationErrors instead of a single FieldError because this can also
// validate Array, Slice and maps fields which may contain more than one error
func (v *Validate) FieldWithValue(val interface{}, field interface{}, tag string) error {
	v.initCheck()

	if len(tag) == 0 || tag == skipValidationTag {
		return nil
	}

	errs := v.errsPool.Get().(ValidationErrors)
	topVal := reflect.ValueOf(val)

	ctag, ok := v.tagCache.Get(tag)
	if !ok {
		v.tagCache.lock.Lock()
		defer v.tagCache.lock.Unlock()

		// could have been multiple trying to access, but once first is done this ensures tag
		// isn't parsed again.
		ctag, ok = v.tagCache.Get(tag)
		if !ok {
			ctag, _ = v.parseFieldTagsRecursive(tag, blank, blank, false)
			v.tagCache.Set(tag, ctag)
		}
	}

	v.traverseField(topVal, topVal, reflect.ValueOf(field), blank, blank, errs, false, false, nil, nil, defaultCField, ctag)

	if len(errs) == 0 {
		v.errsPool.Put(errs)
		return nil
	}

	return errs
}

// StructPartial validates the fields passed in only, ignoring all others.
// Fields may be provided in a namespaced fashion relative to the  struct provided
// i.e. NestedStruct.Field or NestedArrayField[0].Struct.Name and returns nil or ValidationErrors as error
// You will need to assert the error if it's not nil i.e. err.(validator.ValidationErrors) to access the map of errors.
func (v *Validate) StructPartial(current interface{}, fields ...string) error {
	v.initCheck()

	sv, _ := v.ExtractType(reflect.ValueOf(current))
	name := sv.Type().Name()
	m := map[string]struct{}{}

	if fields != nil {
		for _, k := range fields {

			flds := strings.Split(k, namespaceSeparator)
			if len(flds) > 0 {

				key := name + namespaceSeparator
				for _, s := range flds {

					idx := strings.Index(s, leftBracket)

					if idx != -1 {
						for idx != -1 {
							key += s[:idx]
							m[key] = struct{}{}

							idx2 := strings.Index(s, rightBracket)
							idx2++
							key += s[idx:idx2]
							m[key] = struct{}{}
							s = s[idx2:]
							idx = strings.Index(s, leftBracket)
						}
					} else {

						key += s
						m[key] = struct{}{}
					}

					key += namespaceSeparator
				}
			}
		}
	}

	errs := v.errsPool.Get().(ValidationErrors)

	v.ensureValidStruct(sv, sv, sv, blank, blank, errs, true, len(m) != 0, false, m, false)

	if len(errs) == 0 {
		v.errsPool.Put(errs)
		return nil
	}

	return errs
}

// StructExcept validates all fields except the ones passed in.
// Fields may be provided in a namespaced fashion relative to the  struct provided
// i.e. NestedStruct.Field or NestedArrayField[0].Struct.Name and returns nil or ValidationErrors as error
// You will need to assert the error if it's not nil i.e. err.(validator.ValidationErrors) to access the map of errors.
func (v *Validate) StructExcept(current interface{}, fields ...string) error {
	v.initCheck()

	sv, _ := v.ExtractType(reflect.ValueOf(current))
	name := sv.Type().Name()
	m := map[string]struct{}{}

	for _, key := range fields {
		m[name+namespaceSeparator+key] = struct{}{}
	}

	errs := v.errsPool.Get().(ValidationErrors)

	v.ensureValidStruct(sv, sv, sv, blank, blank, errs, true, len(m) != 0, true, m, false)

	if len(errs) == 0 {
		v.errsPool.Put(errs)
		return nil
	}

	return errs
}

// Struct validates a structs exposed fields, and automatically validates nested structs, unless otherwise specified.
// it returns nil or ValidationErrors as error.
// You will need to assert the error if it's not nil i.e. err.(validator.ValidationErrors) to access the map of errors.
func (v *Validate) Struct(current interface{}) error {
	v.initCheck()

	errs := v.errsPool.Get().(ValidationErrors)
	sv := reflect.ValueOf(current)

	//read note 进行结构体校验
	v.ensureValidStruct(sv, sv, sv, blank, blank, errs, true, false, false, nil, false)

	//read note 校验之后，对校验错误进行处理.
	if len(errs) == 0 {
		v.errsPool.Put(errs)
		return nil
	}

	return errs
}

func (v *Validate) ensureValidStruct(topStruct reflect.Value, currentStruct reflect.Value, current reflect.Value, errPrefix string, nsPrefix string, errs ValidationErrors, useStructName bool, partial bool, exclude bool, includeExclude map[string]struct{}, isStructOnly bool) {

	//read note 指针处理成结构体Value(返回其指向的实际结构体)
	if current.Kind() == reflect.Ptr && !current.IsNil() {
		current = current.Elem()
	}

	//read note 结构体类型校验.
	if current.Kind() != reflect.Struct && current.Kind() != reflect.Interface {
		panic("value passed for validation is not a struct")
	}

	//read note 校验结构体
	v.tranverseStruct(topStruct, currentStruct, current, errPrefix, nsPrefix, errs, useStructName, partial, exclude, includeExclude, nil, nil)
}

// tranverseStruct traverses a structs fields and then passes them to be validated by traverseField
func (v *Validate) tranverseStruct(topStruct reflect.Value, currentStruct reflect.Value, current reflect.Value, errPrefix string, nsPrefix string, errs ValidationErrors, useStructName bool, partial bool, exclude bool, includeExclude map[string]struct{}, cs *cStruct, ct *cTag) {

	var ok bool
	first := len(nsPrefix) == 0
	typ := current.Type()

	//read note 结构体校验器_缓存
	cs, ok = v.structCache.Get(typ)
	if !ok {
		cs = v.extractStructCache(current, typ.Name())
	}

	if useStructName {
		errPrefix += cs.Name + namespaceSeparator

		if len(v.fieldNameTag) != 0 {
			nsPrefix += cs.Name + namespaceSeparator
		}
	}

	// structonly tag present don't tranverseFields
	// but must still check and run below struct level validation
	// if present
	//read note 【structonly】标签会忽略掉Field的校验
	if first || ct == nil || ct.typeof != typeStructOnly {

		for _, f := range cs.fields {

			if partial {

				_, ok = includeExclude[errPrefix+f.Name]

				if (ok && exclude) || (!ok && !exclude) {
					continue
				}
			}

			v.traverseField(topStruct, currentStruct, current.Field(f.Idx), errPrefix, nsPrefix, errs, partial, exclude, includeExclude, cs, f, f.cTags)
		}
	}

	// check if any struct level validations, after all field validations already checked.
	//read note 如果结构体层的校验存在的话,需要进行调用.(这边是规则类校验)
	if cs.fn != nil {
		cs.fn(v, &StructLevel{v: v, TopStruct: topStruct, CurrentStruct: current, errPrefix: errPrefix, nsPrefix: nsPrefix, errs: errs})
	}
}

// traverseField validates any field, be it a struct or single field, ensures it's validity and passes it along to be validated via it's tag options
func (v *Validate) traverseField(topStruct reflect.Value, currentStruct reflect.Value, current reflect.Value, errPrefix string, nsPrefix string, errs ValidationErrors, partial bool, exclude bool, includeExclude map[string]struct{}, cs *cStruct, cf *cField, ct *cTag) {

	//read note 处理Ptr、Interface、Invalid(validator自定义)、自定义结构体类型.
	// 返回参数为：值 、 数据源类型、nullable(用来校验 【结构体】的【omiEmpty】标签...)
	current, kind, nullable := v.extractTypeInternal(current, false)
	var typ reflect.Type

	switch kind {
	//read note 处理地址、Interface、Invalid是直接往校验错误池中添加一个Field错误?
	case reflect.Ptr, reflect.Interface, reflect.Invalid:

		if ct == nil {
			return
		}

		if ct.typeof == typeOmitEmpty {
			return
		}

		if ct.hasTag {

			ns := errPrefix + cf.Name

			//read note 如果是Invalid,需要往校验错误池中添加一个Field错误
			if kind == reflect.Invalid {
				errs[ns] = &FieldError{
					FieldNamespace: ns,
					NameNamespace:  nsPrefix + cf.AltName,
					Name:           cf.AltName,
					Field:          cf.Name,
					Tag:            ct.aliasTag,
					ActualTag:      ct.tag,
					Param:          ct.param,
					Kind:           kind,
				}
				return
			}

			errs[ns] = &FieldError{
				FieldNamespace: ns,
				NameNamespace:  nsPrefix + cf.AltName,
				Name:           cf.AltName,
				Field:          cf.Name,
				Tag:            ct.aliasTag,
				ActualTag:      ct.tag,
				Param:          ct.param,
				Value:          current.Interface(),
				Kind:           kind,
				Type:           current.Type(),
			}

			return
		}

	case reflect.Struct:
		typ = current.Type()

		if typ != timeType {

			//read note 对于结构体，这边有一个比较特殊的处理就是cTag是从第二个位置开始，也就是之前一直尝试structonly不成功,就是需要在structonly前面加一个tag才行。
			// 前一个对于struct应该是不生效才对，第二个才会对struct生效
			if ct != nil {
				ct = ct.next
			}

			if ct != nil && ct.typeof == typeNoStructLevel {
				return
			}

			//read note 如果上面的cTag没有生效/没有tagType为【typeNoStructLevel】的标签.就直接进入这个结构体
			v.tranverseStruct(topStruct, current, current, errPrefix+cf.Name+namespaceSeparator, nsPrefix+cf.AltName+namespaceSeparator, errs, false, partial, exclude, includeExclude, cs, ct)
			return
		}
	}

	if !ct.hasTag {
		return
	}

	typ = current.Type()

	//read note 进行字段Field的校验，这边的话是一个for的循环,在不同的标签中去找到对应的处理方式
OUTER:
	for {
		if ct == nil {
			return
		}

		switch ct.typeof {

		case typeExists:
			ct = ct.next
			continue

		case typeOmitEmpty:

			if !nullable && !HasValue(v, topStruct, currentStruct, current, typ, kind, blank) {
				return
			}

			ct = ct.next
			continue

		case typeDive:

			ct = ct.next

			// traverse slice or map here
			// or panic ;)
			switch kind {
			//read note 【dive】标签处理Slice 和 Array：往数组中的每一个元素进行处理
			case reflect.Slice, reflect.Array:

				for i := 0; i < current.Len(); i++ {
					v.traverseField(topStruct, currentStruct, current.Index(i), errPrefix, nsPrefix, errs, partial, exclude, includeExclude, cs, &cField{Name: fmt.Sprintf(arrayIndexFieldName, cf.Name, i), AltName: fmt.Sprintf(arrayIndexFieldName, cf.AltName, i)}, ct)
				}
			//read note 【dive】标签处理 Map：往Map中的每一个元素进行处理
			case reflect.Map:
				for _, key := range current.MapKeys() {
					v.traverseField(topStruct, currentStruct, current.MapIndex(key), errPrefix, nsPrefix, errs, partial, exclude, includeExclude, cs, &cField{Name: fmt.Sprintf(mapIndexFieldName, cf.Name, key.Interface()), AltName: fmt.Sprintf(mapIndexFieldName, cf.AltName, key.Interface())}, ct)
				}
			//read note 【div】只能标注在 Slice 、 Map 、Array 上,不然就是报错.
			default:
				// throw error, if not a slice or map then should not have gotten here
				// bad dive tag
				panic("dive error! can't dive on a non slice or map")
			}

			return

		case typeOr:

			//read note 【|】的处理
			errTag := blank

			for {

				//read note 如果结果是true的话，因为【|】标签只要有一个成立就可以，所以可以直接排除掉其他 【|】标签了
				if ct.fn(v, topStruct, currentStruct, current, typ, kind, ct.param) {

					// drain rest of the 'or' values, then continue or leave
					//read note 排除其他 【|】标签,直到下一个Field的非【|】标签再进入处理.
					for {

						ct = ct.next

						if ct == nil {
							return
						}

						if ct.typeof != typeOr {
							continue OUTER
						}
					}
				}

				errTag += orSeparator + ct.tag

				//read note 没有下一个cTag的处理器.处理一i西安错误信息，返回.
				if ct.next == nil {
					// if we get here, no valid 'or' value and no more tags

					ns := errPrefix + cf.Name

					if ct.hasAlias {
						errs[ns] = &FieldError{
							FieldNamespace: ns,
							NameNamespace:  nsPrefix + cf.AltName,
							Name:           cf.AltName,
							Field:          cf.Name,
							Tag:            ct.aliasTag,
							ActualTag:      ct.actualAliasTag,
							Value:          current.Interface(),
							Type:           typ,
							Kind:           kind,
						}
					} else {
						errs[errPrefix+cf.Name] = &FieldError{
							FieldNamespace: ns,
							NameNamespace:  nsPrefix + cf.AltName,
							Name:           cf.AltName,
							Field:          cf.Name,
							Tag:            errTag[1:],
							ActualTag:      errTag[1:],
							Value:          current.Interface(),
							Type:           typ,
							Kind:           kind,
						}
					}

					return
				}

				ct = ct.next
			}

		default:
			//read note 剩下的标签处理.类似一个循环的处理
			if !ct.fn(v, topStruct, currentStruct, current, typ, kind, ct.param) {

				ns := errPrefix + cf.Name

				errs[ns] = &FieldError{
					FieldNamespace: ns,
					NameNamespace:  nsPrefix + cf.AltName,
					Name:           cf.AltName,
					Field:          cf.Name,
					Tag:            ct.aliasTag,
					ActualTag:      ct.tag,
					Value:          current.Interface(),
					Param:          ct.param,
					Type:           typ,
					Kind:           kind,
				}

				return

			}

			ct = ct.next
		}
	}
}
