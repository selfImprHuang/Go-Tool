/*
 *  @Author : huangzj
 *  @Time : 2020/8/18 14:11
 *  @Description：Lomuto划分,找到一个数组中第k个小的元素.这边只是提供一个示例，具体类型可以根据该算法进行修改
 */

package search

//参数说明：
//list  ：数组
//start : 开始比较的数字下标
//end   ：结束比较的数字下标
//作用说明：
func partition(list []int, start int, end int) int {
	moveSubscript := start
	compareValue := list[start]
	for i := start + 1; i <= end; i++ {
		if list[i] < compareValue {
			moveSubscript++
			list[moveSubscript], list[i] = list[i], list[moveSubscript]
		}
	}
	list[moveSubscript], list[start] = list[start], list[moveSubscript] //交换元素
	return moveSubscript
}

func LomutoQuiteSelect(list []int, pointPosition int) int {
	positionCheck(pointPosition, len(list)) //下标校验

	nowPosition := partition(list, 0, len(list)-1)
	//循环进行lomuto划分处理,知道找到我们需要的那个位置的元素
	for nowPosition != pointPosition-1 {
		if nowPosition > pointPosition-1 {
			nowPosition = partition(list, 0, nowPosition)
		} else {
			nowPosition = partition(list, nowPosition+1, len(list)-1)
		}
	}

	return list[nowPosition]
}

func positionCheck(pointPosition int, len int) {
	if pointPosition <= 0 || pointPosition > len {
		panic("查询的元素下标不能小于0，或者大于数组长度")
	}
}
