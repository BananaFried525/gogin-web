package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func GetConfig(key string) interface{} {
	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}
	return viper.Get(key)
}

func GetPort() interface{} {
	return GetConfig("app.port")
}
