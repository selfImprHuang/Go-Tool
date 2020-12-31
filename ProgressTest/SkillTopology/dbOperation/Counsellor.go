/*
 *  @Author : huangzj
 *  @Time : 2020/12/29 14:49
 *  @Description：
 */

package dbOperation

import (
	"Go-Tool/ProgressTest/SkillTopology/model"
)

type CounsellorSrv struct {
}

func NewCounsellorSrv() *CounsellorSrv {
	return &CounsellorSrv{}
}

//通过模拟的Id来获取对应的技能
func (*CounsellorSrv) GetCounsellorSkill(counsellorId int) *model.Counsellor {
	skillPoint := NewSkillPointSrv()
	return &model.Counsellor{
		Id:    model.CounsellorId,
		Name:  "测试的大臣",
		Skill: skillPoint.GetAllSkillByCounsellorId(model.CounsellorId),
	}
}
