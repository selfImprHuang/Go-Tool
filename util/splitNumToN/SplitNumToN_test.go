/*
 *  @Author : huangZJ
 *  @Time : 2021/6/16 9:50
 *  @Description：
 */

package splitNumToN

import (
	"fmt"
	"testing"
)

func TestSplitNumToNWith2(t *testing.T) {

	fmt.Println("测试两个数的情况")
	for i := 3; i <= 11; i++ {
		fmt.Println()
		fmt.Println(fmt.Sprintf("当前随机数为%d", i))
		for j := 0; j < 20; j++ {
			splitter := NewSplitter(2, i)
			splitter.SplitNumToN()
			for _, num := range splitter.GetRandNumList() {
				fmt.Print(fmt.Sprintf("%d ", num))
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func TestSplitNumToNWith3(t *testing.T) {
	fmt.Println()
	fmt.Println("测试三个数的情况")
	for i := 4; i <= 17; i++ {
		fmt.Println()
		fmt.Println(fmt.Sprintf("当前随机数为%d", i))
		for j := 0; j < 20; j++ {
			splitter := NewSplitter(3, i)
			fmt.Println()
			splitter.SplitNumToN()
			for _, num := range splitter.GetRandNumList() {
				fmt.Print(fmt.Sprintf("%d ", num))
			}
		}
		fmt.Println()
	}
}

func TestSplitNumToNWith4(t *testing.T) {
	fmt.Println()
	fmt.Println("测试四个数的情况")
	for i := 5; i <= 23; i++ {
		fmt.Println()
		fmt.Println(fmt.Sprintf("当前随机数为%d", i))
		for j := 0; j < 100; j++ {
			splitter := NewSplitter(4, i)
			fmt.Println()
			splitter.SplitNumToN()
			for _, num := range splitter.GetRandNumList() {
				fmt.Print(fmt.Sprintf("%d ", num))
			}
		}
		fmt.Println()
	}
}

func TestSplitNumToNWith5(t *testing.T) {
	fmt.Println()
	fmt.Println("测试五个数的情况")
	for i := 6; i <= 29; i++ {
		fmt.Println()
		fmt.Println(fmt.Sprintf("当前随机数为%d", i))
		for j := 0; j < 100; j++ {
			splitter := NewSplitter(5, i)
			fmt.Println()
			splitter.SplitNumToN()
			for _, num := range splitter.GetRandNumList() {
				fmt.Print(fmt.Sprintf("%d ", num))
			}
		}
		fmt.Println()
	}

}

func TestSplitNumToNWith6(t *testing.T) {
	fmt.Println()
	fmt.Println("测试六个数的情况")
	for i := 7; i <= 35; i++ {
		fmt.Println()
		fmt.Println(fmt.Sprintf("当前随机数为%d", i))
		for j := 0; j < 100; j++ {
			splitter := NewSplitter(6, i)
			fmt.Println()
			splitter.SplitNumToN()
			for _, num := range splitter.GetRandNumList() {
				fmt.Print(fmt.Sprintf("%d ", num))
			}
		}
		fmt.Println()
	}
}
