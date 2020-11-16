/*
 *  @Author : huangzj
 *  @Time : 2020/11/13 9:34
 *  @Description：
 */

package jsonEnhance

import (
	"github.com/json-iterator/go"
)

var (
	//config 默认配置
	config = jsoniter.ConfigFastest
)

func MarshalToString(v interface{}) string {

	s, err := config.MarshalToString(v)
	if err != nil {
		panic(err)
	}
	return s
}

//UnmarshalFromString 直接从字符串反序列化
func UnmarshalFromString(str string, v interface{}) {
	err := config.UnmarshalFromString(str, v)
	if err != nil {
		panic(err)
	}
}
