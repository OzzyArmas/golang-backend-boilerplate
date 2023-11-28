package config

type HttpConfig struct {
	ListenPort string `mapstructure:"listenPort"`
}

// type dbConfig struct {
// 	URL string
// }

type Config struct {
	HTTP HttpConfig `mapstructure:"httpConfig"`
	// DB   dbConfig
}
