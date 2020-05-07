/*
 *  @Author : huangzj
 *  @Time : 2020/5/7 13:35
 *  @Description：
 */

package timed

import (
	"Go-Tool/util/timed"
	"fmt"
	"testing"
	"time"
)

func TestTimeCompareUtil(t *testing.T) {
	fmt.Print("\n\n")
	fmt.Println(timed.IsBefore(time.Now().Unix(), time.Now().AddDate(0, 0, 1).Unix()))
	fmt.Println(timed.IsBefore(time.Now().Unix(), time.Now().AddDate(0, 0, -1).Unix()))

	fmt.Print("\n\n")
	fmt.Println(timed.IsAfter(time.Now().Unix(), time.Now().AddDate(0, 0, -1).Unix()))
	fmt.Println(timed.IsAfter(time.Now().Unix(), time.Now().AddDate(0, 0, 1).Unix()))

	fmt.Print("\n\n")
	fmt.Println(timed.IsSameTime(time.Now().Unix(), time.Now().Unix()))
	fmt.Println(timed.IsSameTime(time.Now().Unix(), time.Now().AddDate(0, 0, -1).Unix()))

	fmt.Print("\n\n")
	fmt.Println(timed.IsDiffTime(time.Now().Unix(), time.Now().Unix()))
	fmt.Println(timed.IsDiffTime(time.Now().Unix(), time.Now().AddDate(0, 0, -1).Unix()))

	fmt.Print("\n\n")
	fmt.Println(timed.IsBetween(time.Now().Unix(), time.Now().Unix(), time.Now().Unix()))
	fmt.Println(timed.IsBetween(time.Now().Unix(), time.Now().AddDate(0, 0, -1).Unix(), time.Now().AddDate(0, 0, 1).Unix()))
	fmt.Println(timed.IsBetween(time.Now().Unix(), time.Now().AddDate(0, 0, 1).Unix(), time.Now().AddDate(0, 0, -1).Unix()))
}
