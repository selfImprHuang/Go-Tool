/*
 *  @Author : huangzj
 *  @Time : 2021/3/22 11:09
 *  @Description：
 */

package FullPermutation

type Permutation struct {
	bytes [][]byte
}

func NewPermutation() *Permutation {
	return &Permutation{
		bytes: [][]byte{},
	}
}

func (per *Permutation) Recursion(b []byte, start int) {
	//递归交换到最后一个位置结束
	if start == len(b)-1 {
		bCopy := make([]byte, len(b))
		copy(bCopy, b)
		per.bytes = append(per.bytes, bCopy)
	}

	for i := start; i < len(b); i++ {
		//交换固定位置
		b[i], b[start] = b[start], b[i]
		//固定某一个数字位置，递归后面的其他数字
		per.Recursion(b, start+1)
		//还原固定位置，进行下一次操作
		b[i], b[start] = b[start], b[i]
	}
}

func (per *Permutation) Reset() {
	per.bytes = [][]byte{}
}
