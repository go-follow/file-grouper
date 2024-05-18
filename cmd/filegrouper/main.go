package main

import (
	"fmt"

	"github.com/go-follow/file-grouper/internal/config"
	"github.com/go-follow/file-grouper/internal/grouper"
)

func main() {
	cfg := config.New()

	g := grouper.New(cfg.Directory, cfg.IsRecurse)
	if err := g.GroupFiles(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("success")
}
