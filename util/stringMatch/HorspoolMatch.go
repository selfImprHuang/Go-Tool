/*
 *  @Author : huangzj
 *  @Time : 2020/8/21 16:06
 *  @Description：参考地址：http://www.ifcoding.com/archives/247.html
 *  该算法相当于先构造出对应的字符的移动位置，这边假设只匹配字母，则构造26个字母的移动表
 *  从匹配的字符串的最后一位开始跟被匹配字符串进行匹配，直到没有匹配的字段或者匹配完整之后结束
 *  得到当前匹配不正确的那个位置的字母，以此来找到该字母对应的移动位置，然后把匹配的字符串向后移动对应的位置，使得该字母和后面的字母对上，然后再进行一次匹配，直到找到对应的字符串位置为止，或者没有匹配的字符串返回-1
 *  举例 abcdefg 和  bcd进行匹配
 *    第一次：    abcdefg
 *                 d
 *    发现不匹配并且不匹配的字母是c，所以按照c的移动规则，向后移动一位
 *    第二次：    abcdefg
 *                 cd
 *
 *               abcdefg
 *                bcd
 *   关于另一种更有效的模式匹配算法  Boyer-Moore这边不给出算法实现（比较复杂），可参考：https://www.cnblogs.com/en-heng/p/5095542.html
 */

package stringMatch

func Horspool(allString string, match string) int {
	//生成对应的移动表数据
	move := shiTable(match)
	i := len(match) - 1
	for i < len(allString) {
		k := 0
		for k < len(match) && match[len(match)-1-k] == allString[i-k] {
			k++
		}
		if k == len(match) {
			return i - len(match) + 1
		} else {
			i = i + move[allString[i]-'a']
		}
	}

	return -1
}

//生成对应的移动表数据
func shiTable(s string) []int {
	table := make([]int, 27) //这边只考虑26个字母和一个空字符串的情况
	//所有格子都初始化移动步数为被查找字符串长度.
	for i := 0; i < 27; i++ {
		table[i] = len(s)
	}
	//构造对应的被查找字符串中的字符的移动步数
	for i := 0; i < len(s)-1; i++ {
		table[s[i]-'a'] = len(s) - 1 - i
	}
	return table
}
