/*
 *  @Author : huangzj
 *  @Time : 2020/8/19 11:44
 *  @Description：这边提供一个递归的归并排序算法，非递归的理解成本比较高，并且还有多路归并排序，需要理解对应概念
 *  最坏时间复杂度O(nlogn)   最好时间复杂度O(logn)   空间复杂度O(n)
 */

package sort

func MergeSort(list []int) {
	middle := len(list) / 2
	leftList := copyList(list, 0, middle)
	rightList := copyList(list, middle, len(list))
	if len(list) > 1 {
		MergeSort(leftList)                 //先对左边进行归并排序
		MergeSort(rightList)                //再对右边进行归并排序
		mergeAll(list, leftList, rightList) //对左右排序的结果进行组合
	}
}

func copyList(list []int, start int, end int) []int {
	result := make([]int, 0)
	for i := start; i < end; i++ {
		result = append(result, list[i])
	}
	return result
}

func mergeAll(array []int, left []int, right []int) {
	l, r, i := 0, 0, 0 //左数组的下标、右数组的下标、组合数组的下标
	for l < len(left) && r < len(right) {
		if left[l] > right[r] {
			array[i] = right[r]
			i++
			r++
		} else {
			array[i] = left[l]
			i++
			l++
		}
	}

	for r < len(right) {
		array[i] = right[r]
		r++
		i++
	}

	for l < len(left) {
		array[i] = left[l]
		l++
		i++
	}
}
