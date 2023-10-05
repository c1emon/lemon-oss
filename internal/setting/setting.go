package setting

import "sync"

var cfgInstance *Config
var cfgOnce = sync.Once{}

type Config struct {
	File  string
	LogLv string
	Http  HttpCfg
	DB    DBCfg
}

func GetCfg() *Config {
	cfgOnce.Do(func() {
		cfgInstance = &Config{
			Http: HttpCfg{},
			DB:   DBCfg{},
		}
	})
	return cfgInstance
}
