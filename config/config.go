package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func GetConfig(key string) interface{} {
	viper.SetConfigName("config")
	viper.AddConfigPath("config") // ระบุ path ของ config file
	viper.AutomaticEnv()          // อ่าน value จาก ENV variable
	// แปลง _ underscore ใน env เป็น . dot notation ใน viper
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// อ่าน config
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}
	return viper.Get(key)
}
