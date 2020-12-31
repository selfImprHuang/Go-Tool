/*
 *  @Author : huangzj
 *  @Time : 2020/12/29 14:21
 *  @Description：
 */

package SkillTopology

import (
	"Go-Tool/ProgressTest/SkillTopology/dbOperation"
	"fmt"
)

type Session int

type SkilledWorker struct {
	TX            Session                    //在实际的程序中,这个地方应该是事务，这边示例只用int暂且表示^_^
	CounsellorSrv *dbOperation.CounsellorSrv //在实际的程序中,这个地方是大臣的Srv,以此从数据库中获取大臣信息，这边示例只用int暂且表示
	SkillPointSrv *dbOperation.SkillPointSrv //在实际的程序中,这个地方应该是技能点的Srv,以此获得数据库数据，这边示例只用int暂且表示

	//......      其他相关的属性srv
}

func NewSkilledWorker() *SkilledWorker {
	return &SkilledWorker{
		TX:            0,
		CounsellorSrv: dbOperation.NewCounsellorSrv(),
		SkillPointSrv: dbOperation.NewSkillPointSrv(),
	}
}

//这边的技能模拟对应实际技能,示例只做输出
func (*SkilledWorker) SkillId1() {
	fmt.Println("我是技能1")
}

func (*SkilledWorker) SkillId2() {
	fmt.Println("我是技能2")
}

func (*SkilledWorker) SkillId3() {
	fmt.Println("我是技能3")
}

func (*SkilledWorker) SkillId4() {
	fmt.Println("我是技能4")
}

func (*SkilledWorker) SkillId5() {
	fmt.Println("我是技能5")
}

func (*SkilledWorker) SkillId6() {
	fmt.Println("我是技能6")
}

func (*SkilledWorker) SkillId7() {
	fmt.Println("我是技能7")
}

func (*SkilledWorker) SkillId8() {
	fmt.Println("我是技能8")
}

func (*SkilledWorker) SkillId9() {
	fmt.Println("我是技能9")
}

func (*SkilledWorker) SkillId10() {
	fmt.Println("我是技能10")
}

func (*SkilledWorker) SkillId11() {
	fmt.Println("我是技能11")
}

func (*SkilledWorker) SkillId12() {
	fmt.Println("我是技能12")
}
