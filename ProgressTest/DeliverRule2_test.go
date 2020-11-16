/*
 *  @Author : huangzj
 *  @Time : 2020/11/16 14:44
 *  @Description：
 */

package ProgressTest

import (
	"Go-Tool/util/rand"
	"fmt"
	"math"
	"testing"
)

const (
	MinGroupNum = 10
	MaxGroupNum = 80
)

//需求(直接举例)：
//总共有100个数据，两个变量，大区间(MaxGroupNum) 为80， 小区间(MinGroupNum)为 10 || 100 和 80 对应的参数   都必须是10的倍数
//第一步：先100个数据排序，然后按照80进行划分。
//第二步：100划分剩余的数据组成一组
//第三步：80个数组再一次排序,按照 10 * 4 进行划分，可以分成两组
//第四步：40每组的数据，要分成最后的10 * 4 的4组，按照每次取前5再随机取5个组成一组的方式组成4组
//第五步：不足4组的，按照取前5和随机5个的方式进行分组
func TestDeliver2(t *testing.T) {

	firstOrderRule = FirstOrderRule
	secondOrderRule = SecondOrderRule
	thirdOrderRule = ThirdOrderRule

	list := MakeData()                   //创建数据
	resultList := deliver2(list.Deliver) //进行分组操作
	fmt.Printf(fmt.Sprintf("%v", resultList))
}

func deliver2(delivers []*Deliver) [][]int {

	resultList := make([][]int, 0)
	newDelivers := make([]*Deliver, 0)
	//获得第一次按照大区间分组后剩下的元素数量、第一次分组所有组、分到的大组数量
	leaveNum, partitions, partition := FirstOrder(delivers, newDelivers)

	resultList = withCompletePartition(partitions, resultList)                    //对完整的分组进行处理
	resultList = withLeavePartition(resultList, newDelivers, leaveNum, partition) //对剩余的分组进行处理
	return resultList
}

func withLeavePartition(resultList [][]int, newDelivers []*Deliver, leaveNum int, partition int) [][]int {
	if leaveNum != 0 {
		leaveList := newDelivers[partition*MaxGroupNum:]
		resultList = secondOrder(leaveList, resultList, leaveNum)
	}
	return resultList
}

func withCompletePartition(partitions [][]*Deliver, resultList [][]int) [][]int {
	for _, partition := range partitions {
		resultList = secondOrder(partition, resultList, len(partition))
	}
	return resultList
}

func FirstOrder(delivers []*Deliver, newDelivers []*Deliver) (int, [][]*Deliver, int) {
	newDelivers = firstOrderRule(delivers) //第一次排序
	partition := len(newDelivers) / MaxGroupNum
	leaveNum := len(newDelivers) % MaxGroupNum
	partitions := make([][]*Deliver, 0)

	for i := 0; i < partition; i++ {
		partitions = append(partitions, newDelivers[i*MaxGroupNum:(i+1)*MaxGroupNum])
	}
	return leaveNum, partitions, partition
}

func secondOrder(delivers []*Deliver, resultList [][]int, listNum int) [][]int {
	newDelivers := secondOrderRule(delivers) //第二次排序

	if listNum > 4*MinGroupNum {
		fourInOne := listNum / (4 * MinGroupNum)
		leave := listNum % (4 * MinGroupNum)

		for i := 0; i < fourInOne; i++ {
			resultList = append(resultList, calculateFourInOne(newDelivers[i*4*MinGroupNum:(i+1)*4*MinGroupNum])...) //四倍在一组的计算
		}

		if leave != 0 {
			resultList = append(resultList, calculateByLeft(toList(newDelivers[fourInOne*4*MinGroupNum:]))...)
		}
	} else {
		resultList = append(resultList, calculateByLeft(toList(newDelivers))...)
	}
	return resultList
}

func toList(deliver []*Deliver) [][]int {
	result := make([][]int, 0)
	for _, item := range deliver {
		result = append(result, []int{item.Score, 1})
	}
	return result
}

func calculateFourInOne(delivers []*Deliver) [][]int {
	Ids := thirdOrderRule(delivers) //第三次排序
	result := calculateByLeft(Ids)
	return result
}

func calculateByLeft(leftList [][]int) [][]int {
	result := make([][]int, 0)
	for len(leftList) != 0 {
		newList := make([]int, 0)
		thisGroup := getFirstSome(leftList[0 : MinGroupNum/2]) //先拿到前面的几个

		leftList = leftList[MinGroupNum/2:]
		//然后在随机随取后面的几个
		for i := 0; i < int(math.Ceil(MinGroupNum/2)); i++ {
			newList, leftList = getFrontAndEnd(leftList) //从前部和后部进行组合数组
			thisGroup = append(thisGroup, newList[0])
		}

		result = append(result, thisGroup)
	}
	return result
}

func getFirstSome(list [][]int) []int {
	result := make([]int, 0)
	for _, item := range list {
		result = append(result, item[0])
	}
	return result
}

func getFrontAndEnd(leftList [][]int) ([]int, [][]int) {
	return rand.GetAwardByWeightWithLeftAward(leftList)
}
