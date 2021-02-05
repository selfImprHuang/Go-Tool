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

func TestBinarySearchBetween(t *testing.T) {

	evenNumList := []int{1, 2, 3, 5, 6, 7, 9, 10, 22, 111, 333, 431}

	for _, r := range evenNumList {
		position := BinarySearchBetween(evenNumList, r)
		fmt.Println(position)
	}

	position := BinarySearchBetween(evenNumList, 999)
	fmt.Println(position)

	position = BinarySearchBetween(evenNumList, -10)
	fmt.Println(position)

	position = BinarySearchBetween(evenNumList, 1)
	fmt.Println(position)

	position = BinarySearchBetween(evenNumList, 3211)
	fmt.Println(position)

	position = BinarySearchBetween(evenNumList, 334)
	fmt.Println(position)

	position = BinarySearchBetween(evenNumList, 110)
	fmt.Println(position)

	position = BinarySearchBetween(evenNumList, 3222)
	fmt.Println(position)

	fmt.Println()

	oddNumList := []int{1, 2, 3, 5, 6, 7, 9, 10, 22, 111, 333, 431, 3211}
	for _, r := range oddNumList {
		position := BinarySearchBetween(oddNumList, r)
		fmt.Println(position)
	}

	fmt.Println("测试between")
	fmt.Println()

	position = BinarySearchBetween(oddNumList, 999)
	fmt.Println(position)

	position = BinarySearchBetween(oddNumList, -10)
	fmt.Println(position)

	position = BinarySearchBetween(oddNumList, 1)
	fmt.Println(position)

	position = BinarySearchBetween(oddNumList, 3211)
	fmt.Println(position)

	position = BinarySearchBetween(oddNumList, 334)
	fmt.Println(position)

	position = BinarySearchBetween(oddNumList, 110)
	fmt.Println(position)

	position = BinarySearchBetween(oddNumList, 3222)
	fmt.Println(position)
}
