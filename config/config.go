package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type MySQLConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
	User     string
}

type Config struct {
	AppName string
	MySQL   MySQLConfig
}

var C Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig() ////告诉 Viper 读取配置文件。Viper 将在之前指定的路径中查找名为 config.yml 的文件
	if err != nil {
		fmt.Println("ReadInConfig err:")
		panic(err)
	}
	err = viper.Unmarshal(&C)
	if err != nil {
		fmt.Println("Unmarshal err:")
		panic(err)
	}
}

func GetConfig() Config {
	return C
}
