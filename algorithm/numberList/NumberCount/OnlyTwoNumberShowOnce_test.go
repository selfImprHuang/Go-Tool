/*
 *  @Author : huangzj
 *  @Time : 2021/3/16 18:02
 *  @Description：
 */

package NumberCount

import (
	"fmt"
	"testing"
)

func TestOnlyTwoNumberShowOnce(t *testing.T) {
	fmt.Println(OnlyTwoNumberShowOnce([]int{2, 3, 4, 5, 6, 2, 3, 4, 5, 6, 1, 7}))                           //1,7
	fmt.Println(OnlyTwoNumberShowOnce([]int{2, 3, 4, 5, 6, 2, 3, 4, 5, 6, 1, 7, 8, 9, 0, 19, 19, 0, 8, 9})) //1,7
	fmt.Println(OnlyTwoNumberShowOnce([]int{2, 3, 4, 5, 1, 7, 5, 4, 3, 2}))                                 //1,7
	fmt.Println(OnlyTwoNumberShowOnce([]int{2, 3, 4, 5, 6, 2, 3, 4, 5, 6, 1, 7, 1, 100}))                   //1,100
}
