package setting

import (
	"log"
)
import "github.com/Unknwon/goconfig"

var Config *GlobalConfig
var Redis *RedisConfig
var configfile *goconfig.ConfigFile
var err error

func GetConfigIni(filepath string) error {

	configfile, err = goconfig.LoadConfigFile(filepath)
	if err != nil {
		log.Println("配置文件读取错误,找不到配置文件", err.Error())
		return err
	}

	Config = &GlobalConfig{}
	return nil
}
func GetRedisConfig() error {
	Redis = &RedisConfig{}
	r, err := configfile.GetSection("redis")
	if err != nil {
		return err
	}
	for key, val := range r {
		if key == "ip" {
			Redis.Ip = val
		}
		if key == "port" {
			Redis.Port = val
		}
	}
	return nil
}
