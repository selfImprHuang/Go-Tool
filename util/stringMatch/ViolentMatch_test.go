/*
 *  @Author : huangzj
 *  @Time : 2021/2/25 11:11
 *  @Description：
 */

package stringMatch

import (
	"fmt"
	"testing"
)

func TestViolentMatch(t *testing.T) {
	s := "abcbvxs"
	m := "bcb"
	showResult1(s, m)
	s = "abcbvxs"
	m = "bx"
	showResult1(s, m)
	s = "abcbvxssdaadsad"
	m = "cbvxssdaa"
	showResult1(s, m)
}

func showResult1(s, m string) {
	i := ViolentMatch(s, m)
	if i == -1 {
		fmt.Println("未找到对应字符串")
	} else {
		fmt.Println(fmt.Sprintf("匹配到的字符串是:%v", s[i:len(m)+i]))
	}
}
