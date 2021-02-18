/*
 *  @Author : huangzj
 *  @Time : 2020/10/27 15:19
 *  @Description：按照位数进行存储.这里主要是针对（0，1）状态的场景。比如说奖励领取情况，0表示未领取，1表示领取
 * 	这边有两个参数
 * 					MaxGear：当前存储的最大挡位是什么，在初始化的时候，需要对对应的领取情况进行数组初始化
 *  				GearPickList:挡位领取情况，这边初始化的bit为32，总共有32位，每一位代表一个挡位的领取情况，并且这是一个数组，就代表了 32 - 64 - 96 ...最大 32  * list长度 挡位的所有领取情况
 */

package bitStore

import "fmt"

const (
	bit     = 32 //作为标识，没有实际作用
	realBit = 31 //真实参与计算的每个数字代表的档位数量
)

type BitStore struct {
	MaxGear      int   //存储的最大挡位
	GearPickList []int //用来表示档位的数组(每一个数表示【下标 * bit】位的档位存储结果)
}

//初始化方法，gear表示最大的领取档位是多少.(档位刚开始计算，没有值的情况)
func NewBitStore(gear int, gearPickList []int) *BitStore {
	//校验传入的档位是否正确
	if gear < 1 {
		panic("对应的最大档位不能小于1")
	}
	//初始化不初始化他的长度，而是后面需要的时候再进行动态增长
	return &BitStore{
		MaxGear:      gear,
		GearPickList: gearPickList,
	}
}

//返回指定档位长度的领取状态数组
func (bitStore *BitStore) FindSpecialGearList(gear int) []bool {
	bitStore.checkGearRight(gear)           //校验传进来的长度是否正确
	bitStore.getSizeOrInit()                //如果数组没有初始化，进行初始化
	return bitStore.getGearListByGear(gear) //返回指定长度的领取状态数组
}

//获取当前有记录的档位的Map
func (bitStore *BitStore) FindGearMap() map[int]bool {
	listSize := bitStore.getSizeOrInit()              //获取当前数组长度，如果数组没有初始化，进行初始化
	return bitStore.shortListDeal(listSize, listSize) //短数组的处理(没有达到最大大小的数组长度)
}

//获取所有档位对应的Map
func (bitStore *BitStore) FindAllGearMap() map[int]bool {
	length := bitStore.getGearListLength(bitStore.MaxGear) //获取最大档位的长度
	listSize := bitStore.getSizeOrInit()                   //获取当前数组长度，如果数组没有初始化，进行初始化
	if listSize < length {
		return bitStore.shortListDeal(listSize, length) //短数组的处理(没有达到最大大小的数组长度)
	}
	return bitStore.fullListDeal(listSize, length) //完整的数组的处理
}

//判断当前挡位是否被领取，true表示已经领取，false表示未领取
func (bitStore *BitStore) IsGearReceive(gear int) bool {
	bitStore.checkGearRight(gear) //检验一下档位是否正确
	length := bitStore.getGearListLength(gear)
	if bitStore.getSizeOrInit() < length {
		return false
	}
	position := bitStore.getGearPosition(gear, length)
	return (bitStore.GearPickList[length-1] & (1 << position)) > 0
}

//获取对应档位奖励
func (bitStore *BitStore) ReceiveByGear(gear int) ([]int, map[int]bool) {
	bitStore.checkIfReceive(gear)                      //校验是否被领取了
	length := bitStore.dynamicGrowth(gear)             //数组进行动态增长
	position := bitStore.getGearPosition(gear, length) //获取档位在数组中的位数位置
	bitStore.GearPickList[length-1] = bitStore.GearPickList[length-1] + (1 << uint(position))
	return bitStore.GearPickList, bitStore.FindAllGearMap()
}

//判断档位是否正确
func (bitStore *BitStore) checkGearRight(gear int) {
	if gear > bitStore.MaxGear {
		panic(fmt.Sprintf("档位信息有误，最大档位为：%d", bitStore.MaxGear))
	}
}

//获取当前档位对应的数组长度（这边的length是从1开始的.）
func (bitStore *BitStore) getGearListLength(gear int) int {
	return ((gear - 1) + realBit) / realBit
}

//获取档位所在的二进制位置
func (bitStore *BitStore) getGearPosition(gear, length int) int {
	return gear - realBit*(length-1) - 1
}

func (bitStore *BitStore) getSizeOrInit() int {
	if bitStore.GearPickList == nil {
		bitStore.GearPickList = make([]int, 0)
	}
	return len(bitStore.GearPickList)
}

func (bitStore *BitStore) fullListDeal(listSize int, length int) map[int]bool {
	resultMap := make(map[int]bool, 0)
	for i := 0; i < length; i++ {
		for j := 0; j < realBit && bitStore.MaxGear > i*realBit+j; j++ {
			resultMap[i*realBit+j+1] = (bitStore.GearPickList[i] & (1 << uint(j))) > 0
		}
	}
	return resultMap
}

func (bitStore *BitStore) shortListDeal(listSize, length int) map[int]bool {
	resultMap := make(map[int]bool, 0)
	for i := 0; i < listSize; i++ {
		for j := 0; j < realBit; j++ {
			resultMap[i*realBit+j+1] = (bitStore.GearPickList[i] & (1 << uint(j))) > 0
		}
	}
	for i := listSize; i < length; i++ {
		for j := 0; j < realBit && bitStore.MaxGear > i*realBit+j; j++ {
			resultMap[i*realBit+j+1] = false
		}
	}
	return resultMap
}

func (bitStore *BitStore) dynamicGrowth(gear int) int {
	listSize := bitStore.getSizeOrInit() //获取当前数组长度，如果数组没有初始化，进行初始化
	length := bitStore.getGearListLength(gear)
	//如果数组的长度小于对应的档位长度，这个时候需要去扩展这个数组
	if listSize < length {
		for i := 0; i < (length - listSize); i++ {
			bitStore.GearPickList = append(bitStore.GearPickList, 0)
		}
	}
	return length
}

func (bitStore *BitStore) checkIfReceive(gear int) {
	//判断是否领取已经进行了档位的校验
	if bitStore.IsGearReceive(gear) {
		panic("当前档位已被领取，不可重复领取")
	}
}

func (bitStore *BitStore) getGearListByGear(gear int) []bool {
	gearList := make([]bool, 0)
	listAllGear := len(bitStore.GearPickList) * realBit //当前已领取数组可表示的总档位
	//如果指定的长度大于gear已有数组长度，需要特殊处理
	if listAllGear <= gear {
		for i := 0; i < (listAllGear-1)/realBit+1; i++ {
			for j := 0; j < realBit && realBit*i+j < gear; j++ {
				gearList = append(gearList, (bitStore.GearPickList[i]&(1<<uint(j))) > 0)
			}
		}
		for i := 0; i < gear-listAllGear; i++ {
			for j := 0; j < realBit && realBit*i+j < gear-listAllGear; j++ {
				gearList = append(gearList, false)
			}
		}
		return gearList
	}
	//如果获取的档位比实际记录的档位小
	for i := 0; i < (gear-1)/realBit+1; i++ {
		for j := 0; j < realBit && realBit*i+j < gear; j++ {
			gearList = append(gearList, (bitStore.GearPickList[i]&(1<<uint(j))) > 0)
		}
	}
	return gearList
}
