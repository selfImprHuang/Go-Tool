/*
 *  @Author : huangzj
 *  @Time : 2021/3/16 21:53
 *  @Description：
 */

package NumberCount

type OnlyOneNumberOtherThreeObj struct{}

func (*OnlyOneNumberOtherThreeObj) Doc() string {
	return `
		给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。找出那个只出现了一次的元素。
		要求：线性时间复杂度。 不使用额外空间来实现
	`
}

func OnlyOneNumberOtherThree1(numList []int) int {
	bitList := make([]int, 32)

	for _, num := range numList {
		for i := 0; i < 32; i++ {
			bitList[i] += (num >> i) & 1
		}
	}

	res := 0

	for i := 0; i < 32; i++ {
		if bitList[i]%3 != 0 {
			res += 1 << i
		}
	}

	return res
}

//通过二进制的方式，本质上就是要构造一个相应的回路，即出现num的次数分别从1~3应该是  01 -> 10 -> 00 ...
func OnlyOneNumberOtherThree2(numList []int) int {

	var a uint
	var b uint

	//第一次出现num计算的结果 a = num ,b = 0
	//第二次出现num计算的结果 a = 0 ,b = num
	//第三次出现num计算的结果 a = num ,b = 0
	//因此直接返回a即可
	for _, num := range numList {
		a = a ^ uint(num)&(^b)
		b = b ^ uint(num)&(^a)
	}
	return int(a)
}
