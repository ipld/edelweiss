package codegen

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
	panic("xxx")
}

type GoFile struct {
	Name    string
	Content Generator
}

func (x GoFile) Build() error {
	panic("xxx")
}

type Generator interface {
	WriteTo(io.Writer) (int64, error)
}
