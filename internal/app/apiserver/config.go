package apiserver

type Config struct {
	BindAddr string `toml:"bind_Addr"`
	LogLevel string `toml:log_level`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8000",
		LogLevel: "debug",
	}
}
