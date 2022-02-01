package codegen

import (
	"io"

	"github.com/ipld/edelweiss/def"
)

type GoTypeRef struct {
	PkgPath  string // go pkg path
	TypeName string // go type name
}

func (g GoTypeRef) WriteRef(w io.Writer) (int, error) {
	panic("XXX")
}

type GoTypeImpl interface {
	Def() def.Type
	GoTypeRef() GoTypeRef
	WriteDef(io.Writer) (int, error)
	WriteRef(io.Writer) (int, error)
}
