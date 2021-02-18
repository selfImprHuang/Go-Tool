/*
 *  @Author : huangzj
 *  @Time : 2020/10/27 15:19
 *  @Description：
 */

package bitStore

import (
	"fmt"
	"math"
	"testing"
)

func TestBitStore(t *testing.T) {
	fmt.Println("测试2档的情况")
	testTwo()
	fmt.Println()

	fmt.Println("测试10档的情况")
	testTen()
	fmt.Println()

	fmt.Println("测试18档的情况")
	testEighteen()
	fmt.Println()

	fmt.Println("测试33档的情况")
	testThirtyThree()
	fmt.Println()

	fmt.Println("测试58档的情况")
	testFiftyEight()
	fmt.Println()

	fmt.Println("测试1000档的情况")
	testOneThousand(t)
	fmt.Println()
}

func TestBitStoreFind(t *testing.T) {
	bitStore := NewBitStore(120, nil)
	bitStore.ReceiveByGear(1)
	bitStore.ReceiveByGear(3)
	bitStore.ReceiveByGear(5)
	bitStore.ReceiveByGear(7)
	bitStore.ReceiveByGear(13)
	bitStore.ReceiveByGear(10)
	bitStore.ReceiveByGear(16)
	bitStore.ReceiveByGear(33)
	bitStore.ReceiveByGear(35)

	fmt.Println(bitStore.FindGearMap())
	fmt.Println(bitStore.FindSpecialGearList(1))
	fmt.Println(len(bitStore.FindSpecialGearList(1)))
	fmt.Println(bitStore.FindSpecialGearList(4))
	fmt.Println(len(bitStore.FindSpecialGearList(4)))
	fmt.Println(bitStore.FindSpecialGearList(2))
	fmt.Println(len(bitStore.FindSpecialGearList(2)))
	fmt.Println(bitStore.FindSpecialGearList(10))
	fmt.Println(len(bitStore.FindSpecialGearList(10)))
	fmt.Println(bitStore.FindSpecialGearList(13))
	fmt.Println(len(bitStore.FindSpecialGearList(13)))
	fmt.Println(bitStore.FindSpecialGearList(14))
	fmt.Println(len(bitStore.FindSpecialGearList(14)))
	fmt.Println(bitStore.FindSpecialGearList(20))
	fmt.Println(len(bitStore.FindSpecialGearList(20)))
	fmt.Println(bitStore.FindSpecialGearList(32))
	fmt.Println(len(bitStore.FindSpecialGearList(32)))
	fmt.Println(bitStore.FindSpecialGearList(63))
	fmt.Println(len(bitStore.FindSpecialGearList(63)))
	fmt.Println(bitStore.FindSpecialGearList(64))
	fmt.Println(len(bitStore.FindSpecialGearList(64)))
	fmt.Println(bitStore.FindSpecialGearList(65))
	fmt.Println(len(bitStore.FindSpecialGearList(65)))
}

func newStoreMap(bitStore *BitStore) {
	gearMap := bitStore.FindAllGearMap()
	falseList := make([]int, 0)
	for key, value := range gearMap {
		if value == true {
			falseList = append(falseList, key)
		}
	}
	if len(falseList) == 0 {
		fmt.Println("当前无未领取档位")
	}
}

func getReceiveResult(bitStore *BitStore) {
	gearMap1 := bitStore.FindAllGearMap()
	keyList := make([]int, 0)
	for key, value := range gearMap1 {
		if value == true {
			keyList = append(keyList, key)
		}
	}
	fmt.Println("已经领取的档位信息如下")
	for value := range keyList {
		fmt.Print(fmt.Sprintf("%d,", keyList[value]))
	}
}

