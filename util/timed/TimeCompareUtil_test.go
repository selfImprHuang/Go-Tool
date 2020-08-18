/*
 *  @Author : huangzj
 *  @Time : 2020/5/7 13:35
 *  @Description：
 */

package timed

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeCompareUtil(t *testing.T) {
	fmt.Print("\n\n")
	fmt.Println(IsBefore(time.Now().Unix(), time.Now().AddDate(0, 0, 1).Unix()))
	fmt.Println(IsBefore(time.Now().Unix(), time.Now().AddDate(0, 0, -1).Unix()))

	fmt.Print("\n\n")
	fmt.Println(IsAfter(time.Now().Unix(), time.Now().AddDate(0, 0, -1).Unix()))
	fmt.Println(IsAfter(time.Now().Unix(), time.Now().AddDate(0, 0, 1).Unix()))

	fmt.Print("\n\n")
	fmt.Println(IsSameTime(time.Now().Unix(), time.Now().Unix()))
	fmt.Println(IsSameTime(time.Now().Unix(), time.Now().AddDate(0, 0, -1).Unix()))

	fmt.Print("\n\n")
	fmt.Println(IsDiffTime(time.Now().Unix(), time.Now().Unix()))
	fmt.Println(IsDiffTime(time.Now().Unix(), time.Now().AddDate(0, 0, -1).Unix()))

	fmt.Print("\n\n")
	fmt.Println(IsBetween(time.Now().Unix(), time.Now().Unix(), time.Now().Unix()))
	fmt.Println(IsBetween(time.Now().Unix(), time.Now().AddDate(0, 0, -1).Unix(), time.Now().AddDate(0, 0, 1).Unix()))
	fmt.Println(IsBetween(time.Now().Unix(), time.Now().AddDate(0, 0, 1).Unix(), time.Now().AddDate(0, 0, -1).Unix()))
}
