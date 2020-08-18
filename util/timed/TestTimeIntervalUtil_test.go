/*
 *  @Author : huangzj
 *  @Time : 2020/5/7 11:27
 *  @Description：
 */

package timed

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeIntervalUtil(t *testing.T) {
	fmt.Print("\n\n")
	fmt.Print(GetTimeDefaultFormatString(GetSecondsBefore(time.Now().Unix(), 10)))
	fmt.Print("\n\n")
	fmt.Print(GetTimeDefaultFormatString(GetSecondsAfter(time.Now().Unix(), 10)))
	fmt.Print("\n\n")
	fmt.Print(GetTimeDefaultFormatString(GetMinutesBefore(time.Now().Unix(), 10)))
	fmt.Print("\n\n")
	fmt.Print(GetTimeDefaultFormatString(GetMinutesAfter(time.Now().Unix(), 10)))
	fmt.Print("\n\n")
	fmt.Print(GetTimeDefaultFormatString(GetHoursBefore(time.Now().Unix(), 10)))
	fmt.Print("\n\n")
	fmt.Print(GetTimeDefaultFormatString(GetHoursAfter(time.Now().Unix(), 10)))
	fmt.Print("\n\n")
	fmt.Print(GetTimeDefaultFormatString(GetDaysBefore(time.Now().Unix(), 10)))
	fmt.Print("\n\n")
	fmt.Print(GetTimeDefaultFormatString(GetDaysAfter(time.Now().Unix(), 10)))
	fmt.Print("\n\n")
	fmt.Print(GetTimeDefaultFormatString(GetMonthsBefore(time.Now().Unix(), 10)))
	fmt.Print("\n\n")
	fmt.Print(GetTimeDefaultFormatString(GetMonthsAfter(time.Now().Unix(), 10)))
	fmt.Print("\n\n")
	fmt.Print(GetTimeDefaultFormatString(GetYearsBefore(time.Now().Unix(), 10)))
	fmt.Print("\n\n")
	fmt.Print(GetTimeDefaultFormatString(GetYearsAfter(time.Now().Unix(), 10)))
	fmt.Print("\n\n")
	o1 := GetIntervalBetweenTimes(time.Now().Unix(), time.Now().AddDate(3, 2, 5).Add(5*time.Hour).Add(7*time.Minute).Add(10*time.Second).Unix())
	o2 := GetIntervalBetweenTimes(time.Now().AddDate(3, 2, 5).Add(5*time.Hour).Add(7*time.Minute).Add(10*time.Second).Unix(), time.Now().Unix())
	o3 := GetIntervalBetweenTimes(time.Now().Unix(), time.Now().AddDate(3, -2, 5).Add(5*time.Hour).Add(7*time.Minute).Add(10*time.Second).Unix())

	fmt.Println(o1.IntervalYear, " y ", o1.IntervalMonth, " m ", o1.IntervalDay, " d ", o1.IntervalHour, " H ", o1.IntervalMinute, " mm ", o1.IntervalSecond, " ss ")
	fmt.Println(o2.IntervalYear, " y ", o2.IntervalMonth, " m ", o2.IntervalDay, " d ", o2.IntervalHour, " H ", o2.IntervalMinute, " mm ", o2.IntervalSecond, " ss ")
	fmt.Println(o3.IntervalYear, " y ", o3.IntervalMonth, " m ", o3.IntervalDay, " d ", o3.IntervalHour, " H ", o3.IntervalMinute, " mm ", o3.IntervalSecond, " ss ")
	o4 := GetIntervalBetweenTimesDetail(time.Now().Unix(), time.Now().AddDate(3, -2, 5).Add(5*time.Hour).Add(7*time.Minute).Add(10*time.Second).Unix())
	o5 := GetIntervalBetweenTimesDetail(time.Now().Unix(), time.Now().Unix())
	o6 := GetIntervalBetweenTimesDetail(time.Now().Unix(), time.Now().AddDate(3, 2, 5).Add(5*time.Hour).Add(7*time.Minute).Add(10*time.Second).Unix())

	fmt.Println(o4.Year, " y ", o4.Month, " m ", o4.Week, " w ", o4.Day, " d ", o4.Hour, " H ", o4.Minute, " mm ", o4.Second, " ss ")
	fmt.Println(o5.Year, " y ", o5.Month, " m ", o5.Week, " w ", o5.Day, " d ", o5.Hour, " H ", o5.Minute, " mm ", o5.Second, " ss ")
	fmt.Println(o6.Year, " y ", o6.Month, " m ", o6.Week, " w ", o6.Day, " d ", o6.Hour, " H ", o6.Minute, " mm ", o6.Second, " ss ")
}
