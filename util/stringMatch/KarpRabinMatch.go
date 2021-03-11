/*
 *  @Author : huangzj
 *  @Time : 2021/3/5 13:17
 *  @Description：
 */

package stringMatch

func KarpRabinMatch(allString, modeString string) int {
	hashMode := hash(modeString)
	for i := 0; i < len(allString)-len(modeString)+1; i++ {
		hashKey := hash(allString[i : i+len(modeString)+1])
		if hashMode == hashKey {
			for j := 0; j < len(modeString); j++ {
				if allString[i+j] != modeString[j] {
					break
				}
			}
			return i
		}
	}

	return -1
}

func hash(s string) int {
	return 1
}
