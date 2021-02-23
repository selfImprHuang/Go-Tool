/*
 *  @Author : huangzj
 *  @Time : 2020/12/3 11:56
 *  @Description：比特位操作的工作类，主要参考：https://learnku.com/go/t/23460/bit-operation-of-go 实现的工具
 */

package bitOperation

import (
	"fmt"
)

//判断num的二进制表示某个位置是否为1
func JudgeOneWithPosition(num int, position int) bool {
	if position < 1 {
		panic(fmt.Sprintf("输入有误：Position不能小于1，"+
			"当前出入的Position为%d", position))
	}
	return num&(1<<uint(position-1)) != 0
}

//判断num是不是偶数
func JudgeIfEven(num int) bool {
	return num&1 == 0
}

//判断是不是奇数
func JudgeIfOdd(num int) bool {
	return num&1 != 0
}

//判断一个数的二进制形式是否只有一个位置为1,这个方法可以用来判断是不是2的n次方
func JudgeOnlyOneBitWithOne(num int) bool {
	if num == 0 {
		return false
	}
	return num&(num-1) == 0
}

//设置num的某一个位置为1，如果原来就是1的话，就不发生变化
func SetOneByPosition(num int, position int) int {
	if position < 1 {
		panic(fmt.Sprintf("输入有误：Position不能小于1，"+
			"当前出入的Position为%d", position))
	}
	return num | (1 << uint(position-1))
}

//判断两个数字的加减符号是否相同
func JudgeWithEqualSymbol(num1 int, num2 int) bool {
	return num1^num2 >= 0
}

//设置某一个数num的二进制Position位为0
func SetZeroBuPosition(num int, position int) int {
	if position < 1 {
		panic(fmt.Sprintf("输入有误：Position不能小于1，"+
			"当前出入的Position为%d", position))
	}
	return num &^ (1 << uint(position-1))
}
