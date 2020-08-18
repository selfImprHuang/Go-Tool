/*
 *  @Author : huangzj
 *  @Time : 2020/6/29 17:15
 *  @Description：
 */

package rand

import (
	"fmt"
	"strconv"
	"testing"
)

func TestRand(t *testing.T) {
	fmt.Printf(strconv.Itoa(GetRandInt(100, 2000)))
	for _, r := range GetRandIntNoRepeat(10, 100, 200) {
		fmt.Printf(strconv.Itoa(r))
	}

	for _, r := range GetDiffNum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, 10) {
		fmt.Printf(strconv.Itoa(r))
	}
}
