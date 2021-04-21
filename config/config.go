package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func New() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/opt/jenkins-tracer")
	viper.AddConfigPath("$HOME/.jenkins-tracer")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
