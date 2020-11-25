/*
 *  @Author : huangzj
 *  @Time : 2020/11/24 16:17
 *  @Description：
 */

package linqUse

type Person struct {
	Name string //拥有者的名字
}

type Pet struct {
	Name      string //动物的名字
	OwnerName string //拥有者的名字
}

type Man struct {
	Name string //拥有者的名字
	Pets []Pets //宠物们
}

type Pets struct {
	Name string //动物的名字
}

func MakeInnerData() []Man {
	man := make([]Man, 0)
	man = append(man, Man{
		Name: "康康",
		Pets: []Pets{{Name: "康康的狗"}, {Name: "康康的猫"}},
	})
	man = append(man, Man{
		Name: "老施",
		Pets: []Pets{{Name: "老施的🐟"}, {Name: "老施的鸟"}},
	})
	man = append(man, Man{
		Name: "小明",
		Pets: []Pets{{Name: "小明的🐖"}, {Name: "小明的狗"}},
	})

	return man
}

func MakeJoinData() ([]Person, []Pet) {
	kangkang := Person{Name: "爱吃合合乐的康康"}
	laoshi := Person{Name: "老施"}
	xiaoming := Person{Name: "不要催-小明"}
	expect := Person{Name: "我是没有宠物的人"}

	dog := Pet{Name: "康康的狗", OwnerName: kangkang.Name}
	cat := Pet{Name: "康康的猫", OwnerName: kangkang.Name}
	fish := Pet{Name: "老施的🐟", OwnerName: laoshi.Name}
	pig := Pet{Name: "小明的🐖", OwnerName: xiaoming.Name}

	return []Person{kangkang, laoshi, xiaoming, expect}, []Pet{dog, cat, fish, pig}
}
