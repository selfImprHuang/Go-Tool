/*
 *  @Author : huangzj
 *  @Time : 2020/7/13 11:33
 *  @Description：
 */

package util

import (
	"fmt"
	"testing"
)

//不同方式获取哈希值，结果相差挺大，从输出的结果可以非常直观的看出
func TestHashMode(t *testing.T) {
	fmt.Println(UseBigIntMod("3213444658987654", 9999999))
	//因为big.Int只能接受大数，所以如果这边传入字符串没办法转换成功，会发生报错
	//fmt.Println(UseBigIntMod("asdsadasdasdas", 20))
	fmt.Println(Mod("3213444658987654", 9999999))
	fmt.Println(Crc32Mode("3213444658987654", 9999999))

	fmt.Println()

	fmt.Println(Mod("asdsadasdasdas", 9999999))
	fmt.Println(Crc32Mode("asdsadasdasdas", 9999999))
}
