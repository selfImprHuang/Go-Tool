/*
 *  @Author : huangzj
 *  @Time : 2020/8/19 17:57
 *  @Description：
 */

package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	QuickSort(list, 0, len(list)-1)

	for _, r := range list {
		fmt.Print(r, "   ")
	}
	fmt.Println()

	list1 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	QuickSort(list1, 0, len(list1)-1)

	for _, r := range list1 {
		fmt.Print(r, "   ")
	}
	fmt.Println()

	list2 := []int{432, 412, 213, 456, 7, 213, 4, 78, 13, 87, 312321, 456, 7, 231, 234, 3}
	QuickSort(list2, 0, len(list2)-1)

	for _, r := range list2 {
		fmt.Print(r, "   ")
	}
	fmt.Println()

	list3 := []int{1, 312, 356, 76, 2, 54, 78, 321, 3455, 65732, 123534, 54351, 4213, 123, 56, 6, 732, 123, 44, 51, 5, 61, 543}
	QuickSort(list3, 0, len(list3)-1)

	for _, r := range list3 {
		fmt.Print(r, "   ")
	}
	fmt.Println()
}

func TestQuickSort1(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	QuickSort1(list, 0, len(list)-1)

	for _, r := range list {
		fmt.Print(r, "   ")
	}
	fmt.Println()

	list1 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	QuickSort1(list1, 0, len(list1)-1)

	for _, r := range list1 {
		fmt.Print(r, "   ")
	}
	fmt.Println()

	list2 := []int{432, 412, 213, 456, 7, 213, 4, 78, 13, 87, 312321, 456, 7, 231, 234, 3}
	QuickSort1(list2, 0, len(list2)-1)

	for _, r := range list2 {
		fmt.Print(r, "   ")
	}
	fmt.Println()

	list3 := []int{1, 312, 356, 76, 2, 54, 78, 321, 3455, 65732, 123534, 54351, 4213, 123, 56, 6, 732, 123, 44, 51, 5, 61, 543}
	QuickSort1(list3, 0, len(list3)-1)

	for _, r := range list3 {
		fmt.Print(r, "   ")
	}
	fmt.Println()
}
