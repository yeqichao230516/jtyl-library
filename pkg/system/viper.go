package system

import (
	"github.com/spf13/viper"
)

func Viper() (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return v, nil
}
