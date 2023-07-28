package apiserver

import "github.com/DoomGuy1818/restapi-server/internal/app/store"

type Config struct {
	BindAddr string `toml:"bind_Addr"`
	LogLevel string `toml:log_level`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8000",
		LogLevel: "debug",
		Store: store.NewConfig(),
	}
}
