/*
 *  @Author : huangzj
 *  @Time : 2020/12/2 11:19
 *  @Description：
 */

package testContainer

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestHeapTool(t *testing.T) {
	personList := new(heapTool)
	personList.Push(&Person{
		Name:  "小明",
		Age:   20,
		Money: 100000.99,
	})
	personList.Push(&Person{
		Name:  "小施",
		Age:   30,
		Money: 1002341.99,
	})

	personList.Push(&Person{
		Name:  "小康",
		Age:   10,
		Money: 200.99,
	})

	personList.Push(&Person{
		Name:  "老施",
		Age:   50,
		Money: 10343240000.99,
	})

	personList.Push(&Person{
		Name:  "老康",
		Age:   70,
		Money: 10340.99,
	})
	personList.Push(&Person{
		Name:  "老明",
		Age:   80,
		Money: 13240000.99,
	})
	personList.Push(&Person{
		Name:  "老林",
		Age:   90,
		Money: 10340000.99,
	})

	heap.Init(personList)
	for personList.Len() > 0 {
		pop := heap.Pop(personList)
		fmt.Println(fmt.Sprintf("%v,%v岁，资产：%v", pop.(*Person).Name, pop.(*Person).Age, pop.(*Person).Money))
	}

}
