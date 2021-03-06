/*
 *  @Author : huangzj
 *  @Time : 2020/8/21 9:28
 *  @Description：
 */

package sort

import (
	"fmt"
	"testing"
)

func TestCountingSort(t *testing.T) {
	list1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	CountingSort(list1)
	for _, r := range list1 {
		fmt.Print(r, "  ")
	}
	fmt.Println()

	list2 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	result2 := CountingSort(list2)
	for _, r := range result2 {
		fmt.Print(r, "  ")
	}
	fmt.Println()

	list3 := []int{1, 432, 12, 67, 341, 874, 56332, 43, 6564, 234, 980, 4234, 6589932, 80, 42, 4234, 55}
	result3 := CountingSort(list3)
	for _, r := range result3 {
		fmt.Print(r, "  ")
	}
	fmt.Println()

	list4 := []int{4523325, 21, 43, 1, 265, 7657, 8754, 234, 543, 536, 2, 6543, 772, 432, 5, 6, 7214, 6754}
	result4 := CountingSort(list4)
	for _, r := range result4 {
		fmt.Print(r, "  ")
	}
	fmt.Println()

}
