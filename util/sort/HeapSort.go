/*
 *  @Author : huangzj
 *  @Time : 2020/8/20 16:13
 *  @Description：堆结构的基本特性：
 *  i 结点的父结点下标就为(i – 1) / 2。
 *	i 节点的左右子结点下标分别为2 * i + 1和2 * i + 2。
 *  具体图解可参考：https://www.cnblogs.com/chengxiao/p/6129630.html
 */

package sort

func HeapSort(list []int) {
	//构建堆,构建堆的算法，会根据父节点去构建子节点，并且再往上构建父节点的父节点..所以这边的构建只要到数组长度/2 -1就行.
	//这样就可以涵盖它的两个子节点.
	for i := len(list)/2 - 1; i >= 0; i-- {
		makeHeap(list, i, len(list))
	}

	//构建的是最大堆，每次最大的元素都在第0个位置，所有每一次构建堆，都要把第0个位置的数置换到对应的当前的最后一个位置.
	//上面已经构建了一个完整的堆，所以这边只需要对堆进行调整即可
	for j := len(list) - 1; j > 0; j-- {
		list[j], list[0] = list[0], list[j] //将堆顶元素与末尾元素进行交换
		makeHeap(list, 0, j)                //重新对堆进行调整(这个时候是对除去了上次的最大节点的堆进行重组)
	}
}

//构建堆算法
func makeHeap(list []int, start int, end int) {
	temp := list[start] //先把当前父节点的值保存下来
	for i := 2*start + 1; i < end; i = 2*start + 1 {
		//如果左子树值比右子树小，那么要比较的就是右子树和父节点
		if i < end-1 && list[i] < list[i+1] {
			i++
		}
		//如果被比较的子节点更大，则应该放在父节点的位置.并且开始节点往后移动，比较下一个节点(当前父节点的子节点)的父子节点值.
		if list[i] > temp {
			list[start], list[i] = list[i], list[start]
			start = i
		} else {
			break //如果父节点大于两个自己点，则说明是完整的堆，则退出循环
		}
	}
}
