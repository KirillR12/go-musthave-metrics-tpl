package config

type Config struct {
	Address string
}

func NewConfig() *Config {
	return &Config{
		Address: "localhost:8080",
	}
}
