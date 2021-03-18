/*
 *  @Author : huangzj
 *  @Time : 2021/3/16 17:52
 *  @Description：
 */

package NumberCount

import (
	"fmt"
	"testing"
)

func TestOnlyOneNumberShowOnce(t *testing.T) {
	fmt.Println(OnlyOneNumberShowOnce([]int{1, 2, 3, 2, 3, 4, 4}))        //1
	fmt.Println(OnlyOneNumberShowOnce([]int{1, 2, 3, 2, 3, 4, 4, 1, 5}))  //5
	fmt.Println(OnlyOneNumberShowOnce([]int{1, 2, 3, 2, 3, 4, 4, 1, 10})) //10
	fmt.Println(OnlyOneNumberShowOnce([]int{2, 2, 2, 2, 1, 3, 3, 3, 3}))  //1
}
