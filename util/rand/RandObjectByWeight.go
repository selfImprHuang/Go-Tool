/*
 *  @Author : huangzj
 *  @Time : 2021/2/23 20:50
 *  @Description：
 */

package rand

func GetAwardWeightByWeight(awardWeight []*AwardWeight) *AwardWeight {
	//求总权重
	weight := 0
	for _, award := range awardWeight {
		weight += award.Weight
	}
	randNum := GetRandInt(1, weight)
	sum := 0
	for _, award := range awardWeight {
		sum += award.Weight
		if sum >= randNum {
			return award
		}
	}
	return nil
}

//根据权重不放回的抽取，出参是：抽到的道具、剩余没有抽的道具
func GetByWeightWithLeftAward(awardWeight []*AwardWeight) (*AwardWeight, []*AwardWeight) {
	weight := 0
	leftAwards := make([]*AwardWeight, 0)
	for _, award := range awardWeight {
		weight += award.Weight
	}
	randNum := GetRandInt(1, weight)
	sum := 0
	for index, award := range awardWeight {
		sum += award.Weight
		if sum >= randNum {
			//如果获取到奖励则将当前获取的奖励扣除
			leftAwards = append(awardWeight[0:index], awardWeight[index+1:]...)
			return award, leftAwards
		}
	}
	return nil, leftAwards
}

// 根据权重不放回的抽取X个元素，直接返回这些元素
// 入参：奖池、抽取数量，出参：获得的所有道具，剩余的其他道具
func GetSomeAwardsFromPool(srcAwardPool []*AwardWeight, count int) ([]*AwardWeight, []*AwardWeight) {
	if count > len(srcAwardPool) {
		panic("抽取数量不对")
	}
	var finalAwardList []*AwardWeight
	var award *AwardWeight
	var leftAward = make([]*AwardWeight, len(srcAwardPool))
	copy(leftAward, srcAwardPool)

	for i := 0; i < count; i++ {
		award, leftAward = GetByWeightWithLeftAward(leftAward)
		finalAwardList = append(finalAwardList, award)
	}
	return finalAwardList, leftAward
}
