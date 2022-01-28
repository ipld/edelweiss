package blueprint

import (
	"io"
)

type Builder interface {
	Build() error
}

type GoFileOrDir interface {
	Builder
}

type GoDir struct {
	LocalPath string
	Children  []GoFileOrDir
}

func (x GoDir) Build() error {
	return XXX
}

type GoFile struct {
	Name    string
	Content Generator
}

func (x GoFile) Build() error {
	return XXX
}

type Generator interface {
	WriteTo(io.Writer) (int64, error)
}
