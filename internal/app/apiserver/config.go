package apiserver

// congif ...
type Config struct {
	BindAddr string `toml:"bind_ddr"`
	LogLevel string `toml:"log_level"`
}

// new config
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
