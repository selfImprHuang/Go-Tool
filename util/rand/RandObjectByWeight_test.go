/*
 *  @Author : huangzj
 *  @Time : 2021/2/23 22:08
 *  @Description：
 */

package rand

import (
	"fmt"
	"testing"
)

func TestRandByAwardWeight(t *testing.T) {
	awardList := make([]*AwardWeight, 0)
	awardList = append(awardList, &AwardWeight{
		Award:  []int{1, 2},
		Weight: 10,
	})
	awardList = append(awardList, &AwardWeight{
		Award:  []int{2, 2},
		Weight: 20,
	})
	awardList = append(awardList, &AwardWeight{
		Award:  []int{3, 2},
		Weight: 30,
	})

	award := GetAwardWeightByWeight(awardList)
	fmt.Println(award.Award)

	award1, list3 := GetByWeightWithLeftAward(awardList)
	fmt.Println(award1.Award)
	for _, award := range list3 {
		fmt.Println(award.Award)
	}

	list4, left := GetSomeAwardsFromPool(awardList, 2)
	for _, award := range list4 {
		fmt.Println(award.Award)
	}
	for _, award := range left {
		fmt.Println(award.Award)
	}

}
