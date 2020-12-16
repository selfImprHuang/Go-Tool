/*
 *  @Author : huangzj
 *  @Time : 2020/12/16 11:56
 *  @Description：
 */

package yaml

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

//把文件中的yaml结构转换成结构体
func GetYamlToStructByFile(fileName string, object interface{}) error {
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, object)
	if err != nil {
		return err
	}

	return nil
}

//通过字符串的Yaml转换成结构体
func GetYamlToStructByString(content string, object interface{}) error {
	err := yaml.Unmarshal([]byte(content), object)
	if err != nil {
		return err
	}

	return nil
}

//通过字符串的Yaml转换成map
func GetYamlToMap(content string) (error, map[interface{}]interface{}) {
	m := make(map[interface{}]interface{}, 0)
	err := yaml.Unmarshal([]byte(content), m)
	if err != nil {
		return err, m
	}

	return nil, m
}

//把结构体转换成yaml的字符串
func GetStructToYaml(object interface{}) (error, string) {
	result, err := yaml.Marshal(object)
	if err != nil {
		return err, ""
	}
	return nil, string(result)
}
