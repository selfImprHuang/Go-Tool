/*
 *  @Author : huangzj
 *  @Time : 2020/12/29 16:59
 *  @Description：
 */

package SkillTopology

import (
	"Go-Tool/ProgressTest/SkillTopology/model"
	"fmt"
	"testing"
)

//正确的输出应该是：
//SkillId:1,Parent:[],Children:[2 3 4]

//SkillId:2,Parent:[1],Children:[8]
//SkillId:3,Parent:[1],Children:[5 6]
//SkillId:4,Parent:[1],Children:[10 11]

//SkillId:5,Parent:[3],Children:[7]
//SkillId:6,Parent:[3],Children:[7]

//SkillId:10,Parent:[4],Children:[12]
//SkillId:11,Parent:[4],Children:[12]

//SkillId:8,Parent:[2],Children:[9]
//SkillId:9,Parent:[8],Children:[]

//SkillId:12,Parent:[10 11],Children:[]
//SkillId:7,Parent:[5 6],Children:[]
func TestGetTopologyHeadMap(t *testing.T) {

	worker := NewSkilledWorker()
	skill1 := SimulationSkills(worker)
	//拓扑父关系
	skill1 = BackSteppingHeadWithChildren(skill1)
	for _, skill := range GetTopologyHeadMap(skill1) {
		var pId []int
		for _, p := range skill.Parent {
			pId = append(pId, p.SkillId)
		}
		var cId []int
		for _, c := range skill.Children {
			cId = append(cId, c.SkillId)
		}

		fmt.Println(fmt.Sprintf("SkillId:%v,Parent:%v,Children:%v", skill.SkillId, pId, cId))
	}
}

//正确的输出应该是：
//SkillId:1,Parent:[]

//SkillId:2,Parent:[1]
//SkillId:3,Parent:[1]
//SkillId:4,Parent:[1]

//SkillId:5,Parent:[3]
//SkillId:6,Parent:[3]

//SkillId:11,Parent:[4]
//SkillId:10,Parent:[4]

//SkillId:8,Parent:[2]
//SkillId:9,Parent:[8]

//SkillId:7,Parent:[5 6]
//SkillId:12,Parent:[10 11]
func TestBackSteppingHeadWithChildren(t *testing.T) {
	worker := NewSkilledWorker()
	skill1 := SimulationSkills(worker)
	//拓扑父关系
	skill1 = BackSteppingHeadWithChildren(skill1)
	for _, skill := range GetTopologyHeadMap(skill1) {
		var pId []int
		for _, p := range skill.Parent {
			pId = append(pId, p.SkillId)
		}
		fmt.Println(fmt.Sprintf("SkillId:%v,Parent:%v", skill.SkillId, pId))
	}
}

//正确的输出为：
//
//我是技能11
//父节点Id为4,对应的Update的子节点的Id为：11,Level为：4
//子节点Id为12,对应的Update的父节点的Id为：11,Level为：4
func TestSkillUpdate(t *testing.T) {
	topology := GetCounsellorSkillByDb(NewSkilledWorker(), model.CounsellorId)
	SkillUpdate(topology, 11)

	for _, skill := range GetTopologyHeadMap(topology.HeadSkillNode) {
		for _, child := range skill.Children {
			if child.SkillId == 11 {
				fmt.Println(fmt.Sprintf("父节点Id为%v,对应的Update的子节点的Id为：%v,Level为：%v", skill.SkillId, child.SkillId, child.Level))
			}
		}
		for _, p := range skill.Parent {
			if p.SkillId == 11 {
				fmt.Println(fmt.Sprintf("子节点Id为%v,对应的Update的父节点的Id为：%v,Level为：%v", skill.SkillId, p.SkillId, p.Level))
			}
		}
	}
}
