/*
 *  @Author : huangZJ
 *  @Time : 2021/6/16 10:52
 *  @Description：
 */

package splitNumToN

import (
	"fmt"
	"testing"
)

func TestSplitNumTo3(t *testing.T) {
	for i := 4; i <= 17; i++ {
		fmt.Println()
		fmt.Println(fmt.Sprintf("当前随机数为%d", i))
		for j := 0; j < 100; j++ {
			fmt.Println()
			SplitNumTo3(i)
		}
		fmt.Println()
	}
}
