/*
 *  @Author : huangzj
 *  @Time : 2021/3/18 13:09
 *  @Description：
 */

package NumberCount

import (
	"fmt"
	"testing"
)

func TestOnlyOneNumberOtherK(t *testing.T) {
	fmt.Println(OnlyOneNumberOtherK([]int{1, 2, 3, 4, 2, 3, 4}, 2))                            //1
	fmt.Println(OnlyOneNumberOtherK([]int{1, 2, 3, 4, 2, 3, 4, 2, 3, 4}, 3))                   //1
	fmt.Println(OnlyOneNumberOtherK([]int{1, 2, 3, 4, 4, 5, 6, 2, 2, 3, 5, 5, 3, 4, 6, 6}, 3)) //1
	fmt.Println(OnlyOneNumberOtherK([]int{1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}, 10))               //1

	fmt.Println(OnlyOneNumberOtherK([]int{2, 3, 3000, 4, 2, 3, 4}, 2))                           //3000
	fmt.Println(OnlyOneNumberOtherK([]int{2, 3, 4, 2, 9999, 3, 4, 2, 3, 4}, 3))                  //9999
	fmt.Println(OnlyOneNumberOtherK([]int{2, 3, 4, 4, 5, 6, 321, 2, 2, 3, 5, 5, 3, 4, 6, 6}, 3)) //321
	fmt.Println(OnlyOneNumberOtherK([]int{2, 2, 2, 2, 2, 2, 2, 5891, 2, 2, 2}, 10))              //5891

}
