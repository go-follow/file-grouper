package main

import (
	"fmt"

	"github.com/go-follow/file-grouper/internal/config"
	"github.com/go-follow/file-grouper/internal/grouper"
	logger "github.com/go-follow/file-grouper/pkg"
)

func main() {
	cfg := config.New()
	l := logger.New()

	g := grouper.New(cfg.Directory, cfg.Directory, cfg.IsRecurse)
	count, err := g.GroupFiles()
	if err != nil {
		l.Error(err)
		return
	}

	fmt.Printf("Processing is complete. %d files successfully processed", count)
}
