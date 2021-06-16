/*
 *  @Author : huangZJ
 *  @Time : 2021/6/16 10:50
 *  @Description：
 */

package splitNumToN

import (
	"Go-Tool/util/rand"
	"fmt"
)

//将一个数拆解成三个
func SplitNumTo3(num int) {
	randNumList := make([]int, 0)
	var count int
	for i := 0; i < 2; i++ {
		randNum := getRandValue(num-count, 3-1-i)
		randNumList = append(randNumList, randNum)
		count = count + randNum
	}

	randNumList = append(randNumList, num-count)

	for _, num := range randNumList {
		fmt.Print(fmt.Sprintf("%d ", num))
	}
}

//num 被随机的数
//count count + 1 = 总随机数数量
func getRandValue(num int, count int) int {
	minValue := num - 6*count
	if minValue <= 0 {
		minValue = 1
	}
	maxValue := num - count*1
	if maxValue > 6 {
		maxValue = 6
	}

	return rand.GetRandInt(minValue, maxValue)
}
