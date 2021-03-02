/*
 *  @Author : huangzj
 *  @Time : 2021/2/24 13:19
 *  @Description：KMP字符串匹配算法，参考地址：https://www.cnblogs.com/en-heng/p/5091365.html
 *  http://www.ruanyifeng.com/blog/2013/05/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm.html
 */

package stringMatch

func KmpMatch(allString string, match string) int {
	next := getNext(match)
	i := 0 //被匹配的字符串下标
	j := 0 //模式匹配字符串下标
	sLength := len(allString) - 1
	mLength := len(match) - 1
	for i <= sLength && j <= mLength {
		if j == -1 || allString[i] == match[j] {
			j++
			i++
		} else {
			j = next[j]
		}
	}

	if j == len(match) {
		return i - j
	}
	return -1
}

//失配表函数构建（部分匹配函数）
func getNext(match string) []int {
	next := make([]int, len(match))
	next[0] = -1
	i := 0
	j := -1

	mLength := len(match) - 1

	for i < mLength {
		//处理两种情况： 1.数组第一个数是-1 2. match对应位置字符相等的情况
		if j == -1 || match[i] == match[j] {
			//如果这里是数组第一个数的时候,下标要往后移动
			//如果这里是数组相等的情况，要比较下一个相应的字符，所以下标需要往后移动，而next需要在上一个的基础上去做+1操作
			i++
			j++
			next[i] = j
		} else {
			//如果匹配过程中，出现对应不上的情况，也就是说 match[i] != match[j],
			j = next[j]
		}
	}

	return next
}
