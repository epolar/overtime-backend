package config

import (
	"github.com/spf13/viper"
	"sync"
)

var config *cfg
var configOnce sync.Once

func Config() *cfg {
	configOnce.Do(func() {
		config = &cfg{}
	})
	return config
}

type cfg struct {
}

type CorsConfig struct {
	Origins []string `mapstructure:"origins"`
	Methods []string `mapstructure:"methods"`
	Headers []string `mapstructure:"headers"`
}

func (c *cfg) CorsConfig() *CorsConfig {
	var cors CorsConfig
	if err := viper.UnmarshalKey("cors", &cors); err != nil {
		panic(err)
	}
	return &cors
}
