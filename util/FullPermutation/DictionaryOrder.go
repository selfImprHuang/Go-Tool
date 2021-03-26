/*
 *  @Author : huangzj
 *  @Time : 2021/3/22 15:29
 *  @Description：
 */

package FullPermutation

import (
	sort2 "Go-Tool/util/sort"
)

//字典序获取全排列
func (per *Permutation) DictionaryOrder(b []byte) {
	//先排序
	sort2.QuickSortByte(b, 0, len(b)-1)
	for {
		bef := 0
		afr := 0
		result := make([]byte, len(b))
		copy(result, b)
		per.bytes = append(per.bytes, result)

		//从右往左找到第一个左边比右边小的数字
		for bef = len(b) - 2; ; bef-- {
			if bef < 0 {
				return
			}
			if b[bef] < b[bef+1] {
				break
			}
		}

		//知道q的位置之后
		//从q开始向右边查找比位置q的数字大的第一个数字的位置
		for afr = len(b) - 1; afr > 0; afr-- {
			if b[afr] > b[bef] {
				break
			}
		}

		//交换两个数字
		b[bef], b[afr] = b[afr], b[bef]
		//把后面的数字进行转置成递增顺序
		b = per.Reverse(b, bef+1, len(b)-1)
	}
}

func (per *Permutation) Reverse(b []byte, start, end int) []byte {
	for start <= end {
		b[start], b[end] = b[end], b[start]
		start++
		end--
	}
	return b
}
