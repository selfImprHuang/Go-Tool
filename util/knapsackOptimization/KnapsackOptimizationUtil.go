/*
 *  @Author : huangzj
 *  @Time : 2020/7/31 15:35
 *  @Description：背包最优问题工具类
 *  基本场景：给定n个重量为w1w 1，w2w 2 ,w3w 3 ,…,wnw n ，价值为v1v 1 ,v2v 2 ,v3v 3 ,…,vnv n 的物品和容量为CC的背包，求这个物品中一个最有价值的子集，使得在满足背包的容量的前提下，包内的总价值最大
 *  参考地址：https://blog.csdn.net/chanmufeng/article/details/82955730
 */

package knapsackOptimization

type KnapsackOptimization struct {
	BagCapacity int        //背包的总容量上限
	BagItemList []*BagItem //背包道具列表
	ItemNum     int        //道具的数量
}

func NewKnapsackOptimization(bagItemList []*BagItem, bagCapacity int) *KnapsackOptimization {
	if bagCapacity <= 0 {
		panic("背包总容量不能小于等于0")
	}
	return &KnapsackOptimization{
		BagCapacity: bagCapacity,
		BagItemList: bagItemList,
		ItemNum:     len(bagItemList),
	}
}

//通过递归的方式解决上述基本场景，这种方式效率比较低下
func (cmd *KnapsackOptimization) OptimizePackageByRecursion() int {
	return cmd.knapsackRecursion(cmd.BagItemList, len(cmd.BagItemList)-1, cmd.BagCapacity)
}

func (cmd *KnapsackOptimization) knapsackRecursion(bagItemList []*BagItem, index, capacity int) int {
	if index < 0 || capacity <= 0 {
		return 0
	}
	res := cmd.knapsackRecursion(bagItemList, index-1, capacity) //不放第index个物品所得价值
	//放第index个物品所得价值（前提是：第index个物品可以放得下）
	if bagItemList[index].Weight <= capacity {
		res = max(res, bagItemList[index].Value+cmd.knapsackRecursion(bagItemList, index-1, capacity-bagItemList[index].Weight))
	}
	return res
}

func (cmd *KnapsackOptimization) KnapsackCycle() int {
	dynamicPlan := cmd.initKnapsackParam() //初始化参数，这边通过动态规划的方式处理，需要一个辅助行和辅助列
	//循环所有的道具，进行动态规划二维数组的组装.
	for i := 1; i <= cmd.ItemNum; i++ {
		for j := 1; j <= cmd.BagCapacity; j++ {
			//如果背包重量不够,则不能装下这个道具
			if j < cmd.BagItemList[i].Weight {
				dynamicPlan[i][j] = dynamicPlan[i-1][j]
			} else {
				//如果背包重量足够的话，判断装入这个道具价值更高，还是不装这个道具价值更高
				dynamicPlan[i][j] = max(dynamicPlan[i-1][j], dynamicPlan[i-1][j-cmd.BagItemList[i].Weight]+cmd.BagItemList[i].Value)
			}
		}
	}

	return dynamicPlan[len(cmd.BagItemList)-1][cmd.BagCapacity-1]
}

//通过一维数组实现背包求解
func (cmd *KnapsackOptimization) KnapsackCycleSimple() int {
	dynamicPlan := cmd.initKnapsackParamSimple()
	for i := 1; i <= cmd.ItemNum; i++ {
		for j := cmd.BagCapacity; j >= cmd.BagItemList[i].Weight; j-- {
			dynamicPlan[j] = max(dynamicPlan[j], dynamicPlan[j-cmd.BagItemList[i].Weight]+cmd.BagItemList[i].Value)
		}
	}
	return dynamicPlan[cmd.BagCapacity]
}

func (cmd *KnapsackOptimization) initKnapsackParamSimple() []int {
	itemList := make([]*BagItem, 0)
	itemList = append(itemList, &BagItem{})
	itemList = append(itemList, cmd.BagItemList...)
	cmd.BagItemList = itemList
	return make([]int, cmd.BagCapacity+1)
}

//这里需要组成一个辅助数组去判断动态规划的实现
func (cmd *KnapsackOptimization) initKnapsackParam() [][]int {
	itemList := make([]*BagItem, 0)
	itemList = append(itemList, &BagItem{})
	itemList = append(itemList, cmd.BagItemList...)
	cmd.BagItemList = itemList
	//初始化数组
	dynamicPlan := make([][]int, len(cmd.BagItemList))
	for i := 0; i <= cmd.ItemNum; i++ {
		dynamicPlan[i] = make([]int, cmd.BagCapacity+1)
	}
	return dynamicPlan
}

func max(compareNum, otherCompareNum int) int {
	if compareNum > otherCompareNum {
		return compareNum
	}
	return otherCompareNum
}
