/*
 *  @Author : huangzj
 *  @Time : 2021/3/16 16:23
 *  @Description：
 */

package NumberMoreThan

import (
	"fmt"
	"testing"
)

func TestNumberMoreThanHalf(t *testing.T) {
	fmt.Println(NumberMoreThanHalf([]int{1, 2, 3}))                                     //-1
	fmt.Println(NumberMoreThanHalf([]int{1, 2, 3, 1, 1}))                               //1
	fmt.Println(NumberMoreThanHalf([]int{1, 2, 3, 1}))                                  //-1
	fmt.Println(NumberMoreThanHalf([]int{1, 2, 3, 2, 3, 2, 3, 3}))                      //-1
	fmt.Println(NumberMoreThanHalf([]int{1, 2, 3, 4, 5, 6, 1, 1, 1, 1, 1, 1, 1, 1, 1})) //1
}
