package codegen

import (
	"io"

	"github.com/ipld/edelweiss/def"
)

type GoRef struct {
	PkgPath string // go pkg path
	Name    string // go object name
}

func (g GoRef) Write(ctx GoFileContext, w io.Writer) error {
	return V(ctx.ReferTo(g.PkgPath, g.Name)).Write(ctx, w)
}

type GoTypeRef struct {
	PkgPath  string // go pkg path
	TypeName string // go type name
}

func (g GoTypeRef) Write(ctx GoFileContext, w io.Writer) error {
	return GoRef{g.PkgPath, g.TypeName}.Write(ctx, w)
}

type GoTypeImpl interface {
	ProtoDef() def.Type
	GoTypeRef() GoTypeRef
	GoDef() Blueprint
}

type GoTypeImpls []GoTypeImpl
