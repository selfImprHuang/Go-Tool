/*
 *  @Author : huangzj
 *  @Time : 2020/8/21 9:28
 *  @Description：计数排序,时间复杂度：Ο (n+k)、空间复杂度是 O(k)
 *  限制：1.当数列最大最小值差距过大时，并不适用计数排序。2.当数列元素不是整数，并不适用计数排序。
 *  参考地址：https://blog.csdn.net/csdnnews/article/details/83005778
 */

package sort

import "math"

func CountingSort(list []int) []int {
	min, max := findMinAndMax(list)
	countArray := make([]int, max-min+1)
	for index := range list {
		countArray[list[index]-min]++ //根据偏移量对数组某个下标的元素进行增一，如果这个位置元素存在
	}

	//创建结果数组，这边其实可以直接输出或者使用原来的数组
	result := make([]int, len(list))
	pos := 0
	for index := range countArray {
		for countArray[index] != 0 {
			result[pos] = index + min
			pos++
			countArray[index]--
		}
	}

	return result
}

func findMinAndMax(list []int) (int, int) {
	min := math.MaxInt64
	max := math.MinInt64
	//找到数组中的最大值和最小值
	for _, num := range list {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	return min, max
}
