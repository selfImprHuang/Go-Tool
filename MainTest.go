/*
 *  @Author : huangzj
 *  @Time : 2020/4/24 14:51
 *  @Description：
 */

package main

import (
	"Go-Tool/util/reflectM"
	"fmt"
	"math"
	"reflect"
)

func main() {
	//fmt.Println(util.Crc32Mode("ewqeqweqw", 4))
	//fmt.Println(util.Crc32Mode("sdasdsa", 4))
	//fmt.Println(util.Crc32Mode("dasdsadasdas", 4))
	//fmt.Println(util.Crc32Mode("321321312", 4))
	//fmt.Println(util.Crc32Mode("dsadas", 4))
	//fmt.Println(util.Crc32Mode("dasd23132131", 4))
	//fmt.Println(util.Crc32Mode("eqwewqewqd4312312", 4))
	//fmt.Println(util.Crc32Mode("3123213dasdasdas", 4))
	//fmt.Println(util.Crc32Mode("312321dfdsfsdg", 4))
	//fmt.Println(util.Crc32Mode("ewqrewrqwe", 4))
	fmt.Println(int(math.Ceil(1.1)))
	fmt.Println(int(math.Ceil(1.6)))
	fmt.Println(int(math.Floor(1.1)))
	fmt.Println(int(math.Floor(1.6)))
	//testReflect()
}

func testReflect() {
	entity := Entity{
		Num: 1,
		S:   "1223",
		T:   make(map[string]string),
	}
	structMap := reflectM.GetStructFieldTypeMap(entity)
	fmt.Println("根据结构体获取Map结构")
	for key, val := range structMap {
		fmt.Println(fmt.Sprintf("结构体属性，key: %v ,value: %v ", key, val))
	}

	strMap := reflectM.GetPtrFieldTypeMap(&entity)
	fmt.Println("根据结构体的指针对象获取Map结构")
	for key, val := range strMap {
		fmt.Println(fmt.Sprintf("结构体属性，key: %v ,value: %v ", key, val))
	}

	ptrList := reflectM.GetPtrTypeFieldNameList(reflect.TypeOf(&entity))
	fmt.Println("根据结构体指针类型获取对应的属性名")
	for _, val := range ptrList {
		fmt.Println(fmt.Sprintf("字段名：%s ", val))
	}

	structList := reflectM.GetStructTypeFieldNameList(reflect.TypeOf(entity))
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
