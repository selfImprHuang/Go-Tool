/*
 *  @Author : huangzj
 *  @Time : 2020/8/5 11:50
 *  @Description：
 */

package knapsackOptimization

import "fmt"

type KnapsackSearchAnswer struct {
	BagCapacity int        //背包的总容量上限
	BagItemList []*BagItem //背包道具列表
	ItemNum     int        //道具的数量

	bestPath        []int //最佳路径(怎么存放价值最高)
	currentBestPath []int //最佳路径(当前解)
	bestValue       []int //每个容量最大价值
	currentValue    int   //当前价值
	currentWeight   int   //当前重量
	lastBestValue   int   //目前装载的最优价值
	surplusValue    int   //剩余道具价值
}

func NewKnapsackSearchAnswer(bagItemList []*BagItem, bagCapacity int) *KnapsackSearchAnswer {
	if bagCapacity <= 0 {
		panic("背包总容量不能小于等于0")
	}
	answer := &KnapsackSearchAnswer{
		BagCapacity:     bagCapacity,
		BagItemList:     bagItemList,
		ItemNum:         len(bagItemList),
		bestPath:        make([]int, len(bagItemList)+1),
		currentBestPath: make([]int, len(bagItemList)+1),
		bestValue:       make([]int, len(bagItemList)+1),
	}
	//初始化剩余道具价值等于所有道具总价值
	for _, r := range bagItemList {
		answer.surplusValue = answer.surplusValue + r.Value
	}

	return answer
}

//回溯法计算最佳路径和最大达到价值
func (cmd *KnapsackSearchAnswer) DynamicPlanForBestValue() {
	cmd.initParam()  //初始化第一个数组为空道具
	cmd.BackTrack(1) //从第一个道具开始进行路径搜索
}

func (cmd *KnapsackSearchAnswer) BackTrack(step int) {
	//遍历到最后一个道具
	if step > cmd.ItemNum {
		//如果当前总价值大于上一次装载的最佳价值
		if cmd.currentValue > cmd.lastBestValue {
			//最终的最佳路径用当前最佳路径来替换
			for i := 1; i <= cmd.ItemNum; i++ {
				cmd.bestPath[i] = cmd.currentBestPath[i]
			}

			cmd.lastBestValue = cmd.currentValue
			return
		}
	}

	//装入该道具之后，剩余的价值要减去当前道具的价值
	cmd.surplusValue = cmd.surplusValue - cmd.BagItemList[step].Value
	//当前道具的总重量没有超重
	if cmd.currentWeight+cmd.BagItemList[step].Weight <= cmd.BagCapacity {

		cmd.currentBestPath[step] = 1                                        //当前道具选中了，设置为1
		cmd.currentValue = cmd.currentValue + cmd.BagItemList[step].Value    //计算当前价值
		cmd.currentWeight = cmd.currentWeight + cmd.BagItemList[step].Weight //计算当前价值

		cmd.BackTrack(step + 1) //继续遍历下一个道具

		cmd.currentValue = cmd.currentValue - cmd.BagItemList[step].Value    //遍历完成要恢复成原来的价值
		cmd.currentWeight = cmd.currentWeight - cmd.BagItemList[step].Weight //遍历完成要恢复成原来的重量
	}

	//这边是剪枝的操作，如果剩余价值加上当前价值会超过上一次的最佳价值，那么就要往右子树遍历，否则不操作.
	if cmd.surplusValue+cmd.currentValue > cmd.lastBestValue {
		cmd.currentBestPath[step] = 0 //当前路径置为0，不选择该位置
		cmd.BackTrack(step + 1)

	}

	cmd.surplusValue = cmd.surplusValue + cmd.BagItemList[step].Value
}

//输出最佳路径解
func (cmd *KnapsackSearchAnswer) PrintBestSearchPath() {

	for index, r := range cmd.bestPath {
		if r == 1 {
			fmt.Print("   ")
			fmt.Println(fmt.Sprintf("道具的价值是:%d,道具的重量是%d", cmd.BagItemList[index].Value, cmd.BagItemList[index].Weight))
		}
	}
}

func (cmd *KnapsackSearchAnswer) initParam() {
	itemList := make([]*BagItem, 0)
	itemList = append(itemList, &BagItem{})
	itemList = append(itemList, cmd.BagItemList...)
	cmd.BagItemList = itemList
}

func (cmd *KnapsackSearchAnswer) BestValue() int {
	return cmd.lastBestValue
}
