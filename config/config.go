package config

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Name string `mapstructure:"name"`
	Host struct {
		FccRpcHost string `mapstructure:"fcc_rpc_host"`
		StarlingRpcHost string `mapstructure:"starling_rpc_host"`
	} `mapstructure:"host"`
}

var Global *Config

func InitConfFile() {
	configName := os.Getenv("conf_file")
	if configName == "" {
		configName = "config"
	}

	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")
	viper.SetConfigName(configName)

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&Global)
	if err != nil {
		panic(err)
	}
	Global.Name = configName
}

func InitConfig()  {
	InitConfFile()
	InitGRPClient()
}