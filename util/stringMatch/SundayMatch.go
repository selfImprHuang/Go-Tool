/*
 *  @Author : huangzj
 *  @Time : 2021/3/2 11:21
 *  @Description：
 */

package stringMatch

func ShiftTable(modeString string) []int {
	shiftTable := make([]int, 256)

	//一开始设置所有的字符匹配不上向右移动大小为：模式字符串长度 + 1（因为要移动到下一个字符的后一位）
	for i := range shiftTable {
		shiftTable[i] = len(modeString) + 1
	}

	//设置模式字符串中每一个字符的移动长度.
	for i := range modeString {
		shiftTable[modeString[i]-'a'] = len(modeString) - i
	}

	return shiftTable
}

func SundayMatch(allString, modeString string) int {
	table := ShiftTable(modeString)
	aLength := len(allString)
	mLength := len(modeString)
	i := 0 //这个参数记录的是移动过程中，，模式字符串对应原始字符串的第一个位置.

	for i < aLength-mLength {
		j := 0
		for allString[i+j] == modeString[j] {
			j++
			if j >= len(modeString) {
				return i
			}
		}
		//找到移动后模式字符串对应原始字符串的第一个位置
		move := allString[i+mLength]
		i = i + table[move-'a']
	}

	return -1
}
