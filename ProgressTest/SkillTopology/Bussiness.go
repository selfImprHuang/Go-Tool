/*
 *  @Author : huangzj
 *  @Time : 2020/12/29 17:25
 *  @Description：业务场景
 */

package SkillTopology

import "Go-Tool/ProgressTest/SkillTopology/model"

//获取根据大臣的数据获取一个大臣对应的技能树拓扑的等级对应关系
func GetCounsellorSkillByDb(worker *SkilledWorker, counsellorId int) *model.Topology {
	counsellor := worker.CounsellorSrv.GetCounsellorSkill(counsellorId)
	//根据大臣Id获取大臣技能树拓扑
	topology := SkillTopologyMap[counsellorId]
	//根据大臣当前等级信息,组合技能树技能等级
	topology = MakeSimulationSkillsWithLevel(topology, counsellor.Skill)
	return topology
}
