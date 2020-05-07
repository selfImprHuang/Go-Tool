/*
 *  @Author : huangzj
 *  @Time : 2020/5/7 12:00
 *  @Description： 时间比较工具
 */

package timed

func IsBefore(timestamp, compareStamp int64) bool {
	return timestamp < compareStamp
}

func IsAfter(timestamp, compareStamp int64) bool {
	return timestamp > compareStamp
}

func IsSameTime(timestamp, compareStamp int64) bool {
	return timestamp == compareStamp
}

func IsDiffTime(timestamp, compareStamp int64) bool {
	return !IsSameTime(timestamp, compareStamp)
}

func IsBetween(timeStamp, compareBefore, compareAfter int64) bool {
	return timeStamp >= compareBefore && timeStamp <= compareAfter
}
