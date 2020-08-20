/*
 *  @Author : huangzj
 *  @Time : 2020/8/19 10:24
 *  @Description：
 *
 */

package search

func InsertionSearch(list []int, value int) int {
	low := 0
	high := len(list) - 1
	var mid int

	//插值查找这边如果数字和起始数字差别大的话,计算出来的下标会越界挺多的
	if value > list[high-1] || value < list[0] {
		return -1
	}

	for low < high {
		mid = low + (high-low)*(value-list[low])/(list[high]-list[low])
		if value == list[mid] {
			return mid
		}

		if value < list[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
