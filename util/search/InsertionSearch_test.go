/*
 *  @Author : huangzj
 *  @Time : 2020/8/19 10:38
 *  @Description：
 */

package search

import (
	"fmt"
	"testing"
)

func TestInsertionSearch(t *testing.T) {

	evenNumList := []int{1, 2, 3, 5, 6, 7, 9, 10, 22, 111, 333, 431}

	for _, r := range evenNumList {
		position := InsertionSearch(evenNumList, r)
		fmt.Println(position)
	}

	position := InsertionSearch(evenNumList, 999)
	fmt.Println(position)

	position = InsertionSearch(evenNumList, -10)
	fmt.Println(position)

	fmt.Println()

	oddNumList := []int{1, 2, 3, 5, 6, 7, 9, 10, 22, 111, 333, 431, 3211}
	for _, r := range oddNumList {
		position := InsertionSearch(oddNumList, r)
		fmt.Println(position)
	}

	position = InsertionSearch(oddNumList, 999)
	fmt.Println(position)

	position = InsertionSearch(oddNumList, -10)
	fmt.Println(position)
}
