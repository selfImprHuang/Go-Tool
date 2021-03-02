/*
 *  @Author : huangzj
 *  @Time : 2020/6/29 10:48
 *  @Description：
 */

package rand

import (
	"fmt"
	"math/rand"
)

var R *rand.Rand

//获取min和max之间的一个随机数，返回值范围：[min,max]
func GetRandInt(min int, max int) int {
	if min > max {
		panic(fmt.Sprint("GetRandInt传入参数无效，min = ", min, ",max = ", max))
	}
	//Intn返回的结果是[0,n),所以这边会包含min和max
	value := R.Intn(max-min+1) + min
	return value
}

//  随机算出count个数不重复
func GetRandIntNoRepeat(count, min, max int) []int {
	if count > max-min+1 {
		panic("数据错误")
	}

	randAttr := make([]int, 0)

	// 用于记录以及选中的数
	selected := make(map[int]interface{}, 0)

	for i := 1; i <= count; i++ {
		randIndex := GetRandInt(min, max)
		if _, ok := selected[randIndex]; ok {
			i -= 1
			continue
		}
		selected[randIndex] = nil
		randAttr = append(randAttr, randIndex)
	}

	return randAttr
}

//从一个数组中获取到X个不重复下标的随机数
func GetDiffNum(data []int, count int) []int {
	if count > len(data) {
		panic("获取随机数数量有误")
	}
	diffList := make([]int, 0)
	diffListIndex := GetRandIntNoRepeat(count, 0, len(data)-1)
	for _, i := range diffListIndex {
		diffList = append(diffList, data[i])
	}
	return diffList
}
