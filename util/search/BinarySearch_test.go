/*
 *  @Author : huangzj
 *  @Time : 2020/8/18 17:59
 *  @Description：
 */

package search

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {

	evenNumList := []int{1, 2, 3, 5, 6, 7, 9, 10, 22, 111, 333, 431}

	for _, r := range evenNumList {
		position := BinarySearch(evenNumList, r)
		fmt.Println(position)
	}

	position := BinarySearch(evenNumList, 999)
	fmt.Println(position)

	position = BinarySearch(evenNumList, -10)
	fmt.Println(position)

	fmt.Println()

	oddNumList := []int{1, 2, 3, 5, 6, 7, 9, 10, 22, 111, 333, 431, 3211}
	for _, r := range oddNumList {
		position := BinarySearch(oddNumList, r)
		fmt.Println(position)
	}

	position = BinarySearch(oddNumList, 999)
	fmt.Println(position)

	position = BinarySearch(oddNumList, -10)
	fmt.Println(position)
}
