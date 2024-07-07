package apiserver

// congif ...
type Config struct {
	BindAddr string `toml:"bind_ddr"`
}

// new config
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
	}
}
