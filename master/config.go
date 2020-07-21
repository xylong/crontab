package master

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var Conf *Config

// Config 配置
type Config struct {
	ApiPort         int `json:"api_port"`
	APiReadTimeOut  int `json:"a_pi_read_time_out"`
	APiWriteTimeOut int `json:"a_pi_write_time_out"`
}

// InitConfig 加载配置
func InitConfig(filename string) (err error) {
	var (
		content []byte
		conf    Config
	)
	if content, err = ioutil.ReadFile(filename); err != nil {
		return
	}
	if err = json.Unmarshal(content, &conf); err != nil {
		return
	}
	Conf = &conf
	fmt.Println(Conf)
	return
}
