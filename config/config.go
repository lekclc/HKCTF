package cfg

import (
	"fmt"

	"github.com/docker/docker/client"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Docker *client.Client

func Init() {
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func Get(key string) any {
	return viper.Get(key)
}
