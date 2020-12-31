/*
 *  @Author : huangzj
 *  @Time : 2020/12/29 15:46
 *  @Description：
 */

package SkillTopology

import "Go-Tool/ProgressTest/SkillTopology/model"

var SkillTopologyMap map[int]*model.Topology

func GetSkillTopologyMapById(counsellorId int) *model.Topology {
	return SkillTopologyMap[counsellorId]
}

func init() {
	SkillTopologyMap = make(map[int]*model.Topology, 0)

	SkillTopologyMap[model.CounsellorId] = SimulationParamGetTopology(NewSkilledWorker())
	//....下面是不同大臣的初始化技能树链
}
