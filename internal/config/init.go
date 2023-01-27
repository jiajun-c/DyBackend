package config

import "github.com/spf13/viper"

func Init(filename string) {
	viper.SetConfigName(filename)

	viper.AddConfigPath("../../config/app/")

	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
