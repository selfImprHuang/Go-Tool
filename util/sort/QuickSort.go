/*
 *  @Author : huangzj
 *  @Time : 2020/8/19 17:06
 *  @Description：
 */

package sort

import "Go-Tool/util/search"

//通过hoare划分来实现快速排序
func QuickSort(list []int, start, end int) {
	if start < end {
		position := hoarePartition(list, start, end)
		QuickSort(list, start, position) //注意这边结束的下表是position，而不是position + 1
		QuickSort(list, position+1, end)
	}
}

//通过lomuto划分来实现快速排序
func QuickSort1(list []int, start, end int) {
	if start < end {
		position := search.LomutoPartition(list, start, end)
		QuickSort(list, start, position)
		QuickSort(list, position+1, end)
	}
}

func hoarePartition(list []int, start int, end int) int {
	value := list[start] //比较的元素
	s := start
	e := end
	for s < e {
		//找到数组终止前，比目标数(value)大的数字
		for ; s < end && list[s] < value; s++ {
		}
		//找到数组结束前，比目标数(value)小的数字
		for ; e > start && list[e] >= value; e-- {
		}
		if s < e {
			list[s], list[e] = list[e], list[s]
		}
	}
	list[start], list[e] = list[e], list[start]
	return e
}
