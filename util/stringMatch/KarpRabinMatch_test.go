/*
 *  @Author : huangzj
 *  @Time : 2021/3/5 13:25
 *  @Description：
 */

package stringMatch

import (
	"fmt"
	"testing"
)

func TestKarpRabinMatch(t *testing.T) {
	fmt.Println("该子串的位置为：", KarpRabinMatch("abcdefghijk", "abcd")) //输出0
	fmt.Println("该子串的位置为：", KarpRabinMatch("abcdefghijk", "bcd"))  //输出1
	fmt.Println("该子串的位置为：", KarpRabinMatch("abcdefghijk", "fgh"))  //输出5
	fmt.Println("该子串的位置为：", KarpRabinMatch("abcdefghijk", "hijk")) //输出7
	fmt.Println("该子串的位置为：", KarpRabinMatch("abcdefghijk", "jk"))   //输出9
	fmt.Println("该子串的位置为：", KarpRabinMatch("abcdefghijk", "k"))    //输出10

	fmt.Println("该子串的位置为：", KarpRabinMatch("abcdefghijk", "kbz"))      //输出-1
	fmt.Println("该子串的位置为：", KarpRabinMatch("abcdefghijk", "kdasdsad")) //输出-1
	fmt.Println("该子串的位置为：", KarpRabinMatch("abcdefghijk", "dsadsa"))   //输出-1
}
