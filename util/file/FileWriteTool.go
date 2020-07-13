/*
 *  @Author : huangzj
 *  @Time : 2020/4/26 14:56
 *  @Description： 文件写入工具
 */

package file

import (
	"io/ioutil"
	"os"
)

/*
 * @param filePath 文件位置
 * @param content 文件内容
 * @return
 * @description
 */
func WriteByIoUtil(filePath, content string) error {
	err := ioutil.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

/*
 * @param filePath 文件位置
 * @param content 文件内容
 * @return FileError错误信息
 * @return  File 文件对象，成功写入才会返回文件对象，否则是nil
 * @description
 */
func WriteByFile(filePath, content string) (error, *os.File) {
	fileObj, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err, nil
	}
	defer fileObj.Close()

	if _, err := fileObj.WriteString(content); err != nil {
		return err, nil
	}

	return nil, fileObj
}
