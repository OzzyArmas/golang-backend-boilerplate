package config

type HttpConfig struct {
	ListenPort string `mapstructure:"listenPort"`
}

type Config struct {
	HTTP HttpConfig `mapstructure:"httpConfig"`
}
