/*
 *  @Author : huangzj
 *  @Time : 2020/5/7 10:39
 *  @Description： 获取时间间隔工具类（这边只返回时间戳，如果需要格式化，参考TimeFormatUtil.go）
 */

package timed

import (
	"Go-Tool/util/timed/vo"
	"math"
	"time"
)

//年月日等对应的秒数
var byTime = []int64{365 * 24 * 60 * 60, 30 * 24 * 60 * 60, 7 * 24 * 60 * 60, 24 * 60 * 60, 60 * 60, 60, 1}

/*
 * 获取多少秒之前的时间戳
 */
func GetSecondsBefore(timestamp int64, seconds int) int64 {
	tm := time.Unix(timestamp, 0)
	return tm.Unix() - int64(seconds)
}

/*
 * 获取多少秒之后的时间戳
 */
func GetSecondsAfter(timestamp int64, seconds int) int64 {
	tm := time.Unix(timestamp, 0)
	return tm.Unix() + int64(seconds)
}

/*
 * 获取多少分钟之前的时间戳
 */
func GetMinutesBefore(timestamp int64, minutes int) int64 {
	tm := time.Unix(timestamp, 0)
	return tm.Add(time.Minute * time.Duration(-minutes)).Unix()
}

/*
 * 获取多少分钟之后的时间戳
 */
func GetMinutesAfter(timestamp int64, minutes int) int64 {
	tm := time.Unix(timestamp, 0)
	return tm.Add(time.Minute * time.Duration(minutes)).Unix()
}

/*
 * 获取多少小时之前的时间戳
 */
func GetHoursBefore(timestamp int64, hours int) int64 {
	tm := time.Unix(timestamp, 0)
	return tm.Add(time.Hour * time.Duration(-hours)).Unix()
}

/*
 * 获取多少小时之后的时间戳
 */
func GetHoursAfter(timestamp int64, hours int) int64 {
	tm := time.Unix(timestamp, 0)
	return tm.Add(time.Hour * time.Duration(hours)).Unix()
}

/*
 * 获取多少天之前的时间戳
 */
func GetDaysBefore(timestamp int64, days int) int64 {
	tm := time.Unix(timestamp, 0)
	return tm.AddDate(0, 0, -days).Unix()
}

/*
 * 获取多少天之后的时间戳
 */
func GetDaysAfter(timestamp int64, days int) int64 {
	tm := time.Unix(timestamp, 0)
	return tm.AddDate(0, 0, days).Unix()
}

/*
 * 获取多少个月之前的时间戳
 */
func GetMonthsBefore(timestamp int64, months int) int64 {
	tm := time.Unix(timestamp, 0)
	return tm.AddDate(0, -months, 0).Unix()
}

/*
 * 获取多少个月之后的时间戳
 */
func GetMonthsAfter(timestamp int64, months int) int64 {
	tm := time.Unix(timestamp, 0)
	return tm.AddDate(0, months, 0).Unix()
}

/*
 * 获取多少年之前的时间戳
 */
func GetYearsBefore(timestamp int64, years int) int64 {
	tm := time.Unix(timestamp, 0)
	return tm.AddDate(-years, 0, 0).Unix()
}

/*
 * 获取多少年之后的时间戳
 */
func GetYearsAfter(timestamp int64, years int) int64 {
	tm := time.Unix(timestamp, 0)
	return tm.AddDate(years, 0, 0).Unix()
}

/*
 * @param timestamp 时间戳
 * @param stamp 时间戳
 * @description 获取两个时间戳之间相差的年月日时分秒（这边每个参数分别是独立的不相关联）
 */
func GetIntervalBetweenTimes(timestamp, stamp int64) vo.IntervalObj {
	//时间大小保证前小后大
	if timestamp > stamp {
		timestamp, stamp = stamp, timestamp
	}
	timeBefore := time.Unix(timestamp, 0)
	timeAfter := time.Unix(stamp, 0)
	m := timeAfter.Sub(timeBefore)
	month, year := subMonth(timeBefore, timeAfter)
	return vo.IntervalObj{
		IntervalYear:   year,
		IntervalMonth:  month,
		IntervalDay:    m.Hours() / 24,
		IntervalHour:   m.Hours(),
		IntervalMinute: m.Minutes(),
		IntervalSecond: m.Seconds(),
	}

}

/*
 * @param timestamp 时间戳
 * @param stamp 时间戳
 * @description 获取两个时间戳之间相差的年月日时分秒（对象的所有属性组成对应的相差时间，不是分开表示的)
 */
func GetIntervalBetweenTimesDetail(timestamp, stamp int64) vo.IntervalTimeObj {
	var obj vo.IntervalTimeObj
	//时间大小保证前小后大
	if timestamp > stamp {
		timestamp, stamp = stamp, timestamp
	}
	ct := stamp - timestamp
	for i := 0; i < len(byTime); i++ {
		//如果小于的话，采用直接默认值为0.所以这边直接continue
		if ct < byTime[i] {
			continue
		}
		var temp = math.Floor(float64(ct / byTime[i]))
		ct = ct % byTime[i]
		if temp > 0 {
			makeUpObj(&obj, i, int(temp))

		}
	}

	return obj
}

func makeUpObj(obj *vo.IntervalTimeObj, i int, f int) {
	switch i {
	case 0:
		obj.Year = f
	case 1:
		obj.Month = f
	case 2:
		obj.Week = f
	case 3:
		obj.Day = f
	case 4:
		obj.Hour = f
	case 5:
		obj.Minute = f
	case 6:
		obj.Second = f
	}
}

func subMonth(timestamp, stamp time.Time) (month float64, year float64) {
	// 计算日期相差多少月
	y1 := timestamp.Year()
	y2 := stamp.Year()
	m1 := int(timestamp.Month())
	m2 := int(stamp.Month())
	d1 := timestamp.Day()
	d2 := stamp.Day()

	yearInterval := y2 - y1
	// 如果 d1的 月-日 小于 d2的 月-日 那么 yearInterval-- 这样就得到了相差的年数
	if m1 > m2 || (m1 == m2 && d1 > d2) {
		yearInterval--
	}
	// 获取月数差值
	monthInterval := m2 - m1
	if m2 < m1 {
		monthInterval += 12
	}
	if d1 > d2 {
		monthInterval--
	}
	year = float64(yearInterval) + float64(monthInterval)/12
	month = float64(yearInterval*12 + monthInterval)
	return
}
