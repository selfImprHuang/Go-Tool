/*
 *  @Author : huangzj
 *  @Time : 2021/3/16 18:01
 *  @Description：
 */

package NumberCount

type OnlyTwoNumberShowOnceObj struct {
}

func (*OnlyTwoNumberShowOnceObj) Doc() string {
	return `
		一个整型数组里除了两个数字之外，其他的数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度为O(n)，空间复杂度为O(1)。
	`
}

func OnlyTwoNumberShowOnce(numList []int) (int, int) {
	if len(numList) < 2 {
		panic("数组长度不能小于两个")
	}
	result := 0
	//先进行异或，得到异或的结果
	for _, num := range numList {
		result = result ^ num
	}
	//找到异或后为1的位置
	pos := 0
	for ; pos < 32; pos++ {
		if ((result >> pos) & 1) == 1 {
			break
		}
	}

	firstList := make([]int, 0)
	secondList := make([]int, 0)

	//根据异或位的结果，把数组分成两组
	for _, num := range numList {
		if ((num >> pos) & 1) == 1 {
			firstList = append(firstList, num)
		} else {
			secondList = append(secondList, num)
		}
	}

	first := 0
	for _, num := range firstList {
		first ^= num
	}
	second := 0

	for _, num := range secondList {
		second ^= num
	}

	return first, second
}
