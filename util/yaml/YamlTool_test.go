/*
 *  @Author : huangzj
 *  @Time : 2020/12/16 11:56
 *  @Description：
 */

package yaml

import (
	"fmt"
	"io/ioutil"
	"testing"
)

type MySQLConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DbName   string `yaml:"dbname"`
}
type RedisConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Auth string `yaml:"-"` //标签表示不显示
}
type NginxProxyConfig struct {
	Counter   int      `yaml:"counter"`
	NginxList []string `yaml:"nginx_list"`
}

type ServiceYaml struct {
	MySQL      MySQLConfig      `yaml:"mysql"`
	Redis      RedisConfig      `yaml:"redis"`
	NginxProxy NginxProxyConfig `yaml:"nginx_proxy"`
}

func TestYamlToStructByFile(t *testing.T) {
	config := new(ServiceYaml)
	_ = GetYamlToStructByFile("config.yaml", &config)

	fmt.Printf("MySQL host : %s, port : %d, user : %s, password : %s, db_name : %s \n",
		config.MySQL.Host, config.MySQL.Port, config.MySQL.User, config.MySQL.Password, config.MySQL.DbName)

	fmt.Printf("Redis host : %s, port %d, auth : %s\n",
		config.Redis.Host, config.Redis.Port, config.Redis.Auth)

	fmt.Printf("Vip Counter: %d, Vip List : %v\n", config.NginxProxy.Counter, config.NginxProxy.NginxList)
}

func TestYamlToStructByString(t *testing.T) {
	config := new(ServiceYaml)
	fileByte, _ := ioutil.ReadFile("config.yaml")
	_ = GetYamlToStructByString(string(fileByte), &config)

	fmt.Printf("MySQL host : %s, port : %d, user : %s, password : %s, db_name : %s \n",
		config.MySQL.Host, config.MySQL.Port, config.MySQL.User, config.MySQL.Password, config.MySQL.DbName)

	fmt.Printf("Redis host : %s, port %d, auth : %s\n",
		config.Redis.Host, config.Redis.Port, config.Redis.Auth)

	fmt.Printf("Vip Counter: %d, Vip List : %v\n", config.NginxProxy.Counter, config.NginxProxy.NginxList)
}

func TestYamlToMap(t *testing.T) {
	fileByte, _ := ioutil.ReadFile("config.yaml")
	_, m := GetYamlToMap(string(fileByte))
	for key, value := range m {
		fmt.Println(fmt.Sprintf("key:%v ,value: %v", key, value))
	}
}

func TestStructToYaml(t *testing.T) {
	config := new(ServiceYaml)
	fileByte, _ := ioutil.ReadFile("config.yaml")
	_ = GetYamlToStructByString(string(fileByte), &config)
	_, result := GetStructToYaml(config)
	fmt.Println(result)
}
