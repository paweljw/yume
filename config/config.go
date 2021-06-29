package config

import (
	"github.com/pelletier/go-toml"
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
