/*
 *  @Author : huangzj
 *  @Time : 2020/5/7 9:39
 *  @Description：
 */

package timed

import (
	"fmt"
	"testing"
)

func TestTimeStampUtil(t *testing.T) {
	fmt.Print("\n\n")
	fmt.Println(GetNowHourUnix())

	fmt.Print("\n\n")
	fmt.Println(GetNowUnix())

	fmt.Print("\n\n")
	fmt.Println(GetZeroHourUnix())

	fmt.Print("\n\n")
	fmt.Println(GetNowYearUnix())

	fmt.Print("\n\n")
	fmt.Println(GetNowMonthUnix())

}
