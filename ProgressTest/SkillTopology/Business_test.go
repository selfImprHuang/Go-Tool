/*
 *  @Author : huangzj
 *  @Time : 2020/12/29 17:39
 *  @Description：
 */

package SkillTopology

import (
	"Go-Tool/ProgressTest/SkillTopology/model"
	"fmt"
	"testing"
)

func TestDrawSkill(t *testing.T) {
	topology := GetCounsellorSkillByDb(NewSkilledWorker(), model.CounsellorId)
	DealWithUnLockSkill(topology)
}

func TestUpdateSkill(t *testing.T) {
	//测试一个没有达到升级条件的节点7
	topology := GetCounsellorSkillByDb(NewSkilledWorker(), model.CounsellorId)
	ok := CheckSkillUpdate(topology, 7)

	fmt.Println(fmt.Sprintf("技能7是否满足升级条件：%v", ok))

	//测试一个可以升级的节点：11
	ok = CheckSkillUpdate(topology, 11)

	fmt.Println(fmt.Sprintf("技能11是否满足升级条件：%v", ok))

}
