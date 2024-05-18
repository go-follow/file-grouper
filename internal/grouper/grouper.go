package grouper

import (
	"fmt"
	"os"
	"path"
	"time"
)

type Grouper interface {
	GroupFiles() error
}

type Group struct {
	directory string
	isRecurse bool
}

func New(directory string, isRecurse bool) Grouper {
	return &Group{
		directory: directory,
		isRecurse: isRecurse,
	}
}

func (g *Group) GroupFiles() error {
	entries, err := os.ReadDir(g.directory)
	if err != nil {
		return fmt.Errorf("failed os.ReadDir: %w", err)
	}

	for _, e := range entries {
		// TODO: implement recursive
		if e.IsDir() {
			continue
		}
		fileInfo, err := e.Info()
		if err != nil {
			return fmt.Errorf("failed get Info file: %w", err)
		}
		path := g.makePathGroup(fileInfo.ModTime())
		src := fmt.Sprintf("%s/%s", g.directory, fileInfo.Name())
		dst := fmt.Sprintf("%s/%s", path, fileInfo.Name())
		if err := g.copyFile(src, dst); err != nil {
			return fmt.Errorf("failed copyFile: %w", err)
		}
	}
	return nil
}

func (g *Group) makePathGroup(t time.Time) string {
	return fmt.Sprintf("%s/output/%d/%s", g.directory, t.Year(), t.Format("2006-01-02"))
}

func (g *Group) copyFile(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return fmt.Errorf("failed os.ReadFile: %w", err)
	}

	if err := os.MkdirAll(path.Dir(dst), 0755); err != nil {
		return fmt.Errorf("failed os.MkdirAll: %w", err)
	}

	if err := os.WriteFile(dst, data, 0600); err != nil {
		return fmt.Errorf("failed os.WriteFile: %w", err)
	}
	return nil
}
