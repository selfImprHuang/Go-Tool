/*
 *  @Author : huangzj
 *  @Time : 2020/12/29 14:22
 *  @Description：模拟的一个大臣的技能树拓扑关系
 */

package SkillTopology

import "Go-Tool/ProgressTest/SkillTopology/model"

type NodeRelation struct {
	ParentId   int
	ChildrenId []int
}

var (
	//拓扑父子关系，模拟读配置
	topologyConfig = map[int][]int{
		1:  {2, 3, 4},
		2:  {8},
		3:  {5, 6},
		4:  {10, 11},
		5:  {7},
		6:  {7},
		8:  {9},
		10: {12},
		11: {12},
	}

	headNodeId = 1
)

//模拟其中一个大臣获取技能拓扑
func SimulationParamGetTopology(worker *SkilledWorker) *model.Topology {
	//模拟获取技能链
	skill1 := SimulationSkills(worker)
	//拓扑父关系
	skill1 = BackSteppingHeadWithChildren(skill1)

	topology := &model.Topology{
		HeadSkillNode: skill1,
		CounsellorId:  model.CounsellorId,
		SkillIdMap:    GetTopologyHeadMap(skill1), //获取技能Id和节点的对应关系，方便查找
	}

	return topology
}

func SimulationSkills(worker *SkilledWorker) *model.SkillNode {
	//根节点技能
	skill1 := &model.SkillNode{
		SkillId:   1,
		SkillFunc: worker.SkillId1,
		Parent:    nil, //第一个技能的父节点应该是空的
	}
	skill2 := model.SkillNode{
		SkillId:   2,
		SkillFunc: worker.SkillId2,
	}
	skill3 := model.SkillNode{
		SkillId:   3,
		SkillFunc: worker.SkillId3,
	}
	skill4 := model.SkillNode{
		SkillId:   4,
		SkillFunc: worker.SkillId4,
	}
	skill5 := model.SkillNode{
		SkillId:   5,
		SkillFunc: worker.SkillId5,
	}
	skill6 := model.SkillNode{
		SkillId:   6,
		SkillFunc: worker.SkillId6,
	}
	skill7 := model.SkillNode{
		SkillId:   7,
		SkillFunc: worker.SkillId7,
	}
	skill8 := model.SkillNode{
		SkillId:   8,
		SkillFunc: worker.SkillId8,
	}
	skill9 := model.SkillNode{
		SkillId:   9,
		SkillFunc: worker.SkillId9,
	}
	skill10 := model.SkillNode{
		SkillId:   10,
		SkillFunc: worker.SkillId10,
	}
	skill11 := model.SkillNode{
		SkillId:   11,
		SkillFunc: worker.SkillId11,
	}
	skill12 := model.SkillNode{
		SkillId:   12,
		SkillFunc: worker.SkillId12,
	}
	skillMap := map[int]*model.SkillNode{
		1:  skill1,
		2:  &skill2,
		3:  &skill3,
		4:  &skill4,
		5:  &skill5,
		6:  &skill6,
		7:  &skill7,
		8:  &skill8,
		9:  &skill9,
		10: &skill10,
		11: &skill11,
		12: &skill12,
	}

	//拓扑父子关系
	return MakeTopologyByConfig(skillMap, topologyConfig, headNodeId)
}

func MakeTopologyByConfig(nodes map[int]*model.SkillNode, config map[int][]int, headNodeId int) *model.SkillNode {
	headNode := nodes[headNodeId]
	if headNode == nil {
		return nil
	}
	for nodeId, childrenId := range config {
		childrenNodes := make([]*model.SkillNode, 0)
		for _, childId := range childrenId {
			if node := nodes[childId]; node != nil {
				childrenNodes = append(childrenNodes, node)
			}
		}
		nodes[nodeId].Children = childrenNodes
	}

	return headNode
}

//组装技能所对应的Map节点
func MakeSimulationSkillsWithLevel(topology *model.Topology, skillList []*model.Skill) *model.Topology {
	if topology == nil || topology.HeadSkillNode == nil || len(skillList) == 0 {
		return topology
	}
	skillIdMap := topology.SkillIdMap
	for _, skill := range skillList {
		if v, ok := skillIdMap[skill.Id]; ok {
			v.Level = skill.Level
			skillIdMap[skill.Id] = v
		}
	}

	return topology
}
