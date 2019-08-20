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

func GetConfig() *viper.Viper {
	locations := GetConfigLocations()
	config := &viper.Viper{}
	viper.SetConfigName(DefaultConfigFileName)
	for _, location := range locations {
		viper.AddConfigPath(location)
	}
	viper.SetConfigName(OverrideConfigFileName)
	for _, location := range locations {
		viper.AddConfigPath(location)
	}

	return config
}
