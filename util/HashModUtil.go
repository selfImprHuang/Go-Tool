//通过hash算法获取字符串对应的哈希值
package util

import (
	"bytes"
	"hash/crc32"
	"math/big"
)

func bigInt(s string) (*big.Int, bool) {
	bInt := new(big.Int)
	return bInt.SetString(s, 10)
}

func UseBigIntMod(s string, mod int) int64 {
	b, br := bigInt(s)
	if !br {
		panic("该方法只可用于大数mod计算")
	}

	result := b.Mod(b, new(big.Int).SetInt64(int64(mod)))
	return result.Int64()
}

func calBytesSum(s string) int {
	var sum int
	for _, b1 := range bytes.NewBufferString(s).Bytes() {
		sum = sum + int(b1)
	}

	return sum
}

func Mod(s string, mod int) int {
	sum := calBytesSum(s)
	return sum % mod
}

func crc32Change(s string) uint32 {
	return crc32.ChecksumIEEE([]byte(s))
}

func Crc32Mode(s string, mod int) int {
	u := crc32Change(s)
	return int(u) % mod
}
