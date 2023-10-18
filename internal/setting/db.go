package setting

type DBCfg struct {
	Driver string `mapstructure:"driver"`
	Source string `mapstructure:"source"`
}
