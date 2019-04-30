package util

import "github.com/spf13/viper"

func ReadKey(key string, obj interface{}) {
	err := viper.UnmarshalKey(key, &obj)
	if err != nil {
		panic("unable to decode key into struct.")
	}
}