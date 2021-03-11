/*
 *  @Author : huangzj
 *  @Time : 2021/3/8 14:51
 *  @Description：
 */

package stringMatch

import (
	"fmt"
	"testing"
)

func TestShiftAndMatch(t *testing.T) {
	//fmt.Println("该子串的位置为：", ShiftAndMatch("abcdefghijk", "abcd")) //输出0
	//fmt.Println("该子串的位置为：", ShiftAndMatch("abcdefghijk", "bcd"))  //输出1
	//fmt.Println("该子串的位置为：", ShiftAndMatch("abcdefghijk", "fgh"))  //输出5
	//fmt.Println("该子串的位置为：", ShiftAndMatch("abcdefghijk", "hijk")) //输出7
	//fmt.Println("该子串的位置为：", ShiftAndMatch("abcdefghijk", "jk"))   //输出9
	//fmt.Println("该子串的位置为：", ShiftAndMatch("abcdefghijk", "k"))    //输出10
	//
	//fmt.Println("该子串的位置为：", ShiftAndMatch("abcdefghijk", "kbz"))      //输出-1
	//fmt.Println("该子串的位置为：", ShiftAndMatch("abcdefghijk", "kdasdsad")) //输出-1
	//fmt.Println("该子串的位置为：", ShiftAndMatch("abcdefghijk", "dsadsa"))   //输出-1

	fmt.Println("该子串的位置为：", ShiftAndMatch("cbcbcbaefd", "cbcba"))
}

func TestShiftOrMatch(t *testing.T) {
	//fmt.Println("该子串的位置为：", ShiftOrMatch("abcdefghijk", "abcd")) //输出0
	//fmt.Println("该子串的位置为：", ShiftOrMatch("abcdefghijk", "bcd"))  //输出1
	//fmt.Println("该子串的位置为：", ShiftOrMatch("abcdefghijk", "fgh"))  //输出5
	//fmt.Println("该子串的位置为：", ShiftOrMatch("abcdefghijk", "hijk")) //输出7
	//fmt.Println("该子串的位置为：", ShiftOrMatch("abcdefghijk", "jk"))   //输出9
	//fmt.Println("该子串的位置为：", ShiftOrMatch("abcdefghijk", "k"))    //输出10
	//
	//fmt.Println("该子串的位置为：", ShiftOrMatch("abcdefghijk", "kbz"))      //输出-1
	//fmt.Println("该子串的位置为：", ShiftOrMatch("abcdefghijk", "kdasdsad")) //输出-1
	//fmt.Println("该子串的位置为：", ShiftOrMatch("abcdefghijk", "dsadsa"))   //输出-1

	fmt.Println("该子串的位置为：", ShiftOrMatch("cbcbcbaefd", "cbcba"))
}
