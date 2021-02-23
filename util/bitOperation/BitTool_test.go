/*
 *  @Author : huangzj
 *  @Time : 2020/12/3 14:11
 *  @Description：
 */

package bitOperation

import (
	"fmt"
	"testing"
)

func TestSetZeroBuPosition(t *testing.T) {
	fmt.Println("设置某个位置上的元素为0")
	fmt.Println(fmt.Sprintf("原始: %16b", 1))
	fmt.Println(fmt.Sprintf("结果：%16b", SetZeroBuPosition(1, 1)))
	fmt.Println(fmt.Sprintf("原始: %16b", 10))
	fmt.Println(fmt.Sprintf("结果：%16b", SetZeroBuPosition(10, 2)))
	fmt.Println(fmt.Sprintf("原始: %16b", 20))
	fmt.Println(fmt.Sprintf("结果：%16b", SetZeroBuPosition(20, 5)))
	fmt.Println(fmt.Sprintf("原始: %16b", 101))
	fmt.Println(fmt.Sprintf("结果：%16b", SetZeroBuPosition(101, 6)))
	fmt.Println(fmt.Sprintf("原始: %16b", 20121))
	fmt.Println(fmt.Sprintf("结果：%16b", SetZeroBuPosition(20121, 8)))
	fmt.Println(fmt.Sprintf("原始: %16b", 821331))
	fmt.Println(fmt.Sprintf("结果：%16b", SetZeroBuPosition(821331, 12)))
	fmt.Println(fmt.Sprintf("原始: %16b", 821331))
	fmt.Println(fmt.Sprintf("结果：%16b", SetZeroBuPosition(821331, 3)))
	fmt.Println()
	fmt.Println()
}

func TestSetOneByPosition(t *testing.T) {
	fmt.Println("设置某个位置上的元素为1")
	fmt.Println(fmt.Sprintf("原始: %16b", 2))
	fmt.Println(fmt.Sprintf("结果：%16b", SetOneByPosition(2, 1)))
	fmt.Println(fmt.Sprintf("原始: %16b", 10))
	fmt.Println(fmt.Sprintf("结果：%16b", SetOneByPosition(10, 1)))
	fmt.Println(fmt.Sprintf("原始: %16b", 20))
	fmt.Println(fmt.Sprintf("结果：%16b", SetOneByPosition(20, 1)))
	fmt.Println(fmt.Sprintf("原始: %16b", 101))
	fmt.Println(fmt.Sprintf("结果：%16b", SetOneByPosition(101, 4)))
	fmt.Println(fmt.Sprintf("原始: %16b", 20121))
	fmt.Println(fmt.Sprintf("结果：%16b", SetOneByPosition(20121, 6)))
	fmt.Println(fmt.Sprintf("原始: %16b", 821331))
	fmt.Println(fmt.Sprintf("结果：%16b", SetOneByPosition(821331, 11)))
	fmt.Println(fmt.Sprintf("原始: %16b", 821331))
	fmt.Println(fmt.Sprintf("结果：%16b", SetOneByPosition(821331, 2)))

	fmt.Println()
	fmt.Println()
}

func TestJudgeWithEqualSymbol(t *testing.T) {
	fmt.Println("判断两个数正负符号是否一致")
	fmt.Println(JudgeWithEqualSymbol(10, 20))
	fmt.Println(JudgeWithEqualSymbol(10, -20))
	fmt.Println(JudgeWithEqualSymbol(-10, -20))
	fmt.Println()
	fmt.Println()
}

func TestJudgeOneWithPosition(t *testing.T) {
	fmt.Println("判断二进制表示的某个位上是否为1")
	fmt.Println(fmt.Sprintf("%16b", 1))
	fmt.Println(JudgeOneWithPosition(1, 1))
	fmt.Println(JudgeOneWithPosition(1, 2))
	fmt.Println(fmt.Sprintf("%16b", 10))
	fmt.Println(JudgeOneWithPosition(10, 2))
	fmt.Println(JudgeOneWithPosition(10, 3))
	fmt.Println(fmt.Sprintf("%16b", 501))
	fmt.Println(JudgeOneWithPosition(501, 7))
	fmt.Println(JudgeOneWithPosition(501, 2))
	fmt.Println()
	fmt.Println()
}

func TestJudgeIfOdd(t *testing.T) {
	fmt.Println("奇数的判断")
	fmt.Println(JudgeIfOdd(1))
	fmt.Println(JudgeIfOdd(2))
	fmt.Println(JudgeIfOdd(101))
	fmt.Println(JudgeIfOdd(200))
	fmt.Println(JudgeIfOdd(91231))
	fmt.Println(JudgeIfOdd(8321312))
	fmt.Println()
	fmt.Println()
}

func TestJudgeIfEven(t *testing.T) {
	fmt.Println("偶数的判断")
	fmt.Println(JudgeIfEven(1))
	fmt.Println(JudgeIfEven(2))
	fmt.Println(JudgeIfEven(101))
	fmt.Println(JudgeIfEven(200))
	fmt.Println(JudgeIfEven(91231))
	fmt.Println(JudgeIfEven(8321312))
	fmt.Println()
	fmt.Println()
}

func TestJudgeOnlyOneBitWithOne(t *testing.T) {
	fmt.Println("判断二进制表示是不是只有一个位为1")
	for i := 0; i <= 32; i++ {
		fmt.Println(fmt.Sprintf("%d是不是2的n次方，结果为%v", i, JudgeOnlyOneBitWithOne(i)))
	}
	fmt.Println()
	fmt.Println(JudgeOnlyOneBitWithOne(-1))
	fmt.Println(JudgeOnlyOneBitWithOne(-2))
	fmt.Println(JudgeOnlyOneBitWithOne(-3))
	fmt.Println(JudgeOnlyOneBitWithOne(-4))
	fmt.Println(JudgeOnlyOneBitWithOne(-56))
	fmt.Println(JudgeOnlyOneBitWithOne(-64))
	fmt.Println()
	fmt.Println()
}
