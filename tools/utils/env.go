package utils

import "log"
import "github.com/spf13/viper"

func NewEnv(e interface{}) {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&e)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}
}
