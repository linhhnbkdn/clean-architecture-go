package config

import (
	"github.com/kelseyhightower/envconfig"
)

type envConfig struct {
	APP_ENV string `env:"AppEnv" required:"true"`
}

var CfgEnv envConfig

func LoadCfgEnv() {
	envconfig.MustProcess("", &CfgEnv)
}
