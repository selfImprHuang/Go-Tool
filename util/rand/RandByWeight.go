/*
 *  @Author : huangzj
 *  @Time : 2020/6/29 10:23
 *  @Description：这边的随机获取没有加锁，对于一些需要加锁的场景可以在外面再包装一层
 */

package rand

import (
	"Go-Tool/util/array"
	"math/rand"
	"time"
)

//传入一组包含权重的数组，计算权重，返回对应权重数组的下标(即对应权重获取的结果)
func GetRandValueByWeight(weightList []int) int {
	//计算总的权重
	totalWeight := 0
	for _, value := range weightList {
		totalWeight += value
	}
	nowRand := R.Intn(totalWeight)
	rate := 0
	for i := 0; i < len(weightList); i++ {
		rate += weightList[i]
		if nowRand < rate {
			return i
		}
	}
	return 0
}

//通过权重获取奖励结果,这边的awardList的结构应该是 [[道具Id，道具数量，权重],....]
//返回结果的第一个元素是不包含权重的数组，第二个元素是包含权重的数组
func GetAwardByWeight(awardList [][]int) ([]int, []int) {
	//这边应该加一个数组的检查.
	weightListCheck(awardList)

	//最后一位一定是权重.
	weight := 0
	for _, v := range awardList {
		weight += v[len(v)-1]
	}
	randNum := GetRandInt(1, weight)
	sum := 0
	for _, v := range awardList {
		sum += v[len(v)-1]
		if sum >= randNum {
			result := make([]int, 0)
			weightResult := make([]int, 0)
			return append(result, v[:len(v)-1]...), append(weightResult, v[:]...)
		}
	}
	return nil, nil
}

//根据权重从列表获取元素
//入参分别是抽取的数字列表、权重列表
func GetByWeight(targetList []int, weightList []int) int {
	//一开始就检查权重和对应的元素，长度是否一致
	if len(targetList) != len(weightList) {
		panic("权重和目标列表元素数量不匹配")
	}
	weight := 0
	for _, v := range weightList {
		weight += v
	}
	randNum := GetRandInt(1, weight)
	sum := 0
	for index, v := range weightList {
		sum += v
		if sum >= randNum {
			return targetList[index]
		}
	}
	return 0
}

//根据权重不放回的抽取，出参是：抽到的道具、剩余没有抽的道具
func GetAwardByWeightWithLeftAward(srcAwardList [][]int) ([]int, [][]int) {
	//这边应该加一个数组的检查.
	weightListCheck(srcAwardList)

	awardList := array.DeepCopyIntSlice2(srcAwardList)

	weight := 0
	award := make([]int, 0)
	leftAwards := make([][]int, 0)
	for _, v := range awardList {
		weight += v[len(v)-1]
	}
	randNum := GetRandInt(1, weight)
	sum := 0
	for index, v := range awardList {
		sum += v[len(v)-1]
		if sum >= randNum {
			//如果获取到奖励则将当前获取的奖励扣除
			leftAwards = append(awardList[0:index], awardList[index+1:]...)
			award = v[:len(v)-1]
			return award, leftAwards
		}
	}
	return award, leftAwards
}

//根据权重不放回的抽取X个元素，直接返回这些元素
// 入参：奖池、抽取数量，出参：获得的所有道具
func GetCountAwardsFromPool(srcAwardPool [][]int, count int) [][]int {
	awardPool := array.DeepCopyIntSlice2(srcAwardPool)

	if count > len(awardPool) {
		panic("抽取数量不对")
	}
	var finalAwardList [][]int
	var award []int
	var leftAward = make([][]int, len(awardPool))
	copy(leftAward, awardPool)

	for i := 0; i < count; i++ {
		award, leftAward = GetAwardByWeightWithLeftAward(leftAward)
		finalAwardList = append(finalAwardList, award)
	}
	return finalAwardList
}

//通过百分比的方式来获取对应的数组，也就是说权重的总值可能超过传入的百分比数，此时取的只是0-百分比数的范围内
//所以这个方法可以按照概率取一个范围内的元素
func GetAwardByPercentage(srcAwardList [][]int, fullPercent int) []int {
	weightListCheck(srcAwardList)
	if fullPercent <= 0 {
		panic("抽取概率不能小于等于0")
	}

	awardList := array.DeepCopyIntSlice2(srcAwardList)

	randNum := GetRandInt(1, fullPercent)
	sum := 0
	for _, v := range awardList {
		sum += v[len(v)-1]
		if sum >= randNum {
			result := make([]int, 0)
			return append(result, v[:len(v)-1]...)
		}
	}
	return []int{}
}

func weightListCheck(awardList [][]int) {
	for _, r := range awardList {
		if len(r) == 0 {
			panic("一维数组长度不能为0，必须存在权重")
		}
	}
}

func InitRand() {
	R = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func init() {
	InitRand()
}
