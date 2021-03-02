/*
 *  @Author : huangzj
 *  @Time : 2021/2/24 17:52
 *  @Description：
 */

package stringMatch

import (
	"fmt"
	"testing"
)

func TestKmpMatch(t *testing.T) {
	s := "abcbvxs"
	m := "bcb"
	showResult(s, m)
	s = "abcbvxs"
	m = "bx"
	showResult(s, m)
	s = "abcbvxssdaadsad"
	m = "cbvxssdaa"
	showResult(s, m)
}

func TestMatchNum(t *testing.T) {
	s := "ababababca"
	m := "abababca"
	showResult(s, m)
}

func showResult(s, m string) {
	i := KmpMatch(s, m)
	if i == -1 {
		fmt.Println("未找到对应字符串")
	} else {
		fmt.Println(fmt.Sprintf("匹配到的字符串是:%v", s[i:len(m)+i]))
	}
}

func TestGetNext(t *testing.T) {
	next := getNext("abababca")
	fmt.Println(next)
	next1 := getNext("abababc")
	fmt.Println(next1)
	next2 := getNext("abababcxy")
	fmt.Println(next2)
	next3 := getNext("abababcxyxyy")
	fmt.Println(next3)
	next4 := getNext("abcdabd")
	fmt.Println(next4)
	next5 := getNext("abcxabcxabcyabcz")
	fmt.Println(next5)
}
