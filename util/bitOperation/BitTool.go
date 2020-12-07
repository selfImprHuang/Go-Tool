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
		panic(fmt.Sprintf("输入有误：Position不能小于1，当前出入的Position为%d", position))
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
	//解析：num & (num -1) 当num为2的N次方的时候, 2N次方 & (2N次方 - 1) ，因为做了最高位的右移（2N次方只有一个1，2N次方-1除了2N次方的那一个1的位置为0，其他都为1）
	//所以&的结果就是 0 ，而0 & 2N次方一定是0，所以如果结果是0，那么一定是2N次方，只有一个位置有1
	//那么 num & (num -1) 当num不是2的N次方的时候,他的结果原来有1的最高位置，一定为1，所以再和num进行 & 操作的时候,结果也一定是大于0，所以不止一个1
	return num&(num&(num-1)) == 0
}

//设置num的某一个位置为1，如果原来就是1的话，就不发生变化
func SetOneByPosition(num int, position int) int {
	if position < 1 {
		panic(fmt.Sprintf("输入有误：Position不能小于1，当前出入的Position为%d", position))
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
		panic(fmt.Sprintf("输入有误：Position不能小于1，当前出入的Position为%d", position))
	}
	return num &^ (1 << uint(position-1))
}
