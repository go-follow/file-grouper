package grouper

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"strings"
	"time"
)

type Grouper interface {
	GroupFiles() (int, error)
}

type Group struct {
	directoryInput  string
	direcotryOutput string
	isRecurse       bool
}

func New(dirInput, dirOutput string, isRecurse bool) Grouper {
	return &Group{
		directoryInput:  dirInput,
		direcotryOutput: dirOutput,
		isRecurse:       isRecurse,
	}
}

func (g *Group) GroupFiles() (int, error) {
	entries, err := os.ReadDir(g.directoryInput)
	if err != nil {
		return -1, fmt.Errorf("failed os.ReadDir: %w", err)
	}
	count := 0
	for _, e := range entries {
		if e.IsDir() {
			countR, err := g.recourceGroup(e.Name())
			if err != nil {
				return -1, err
			}
			count += countR
			continue
		}
		if strings.HasPrefix(e.Name(), ".") {
			continue
		}
		fileInfo, err := e.Info()
		if err != nil {
			return -1, fmt.Errorf("failed get Info file: %w", err)
		}
		path := g.makePathGroup(fileInfo.ModTime())
		src := fmt.Sprintf("%s/%s", g.directoryInput, fileInfo.Name())
		dst := fmt.Sprintf("%s/%s", path, fileInfo.Name())
		if err := g.moveFile(src, dst); err != nil {
			return -1, fmt.Errorf("failed copyFile: %w", err)
		}
		count++
	}
	return count, nil
}

func (g *Group) recourceGroup(dirName string) (int, error) {
	if !g.isRecurse {
		return 0, nil
	}
	inputDir := fmt.Sprintf("%s/%s", g.directoryInput, dirName)
	groupR := New(inputDir, g.direcotryOutput, g.isRecurse)
	countR, err := groupR.GroupFiles()
	if err != nil {
		return -1, fmt.Errorf("failed recurce GroupFiles: %w", err)
	}
	entrs, err := os.ReadDir(inputDir)
	if err != nil {
		return -1, fmt.Errorf("failed recurce os.ReadDir: %w", err)
	}
	if g.isEmptyEntries(entrs) {
		if err := os.RemoveAll(inputDir); err != nil {
			return -1, fmt.Errorf("failed recurce os.Remove: %w", err)
		}
	}
	return countR, nil
}

func (g *Group) isEmptyEntries(entrs []fs.DirEntry) bool {
	for _, e := range entrs {
		if !strings.HasPrefix(e.Name(), ".") {
			return false
		}
	}
	return true
}

func (g *Group) makePathGroup(t time.Time) string {
	return fmt.Sprintf("%s/%d/%s", g.direcotryOutput, t.Year(), t.Format("2006-01-02"))
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
