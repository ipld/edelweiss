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

func (g GoRef) Append(suffix string) GoRef {
	return GoRef{PkgPath: g.PkgPath, Name: g.Name + suffix}
}

func (g GoRef) Prepend(prefix string) GoRef {
	return GoRef{PkgPath: g.PkgPath, Name: prefix + g.Name}
}

type GoTypeRef struct {
	PkgPath  string // go pkg path
	TypeName string // go type name
}

func (g GoTypeRef) Write(ctx GoFileContext, w io.Writer) error {
	return GoRef{g.PkgPath, g.TypeName}.Write(ctx, w)
}

func (g GoTypeRef) Append(suffix string) GoTypeRef {
	return GoTypeRef{PkgPath: g.PkgPath, TypeName: g.TypeName + suffix}
}

func (g GoTypeRef) Prepend(prefix string) GoTypeRef {
	return GoTypeRef{PkgPath: g.PkgPath, TypeName: prefix + g.TypeName}
}

type GoTypeImpl interface {
	ProtoDef() def.Def
	GoTypeRef() GoTypeRef
	GoDef() Blueprint
}

type GoTypeImpls []GoTypeImpl
