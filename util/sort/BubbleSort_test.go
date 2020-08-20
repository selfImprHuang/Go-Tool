/*
 *  @Author : huangzj
 *  @Time : 2020/8/20 11:21
 *  @Description：
 */

package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	list1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	BubbleSort(list1)
	for _, r := range list1 {
		fmt.Print(r, "  ")
	}
	fmt.Println()

	list2 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	BubbleSort(list2)
	for _, r := range list2 {
		fmt.Print(r, "  ")
	}
	fmt.Println()

	list3 := []int{1, 432, 12, 67, 341, 874, 56332, 43, 6564, 234, 980, 4234, 6589932, 80, 42, 4234, 55}
	BubbleSort(list3)
	for _, r := range list3 {
		fmt.Print(r, "  ")
	}
	fmt.Println()

	list4 := []int{4523325, 21, 43, 1, 265, 7657, 8754, 234, 543, 536, 2, 6543, 772, 432, 5, 6, 7214, 6754}
	BubbleSort(list4)
	for _, r := range list4 {
		fmt.Print(r, "  ")
	}
	fmt.Println()

}
