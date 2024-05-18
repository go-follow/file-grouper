package config

import "flag"

type Config struct {
	Directory string
	IsRecurse bool
}

func New() *Config {
	cfg := &Config{}
	flag.StringVar(&cfg.Directory, "directory", ".", "Directory to search files")
	flag.BoolVar(&cfg.IsRecurse, "recurse", false, "Recurse into subdirectories")
	flag.Parse()
	return cfg
}
