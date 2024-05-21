package main

import (
	"github.com/go-follow/file-grouper/internal/config"
	"github.com/go-follow/file-grouper/internal/grouper"
	logger "github.com/go-follow/file-grouper/pkg/logger"
)

func main() {
	l := logger.New()
	cfg, err := config.New()
	if err != nil {
		l.Error(err)
		return
	}

	g := grouper.New(cfg.Directory, cfg.Directory, cfg.IsRecurse)
	count, err := g.GroupFiles()
	if err != nil {
		l.Error(err)
		return
	}

	l.Infof("Processing is complete. %d files successfully processed", count)
}
