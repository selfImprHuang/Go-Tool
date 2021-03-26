/*
 *  @Author : huangzj
 *  @Time : 2021/3/22 15:29
 *  @Description：
 */

package FullPermutation

import (
	"fmt"
	"testing"
)

func TestDictionaryOrder(t *testing.T) {
	p := NewPermutation()

	p.DictionaryOrder([]byte("1234"))
	fmt.Println()
	for _, bs := range p.bytes {
		fmt.Println(string(bs))
	}
	p.Reset()

	p.DictionaryOrder([]byte("12345"))
	fmt.Println()
	for _, bs := range p.bytes {
		fmt.Println(string(bs))
	}
	p.Reset()

	p.DictionaryOrder([]byte("abcde"))
	fmt.Println()
	for _, bs := range p.bytes {
		fmt.Println(string(bs))
	}
	p.Reset()

	p.DictionaryOrder([]byte("abcdea"))
	fmt.Println()
	for _, bs := range p.bytes {
		fmt.Println(string(bs))
	}
	p.Reset()
}

func TestPermutation_Reverse(t *testing.T) {
	p := NewPermutation()
	fmt.Println(p.Reverse([]byte{1, 2, 3, 4}, 0, 3))
	fmt.Println(p.Reverse([]byte{1, 2, 3, 4, 5}, 0, 4))
	fmt.Println(p.Reverse([]byte{1, 2, 3, 4, 5, 6}, 0, 5))
	fmt.Println(p.Reverse([]byte{1, 2, 3, 4, 5, 6, 7}, 0, 6))
}
