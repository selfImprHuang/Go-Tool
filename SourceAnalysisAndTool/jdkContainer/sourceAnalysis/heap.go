// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package heap provides heap operations for any type that implements
// heap.Interface. A heap is a tree with the property that each node is the
// minimum-valued node in its subtree.
//
// The minimum element in the tree is the root, at index 0.
//
// A heap is a common way to implement a priority queue. To build a priority
// queue, implement the Heap interface with the (negative) priority as the
// ordering for the Less method, so Push adds items while Pop removes the
// highest-priority item from the queue. The Examples include such an
// implementation; the file example_pq_test.go has the complete source.
//
//read note 拷贝一个代码过来解析.感觉在go里面拷贝代码简单多了，依赖可以直接使用。我在这边使用read note(自定义todo标识)来标识我的解析.
package sourceAnalysis

import "sort"

// The Interface type describes the requirements
// for a type using the routines in this package.
// Any type that implements it may be used as a
// min-heap with the following invariants (established after
// Init has been called or if the data is empty or sorted):
//
//	!h.Less(j, i) for 0 <= i < h.Len() and 2*i+1 <= j <= 2*i+2 and j < h.Len()
//
// Note that Push and Pop in this interface are for package heap's
// implementation to call. To add and remove things from the heap,
// use heap.Push and heap.Pop.
type Interface interface {
	sort.Interface
	Push(x interface{}) // add x as element Len()
	Pop() interface{}   // remove and return element Len() - 1.
}

// Init establishes the heap invariants required by the other routines in this package.
// Init is idempotent with respect to the heap invariants
// and may be called whenever the heap invariants may have been invalidated.
// The complexity is O(n) where n = h.Len().
//read note 对整个Interface 进行重构，时间复杂度是 O(n)
func Init(h Interface) {
	// heapify
	n := h.Len()
	//read note 从最小父节点，到第0位的根节点，分别向下进行重构处理(之所以要用循环是因为一次向下的重构，只能对一条连续的分支进行重构，不彻底.)
	for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n)
	}
}

// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
// read note 这个Push方法是往 Interface里面去新增一个元素.和我们继承的Push方法有差别.这边同时候做了重构操作.
//  一次Push的时间复杂度是 O(logN),一次Init的时间复杂度是O(N),按道理Push的元素越多，Init的效率越高
func Push(h Interface, x interface{}) {
	//read note 先把元素添加进去
	h.Push(x)
	//read note 然后在从下往上进行重构处理，正常情况下Push都是把元素添加在最后一个位置，如果实现的方法，把Push添加到其他位置如果不进行Init，应该是会有问题.
	up(h, h.Len()-1)
}

// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
// Pop is equivalent to Remove(h, 0).
//read note 把最小的元素输出，需要注意的是真正的Pop必须是调用这个方法，而不是我们继承的那个方法.
// 时间复杂度是O(logN)
func Pop(h Interface) interface{} {
	//read note 真正的Pop输出的是最小的元素，也就是最小堆的第0位置的元素
	// 所以这边的操作是把第0位的数组放到最后一位，然后从位置0开始，到N-1的位置，对所有元素进行down(父子节点比较交换)的操作
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop()
}

// Remove removes and returns the element at index i from the heap.
// The complexity is O(log n) where n = h.Len().
//read note 移除某个位置的元素，时间复杂度是 o(logn)
func Remove(h Interface, i int) interface{} {
	//read note 判断移除的下标不等于最后一个元素位置，要特殊处理
	n := h.Len() - 1
	if n != i {
		//read note 交换第i个元素和最后一个元素
		h.Swap(i, n)
		//read note Fix的处理操作，这边只会处理到移除一个元素后的位置，也就是说被移除的那个元素(在最后的位置)不会参与重构数组的操作
		if !down(h, i, n) {
			up(h, i)
		}
	}
	//read note 调用Pop，把最后一个元素返回回去
	return h.Pop()
}

// Fix re-establishes the heap ordering after the element at index i has changed its value.
// Changing the value of the element at index i and then calling Fix is equivalent to,
// but less expensive than, calling Remove(h, i) followed by a Push of the new value.
// The complexity is O(log n) where n = h.Len().
//read note 当下标为i的元素发生改变，需要进行一次Fix的处理,时间复杂度是 o(logn)
func Fix(h Interface, i int) {
	//read note 想从下标i的这个元素往下查找处理，如果往下没有交换元素，再往上进行处理.
	if !down(h, i, h.Len()) {
		up(h, i)
	}
}

//read note  h:  对应的数组数据
//read note  j:子节点的下标
//read note  方法作用
func up(h Interface, j int) {
	for {
		//read note 拿到对应的父节点的下标
		i := (j - 1) / 2 // parent
		//read note 找到最后一个父节点 || 父节点比子节点小，则跳出循环
		if i == j || !h.Less(j, i) {
			break
		}
		//read note 否则就是父节点比子节点的值大，需要交换对应的元素，然后找到父节点的下标，再往上找其父节点的关系
		h.Swap(i, j)
		j = i
	}
}

//read note  h:  对应的数组数据
//read note  i0: 需要下发处理的坐标，这边也就是左右子节点的父节点下标.
//read note  n:  数组对应的总长度
//read note  方法作用：从父节点开始，循环向下判断对应的元素是否在对应的环境上
func down(h Interface, i0, n int) bool {
	//read note 把父节点的下标拿出来
	i := i0

	//read note 循环的结束条件是：
	for {
		//read note 找到左边子节点下标
		j1 := 2*i + 1
		//read note 退出条件1：超过最大长度或者是负值(这个应该是针对传入就有问题的处理)
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		//read note 拿到左孩子和右孩子中比较小的那个元素（的下标）
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}

		//read note 判断父节点和子节点的大小关系，小的元素应该在父节点，所以如果子节点本身就比较小，直接退出循环，否则交换元素
		//read note 然后再把下标移动到被交换的这个元素上，计算被交换的这个元素和它的左右子节点的大小关系，进入下一个循环
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
	//read note 判断是否发生了交换操作
	return i > i0
}
