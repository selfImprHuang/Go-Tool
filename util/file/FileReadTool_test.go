/*
 *  @Author : huangzj
 *  @Time : 2020/4/30 14:55
 *  @Description： 文件读取工具测试
 */

package file

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestReadFileLineNum(t *testing.T) {
	fmt.Print(ReadFileLineNum("F:\\Go个人代码\\src\\Go-Tool\\test\\emptyDirTest\\1\\11.txt"))
}

func TestReadFileLineNumExceptEmptyLine(t *testing.T) {
	fmt.Println(ReadFileLineNumExceptEmptyLine("F:\\Go个人代码\\src\\Go-Tool\\test\\emptyDirTest\\1\\11.txt"))
}

func TestReadJsonFile(t *testing.T) {
	content, _ := ReadJsonFile("F:\\Go个人代码\\src\\Go-Tool\\util\\file\\json_file_test_mode.json")
	fmt.Println(content)
	var testMode JsonTestMode
	_ = json.Unmarshal([]byte(content), &testMode)
	for _, r := range testMode.RareNumAward {
		fmt.Println(r.Num)
		fmt.Println()
		for _, r1 := range r.Award {
			fmt.Println(fmt.Sprintf("%d    %d", r1[0], r1[1]))
		}
	}
}

//被读取Json文件的测试结构体
type JsonTestMode struct {
	CandleLightingReward CandleLightingReward

	Cake []Cake

	RareNumAward []RareNumAward

	CommonNumAward [][]int

	DailyAwardPool [][]int

	DailyExtraction []DailyExtraction

	FinalAward [][]int

	CandleNumStart int

	CandleNumInterval int

	DiamondNum int

	ThemeId int
}

type CandleLightingReward struct {
	Num       int
	AwardPool [][]int
}
type Cake struct {
	Level     int
	CandleNum int
	Award     [][]int
}
type RareNumAward struct {
	Num   int
	Award [][]int
}
type DailyExtraction struct {
	Day int
	Num int
}
