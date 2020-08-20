/*
 *  @Author : huangzj
 *  @Time : 2020/8/20 11:28
 *  @Description：
 */

package sort

func SelectSort(list []int) {
	var minPoint int
	for i := 0; i < len(list); i++ {
		minPoint = i
		for j := i + 1; j < len(list); j++ {
			if list[j] < list[minPoint] {
				minPoint = j
			}
		}
		list[i], list[minPoint] = list[minPoint], list[i]
	}
}
