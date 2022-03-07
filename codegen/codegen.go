package codegen

import (
	"fmt"

	"github.com/ipld/edelweiss/defs"
)

// compilation structures

type LookupDepGoRef interface {
	LookupDepGoRef(defs.Def) GoTypeRef
}

type DefToGoTypeRef map[defs.Def]GoTypeRef

func (m DefToGoTypeRef) LookupDepGoRef(t defs.Def) GoTypeRef {
	r, ok := m[t]
	if !ok {
		panic(fmt.Sprintf("missing dependency %#v", t))
	}
	return r
}

// file generation

type GoFileContext interface {
	RequireImport(pkgPath string) *GoFileImport
	ReferTo(pkgPath string, symbol string) string
}

type GoFileImport struct {
	PkgPath string
	Alias   string
}
