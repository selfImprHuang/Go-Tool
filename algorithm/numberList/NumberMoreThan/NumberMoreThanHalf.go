/*
 *  @Author : huangzj
 *  @Time : 2021/3/16 15:45
 *  @Description：
 */

package NumberMoreThan

type NumberMoreThanHalfObj struct{}

func (n *NumberMoreThanHalfObj) Doc() string {
	return `
		题目：
			
			数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。
			
			例如输入一个长度为9的数组{1,2,3,2,2,2,5,4,2}。由于数字2在数组中出现了5次，超过数组长度的一半，因此输出2。

			要求：时间复杂度O(N),空间复杂度O(1)
		`
}

func NumberMoreThanHalf(numList []int) int {
	var number, count int
	for _, j := range numList {
		if number == j {
			count++
		} else if number != j && count <= 1 {
			number = j
		} else {
			//number !=j && count >1
			count--
		}
	}

	//最后验证，确保一定超过
	countValid := 0
	for _, j := range numList {
		if j == number {
			countValid++
		}
	}
	if countValid <= len(numList)/2 {
		return -1
	}

	return number
}
