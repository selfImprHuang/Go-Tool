/*
 *  @Author : huangzj
 *  @Time : 2020/8/19 11:45
 *  @Description：
 */

package sort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	list1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	MergeSort(list1)
	for _, r := range list1 {
		fmt.Print(r, "  ")
	}
	fmt.Println()

	list2 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	MergeSort(list2)
	for _, r := range list2 {
		fmt.Print(r, "  ")
	}
	fmt.Println()

	list3 := []int{100, 5, 222, 312, 552, 354, 8, 21, 90, 2, 0, 12, 345, 67}
	MergeSort(list3)
	for _, r := range list3 {
		fmt.Print(r, "  ")
	}
	fmt.Println()

	list4 := []int{1, 19, 42, 7, 3131, 44, 56, 123, 589, 21, 975, 123, 5467, 143, 7, 321, 77, 3}
	MergeSort(list4)
	for _, r := range list4 {
		fmt.Print(r, "  ")
	}
	fmt.Println()
}
