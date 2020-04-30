/*
 *  @Author : huangzj
 *  @Time : 2020/4/30 14:55
 *  @Description： 文件读取工具测试
 */

package file

import (
	"Go-Tool/util/file"
	"fmt"
	"testing"
)

func TestReadFileLineNum(t *testing.T) {
	fmt.Print(file.ReadFileLineNum("F:\\Go个人代码\\src\\Go-Tool\\test\\emptyDirTest\\1\\11.txt"))
}

func TestReadFileLineNumExceptEmptyLine(t *testing.T) {
	fmt.Println(file.ReadFileLineNumExceptEmptyLine("F:\\Go个人代码\\src\\Go-Tool\\test\\emptyDirTest\\1\\11.txt"))
}
