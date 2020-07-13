/*
 *  @Author : huangzj
 *  @Time : 2020/4/26 14:55
 *  @Description： 文件读取工具
 */

package file

import (
	"bufio"
	"encoding/json"
	"errors"
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
func ReadInOnce(filePath string) (string, error) {
	err, is, _ := IsFileExist(filePath)
	if err != nil {
		return "", err
	}
	if !is {
		return "", errors.New("非文件")
	}

	fd, iErr := ioutil.ReadFile(filePath)
	if iErr != nil {
		return "", iErr
	}

	return string(fd), nil
}

/*
 * @param filePath 文件地址
 * @return string 读取的文件内容
 * @return FileError 读取错误，如果没有错误返回nil
 * @description 通过for循环去读取文件，这边提供一种读取方式，通过这种方式来处理对应读取的操作
 */
func ReadByCircle(filePath string) (string, error) {
	var fs string
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
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

func ReadFileLineNum(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
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

func ReadFileLineNumExceptEmptyLine(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
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

//读取json文件内容
//这边发现就是说如果按照正常的文件读取方式来读取json文件，读取到的内容是乱码的
//所以这边通过加解码的方式来读取，返回一个json串出去
//外部的调用者通过解析转换成对应的结构体对象
//@param filePath 文件路径
func ReadJsonFile(filePath string) (string, error) {
	filePtr, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer filePtr.Close()

	decoder := json.NewDecoder(filePtr)
	sMap := make(map[string]interface{}, 0)
	err = decoder.Decode(&sMap)
	if err != nil {
		return "", err
	}
	bs, err := json.Marshal(sMap)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}