func testOneThousand(t *testing.T) {
	bitStore := NewBitStore(1000, nil)

	newStoreMap(bitStore) //初始化的校验

	fmt.Println()
	bitStore.ReceiveByGear(3)
	fmt.Println(fmt.Sprintf("%b", bitStore.GearPickList[0]))
	fmt.Println(fmt.Sprintf("这个时候领取结果  3应该返回true，结果是：%v", bitStore.IsGearReceive(3)))
	bitStore.ReceiveByGear(8)
	fmt.Println(fmt.Sprintf("%b", bitStore.GearPickList[0]))
	fmt.Println(fmt.Sprintf("这个时候领取结果  8应该返回true，结果是：%v", bitStore.IsGearReceive(8)))
	bitStore.ReceiveByGear(16)
	fmt.Println(fmt.Sprintf("%b", bitStore.GearPickList[0]))
	fmt.Println(fmt.Sprintf("这个时候领取结果  16应该返回true，结果是：%v", bitStore.IsGearReceive(16)))
	bitStore.ReceiveByGear(18)
	fmt.Println(fmt.Sprintf("%b", bitStore.GearPickList[0]))
	fmt.Println(fmt.Sprintf("这个时候领取结果  18应该返回true，结果是：%v", bitStore.IsGearReceive(18)))
	bitStore.ReceiveByGear(27)
	fmt.Println(fmt.Sprintf("%b", bitStore.GearPickList[0]))
	fmt.Println(fmt.Sprintf("这个时候领取结果  27应该返回true，结果是：%v", bitStore.IsGearReceive(27)))
	bitStore.ReceiveByGear(32)
	fmt.Println(fmt.Sprintf("%b", bitStore.GearPickList[1]))
	fmt.Println(fmt.Sprintf("这个时候领取结果  32应该返回true，结果是：%v", bitStore.IsGearReceive(32)))
	bitStore.ReceiveByGear(33)
	fmt.Println(fmt.Sprintf("%b", bitStore.GearPickList[1]))
	fmt.Println(fmt.Sprintf("这个时候领取结果  33应该返回true，结果是：%v", bitStore.IsGearReceive(33)))
	bitStore.ReceiveByGear(53)
	fmt.Println(fmt.Sprintf("%b", bitStore.GearPickList[1]))
	fmt.Println(fmt.Sprintf("这个时候领取结果  53应该返回true，结果是：%v", bitStore.IsGearReceive(53)))
	bitStore.ReceiveByGear(44)
	fmt.Println(fmt.Sprintf("%b", bitStore.GearPickList[1]))
	fmt.Println(fmt.Sprintf("这个时候领取结果  44应该返回true，结果是：%v", bitStore.IsGearReceive(44)))
	bitStore.ReceiveByGear(48)
	fmt.Println(fmt.Sprintf("%b", bitStore.GearPickList[1]))
	fmt.Println(fmt.Sprintf("这个时候领取结果  48应该返回true，结果是：%v", bitStore.IsGearReceive(48)))

	fmt.Println(fmt.Sprintf("判断当前数组是否动态增长，当前数组长度为:%d", len(bitStore.GearPickList)))

	bitStore.ReceiveByGear(123)
	fmt.Println(fmt.Sprintf("这个时候领取结果  123应该返回true，结果是：%v", bitStore.IsGearReceive(123)))

	fmt.Println(fmt.Sprintf("判断当前数组是否动态增长，当前数组长度为:%d", len(bitStore.GearPickList)))
	getReceiveResult(bitStore)
	fmt.Println()

	bitStore.ReceiveByGear(376)
	fmt.Println(fmt.Sprintf("这个时候领取结果  376应该返回true，结果是：%v", bitStore.IsGearReceive(376)))
	fmt.Println(fmt.Sprintf("判断当前数组是否动态增长，当前数组长度为:%d", len(bitStore.GearPickList)))
	bitStore.ReceiveByGear(588)
	fmt.Println(fmt.Sprintf("这个时候领取结果  588应该返回true，结果是：%v", bitStore.IsGearReceive(588)))
	fmt.Println(fmt.Sprintf("判断当前数组是否动态增长，当前数组长度为:%d", len(bitStore.GearPickList)))
	bitStore.ReceiveByGear(997)
	fmt.Println(fmt.Sprintf("这个时候领取结果  997应该返回true，结果是：%v", bitStore.IsGearReceive(997)))

	fmt.Println()
	getReceiveResult(bitStore)
}

