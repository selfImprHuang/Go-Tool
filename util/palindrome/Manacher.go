/*
 *  @Author : huangzj
 *  @Time : 2021/3/13 21:58
 *  @Description：
 */

package palindrome

import "fmt"

//返回字符串中最大回文的数量和回文字符串
func Manacher(allString string) (int, string) {
	maxMid := 0    //当前最大回文串中部下标
	radiusPos := 0 //最大回文半径对应的右边的下标位置
	result := generateString(allString)
	point := make([]int, len(result)) //以j为中心的最长回文半径
	maxLen := -1
	var palindrome string

	for i := 1; i < len(result); i++ {
		j := 2*maxMid - i //与i下标对称的左边的下标
		if i >= radiusPos {
			point[i] = 1
		} else if radiusPos-i > point[j] {
			point[i] = point[j]
		} else {
			// point[j] >= radiusPos-i
			point[i] = radiusPos - i
		}
		//从后面的位置继续开始找对应的字符
		for (i-point[i] >= 0 && i+point[i] < len(result)) && result[i+point[i]] == result[i-point[i]] {
			point[i]++
		}
		//更新最长回文串信息
		if i+point[i] > radiusPos {
			maxMid = i
			radiusPos = i + point[i]
		}
		if maxLen < point[i]-1 && point[i]-1 > 1 {
			maxLen = point[i] - 1
			palindrome = allString[maxMid/2-maxLen/2 : radiusPos/2]
		}
	}

	return maxLen, palindrome
}

//把字符串转换成特殊的字符串
func generateString(allString string) string {
	var result string
	for i := range allString {
		result = fmt.Sprintf("%s%s%c", result, "#", allString[i])
	}
	return fmt.Sprintf("%s%s", result, "#")
}
