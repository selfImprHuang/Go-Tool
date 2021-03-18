/*
 *  @Author : huangzj
 *  @Time : 2021/3/16 17:51
 *  @Description：
 */

package NumberCount

type OnlyOneNumberShowOnceObj struct{}

func (*OnlyOneNumberShowOnceObj) Doc() string {
	return `
		一个整型数组里除了一个数字之外，其他的数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度为O(n)，空间复杂度为O(1)。
	`
}

//通过异或解决
func OnlyOneNumberShowOnce(numberList []int) int {
	result := 0
	for _, num := range numberList {
		result ^= num
	}

	return result
}
