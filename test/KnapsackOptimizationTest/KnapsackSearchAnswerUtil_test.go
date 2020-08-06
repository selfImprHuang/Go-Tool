/*
 *  @Author : huangzj
 *  @Time : 2020/8/5 15:40
 *  @Description：
 */

package KnapsackOptimizationTest

import (
	"Go-Tool/util/KnapsackOptimization"
	"fmt"
	"testing"
)

func TestKnapsackSearchAnswerUtil(t *testing.T) {
	bagItem := make([]*KnapsackOptimization.BagItem, 0)
	bagItem = append(bagItem, &KnapsackOptimization.BagItem{
		Value:  3,
		Weight: 2,
	})
	bagItem = append(bagItem, &KnapsackOptimization.BagItem{
		Value:  3,
		Weight: 3,
	})
	bagItem = append(bagItem, &KnapsackOptimization.BagItem{
		Value:  4,
		Weight: 4,
	})
	bagItem = append(bagItem, &KnapsackOptimization.BagItem{
		Value:  5,
		Weight: 5,
	})
	bagItem = append(bagItem, &KnapsackOptimization.BagItem{
		Value:  6,
		Weight: 2,
	})
	bagItem = append(bagItem, &KnapsackOptimization.BagItem{
		Value:  7,
		Weight: 2,
	})
	//通过回溯解背包问题
	cmd := KnapsackOptimization.NewKnapsackSearchAnswer(bagItem, 12)
	cmd.DynamicPlanForBestValue()
	cmd.PrintBestSearchPath()
	fmt.Println(fmt.Sprintf("最佳价值为：%d", cmd.BestValue()))
}
