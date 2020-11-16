/*
 *  @Author : huangzj
 *  @Time : 2020/11/16 18:00
 *  @Description：
 */

package ProgressTest

import "github.com/ahmetalpbalkan/go-linq"

type orderRule = func([]*Deliver) []*Deliver
type intOrderRule = func([]*Deliver) [][]int

var (
	firstOrderRule  orderRule
	secondOrderRule orderRule
	thirdOrderRule  intOrderRule
)

func FirstOrderRule(delivers []*Deliver) []*Deliver {
	var newDelivers []*Deliver
	linq.From(delivers).OrderByDescending(func(d interface{}) interface{} {
		return d.(*Deliver).Score
	}).ToSlice(&newDelivers)
	return newDelivers
}

func SecondOrderRule(delivers []*Deliver) []*Deliver {
	newDeliver := make([]*Deliver, 0)
	//第二次排序
	linq.From(delivers).OrderByDescending(func(d interface{}) interface{} {
		return d.(*Deliver).Score
	}).ThenBy(func(d interface{}) interface{} {
		return d.(*Deliver).Id
	}).ToSlice(&newDeliver)

	return newDeliver
}

func ThirdOrderRule(delivers []*Deliver) [][]int {
	Ids := make([][]int, 0)
	//第三次排序
	linq.From(delivers).OrderByDescending(func(d interface{}) interface{} {
		return d.(*Deliver).Score
	}).Select(func(i interface{}) interface{} {
		return []int{i.(*Deliver).Score, 1}
	}).ToSlice(&Ids)
	return Ids
}
