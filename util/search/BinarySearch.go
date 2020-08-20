/*
 *  @Author : huangzj
 *  @Time : 2020/8/18 17:45
 *  @Description：二分查找法，对于有序数组进行查找
 */

package search

func BinarySearch(list []int, value int) int {
	low := 0
	high := len(list) - 1
	var mid int

	for low <= high {
		mid = (low + high) >> 1 //相当于(low + high) / 2
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
