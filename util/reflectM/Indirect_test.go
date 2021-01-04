/*
 *  @Author : huangzj
 *  @Time : 2021/1/4 15:46
 *  @Description：
 */

package reflectM

import (
	"fmt"
	"reflect"
	"testing"
)

type Test struct {
	testA int
	testB string
}

func TestIndirect(t *testing.T) {
	i := 20
	s := "测试"
	fmt.Println("测试type")
	test := Test{
		testA: 10000,
		testB: "我是测试属性",
	}

	fmt.Println(IndirectType(reflect.TypeOf(i)))
	fmt.Println(IndirectType(reflect.TypeOf(s)))
	fmt.Println(IndirectType(reflect.TypeOf(test)))
	fmt.Println(IndirectType(reflect.TypeOf(&test)))
	fmt.Println(IndirectType(reflect.TypeOf([]Test{test})))
	fmt.Println(IndirectType(reflect.TypeOf([]*Test{&test, &test, &test})))
	fmt.Println(IndirectType(reflect.TypeOf(map[string]Test{
		"a": test,
	})))
	fmt.Println(IndirectType(reflect.TypeOf(map[string]*Test{
		"a": &test,
	})))

	fmt.Println("测试value")
	fmt.Println()
	fmt.Println()

	fmt.Println(Indirect(reflect.ValueOf(i)))
	fmt.Println(Indirect(reflect.ValueOf(s)))
	fmt.Println(Indirect(reflect.ValueOf(test)))
	fmt.Println(Indirect(reflect.ValueOf(&test)))
	fmt.Println(Indirect(reflect.ValueOf([]Test{test})))
	fmt.Println(Indirect(reflect.ValueOf([]*Test{&test})))
	fmt.Println(Indirect(reflect.ValueOf(map[string]Test{
		"a": test,
	})))
	fmt.Println(Indirect(reflect.ValueOf(map[string]*Test{
		"a": &test,
	})))
}
