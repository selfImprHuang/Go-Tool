/*
 *  @Author : huangzj
 *  @Time : 2020/5/7 9:39
 *  @Description：
 */

package timed

import (
	"Go-Tool/util/timed"
	"fmt"
	"testing"
)

func TestTimeStampUtil(t *testing.T) {
	fmt.Print("\n\n")
	fmt.Println(timed.GetNowHourUnix())

	fmt.Print("\n\n")
	fmt.Println(timed.GetNowUnix())

	fmt.Print("\n\n")
	fmt.Println(timed.GetZeroHourUnix())

	fmt.Print("\n\n")
	fmt.Println(timed.GetNowYearUnix())

	fmt.Print("\n\n")
	fmt.Println(timed.GetNowMonthUnix())

}
