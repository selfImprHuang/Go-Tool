/*
 *  @Author : huangzj
 *  @Time : 2021/3/23 19:53
 *  @Description：
 */

package FullPermutation

import (
	"fmt"
	"testing"
)

func TestDecreaseBit(t *testing.T) {

	bit := []int{0, 0}
	for i := 0; i < 6; i++ {
		fmt.Println(bit)
		bit = decreaseBit(bit)
	}

	fmt.Println()
	bit = []int{0, 0, 0}
	for i := 0; i < 24; i++ {
		fmt.Println(bit)
		bit = decreaseBit(bit)
	}

	fmt.Println()
	bit = []int{0, 0, 0, 0}
	for i := 0; i < 120; i++ {
		fmt.Print(bit)
		bit = decreaseBit(bit)
	}

}
func TestGetRangeBySpace1(t *testing.T) {
	fmt.Println(getRangeBySpace1([]int{1, 2, 3}, []int{0, 0}))
	fmt.Println(getRangeBySpace1([]int{1, 2, 3}, []int{1, 0}))
	fmt.Println(getRangeBySpace1([]int{1, 2, 3}, []int{2, 0}))
	fmt.Println(getRangeBySpace1([]int{1, 2, 3}, []int{0, 1}))
	fmt.Println(getRangeBySpace1([]int{1, 2, 3}, []int{1, 1}))
	fmt.Println(getRangeBySpace1([]int{1, 2, 3}, []int{2, 1}))

	fmt.Println(getRangeBySpace1([]int{1, 2, 3, 4, 5}, []int{0, 0, 0, 0}))
	fmt.Println(getRangeBySpace1([]int{1, 2, 3, 4, 5}, []int{0, 0, 0, 1}))
	fmt.Println(getRangeBySpace1([]int{1, 2, 3, 4, 5}, []int{4, 3, 2, 1}))
	fmt.Println(getRangeBySpace1([]int{1, 2, 3, 4, 5}, []int{3, 0, 0, 0}))
	fmt.Println(getRangeBySpace1([]int{1, 2, 3, 4, 5}, []int{1, 3, 2, 1}))

	fmt.Println(getRangeBySpace1([]int{1, 2, 3, 4, 5, 6, 7}, []int{3, 4, 2, 2, 2, 1}))
}

func TestDigitDecrease(t *testing.T) {
	fmt.Println(DigitDecreases([]int{1, 2, 3}))
	fmt.Println(DigitDecreases([]int{20, 50, 30, 67}))
	fmt.Println(DigitDecreases([]int{20, 50, 30, 99, 18213}))
	fmt.Println(DigitDecreases([]int{20, 50, 30, 78321, 2321, 432}))
}
