/*
 *  @Author : huangzj
 *  @Time : 2020/8/18 14:54
 *  @Description：
 */

package search

import (
	"fmt"
	"testing"
)

func TestPartition(t *testing.T) {
	list := []int{99, 1, 2, 3, 4, 100, 200, 90, 5}
	s := partition(list, 0, 9-1)
	fmt.Println(fmt.Sprintf("当前第一个元素应该存在位置为：%d", s))
	for _, r := range list {
		fmt.Print(r, " ")
	}

	s = partition(list, 0, 9-1)
	fmt.Println(fmt.Sprintf("当前第一个元素应该存在位置为：%d", s))
	for _, r := range list {
		fmt.Print(r, " ")
	}
}

func TestLomutoQuiteSelect(t *testing.T) {
	for k := 1; k <= 9; k++ {
		s := LomutoQuiteSelect([]int{99, 1, 2, 3, 4, 100, 200, 90, 5}, k)
		fmt.Println(fmt.Sprintf("当前第%d个小的元素是：%d", k, s))

		fmt.Println()
	}

}
