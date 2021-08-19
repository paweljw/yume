package config

import (
	"github.com/pelletier/go-toml"
	"strings"
)

var Config *toml.Tree

func LoadConfiguration() {
	var err error

	Config, err = toml.LoadFile("resources/config.toml")

	if err != nil {
		panic(err)
	}
}

func GetMessage(msg string) string {
	return Config.Get("messages." + msg).(string)
}

func IsBadName(name string) bool {
	badNames := Config.Get("config.bad_names").([]interface{})

	for _, badName := range badNames {
		if strings.Contains(name, badName.(string)) {
			return true
		}
	}

	return false
}
