/*
 *  @Author : huangzj
 *  @Time : 2020/4/30 16:06
 *  @Description：数组工具类测试
 */

package array

import (
	"fmt"
	"testing"
)

type MM struct {
	a int
	b string
	n NN
}

type NN struct {
	a int
	b float64
}

func TestArray(t *testing.T) {
	n := NN{
		a: 12,
		b: 2.22,
	}
	n1 := NN{
		a: 12,
		b: 2.22,
	}
	n2 := NN{
		a: 12,
		b: 2.212,
	}
	x := MM{
		a: 1,
		b: "1",
		n: n,
	}
	y := MM{
		a: 1,
		b: "1",
		n: n1,
	}
	z := MM{
		a: 1,
		b: "1",
		n: n2,
	}

	fmt.Print(x == y)
	fmt.Print("\n\n测试一下")
	fmt.Print(x == z)
}
