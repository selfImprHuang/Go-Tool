/*
 *  @Author : huangZJ
 *  @Time : 2021/6/15 11:20
 *  @Description：将一个数拆分成其他N个数，这N个数相加的结果等于这个数
 */

package splitNumToN

import "Go-Tool/util/rand"

type Splitter struct {
	NumCount int //拆分成随机数的数量
	Num      int //被拆分的总数
	MinValue int //被拆分的最小值
	MaxValue int //被拆分的最大值

	SplitNumList []int //存放拆解之后的结果
}

func NewSplitter(numCount, num int) *Splitter {
	//给定默认的限制的最大最小值
	splitter := &Splitter{
		NumCount:     numCount,
		Num:          num,
		MinValue:     1,
		MaxValue:     6,
		SplitNumList: make([]int, 0),
	}
	//校验数值是否满足
	splitter.checkValue()
	return splitter
}

func (splitter *Splitter) checkValue() {
	if splitter.MinValue*splitter.NumCount > splitter.Num {
		panic("被拆分的数值过小")
	}
	if splitter.MaxValue*splitter.NumCount < splitter.Num {
		panic("被拆分的数值过大")
	}
}

//将一个数拆成N个，并且保证N个数相加值等于被拆分的这个数
func (splitter *Splitter) SplitNumToN() {
	var count int
	for i := 0; i < splitter.NumCount-1; i++ {
		randNum := splitter.getRandValue(splitter.Num-count, splitter.NumCount-1-i)
		splitter.SplitNumList = append(splitter.SplitNumList, randNum)
		count = count + randNum
	}

	splitter.SplitNumList = append(splitter.SplitNumList, splitter.Num-count)
}

func (splitter *Splitter) GetRandNumList() []int {
	return splitter.SplitNumList
}

//num 被随机的数
//count count + 1 = 总随机数数量
func (splitter *Splitter) getRandValue(num int, count int) int {
	minValue := num - splitter.MaxValue*count
	if minValue <= 0 {
		minValue = splitter.MinValue
	}
	maxValue := num - count*1
	if maxValue > splitter.MaxValue {
		maxValue = splitter.MaxValue
	}

	return rand.GetRandInt(minValue, maxValue)
}
