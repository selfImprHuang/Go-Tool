/*
 *  @Author : huangzj
 *  @Time : 2021/3/22 11:22
 *  @Description：
 */

package FullPermutation

import (
	"fmt"
	"testing"
)

func TestRecursion(t *testing.T) {
	p := NewPermutation()

	p.Recursion([]byte("123"), 0)
	fmt.Println()
	for _, bs := range p.bytes {
		fmt.Println(string(bs))
	}
	p.Reset()

	p.Recursion([]byte("12345"), 0)
	fmt.Println()
	for _, bs := range p.bytes {
		fmt.Println(string(bs))
	}
	p.Reset()

	p.Recursion([]byte("abcde"), 0)
	fmt.Println()
	for _, bs := range p.bytes {
		fmt.Println(string(bs))
	}
	p.Reset()

	p.Recursion([]byte("abcdea"), 0)
	fmt.Println()
	for _, bs := range p.bytes {
		fmt.Println(string(bs))
	}
	p.Reset()

}
