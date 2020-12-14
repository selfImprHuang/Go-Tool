/*
 *  @Author : huangzj
 *  @Time : 2020/12/11 16:16
 *  @Description：
 */

package testValidator

import (
	"Go-Tool/SourceAnalysisAndTool/validator.v8/sourceAnalysis/gopkg.in/go-playground/validator.v8"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestValidator(t *testing.T) {
	fmt.Println("测试别名校验器")
	testAlias()

	fmt.Println()
	fmt.Println("测试默认和自定义tag校验器")
	testTag()

	fmt.Println()
	fmt.Println("测试字段校验器")
	testField()

	fmt.Println()
	fmt.Println("测试结构体校验器")
	testStruct()

}

func testStruct() {
	//添加结构体校验，对结构体进行处理之后，validator的标签还是会起作用
	config := &validator.Config{
		TagName: "validate",
	}
	valid := validator.New(config)
	valid.RegisterStructValidation(func(v *validator.Validate, structLevel *validator.StructLevel) {
		c := structLevel.CurrentStruct.Interface().(Cat)
		if c.Age < 100 {
			fmt.Println("猫的岁数不能小于100")
		}
	}, Cat{})
	p := Person{
		Name:    "192.168.0.1",
		M:       map[string]int{"A": 1},
		H:       2000,
		Exist:   false,
		Require: false,
		Cat: Cat{
			Age:  0,
			Ower: Ower{},
		},
	}

	err := valid.Struct(p)
	fmt.Println(err)
}

func testField() {
	//如果注册了对应的Field，则会调用自定义的方法，不会再使用validator定义的标签
	//这边测试了结构体和基本类型，都是如上述描述的这么处理
	config := &validator.Config{
		TagName: "validate",
	}
	valid := validator.New(config)
	valid.RegisterAliasValidation("ipe", "ip|ipv4|ipv6")

	valid.RegisterCustomTypeFunc(func(field reflect.Value) interface{} {
		d := field.Interface().(Cat)
		if d.Age < 50 {
			fmt.Println("狗的岁数不能小于50")
		}
		return field
	}, Cat{})

	p := Person{
		Name:    "192.168.0.1",
		M:       map[string]int{"A": 1},
		H:       2000,
		Exist:   false,
		Require: false,
		Cat: Cat{
			Age:  0,
			Ower: Ower{},
		},
	}
	err := valid.Struct(p)
	fmt.Println(err)

	valid.RegisterCustomTypeFunc(func(field reflect.Value) interface{} {
		d := field.Interface().(int)
		if d != 10 {
			fmt.Println("测试一下")
		}
		return field
	}, 1)

	err1 := valid.Struct(p)
	fmt.Println(err1)

}

func testTag() {
	w := Woman{
		Name: "my name is Woman",
	}

	w1 := Woman{
		Name: "my name is wo-man",
	}

	config := &validator.Config{
		TagName: "validate",
	}
	valid := validator.New(config)
	_ = valid.RegisterValidation("diy", func(v *validator.Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldtype reflect.Type, fieldKind reflect.Kind, param string) bool {
		s := field.Interface().(string)

		i, _ := strconv.ParseInt(param, 0, 64)
		if int64(len(s)) < i || strings.Contains(s, "Woman") {
			fmt.Println("校验一下")
			return false
		}
		return true
	})

	err := valid.Struct(w)
	fmt.Println(err)

	err1 := valid.Struct(w1)
	fmt.Println(err1)
}

func testAlias() {
	m := Man{
		Name: "192.168.0.1",
	}

	m1 := Man{
		Name: "213123123",
	}

	config := &validator.Config{
		TagName: "validate",
	}
	valid := validator.New(config)
	valid.RegisterAliasValidation("ipe", "ip|ipv4|ipv6")
	err := valid.Struct(m)
	fmt.Println(err)

	err1 := valid.Struct(m1)
	fmt.Println(err1)
}
