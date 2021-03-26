/*
 *  @Author : huangzj
 *  @Time : 2021/3/23 17:06
 *  @Description：
 */

package FullPermutation

func DigitDecreases(b []int) [][]int {
	result := make([][]int, 0)
	bit := make([]int, len(b)-1)
	sum := 1
	//计算中的结果数,阶乘
	for i := 1; i <= len(b); i++ {
		sum = sum * i
	}
	for i := 0; i < sum; i++ {
		result = append(result, getRangeBySpace1(b, bit))
		bit = decreaseBit(bit)
	}
	return result
}

//根据递增进制位,并且根据【数空格法】在原来的数组中求得全排列的一种可能实现
//origin 原始数组
//increaseBit 递增进制位，比原始数组长度小1，这里通过数组实现，更简单
func getRangeBySpace1(origin []int, increaseBit []int) []int {
	permutation := make([]int, len(origin))
	flag := make(map[int]bool, 0)
	//排列数组从左往右，即从大到小
	for i := 0; i < len(origin)-1; i++ {
		p := increaseBit[i]  //从右到左，对应数字应该填充的位置
		j := len(origin) - 1 //该数字在本次排列的位置
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
			permutation[i] = origin[len(origin)-1]
		}
	}

	return permutation
}

//递减进制位
//每次加一,计算新的数组结果
func decreaseBit(bit []int) []int {

	//数组下标从小到大,进制位从大到小
	for i := 0; i < len(bit); i++ {
		if bit[i]+1 >= len(bit)+1-i {
			bit[i] = (bit[i] + 1) % (len(bit) + 1 - i)
		} else {
			bit[i] += 1
			break
		}
	}

	return bit
}
