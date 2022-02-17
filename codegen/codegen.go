package codegen

import (
	"fmt"

	"github.com/ipld/edelweiss/def"
)

// compilation structures

type LookupDepGoRef interface {
	LookupDepGoRef(def.Type) GoTypeRef
}

type DefToGoTypeRef map[def.Type]GoTypeRef

func (m DefToGoTypeRef) LookupDepGoRef(t def.Type) GoTypeRef {
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