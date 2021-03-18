/*
 *  @Author : huangzj
 *  @Time : 2021/3/18 13:08
 *  @Description：
 */

package NumberCount

type OnlyOneNumberOtherKObj struct{}

func (*OnlyOneNumberOtherKObj) Doc() string {
	return `
		给定一个整型数组 arr和一个大于1的整数k。已知 arr中只有1个数出现了1次，其他的数都出现了k次，请返回只出现了1次的数。

		【要求】 时间复杂度为 O（N），额外空间复杂度为 O（1）。
	`
}

func OnlyOneNumberOtherK(numList []int, k int) int {
	bitList := make([]int, 32)
	for _, num := range numList {
		//把每个数字转换成二进制
		store := make([]int, 32)
		for i := 0; num != 0; i++ {
			store[i] = num % k
			num = num / k
		}

		for j := 0; j < 32; j++ {
			//不进位加法
			bitList[j] = (store[j] + bitList[j]) % k
		}
	}

	//k进制转换回十进制
	power := 1  //k的次方值
	result := 0 //只出现一次的数字
	for i, bit := range bitList {
		power = 1
		for j := 0; j < i; j++ {
			power = power * k
		}

		result = result + power*bit
	}

	return result
}
