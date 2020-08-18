/*
 *  @Author : huangzj
 *  @Time : 2020/4/30 14:56
 *  @Description： 文件工具测试
 */

package file

import (
	"fmt"
	"testing"
)

var (
	filePaths = [5]string{"F:\\Go个人代码\\src\\Go-Tool\\test\\emptyDirTest\\1\\11.txt",
		"F:\\Go个人代码\\src\\Go-Tool\\test\\emptyDirTest\\2\\1\\11.txt",
		"F:\\Go个人代码\\src\\Go-Tool\\test\\emptyDirTest\\2\\2222\\4\\4\\111.txt",
		"F:\\Go个人代码\\src\\Go-Tool\\test\\emptyDirTest\\2\\2222\\4\\4\\4\\4\\1211.txt",
		"F:\\Go个人代码\\src\\Go-Tool\\test\\emptyDirTest\\3\\11.txt"}
)

func TestGetAllFileNameFromDir(t *testing.T) {
	s := GetAllFileNameFromDir("F:\\Go个人代码\\src\\Go-Tool\\test\\emptyDirTest")

	for _, row := range s {
		fmt.Print(row)
	}
}

func TestGetAllEmptyDir(t *testing.T) {
	_, s := GetAllEmptyDir("F:\\Go个人代码\\src\\Go-Tool\\test\\emptyDirTest")

	for _, row := range s {
		fmt.Print(fmt.Sprint(row, "\n\n"))
	}
}

func TestGetAllDir(t *testing.T) {
	_, s := GetAllDir("F:\\Go个人代码\\src\\Go-Tool\\test\\emptyDirTest")

	for _, row := range s {
		fmt.Print(row)
	}
}

func TestGetAllNotEmptyDir(t *testing.T) {
	_, s := GetAllNotEmptyDir("F:\\Go个人代码\\src\\Go-Tool\\test\\emptyDirTest")

	for _, row := range s {
		fmt.Print(row)
	}
}
