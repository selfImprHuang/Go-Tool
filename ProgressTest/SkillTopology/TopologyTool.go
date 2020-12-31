/*
 *  @Author : huangzj
 *  @Time : 2020/12/29 16:20
 *  @Description：拓扑工具
 */

package SkillTopology

import "Go-Tool/ProgressTest/SkillTopology/model"

//设置技能提升，执行对应的方法
func SkillUpdate(topology *model.Topology, skillId int) {
	//如果技能满足提升的条件
	if CheckSkillUpdate(topology, skillId) {
		skill := topology.SkillIdMap[skillId]
		skill.Level = skill.Level + 1
		//执行技能方法
		skill.SkillFunc()
		//更新节点信息
		topology.SkillIdMap[skillId] = skill
		//刷新对应父节点的属性
		RefreshParent(skill)
		//刷新对应子节点的属性
		RefreshChildren(skill)
		return
	}

	panic("技能未达到解锁要求")
}

//当技能发生变化的时候，刷新父节点对应的该技能的属性
func RefreshParent(skill *model.SkillNode) {
	for _, parent := range skill.Parent {
		children := make([]*model.SkillNode, 0)
		for _, child := range parent.Children {
			if child.SkillId != skill.SkillId {
				children = append(children, child)
			} else {
				children = append(children, skill)
			}
		}
		parent.Children = children
	}
}

//当技能发生变化的时候,刷新子节点对应的技能的属性
func RefreshChildren(skill *model.SkillNode) {
	for _, child := range skill.Children {
		parent := make([]*model.SkillNode, 0)
		for _, p := range child.Parent {
			if child.SkillId != skill.SkillId {
				parent = append(parent, p)
			} else {
				parent = append(parent, skill)
			}
		}
		child.Parent = parent
	}
}

//校验技能升级是否合法
func CheckSkillUpdate(topology *model.Topology, skillId int) bool {
	skill := topology.SkillIdMap[skillId]
	if topology == nil || topology.HeadSkillNode == nil || skill == nil {
		return false
	}
	//父节点升级，直接返回true
	if skill.Parent == nil {
		return true
	}
	for _, parent := range skill.Parent {
		if parent.Level == 0 {
			return false
		}
	}

	return true
}

//执行所有等级大于0的技能效果
func DealWithUnLockSkill(topology *model.Topology) {
	if topology == nil || topology.HeadSkillNode == nil {
		return
	}
	node := topology.HeadSkillNode
	travelDealSkillNode(node)
}

//获取技能Id对应的技能节点
func GetTopologyHeadMap(node *model.SkillNode) map[int]*model.SkillNode {
	skillIdMap := make(map[int]*model.SkillNode, 0)
	if node == nil {
		return skillIdMap
	}
	//头节点的处理
	skillIdMap[node.SkillId] = node
	//遍历所有的子节点
	travelSkillNode(node, skillIdMap)
	return skillIdMap
}

//已经设置好子节点属性（Children）,反推父节点属性（Parent）
func BackSteppingHeadWithChildren(node *model.SkillNode) *model.SkillNode {
	if node == nil {
		return node
	}
	//头节点的处理
	travelNodeSetParent(node)
	return node
}

func travelDealSkillNode(node *model.SkillNode) {
	//执行对应的方法
	if node.SkillFunc != nil && node.Level > 0 {
		node.SkillFunc()
	}
	//递归遍历所有的子节点
	for _, child := range node.Children {
		travelDealSkillNode(child)
	}
}

func travelNodeSetParent(node *model.SkillNode) {
	for _, child := range node.Children {
		if child != nil {
			//有多个父的情况
			if child.Parent != nil || len(child.Parent) != 0 {
				child.Parent = append(child.Parent, node)
			} else {
				child.Parent = []*model.SkillNode{node}
			}
			//再深入到子节点的内部
			travelNodeSetParent(child)
		}
	}
}

func travelSkillNode(node *model.SkillNode, skillIdMap map[int]*model.SkillNode) {
	for _, child := range node.Children {
		if child != nil {
			skillIdMap[child.SkillId] = child
			travelSkillNode(child, skillIdMap)
		}
	}
}
