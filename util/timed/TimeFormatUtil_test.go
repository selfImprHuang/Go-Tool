/*
 *  @Author : huangzj
 *  @Time : 2020/5/7 10:04
 *  @Description：
 */

package timed

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeFormatUtil(t *testing.T) {
	fmt.Print("\n\n")
	fmt.Print(GetTimeDefaultFormatString(GetNowYearUnix()))
	fmt.Print("\n\n")
	fmt.Print(GetTimeFormatString(GetNowHourUnix(), time.ANSIC))
	fmt.Print("\n\n")
	fmt.Print(GetNowDefaultFormatString())
	fmt.Print("\n\n")
	fmt.Println(GetMonthFormatString(GetNowHourUnix()))
	fmt.Print("\n\n")
	fmt.Println(GetSimpleTimeFormatString(time.Now().Unix()))
}