func testFiftyEight() {
	bitStore := NewBitStore(58, nil)

	newStoreMap(bitStore) //初始化的校验

	fmt.Println()
	bitStore.ReceiveByGear(3)
	fmt.Println(fmt.Sprintf("这个时候领取结果  3应该返回true，结果是：%v", bitStore.IsGearReceive(3)))
	bitStore.ReceiveByGear(8)
	fmt.Println(fmt.Sprintf("这个时候领取结果  8应该返回true，结果是：%v", bitStore.IsGearReceive(8)))
	bitStore.ReceiveByGear(16)
	fmt.Println(fmt.Sprintf("这个时候领取结果  16应该返回true，结果是：%v", bitStore.IsGearReceive(16)))
	bitStore.ReceiveByGear(18)
	fmt.Println(fmt.Sprintf("这个时候领取结果  18应该返回true，结果是：%v", bitStore.IsGearReceive(18)))
	bitStore.ReceiveByGear(27)
	fmt.Println(fmt.Sprintf("这个时候领取结果  27应该返回true，结果是：%v", bitStore.IsGearReceive(27)))
	bitStore.ReceiveByGear(32)
	fmt.Println(fmt.Sprintf("这个时候领取结果  32应该返回true，结果是：%v", bitStore.IsGearReceive(32)))
	bitStore.ReceiveByGear(33)
	fmt.Println(fmt.Sprintf("这个时候领取结果  33应该返回true，结果是：%v", bitStore.IsGearReceive(33)))
	bitStore.ReceiveByGear(53)
	fmt.Println(fmt.Sprintf("这个时候领取结果  53应该返回true，结果是：%v", bitStore.IsGearReceive(53)))
	bitStore.ReceiveByGear(44)
	fmt.Println(fmt.Sprintf("这个时候领取结果  44应该返回true，结果是：%v", bitStore.IsGearReceive(44)))
	bitStore.ReceiveByGear(48)
	fmt.Println(fmt.Sprintf("这个时候领取结果  48应该返回true，结果是：%v", bitStore.IsGearReceive(48)))
	fmt.Println()
	gearMap1 := bitStore.FindAllGearMap()
	for key, value := range gearMap1 {
		fmt.Println(fmt.Sprintf("当前map的key为%d,当前map的value为%v", key, value))
	}
}

func testThirtyThree() {
	bitStore := NewBitStore(33, nil)

	newStoreMap(bitStore) //初始化的校验
	fmt.Println()
	bitStore.ReceiveByGear(3)
	fmt.Println(fmt.Sprintf("这个时候领取结果  3应该返回true，结果是：%v", bitStore.IsGearReceive(3)))
	bitStore.ReceiveByGear(8)
	fmt.Println(fmt.Sprintf("这个时候领取结果  8应该返回true，结果是：%v", bitStore.IsGearReceive(8)))
	bitStore.ReceiveByGear(16)
	fmt.Println(fmt.Sprintf("这个时候领取结果  16应该返回true，结果是：%v", bitStore.IsGearReceive(16)))
	bitStore.ReceiveByGear(18)
	fmt.Println(fmt.Sprintf("这个时候领取结果  18应该返回true，结果是：%v", bitStore.IsGearReceive(18)))
	bitStore.ReceiveByGear(27)
	fmt.Println(fmt.Sprintf("这个时候领取结果  27应该返回true，结果是：%v", bitStore.IsGearReceive(27)))
	bitStore.ReceiveByGear(32)
	fmt.Println(fmt.Sprintf("这个时候领取结果  32应该返回true，结果是：%v", bitStore.IsGearReceive(32)))
	bitStore.ReceiveByGear(33)
	fmt.Println(fmt.Sprintf("这个时候领取结果  33应该返回true，结果是：%v", bitStore.IsGearReceive(33)))

	fmt.Println()
	gearMap1 := bitStore.FindAllGearMap()
	for key, value := range gearMap1 {
		fmt.Println(fmt.Sprintf("当前map的key为%d,当前map的value为%v", key, value))
	}
}

