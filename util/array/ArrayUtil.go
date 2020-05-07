/*
 *  @Author : huangzj
 *  @Time : 2020/4/30 15:53
 *  @Description: 数组工具
 *  这边测试了一下，go的结构体或者是非结构体类型，通过==比较的话，会校验到对应的每一个属性的值。
 *  也就是说就算两个结构体的指针地址不一样，但是他的所有属性值都是一样的话，==返回的结果也是true
 *  这边直接在入口的地方进行了map、slice、interface、Func、chan、ptr、UnsafePointer 类型的拦截
 *	但是如果结构体中包含不能比较的类型，应该也是会报错
 */

package array

import "reflect"

/*
 * 返回数组中是否包含对应元素的结果
 */
func Contains(list []interface{}, ele interface{}) bool {
	for _, row := range list {
		if reflect.DeepEqual(row, ele) {
			return true
		}
	}
	return false
}

/*
 * 返回数组中是否不包含对应元素的结果
 */
func NotContains(list []interface{}, ele interface{}) bool {
	return !Contains(list, ele)
}

/*
 * 返回两个数组的所有元素是否相等的结果
 */
func IsSameList(list, eList []interface{}) bool {
	if len(list) != len(eList) {
		return false
	}

	return reflect.DeepEqual(list, eList)

}

/*
 * @param fList 传父集合
 * @param sList 传子集合
 * @return sList是不是fList的子集
 */
func IsSubSet(fList, sList []interface{}) bool {
	if len(fList) < len(sList) {
		return false
	}
	for _, r := range sList {
		if NotContains(fList, r) {
			return false
		}
	}
	return true
}

/*
 * 获取两个集合的交集
 */
func Intersection(list, otherList []interface{}) []interface{} {
	result := make([]interface{}, 0)
	for _, row := range list {
		if Contains(otherList, row) {
			result = append(result, row)
		}
	}
	return result
}

/*
 * 返回前一个数组包含，后一个数组不包含的结果(不是数量大的数组-数量小的数组)
 */
func DiffSet(cList, nList []interface{}) []interface{} {
	result := make([]interface{}, 0)
	for _, row := range cList {
		if NotContains(nList, row) {
			result = append(result, row)
		}
	}
	return result
}

/*
 * 返回两个数组的并集,不包含重复的数据
 */
func Union(aList, bList []interface{}) []interface{} {
	result := make([]interface{}, 0)

	for _, row := range aList {
		if NotContains(result, row) {
			result = append(result, row)
		}
	}
	for _, row := range bList {
		if NotContains(result, row) {
			result = append(result, row)
		}
	}
	return result
}
