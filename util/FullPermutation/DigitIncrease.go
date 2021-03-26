/*
 *  @Author : huangzj
 *  @Time : 2021/3/23 17:05
 *  @Description：
 */

package FullPermutation

func DigitIncrease(b []int) [][]int {
	result := make([][]int, 0)
	bit := make([]int, len(b)-1)
	sum := 1
	//计算中的结果数,阶乘
	for i := 1; i <= len(b); i++ {
		sum = sum * i
	}
	for i := 0; i < sum; i++ {
		result = append(result, getRangeBySpace(b, bit))
		bit = increaseBit(bit)
	}
	return result
}

//根据递增进制位,并且根据【数空格法】在原来的数组中求得全排列的一种可能实现
//origin 原始数组
//increaseBit 递增进制位，比原始数组长度小1，这里通过数组实现，更简单
func getRangeBySpace(origin []int, increaseBit []int) []int {
	permutation := make([]int, len(origin))
	flag := make(map[int]bool, 0)
	//排列数组从左往右，即从大到小
	for i := len(origin) - 1; i > 0; i-- {
		p := increaseBit[i-1] //从右到左，对应数字应该填充的位置
		j := len(origin) - 1  //该数字在本次排列的位置
		//中介数计算从左到右，即从大到小
		for ; ; j-- {
			//该位置未被填充
			if !flag[j] {
				p--
			}
			//找到对应位置需要推出循环
			if p < 0 {
				flag[j] = true
				break
			}
		}
		permutation[j] = origin[i]
	}
	for i := 0; i < len(origin); i++ {
		if !flag[i] {
			permutation[i] = origin[0]
		}
	}

	return permutation
}

//递增进制位
//根据递增进制位进行+1操作，返回递增的结果
func increaseBit(bit []int) []int {
	//最开始是从二进制开始
	for i, num := range bit {
		if num+1 >= i+2 {
			bit[i] = (num + 1) % (i + 2)
		} else {
			bit[i] = num + 1
			break
		}
	}

	return bit
}
