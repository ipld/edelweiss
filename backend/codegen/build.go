package codegen

import (
	"fmt"
)

type Builder interface {
	Build() error
}

type Builders []Builder

func (bs Builders) Build() error {
	for _, b := range bs {
		if err := b.Build(); err != nil {
			return fmt.Errorf("builder %#v (%w)", err)
		}
	}
	return nil
}
