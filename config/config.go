package config

type Config struct {
	AcquTime uint32
}

func NewConfig() *Config {
	return new(Config)
}

func (config *Config) Init() {
	//static init FIXME
	config.AcquTime = 30
}