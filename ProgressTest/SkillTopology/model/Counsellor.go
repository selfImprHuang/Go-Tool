/*
 *  @Author : huangzj
 *  @Time : 2020/12/29 15:08
 *  @Description：
 */

package model

type Counsellor struct {
	Id    int
	Name  string
	Skill []*Skill
}

type Skill struct {
	Id          int
	Level       int
	Description string
}
