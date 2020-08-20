/*
 *  @Author : huangzj
 *  @Time : 2020/8/20 11:21
 *  @Description：冒泡排序
 */

package sort

func BubbleSort(list []int) {
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}
