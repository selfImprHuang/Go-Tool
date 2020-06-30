/*
 *  @Author : huangzj
 *  @Time : 2020/4/30 16:06
 *  @Description：数组工具类测试
 */

package array

import (
	"Go-Tool/util/array"
	"encoding/json"
	"fmt"
	"testing"
)

type MM struct {
	A int
	B string
	N NN
}

type NN struct {
	A int
	B float64
}

var (
	mm  []interface{}
	mm1 []interface{}
	mm2 []interface{}
	nn  []interface{}
	nn1 []interface{}
)

func TestBool(t *testing.T) {
	fmt.Print("\n\nContains\n\n")
	fmt.Println(array.Contains(mm, mm))
	fmt.Println(array.Contains(mm, mm1))
	fmt.Println(array.Contains(mm, MM{
		A: 1,
		B: "2",
	}))
	fmt.Print("\n\nNotContains\n\n")
	fmt.Println(array.NotContains(mm, mm))
	fmt.Println(array.NotContains(mm, mm1))
	fmt.Println(array.NotContains(mm, MM{
		A: 1,
		B: "2",
	}))
	fmt.Print("\n\nIsSameList\n\n")
	fmt.Println(array.IsSameList(mm, mm))
	fmt.Println(array.IsSameList(mm, nn1))
	fmt.Println(array.IsSameList(mm, mm2))
	fmt.Println(array.IsSameList(mm, nn))

	fmt.Print("\n\nIsSubSet\n\n")
	fmt.Println(array.IsSubSet(mm, mm))
	fmt.Println(array.IsSubSet(mm, mm1))
	fmt.Println(array.IsSubSet(mm, mm2))
	fmt.Println(array.IsSubSet(mm, nn))
	fmt.Printf("\n\n")

}

func TestJoin(t *testing.T) {
	fmt.Print("\n\nmm-nn\n\n")
	intersection(mm, nn)
	fmt.Print("\n\nmm-mm\n\n")
	intersection(mm, mm)
	fmt.Print("\n\nmm-mm1\n\n")
	intersection(mm, mm1)
	fmt.Print("\n\nmm-mm2\n\n")
	intersection(mm, mm2)

	fmt.Print("\n\ndiff\n\n")
	diffSet(mm, nn)
	fmt.Print("\n\n\n\n")
	diffSet(mm, mm)
	fmt.Print("\n\n\n\n")
	diffSet(mm, mm1)
	fmt.Print("\n\n\n\n")
	diffSet(mm, mm2)

	fmt.Print("\n\nunion\n\n")
	union(mm, nn)
	fmt.Print("\n\n\n\n")
	union(mm, mm)
	fmt.Print("\n\n\n\n")
	union(mm, mm1)
	fmt.Print("\n\n\n\n")
	union(mm, mm2)
}

func TestEqual(t *testing.T) {
	DeepCopyIntSlice()
	DeepCopyIntSlice2()
}

func intersection(mm, nn []interface{}) {
	i1 := array.Intersection(mm, nn)
	for _, row := range i1 {
		b, err := json.Marshal(row)
		fmt.Println(err)
		fmt.Print(string(b))
	}
}

func diffSet(mm, nn []interface{}) {
	i1 := array.DiffSet(mm, nn)
	for _, row := range i1 {
		b, _ := json.Marshal(row)
		fmt.Print(string(b))
	}
}

func union(mm, nn []interface{}) {
	i1 := array.Union(mm, nn)
	for _, row := range i1 {
		b, _ := json.Marshal(row)
		fmt.Print(string(b))
	}
}

func DeepCopyIntSlice() {
	src := []int{1, 2, 3, 4, 5, 6}
	src1 := array.DeepCopyIntSlice(src)
	fmt.Println(&src == &src1)
}

func DeepCopyIntSlice2() {
	src := [][]int{{1, 2, 3}, {4, 5, 6}}
	src1 := array.DeepCopyIntSlice2(src)
	fmt.Println(&src == &src1)
}

func init() {
	mm = make([]interface{}, 0)
	mm1 = make([]interface{}, 0)
	nn = make([]interface{}, 0)
	nn1 = make([]interface{}, 0)
	mm = append(mm, MM{
		A: 1,
		B: "1",
	})
	mm = append(mm, MM{
		A: 1,
		B: "2",
	})
	mm1 = append(mm1, MM{
		A: 1,
		B: "2",
	})
	mm2 = append(mm2, MM{
		A: 1,
		B: "2",
	})
	mm2 = append(mm2, MM{
		A: 1,
		B: "3",
	})

	nn = append(nn, NN{
		B: 1.2,
		A: 1,
	})
	nn = append(nn, NN{
		B: 2.9,
		A: 20,
	})

	nn1 = append(nn1, NN{
		A: 30,
		B: 2.2,
	})
}
