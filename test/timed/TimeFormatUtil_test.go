/*
 *  @Author : huangzj
 *  @Time : 2020/5/7 10:04
 *  @Description：
 */

package timed

import (
	"Go-Tool/util/timed"
	"fmt"
	"testing"
	"time"
)

func TestTimeFormatUtil(t *testing.T) {
	fmt.Print("\n\n")
	fmt.Print(timed.GetTimeDefaultFormatString(timed.GetNowYearUnix()))
	fmt.Print("\n\n")
	fmt.Print(timed.GetTimeFormatString(timed.GetNowHourUnix(), time.ANSIC))
	fmt.Print("\n\n")
	fmt.Print(timed.GetNowDefaultFormatString())
	fmt.Print("\n\n")
	fmt.Println(timed.GetMonthFormatString(timed.GetNowHourUnix()))
	fmt.Print("\n\n")
	fmt.Println(timed.GetSimpleTimeFormatString(time.Now().Unix()))
}
