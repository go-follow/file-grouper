package main

import (
	"fmt"

	"github.com/go-follow/file-grouper/internal/config"
	"github.com/go-follow/file-grouper/internal/grouper"
)

func main() {
	cfg := config.New()

	g := grouper.New(cfg.Directory, cfg.IsRecurse)
	count, err := g.GroupFiles()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Processing is complete. %d files successfully processed", count)
}
