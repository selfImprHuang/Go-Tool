/*
 *  @Author : huangzj
 *  @Time : 2021/3/23 19:54
 *  @Description：
 */

package FullPermutation

import (
	"fmt"
	"testing"
)

func TestGetRangeBySpace(t *testing.T) {
	fmt.Println(getRangeBySpace([]int{1, 2, 3}, []int{0, 0})) //123
	fmt.Println(getRangeBySpace([]int{1, 2, 3}, []int{1, 0})) //
	fmt.Println(getRangeBySpace([]int{1, 2, 3}, []int{0, 1})) //
	fmt.Println(getRangeBySpace([]int{1, 2, 3}, []int{1, 1})) //
	fmt.Println(getRangeBySpace([]int{1, 2, 3}, []int{0, 2})) //
	fmt.Println(getRangeBySpace([]int{1, 2, 3}, []int{1, 2})) //321

	//fmt.Println(getRangeBySpace([]int{1, 2, 3, 4, 5}, []int{0, 0, 0, 0})) //12345
	//fmt.Println(getRangeBySpace([]int{1, 2, 3, 4, 5}, []int{0, 0, 0, 1})) //21354
	fmt.Println(getRangeBySpace([]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4})) //54321
	fmt.Println(getRangeBySpace([]int{1, 2, 3, 4, 5}, []int{0, 0, 0, 3})) //12534
	fmt.Println(getRangeBySpace([]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 1})) //43251

	fmt.Println(getRangeBySpace([]int{1, 2, 3, 4, 5, 6, 7}, []int{1, 2, 2, 2, 4, 3})) //3647521
}

func TestIncreaseBit(t *testing.T) {
	bit := []int{0, 0}
	for i := 0; i < 6; i++ {
		fmt.Println(bit)
		bit = increaseBit(bit)
	}
	fmt.Println()

	bit = []int{0, 0, 0, 0}
	for i := 0; i < 120; i++ {
		bit = increaseBit(bit)
		fmt.Println(bit)
	}

	bit = []int{0, 0, 0}
	for i := 0; i < 24; i++ {
		bit = increaseBit(bit)
		fmt.Println(bit)
	}
}

func TestDigitIncrease(t *testing.T) {
	fmt.Println(DigitIncrease([]int{1, 2, 3}))
	fmt.Println(DigitIncrease([]int{20, 50, 30, 67}))
	fmt.Println(DigitIncrease([]int{20, 50, 30, 99, 18213}))
	fmt.Println(DigitIncrease([]int{20, 50, 30, 78321, 2321, 432}))
}
