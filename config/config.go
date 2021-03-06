package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

var config *viper.Viper
var GlobalConfig *ClientConfig

type ClientConfig struct {
	Configurations map[string]SystemConfiguration `json:"configurations"`
}

type SystemConfiguration struct {
	Value map[string]SystemConfigurationValue `json:"value"`
}

type SystemConfigurationValue struct {
	Value     interface{} `json:"value"`
	TypeValue int         `json:"type_value"`
}

type SystemConfigurationValueObject struct {
	Key            interface{} `json:"key"`
	Value          interface{} `json:"value"`
	ValueDisplayUi interface{} `json:"valueDisplayUI"`
}

func GetConfigurationString(clientUuid string, key string) string {
	c, ok := GlobalConfig.Configurations[clientUuid]
	if !ok {
		return ""
	}

	v, ok := c.Value[key]
	if !ok {
		return ""
	}

	return fmt.Sprintf("%v", v.Value)
}

func GetConfigurationInt(clientUuid string, key string) int {
	c, ok := GlobalConfig.Configurations[clientUuid]
	if !ok {
		return 0
	}

	v, ok := c.Value[key]
	if !ok {
		return 0
	}

	s, err := strconv.Atoi(fmt.Sprintf("%v", v.Value))
	if err != nil {
		return 0
	}

	return s
}

func init() {
	env := os.Getenv("APP_ENV")
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("config/")
	config.AddConfigPath("../config/")
	config.AutomaticEnv()
	config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := config.ReadInConfig(); err != nil {
		log.Fatal(err.Error())
	}
}

func GetConfig() *viper.Viper {
	return config
}
