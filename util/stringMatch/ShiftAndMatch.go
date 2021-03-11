/*
 *  @Author : huangzj
 *  @Time : 2021/3/8 14:42
 *  @Description：
 */

package stringMatch

import "fmt"

func ShiftAndMatch(allString, modeString string) int {
	if len(modeString) > 32 {
		panic("暂只支持32位大小")
	}

	bitap := GenerateTable(modeString)

	fmt.Println("bitap的二进制：")
	for i := 0; i < len(modeString); i++ {
		fmt.Println(fmt.Sprintf("当前字符:%c,对应二进制:%08b", modeString[i], bitap[modeString[i]-'a']))
	}

	fmt.Println("")

	var status int
	for i := 0; i < len(allString); i++ {
		status = ((status << 1) | 1) & bitap[allString[i]-'a']
		fmt.Println(fmt.Sprintf("当前字符：%c,对应的二进制结果：%08b", allString[i], status))
		if status&(1<<(len(modeString)-1)) > 0 {
			return i - len(modeString) + 1
		}
	}
	return -1
}

func GenerateTable(modeString string) []int {
	bitap := make([]int, 32)
	for i := 0; i < len(modeString); i++ {
		bitap[modeString[i]-'a'] |= 1 << i
	}
	return bitap
}

func ShiftOrMatch(allString, modeString string) int {
	if len(modeString) > 32 {
		panic("暂只支持32位大小")
	}

	bitap := make([]int, 32)
	for i := 0; i < 32; i++ {
		bitap[i] = ^0
	}

	var shift = 1
	for i := 0; i < len(modeString); i++ {
		bitap[modeString[i]-'a'] &= ^shift
		shift <<= 1
	}

	fmt.Println("bitap的二进制：")
	for i := 0; i < len(modeString); i++ {
		fmt.Println(fmt.Sprintf("当前字符:%c,对应二进制:%08b", modeString[i], uint32(bitap[modeString[i]-'a'])))
	}
	fmt.Println()

	status := ^0
	mask := ^(1 << (len(modeString) - 1))
	fmt.Println(fmt.Sprintf("mask:%08b", uint32(mask)))
	for i := 0; i < len(allString); i++ {
		status = (status << 1) | bitap[allString[i]-'a']
		fmt.Println(fmt.Sprintf("当前字符：%c,对应的二进制结果：%08b", allString[i], uint32(status)))
		if ^(status | mask) > 0 {
			return i - len(modeString) + 1
		}
	}

	return -1
}
