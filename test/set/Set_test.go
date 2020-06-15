/*
 *  @Author : huangzj
 *  @Time : 2020/6/15 17:17
 *  @Description：
 */

package set

import (
	set2 "Go-Tool/util/set"
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	set := new(set2.Set)
	set.Add(1)
	set.Add(2)
	set.Add("123")
	fmt.Println(set.Contains(3))
	fmt.Println(set.Contains("456"))
	fmt.Println(set.Contains("123"))
	fmt.Println(set.Contains(1))

	set.Remove(1)
	set.Remove(3)
	for _, v := range set.GetAllSet() {
		fmt.Println(v)
	}
}
