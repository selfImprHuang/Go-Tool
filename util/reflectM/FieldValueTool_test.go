/*
 *  @Author : huangzj
 *  @Time : 2020/12/16 11:51
 *  @Description：
 */

package reflectM

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflect(t *testing.T) {
	entity := Entity{
		Num: 1,
		S:   "1223",
		T:   make(map[string]string),
	}
	structMap := GetStructFieldTypeMap(entity)
	fmt.Println("根据结构体获取Map结构")
	for key, val := range structMap {
		fmt.Println(fmt.Sprintf("结构体属性，key: %v ,value: %v ", key, val))
	}

	strMap := GetPtrFieldTypeMap(&entity)
	fmt.Println("根据结构体的指针对象获取Map结构")
	for key, val := range strMap {
		fmt.Println(fmt.Sprintf("结构体属性，key: %v ,value: %v ", key, val))
	}

	ptrList := GetPtrTypeFieldNameList(reflect.TypeOf(&entity))
	fmt.Println("根据结构体指针类型获取对应的属性名")
	for _, val := range ptrList {
		fmt.Println(fmt.Sprintf("字段名：%s ", val))
	}

	structList := GetStructTypeFieldNameList(reflect.TypeOf(entity))
	fmt.Println("根据结构体类型获取对应所有字段名称")
	for _, val := range structList {
		fmt.Println(fmt.Sprintf("字段名：%s ", val))
	}
}

type Entity struct {
	Num int
	S   string
	T   map[string]string
}
