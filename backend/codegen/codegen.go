package codegen

import (
	"github.com/ipld/edelweiss/def"
)

// compilation structures

type NameToDef map[string]def.Type

type DefToGoTypeRef map[def.Type]GoTypeRef

type RefToGoTypeRef map[def.Ref]GoTypeRef

type DefToGoTypeImpl map[def.Type]GoTypeImpl

type GoTypeImplPlan struct {
	DefToGoTypeRef // definitions that must be code-generated
	RefToGoTypeRef // references used throughout definitions
}

// file generation

type GoFileContext interface {
	RequireImport(pkgPath string) *GoFileImport
}

type GoFileImport struct {
	PkgPath string
	Alias   string
}
