/*
 *  @Author : huangzj
 *  @Time : 2021/3/16 16:46
 *  @Description：
 */

package NumberMoreThan

type NumberMoreThanKObj struct{}

func (*NumberMoreThanKObj) Doc() string {
	return `
		给定一个整型数组arr(数组长度为N) 与一个整数k，打印所有出现次数大于 N/K 的数。如果没有这样的数，返回-1。

		要求：时间复杂度为O（N*K）,额外空间复杂度为O（K）。
	`
}

func NumberMoreThanK(numList []int, k int) []int {
	//小于2的情况是肯定不会超过.
	if k < 2 {
		return []int{}
	}

	//用来保存数字与其出现次数
	numMap := make(map[int]int, 0)
	for _, num := range numList {
		//如果不存在则数值设置为1
		if _, ok := numMap[num]; !ok {
			numMap[num] = 1
		} else {
			//存在则出现次数+1
			numMap[num] = numMap[num] + 1
		}
		//当容器的大小为k
		if len(numMap) == k {
			for key, value := range numMap {
				//个数正好等于1的要删掉，因为本次减一之后，就等于0了
				if value == 1 {
					delete(numMap, key)
				}
			}
			continue
		}

	}

	//map的key是可能出现次数超过N/k的数
	//这里主要是为了拿到key
	for key := range numMap {
		numMap[key] = 0
	}
	//重新计算对应数字的出现次数
	for _, num := range numList {
		if _, ok := numMap[num]; ok {
			numMap[num]++
		}
	}
	//出现次数超过k次的判断
	result := make([]int, 0)
	for key, value := range numMap {
		if value > len(numList)/k {
			result = append(result, key)
		}
	}

	return result
}
