/*
 *  @Author : huangzj
 *  @Time : 2020/6/29 17:15
 *  @Description：
 */

package rand

import (
	"Go-Tool/util/rand"
	"fmt"
	"strconv"
	"testing"
)

func TestRand(t *testing.T) {
	fmt.Printf(strconv.Itoa(rand.GetRandInt(100, 2000)))
	for _, r := range rand.GetRandIntNoRepeat(10, 100, 200) {
		fmt.Printf(strconv.Itoa(r))
	}

	for _, r := range rand.GetDiffNum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, 10) {
		fmt.Printf(strconv.Itoa(r))
	}
}
