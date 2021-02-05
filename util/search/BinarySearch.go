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

//二分查找在两个数值之间的数据
func BinarySearchBetween(list []int, value int) int {
	low := 0
	high := len(list) - 1
	var mid int

	for low <= high {
		//求mid
		mid = (low + high) >> 1
		//要算后面一个元素，所以下标要判断是否超过 并且需要考虑超过数组最大值的情况
		if (mid+1 < len(list) && value > list[mid] && value < list[mid+1]) || value == list[mid] || (mid+1 >= len(list) && value > list[mid]) {
			return mid
		}
		//位移
		if value < list[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	//未找到元素
	return -1
}
