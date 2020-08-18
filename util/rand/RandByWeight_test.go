/*
 *  @Author : huangzj
 *  @Time : 2020/6/29 17:16
 *  @Description：
 */

package rand

import (
	"fmt"
	"strconv"
	"testing"
)

func TestRandByWeight(t *testing.T) {
	fmt.Printf(strconv.Itoa(GetRandValueByWeight([]int{10, 20, 30, 50})))

	list, list1 := GetAwardByWeight([][]int{{2, 2, 10}, {3, 3, 4}})
	for _, r := range list {
		fmt.Println(strconv.Itoa(r))
	}
	for _, r := range list1 {
		fmt.Println(strconv.Itoa(r))
	}

	fmt.Println(GetByWeight([]int{10, 11, 12, 13}, []int{20, 30, 40, 50}))

	list2, list3 := GetAwardByWeightWithLeftAward([][]int{{2, 3, 30}, {3, 3, 40}})
	for _, r := range list2 {
		fmt.Println(strconv.Itoa(r))
	}
	for _, r := range list3 {
		for _, r1 := range r {
			fmt.Println(strconv.Itoa(r1))
		}
	}

	list4 := GetCountAwardsFromPool([][]int{{1, 2, 30}, {2, 2, 40}, {3, 2, 50}}, 2)
	for _, r := range list4 {
		for _, r1 := range r {
			fmt.Println(strconv.Itoa(r1))
		}
	}

	list5 := GetAwardByPercentage([][]int{{20, 3, 99}, {90, 3, 10}, {12, 3, 22}}, 100)
	for _, r := range list5 {
		fmt.Println(strconv.Itoa(r))
	}

}
