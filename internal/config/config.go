package config

import (
	"flag"
	"fmt"

	"github.com/go-follow/file-grouper/pkg/validation"
)

type Config struct {
	Directory string `validate:"required,dir"`
	IsRecurse bool
}

func New() (*Config, error) {
	cfg := &Config{}
	flag.StringVar(&cfg.Directory, "d", "", "Directory to search files")
	flag.BoolVar(&cfg.IsRecurse, "r", false, "Recurse into subdirectories")
	flag.Parse()
	if err := validation.New().Validate(cfg); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}
	return cfg, nil
}
