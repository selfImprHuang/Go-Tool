/*
 *  @Author : huangzj
 *  @Time : 2020/12/29 14:48
 *  @Description：
 */

package dbOperation

import (
	"Go-Tool/ProgressTest/SkillTopology/model"
	"Go-Tool/util/file"
	"encoding/json"
)

type SkillPointSrv struct {
	//.... 需要的所有属性
}

func NewSkillPointSrv() *SkillPointSrv {
	return &SkillPointSrv{}
}

func (*SkillPointSrv) RacePoint() {

}

type tmp struct {
	Skills []*model.Skill
}

//模拟数据，未激活的技能（Level = 0 ）技能设置为：6、7、8、10、12，其他为激活的技能,没有等级的技能在数据库不体现
func (*SkillPointSrv) GetAllSkillByCounsellorId(counsellorId int) []*model.Skill {
	var tmp tmp
	content, _ := file.ReadJsonFile("F://Go_BySelf//src//Go-Tool//ProgressTest//SkillTopology//skill.json")
	_ = json.Unmarshal([]byte(content), &tmp)

	return tmp.Skills
}
