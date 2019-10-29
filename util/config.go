package util

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

type Conf struct {
	Data map[interface{}]interface{}
}

var Config *Conf

//初始化参数，提供默认参数
func init() {

	yamlFile, err := ioutil.ReadFile("./conf/conf.yml")
	if err != nil {
		fmt.Println(err.Error())
	}
	Config = &Conf{}
	err = yaml.Unmarshal(yamlFile, &Config.Data)
	if err != nil {
		fmt.Println("配置文件加载失败", err.Error())
	}
}

func (c *Conf) Get(name string) interface{} {
	path := strings.Split(name, ".")
	data := c.Data
	for key, value := range path {
		v, ok := data[value]
		if !ok {
			break
		}
		if (key + 1) == len(path) {
			return v
		}
		//判断下级 是否是一个map，继续往下一层搜索
		if reflect.TypeOf(v).String() == "map[interface {}]interface {}" {
			data = v.(map[interface{}]interface{})
		}
	}
	return nil
}

// 从配置文件中获取string类型的值
func (c *Conf) GetString(param string) string {
	value := c.Get(param)
	switch value := value.(type) {
	case string:
		return value
	case bool, float64, int:
		return fmt.Sprint(value)
	default:
		return ""
	}
}

// 从配置文件中获取int类型的值
func (c *Conf) GetInt(param string) int {
	value := c.Get(param)
	switch value := value.(type) {
	case string:
		i, _ := strconv.Atoi(value)
		return i
	case int:
		return value
	case bool:
		if value {
			return 1
		}
		return 0
	case float64:
		return int(value)
	default:
		return 0
	}
}
