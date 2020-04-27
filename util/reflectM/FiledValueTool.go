/*
 *  @Author : huangzj
 *  @Time : 2020/4/24 12:00
 *  @Description：
 */

package reflectM

import "reflect"

/*
 *
 *	方法仅支持【结构体对象】
 *
 * @param 结构体对象
 * @return 字段名和字段值对应的map
 * @description
 */
func GetStructFieldTypeMap(i interface{}) map[string]interface{} {
	//应该在最前面这边有一个判断才对
	if GetValueType(i) != reflect.Struct {
		panic("该方法只支持结构体对象")
	}

	var m map[string]interface{}
	m = make(map[string]interface{})

	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	for i := 0; i < t.NumField(); i++ {
		fn := t.Field(i).Name
		val := v.Field(i).Interface()
		m[fn] = val
	}
	return m
}

/*
 *
 * 	方法仅支持【结构体的指针对象】
 *
 * @param 结构体的指针对象
 * @return 字段名和字段值对应的map
 * @description
 */
func GetPtrFieldTypeMap(i interface{}) map[string]interface{} {
	if GetValueType(i) != reflect.Ptr {
		panic("该方法只支持结构体指针对象")
	}

	var m map[string]interface{}
	m = make(map[string]interface{})

	ind := reflect.Indirect(reflect.ValueOf(i))
	t := ind.Type()
	for i := 0; i < t.NumField(); i++ {
		fn := t.Field(i).Name
		val := ind.Field(i).Interface()
		m[fn] = val
	}

	return m
}

/*
 *
 * 根据结构体类型获取结构体字段的名称数组
 *
 * @param
 * @return
 * @description
 */
func GetStructTypeFieldNameList(t reflect.Type) []string {
	if t.Kind() != reflect.Struct {
		panic("该方法只支持结构体类型")
	}
	list := make([]string, 0)

	for i := 0; i < t.NumField(); i++ {
		list = append(list, t.Field(i).Name)

	}
	return list
}

/*
 *
 * 根据结构体类型的指针获取结构体字段的名称数组
 *
 * @param
 * @return
 * @description
 */
func GetPtrTypeFieldNameList(t reflect.Type) []string {
	if t.Kind() != reflect.Ptr {
		panic("该方法只支持结构体指针类型")
	}
	list := make([]string, 0)

	typ := t.Elem()

	for i := 0; i < typ.NumField(); i++ {
		list = append(list, typ.Field(i).Name)

	}
	return list
}

/*
 * @description 判断传进来的参数是什么类型的
 */
func GetValueType(value interface{}) reflect.Kind {
	return reflect.TypeOf(value).Kind()
}
