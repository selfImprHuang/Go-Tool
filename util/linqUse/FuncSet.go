/*
 *  @Author : huangzj
 *  @Time : 2020/11/12 14:43
 *  @Description：测试方法集合
 */

package linqUse

import "time"

//书籍发布时间是否早于--.--.--
var PublishTimeBeforeFunc = func(thisBook interface{}) bool {
	if thisBook.(Book).PublishTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		return true
	}
	return false
}

//书籍发布时间是否晚于--.--.--
var PublishTimeAfterFunc = func(thisBook interface{}) bool {
	if thisBook.(Book).PublishTime.After(time.Date(2018, 1, 1, 0, 0, 0, 0, time.Local)) {
		return true
	}
	return false
}

//书籍发布时间是否晚于--.--.--
var PublishTimeAfterFunc2 = func(thisBook interface{}) bool {
	if thisBook.(Book).PublishTime.After(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local)) {
		return true
	}
	return false
}

//自定义聚合操作方法
var AggregateFunc = func(first interface{}, second interface{}) interface{} {
	if first.(Book).PublishTime.Before(second.(Book).PublishTime) {
		return first
	}
	return second
}
