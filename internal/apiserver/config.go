package apiserver

type Config struct {
	BindAdress string `toml:"bind_adress"`
}

// Configure new APIServer config and return pointer of it.
func NewConfig() *Config {
	return &Config{}
}
