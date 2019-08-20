package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

const DefaultConfigFileName = "default"
const OverrideConfigFileName = "local"
func GetConfigLocations() [2]string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	locations := [...]string{"%s", "%s/configs"}
	for i, template := range locations {
		locations[i] = fmt.Sprintf(template, dir)
	}
	return locations
}

func GetConfig() Config {
	locations := GetConfigLocations()
	runtimeViper := viper.New()
	runtimeViper.SetConfigName(DefaultConfigFileName)
	for _, location := range locations {
		runtimeViper.AddConfigPath(location)
	}
	defaultErr := runtimeViper.ReadInConfig()
	if defaultErr != nil {
		if _, ok := defaultErr.(viper.ConfigFileNotFoundError); !ok {
			panic(defaultErr)
		}
	}
	runtimeViper.SetConfigName(OverrideConfigFileName)
	for _, location := range locations {
		runtimeViper.AddConfigPath(location)
	}
	err := runtimeViper.MergeInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok || defaultErr != nil {
			panic(err)
		}
	}
	config, err := newConfig(runtimeViper)
	if err != nil {
		panic(err)
	}

	return config
}
