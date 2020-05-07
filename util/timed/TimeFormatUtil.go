/*
 *  @Author : huangzj
 *  @Time : 2020/5/7 9:51
 *  @Description： 时间戳格式化工具
 */

package timed

import (
	"Go-Tool/util/timed/enum"
	"time"
)

const defaultFormat = "2006-01-02 15:04:05"
const simpleFormat = "15:04:05"

/*
 * 获取对应时间的传入格式对应的字符串时间格式
 */
func GetTimeFormatString(timestamp int64, format string) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format(format)
}

/*
 * 获取对应时间的默认格式的时间字符串.
 */
func GetTimeDefaultFormatString(timestamp int64) string {
	return GetTimeFormatString(timestamp, defaultFormat)
}

/*
 * 获取当前时间的默认格式的字符串
 */
func GetNowDefaultFormatString() string {
	return GetTimeFormatString(time.Now().Unix(), defaultFormat)
}

/*
 * 获取时间月份对应的月份信息
 */
func GetMonthFormatString(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return enum.MonthMap[tm.Month().String()]
}

/*
 * 获取时间的简单格式，比如： 小时：分钟：秒钟
 */
func GetSimpleTimeFormatString(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return GetTimeFormatString(tm.Unix(), simpleFormat)
}
