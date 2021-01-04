/*
 *  @Author : huangzj
 *  @Time : 2021/1/4 15:44
 *  @Description：
 */

package reflectM

import "reflect"

//获取value值，包含ptr的处理
func Indirect(reflectValue reflect.Value) reflect.Value {
	for reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}
	return reflectValue
}

//获取Type值，包含ptr的处理
func IndirectType(reflectType reflect.Type) reflect.Type {
	for reflectType.Kind() == reflect.Ptr || reflectType.Kind() == reflect.Slice {
		reflectType = reflectType.Elem()
	}
	return reflectType
}
