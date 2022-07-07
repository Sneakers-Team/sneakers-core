package apiserver

type Config struct {
	BindAdress string `toml:"bind_adress"`
}

func NewConfig() *Config {
	return &Config{}
}
