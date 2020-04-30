/*
 *  @Author : huangzj
 *  @Time : 2020/4/26 14:55
 *  @Description： 文件读取工具
 */

package file

import (
	err2 "Go-Tool/err"
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

/*
 * @param 文件地址
 * @return string 文件内容
 * @return FileError 报错信息
 * @description 一次直接读取文件中的内容不进行中间处理
 */
func ReadInOnce(filePath string) (string, *err2.FileError) {
	err, is, _ := IsFileExist(filePath)
	if err != nil {
		return "", err2.ENewFileError(err)
	}
	if !is {
		return "", err2.NewFileError("非文件")
	}

	fd, iErr := ioutil.ReadFile(filePath)
	if iErr != nil {
		return "", err2.ENewFileError(iErr)
	}

	return string(fd), nil
}

/*
 * @param filePath 文件地址
 * @return string 读取的文件内容
 * @return FileError 读取错误，如果没有错误返回nil
 * @description 通过for循环去读取文件，这边提供一种读取方式，通过这种方式来处理对应读取的操作
 */
func ReadByCircle(filePath string) (string, *err2.FileError) {
	var fs string
	file, err := os.Open(filePath)
	if err != nil {
		return "", err2.ENewFileError(err)
	}
	defer file.Close()

	inputReader := bufio.NewReader(file)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			if inputString != "" {
				fs = fmt.Sprint(fs, inputString)
			}
			break
		}
		fs = fmt.Sprint(fs, inputString)
	}

	return fs, nil
}

func ReadFileLineNum(filePath string) (int, *err2.FileError) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err2.ENewFileError(err)
	}
	defer file.Close()

	num := 0
	inputReader := bufio.NewReader(file)
	for {
		_, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			num++
			break
		}
		num++
	}
	return num, nil
}

func ReadFileLineNumExceptEmptyLine(filePath string) (int, *err2.FileError) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err2.ENewFileError(err)
	}
	defer file.Close()

	num := 0
	inputReader := bufio.NewReader(file)
	for {
		input, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			if input != "\r\n" && input != "" {
				num++
			}
			break
		}
		if input != "\r\n" && input != "" {
			num++
		}
	}
	return num, nil
}
