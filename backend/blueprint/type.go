package blueprint

import (
	"io"

	"github.com/ipld/edelweiss/def"
)

type GoType interface {
	Def() GoTypeDef
	WriteDef(io.Writer)
	WriteRef(io.Writer) // XXX: import context for aliases
}

type GoTypeDef struct {
	Def  def.Type
	Ref  GoTypeRef
	Deps []GoType
}

type GoTypeRef struct {
	PkgPath  string // go pkg path
	TypeName string // go type name
}
