/*
 *  @Author : huangzj
 *  @Time : 2020/7/31 15:35
 *  @Description：背包最优问题工具类
 *  基本场景：给定n个重量为w1w 1，w2w 2 ,w3w 3 ,…,wnw n ，价值为v1v 1 ,v2v 2 ,v3v 3 ,…,vnv n 的物品和容量为CC的背包，求这个物品中一个最有价值的子集，使得在满足背包的容量的前提下，包内的总价值最大
 *  参考地址：https://blog.csdn.net/chanmufeng/article/details/82955730
 */

package KnapsackOptimizationTest

import (
	"Go-Tool/util/KnapsackOptimization"
	"fmt"
	"testing"
)

func TestKnapsackOptimizationUtil(t *testing.T) {
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
	cmd := KnapsackOptimization.NewKnapsackOptimization(bagItem, 12)
	cmd1 := KnapsackOptimization.NewKnapsackOptimization(bagItem, 12)
	cmd2 := KnapsackOptimization.NewKnapsackOptimization(bagItem, 12)
	fmt.Print(cmd.OptimizePackageByRecursion()) //通过递归解背包问题

	fmt.Println(cmd1.KnapsackCycle()) //通过逆序解背包问题.

	fmt.Println(cmd2.KnapsackCycleSimple()) //通过一维数组来解决背包问题
}
