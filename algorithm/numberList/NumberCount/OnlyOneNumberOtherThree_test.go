/*
 *  @Author : huangzj
 *  @Time : 2021/3/16 21:55
 *  @Description：
 */

package NumberCount

import (
	"fmt"
	"testing"
)

func TestOnlyOneNumberOtherThree1(t *testing.T) {
	fmt.Println(OnlyOneNumberOtherThree1([]int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4}))                                     //4
	fmt.Println(OnlyOneNumberOtherThree1([]int{1, 2, 3, 1, 2, 3, 1, 2, 3, 4}))                                     //4
	fmt.Println(OnlyOneNumberOtherThree1([]int{1, 15, 23, 67, 15, 67, 67, 23, 23, 14, 15, 1, 1}))                  //14
	fmt.Println(OnlyOneNumberOtherThree1([]int{100, 200, 300, 4, 6, 100, 100, 6, 6, 4, 200, 300, 200, 300, 5, 4})) //5
}

func TestOnlyOneNumberOtherThree2(t *testing.T) {
	fmt.Println(OnlyOneNumberOtherThree2([]int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4}))                                     //4
	fmt.Println(OnlyOneNumberOtherThree2([]int{1, 2, 3, 1, 2, 3, 1, 2, 3, 4}))                                     //4
	fmt.Println(OnlyOneNumberOtherThree2([]int{1, 15, 23, 67, 15, 67, 67, 23, 23, 14, 15, 1, 1}))                  //14
	fmt.Println(OnlyOneNumberOtherThree2([]int{100, 200, 300, 4, 6, 100, 100, 6, 6, 4, 200, 300, 200, 300, 5, 4})) //5
}
