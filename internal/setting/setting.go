package setting

import "sync"

var cfgInstance *Config
var cfgOnce = sync.Once{}

type Config struct {
	File       string
	LogLv      string        `mapstructure:"log"`
	HttpServer HttpServerCfg `mapstructure:"server"`
	DB         DBCfg         `mapstructure:"db"`
}

func GetCfg() *Config {
	cfgOnce.Do(func() {
		cfgInstance = &Config{
			HttpServer: HttpServerCfg{},
			DB:         DBCfg{},
		}
	})
	return cfgInstance
}
