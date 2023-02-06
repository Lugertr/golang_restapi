package apiserver

type Config struct {
	BindAddr string `toml:"bind_addr"`
}

// NewConfig ...
func newConfig() *Config {
	return &Config{
		BindAddr: ":8080",
	}
}
