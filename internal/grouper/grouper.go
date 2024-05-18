package grouper

import (
	"fmt"
	"os"
	"path"
	"time"
)

type Grouper interface {
	GroupFiles() (int, error)
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

func (g *Group) GroupFiles() (int, error) {
	entries, err := os.ReadDir(g.directory)
	if err != nil {
		return -1, fmt.Errorf("failed os.ReadDir: %w", err)
	}
	count := 0
	for _, e := range entries {
		// TODO: implement recursive
		if e.IsDir() {
			continue
		}
		fileInfo, err := e.Info()
		if err != nil {
			return -1, fmt.Errorf("failed get Info file: %w", err)
		}
		path := g.makePathGroup(fileInfo.ModTime())
		src := fmt.Sprintf("%s/%s", g.directory, fileInfo.Name())
		dst := fmt.Sprintf("%s/%s", path, fileInfo.Name())
		if err := g.moveFile(src, dst); err != nil {
			return -1, fmt.Errorf("failed copyFile: %w", err)
		}
		count++
	}
	return count, nil
}

func (g *Group) makePathGroup(t time.Time) string {
	return fmt.Sprintf("%s/%d/%s", g.directory, t.Year(), t.Format("2006-01-02"))
}

func (g *Group) moveFile(src, dst string) error {
	if err := os.MkdirAll(path.Dir(dst), 0755); err != nil {
		return fmt.Errorf("failed os.MkdirAll: %w", err)
	}
	if err := os.Rename(src, dst); err != nil {
		return fmt.Errorf("failed os.Rename: %w", err)
	}
	return nil
}