func testEighteen() {
	bitStore := NewBitStore(18, nil)

	newStoreMap(bitStore) //初始化的校验
	fmt.Println()
	bitStore.ReceiveByGear(3)
	fmt.Println(fmt.Sprintf("这个时候领取结果  3应该返回true，结果是：%v", bitStore.IsGearReceive(3)))
	bitStore.ReceiveByGear(8)
	fmt.Println(fmt.Sprintf("这个时候领取结果  8应该返回true，结果是：%v", bitStore.IsGearReceive(8)))
	bitStore.ReceiveByGear(16)
	fmt.Println(fmt.Sprintf("这个时候领取结果  16应该返回true，结果是：%v", bitStore.IsGearReceive(16)))
	bitStore.ReceiveByGear(18)
	fmt.Println(fmt.Sprintf("这个时候领取结果  18应该返回true，结果是：%v", bitStore.IsGearReceive(18)))

	fmt.Println()
	gearMap1 := bitStore.FindAllGearMap()
	for key, value := range gearMap1 {
		fmt.Println(fmt.Sprintf("当前map的key为%d,当前map的value为%v", key, value))
	}
}

func testTen() {
	bitStore := NewBitStore(10, nil)

	newStoreMap(bitStore) //初始化的校验
	fmt.Println()
	bitStore.ReceiveByGear(3)
	fmt.Println(fmt.Sprintf("这个时候领取结果  3应该返回true，结果是：%v", bitStore.IsGearReceive(3)))

	fmt.Println()
	bitStore.ReceiveByGear(7)
	fmt.Println(fmt.Sprintf("这个时候领取结果  7应该返回true，结果是：%v", bitStore.IsGearReceive(7)))

	fmt.Println()
	gearMap1 := bitStore.FindAllGearMap()
	for key, value := range gearMap1 {
		fmt.Println(fmt.Sprintf("当前map的key为%d,当前map的value为%v", key, value))
	}
}

func testTwo() {
	bitStore := NewBitStore(2, nil)

	newStoreMap(bitStore) //初始化的校验
	fmt.Println()
	bitStore.ReceiveByGear(1)
	fmt.Println(fmt.Sprintf("这个时候领取结果  1应该返回true，结果是：%v", bitStore.IsGearReceive(1)))

	fmt.Println()
	gearMap1 := bitStore.FindAllGearMap()
	for key, value := range gearMap1 {
		fmt.Println(fmt.Sprintf("当前map的key为%d,当前map的value为%v", key, value))
	}
}

//测试溢出
func TestOverflow(t *testing.T) {
	bitStore := NewBitStore(31, nil)
	for i := 1; i <= 31; i++ {
		bitStore.ReceiveByGear(i)
	}

	fmt.Println(fmt.Sprintf("MaxInt32 = %v", math.MaxInt32))

	for _, r := range bitStore.GearPickList {
		fmt.Println(r)
	}

	bitStore = NewBitStore(32, nil)
	for i := 1; i <= 32; i++ {
		bitStore.ReceiveByGear(i)
	}

	for _, r := range bitStore.GearPickList {
		fmt.Println(r)
	}

	bitStore = NewBitStore(33, nil)
	for i := 1; i <= 33; i++ {
		bitStore.ReceiveByGear(i)
	}

	fmt.Println("数组信息")
	for _, r := range bitStore.GearPickList {
		fmt.Println(r)
	}

	for key, value := range bitStore.FindAllGearMap() {
		fmt.Println(key, value)
	}
}
