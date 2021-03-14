/*
 *  @Author : huangzj
 *  @Time : 2021/3/13 21:58
 *  @Description：
 */

package palindrome

import (
	"fmt"
	"testing"
)

func TestGenerateString(t *testing.T) {
	fmt.Println(generateString("abcdef"))
	fmt.Println(generateString("abc2342def"))
	fmt.Println(generateString("abcdef3213"))
	fmt.Println(generateString("abcderewrewr312f"))
	fmt.Println(generateString("312312abcdef321e"))
}

func TestPalindrome(t *testing.T) {
	fmt.Println(Manacher("abcabccbacba"))
	fmt.Println(Manacher("mabcabccbacbammm"))
	fmt.Println(Manacher("habcabccbacbammm"))
	fmt.Println(Manacher("hmmmabcabccbacbammm"))
	fmt.Println(Manacher("abcdefghij"))
}
