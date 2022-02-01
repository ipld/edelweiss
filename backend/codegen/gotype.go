package codegen

import (
	"fmt"
	"io"

	"github.com/ipld/edelweiss/def"
)

type GoTypeRef struct {
	PkgPath  string // go pkg path
	TypeName string // go type name
}

func (g GoTypeRef) WriteRef(ctx GoFileContext, w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s.%s", ctx.RequireImport(g.PkgPath).Alias, g.TypeName)
	return err
}

type GoTypeImpl interface {
	Def() def.Type
	GoTypeRef() GoTypeRef
	WriteDef(GoFileContext, io.Writer) error
	WriteRef(GoFileContext, io.Writer) error
}
