package codegen

import (
	"github.com/ipld/edelweiss/def"
)

// compilation structures

type NameToDef map[string]def.Type

type LookupDefToGoTypeRef interface {
	LookupDefToGoTypeRef(def.Type) GoTypeRef
}

type DefToGoTypeRef map[def.Type]GoTypeRef

func (m DefToGoTypeRef) LookupDefToGoTypeRef(t def.Type) GoTypeRef {
	return m[t]
}

type RefToGoTypeRef map[def.Ref]GoTypeRef

type DefToGoTypeImpl map[def.Type]GoTypeImpl

type GoTypeImplPlan struct {
	DefToGoTypeRef // definitions that must be code-generated
	RefToGoTypeRef // references used throughout definitions
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
