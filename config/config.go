package config

type Config struct {
	AcquTime uint32
	UDPEnabled bool
	UDPAddress string
	InverseLogic bool
	WdStartTime uint8
	WdActive	bool
}

func NewConfig() *Config {
	return new(Config)
}

func (config *Config) Init() {
	//static init FIXME
	config.AcquTime = 15
	config.UDPEnabled = true
	config.UDPAddress = "127.0.0.1:9220"
	config.InverseLogic = true

	config.WdActive = true
	config.WdStartTime = 12
}