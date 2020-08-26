/*
 *  @Author : huangzj
 *  @Time : 2020/8/21 16:07
 *  @Description：
 */

package stringMatch

import (
	"fmt"
	"testing"
)

func TestHorspool(t *testing.T) {
	fmt.Println("该子串的位置为：", Horspool("abcdefghijk", "abcd")) //输出0
	fmt.Println("该子串的位置为：", Horspool("abcdefghijk", "bcd"))  //输出1
	fmt.Println("该子串的位置为：", Horspool("abcdefghijk", "fgh"))  //输出6
	fmt.Println("该子串的位置为：", Horspool("abcdefghijk", "hijk")) //输出8
	fmt.Println("该子串的位置为：", Horspool("abcdefghijk", "jk"))   //输出10
	fmt.Println("该子串的位置为：", Horspool("abcdefghijk", "k"))    //输出11

	fmt.Println("该子串的位置为：", Horspool("abcdefghijk", "kbz"))      //输出-1
	fmt.Println("该子串的位置为：", Horspool("abcdefghijk", "kdasdsad")) //输出-1
	fmt.Println("该子串的位置为：", Horspool("abcdefghijk", "dsadsa"))   //输出-1

}
