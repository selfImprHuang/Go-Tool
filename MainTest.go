/*
 *  @Author : huangzj
 *  @Time : 2020/4/24 14:51
 *  @Description：
 */

package main

import (
	"fmt"
	"github.com/ahmetb/go-linq"
)

func main() {
	var a = []int{1, 5, 3, 4}
	linq.From(a).Sort(func(i, j interface{}) bool {
		return i.(int) < j.(int)
	}).ToSlice(&a)

	fmt.Println(a)
}
