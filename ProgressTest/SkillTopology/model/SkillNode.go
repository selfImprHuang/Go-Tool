/*
 *  @Author : huangzj
 *  @Time : 2020/12/29 15:00
 *  @Description：技能树节点
 */

package model

type SkillNode struct {
	SkillId  int          //技能Id，对应数据库技能Id字段 和 Skiller技能func
	Parent   []*SkillNode //对应父关系节点，根节点是nil
	Children []*SkillNode //对应子关系节点，没有则为nil
	Level    int          //技能等级，通过技能等级判断该技能是否解锁，省去UnLock字段

	SkillFunc func() //技能节点对应的技能提升效果
}
