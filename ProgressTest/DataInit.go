/*
 *  @Author : huangzj
 *  @Time : 2020/11/13 14:55
 *  @Description：
 */

package ProgressTest

import (
	"Go-Tool/util/file"
	"Go-Tool/util/jsonEnhance"
)

type DeliverList struct {
	Deliver []*Deliver
}

type Deliver struct {
	Score    int
	AllScore int
	Id       int
}

func MakeData() DeliverList {
	var deliver DeliverList
	content, _ := file.ReadJsonFile("F:\\Go_BySelf\\src\\Go-Tool\\ProgressTest\\deliverData.json")
	jsonEnhance.UnmarshalFromString(content, &deliver)
	return deliver
}
