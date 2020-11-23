/*
 *  @Author : huangzj
 *  @Time : 2020/11/13 12:00
 *  @Description：在工程中有一个需求是对 指定数量的数组进行划分，按照每一层的规则分成对应数量，这边做一个测试.(这种方式写的代码太麻烦了，直接废弃，后面有机会再弄)
 */

package ProgressTest

import (
	"Go-Tool/util/rand"
	"fmt"
	"github.com/ahmetb/go-linq"
	"testing"
)

const (
	oneGroupNum = 4
)

//
//未完成
//需求：有一组结构体【Deliver】,现在需要按照一定的规则把这么多的结构体划分为每组X个元素多组，假设这边X = 10 (前提是Deliver可以整除10)
//     第一次按照Score排序，把Deliver划分为每组 8 * 10个元素的多个数组，不足80的按照一组算(如果分数相同，按照Id自然数排序)
//	   第二次按照Id在每个8 * 10的数组中排序，分成 4 * 10 的两个数组
//	   第三次按照AllScore进行排序，把 4 * 10的数组，分成两个2 * 10的数组
//	   第四次在2 * 10的数组中获取前10/2，然后随机在剩下的 2 * 10 - 10/2 的数组中随 10/2 个元素和第一个前五组成一组
//
//	   在第一次划分中，不满8 * 10的数组,按照数组划分往下计算
//
func TestDeliver(t *testing.T) {
	list := MakeData()    //创建数据
	deliver(list.Deliver) //进行分组操作
}

func deliver(list []*Deliver) {
	group := len(list) / (8 * oneGroupNum) //求组数
	//lastNum := len(list) % (8*oneGroupNum) //求最后一组的数量(余数)，如果为0则说明除得尽,不需要额外操作
	query := linq.From(list).OrderByDescending(func(i interface{}) interface{} {
		return i.(*Deliver).Score
	})
	queryList := make([]linq.Query, 0)
	queryList2 := make([]linq.Query, 0)
	queryList3 := make([]linq.Query, 0)

	//第一个步骤：按照Score排序并划分
	for i := 0; i < group; i++ {
		//每次数组的创建都要先去掉前面的8 * 10个元素，再取此时的8*10个元素作为当前组.因为linq没有对应取区间的方法..
		queryList = append(queryList, query.Skip(i*8*oneGroupNum).Take(8*oneGroupNum))
	}

	//第二个步骤：按照Id排序并划分
	for i, item := range queryList {
		for j := 0; j < len(queryList[i].Results())/(4*oneGroupNum); j++ {
			queryList2 = append(queryList2, item.OrderByDescending(func(result interface{}) interface{} {
				return result.(*Deliver).Id
			}).Skip(j*4*oneGroupNum).Take(4*oneGroupNum))
		}
		fmt.Println()
	}

	//第三个步骤：按照AllScore进行排序并划分
	for i, item := range queryList2 {
		queryList3 = append(queryList3, item.OrderByDescending(func(i interface{}) interface{} {
			return i.(*Deliver).AllScore
		}).Skip(i*2*oneGroupNum).Take(2*oneGroupNum))
	}

	for _, item := range queryList3 {
		deliver := make([]*Deliver, 0)
		fmt.Println(item.Results())
		item.ToSlice(&deliver)
		delivers := toIntList(deliver[0 : oneGroupNum/2])
		left := toIntList(deliver[oneGroupNum/2:])
		for i := 0; i < oneGroupNum/2; i++ {
			result := make([]int, 0)
			result, left = rand.GetAwardByWeightWithLeftAward(left)
			delivers = append(delivers, result)
		}
		fmt.Println(delivers)
		fmt.Println(left)
	}

	//第四个步骤：排序划分然后进行随机组合.

	fmt.Printf("")

}

func toIntList(delivers []*Deliver) [][]int {
	list := make([][]int, 0)
	for _, item := range delivers {
		list = append(list, []int{item.Id, item.AllScore, 1})
	}
	return list
}
