/*
 *  @Author : huangzj
 *  @Time : 2020/12/29 15:52
 *  @Description：
 */

package model

type Topology struct {
	HeadSkillNode *SkillNode         //拓扑头结点
	CounsellorId  int                //拓扑对应的大臣Id
	SkillIdMap    map[int]*SkillNode //技能ID对应的技能节点，方便查找
}
