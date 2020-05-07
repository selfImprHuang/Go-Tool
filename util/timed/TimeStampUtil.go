/*
 *  @Author : huangzj
 *  @Time : 2020/5/7 9:20
 *  @Description： 时间戳工具类
 */

package timed

import "time"

/*
 * 获取当月月初的时间戳
 */
func GetNowMonthUnix() int64 {
	now := time.Now()
	month := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	return month.Unix()
}

/*
 * 获取当天零时的时间戳
 */
func GetZeroHourUnix() int64 {
	now := time.Now()
	zeroHour := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return zeroHour.Unix()
}

/*
 * 获取当前时间戳
 */
func GetNowUnix() int64 {
	return time.Now().Unix()
}

/*
 * 获取当前小时时间戳
 */
func GetNowHourUnix() int64 {
	now := time.Now()
	hour := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, now.Location())
	return hour.Unix()
}

/*
 * 获取今年年初的时间戳
 */
func GetNowYearUnix() int64 {
	now := time.Now()
	tm := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
	return tm.Unix()
}
