/*
 *  @Author : huangzj
 *  @Time : 2021/3/16 16:59
 *  @Description：
 */

package NumberMoreThan

import (
	"fmt"
	"testing"
)

func TestNumberMoreThanK(t *testing.T) {
	fmt.Println(NumberMoreThanK([]int{1, 2, 3, 1, 2, 3, 1, 2, 3, 1}, 3)) //1
	fmt.Println(NumberMoreThanK([]int{1, 2, 3, 1, 2, 3, 1, 2, 3, 1}, 4)) //1,2,3
	fmt.Println(NumberMoreThanK([]int{1, 2, 3, 1, 2, 3, 1, 2, 3}, 2))    //
	fmt.Println(NumberMoreThanK([]int{1, 2, 3}, 2))                      //
}
