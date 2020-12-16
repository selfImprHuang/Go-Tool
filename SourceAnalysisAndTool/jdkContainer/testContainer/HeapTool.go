/*
 *  @Author : huangzj
 *  @Time : 2020/12/2 15:39
 *  @Description：
 */

package testContainer

type Person struct {
	Name  string  //名字
	Age   int     //年龄
	Money float64 //身价
}

type heapTool []*Person

func (h *heapTool) Less(i, j int) bool {
	return (*h)[i].Age < (*h)[j].Age
}

func (h *heapTool) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *heapTool) Len() int {
	return len(*h)
}

func (h *heapTool) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *heapTool) Push(v interface{}) {
	*h = append(*h, v.(*Person))
}
