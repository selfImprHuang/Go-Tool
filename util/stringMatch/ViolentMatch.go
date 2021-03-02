/*
 *  @Author : huangzj
 *  @Time : 2021/2/25 11:11
 *  @Description：
 */

package stringMatch

func ViolentMatch(allString, match string) int {
	i := 0
	j := 0
	sLength := len(allString)
	mLenght := len(match)

	for i < sLength && j < mLenght {
		//当前字符如果匹配,继续匹配下一个字符，坐标移动
		if allString[i] == match[j] {
			j++
			i++
		} else {
			//如果字符不匹配，match从头开始，被匹配的字符串从之前匹配的下一个字符开始，坐标修改
			i = i - j + 1
			j = 0
		}
	}

	if j == len(match) {
		return i - j
	}

	return -1
}
