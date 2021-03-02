/*
 *  @Author : huangzj
 *  @Time : 2021/3/2 11:21
 *  @Description：
 */

package stringMatch

import (
	"fmt"
	"testing"
)

func TestSundayMatch(t *testing.T) {
	s := "abcbvxs"
	m := "bcb"
	showResult2(s, m)
	s = "abcbvxs"
	m = "bx"
	showResult2(s, m)
	s = "abcbvxssdaadsad"
	m = "cbvxssdaa"
	showResult2(s, m)
}

func showResult2(s, m string) {
	i := SundayMatch(s, m)
	if i == -1 {
		fmt.Println("未找到对应字符串")
	} else {
		fmt.Println(fmt.Sprintf("匹配到的字符串是:%v", s[i:len(m)+i]))
	}
}
