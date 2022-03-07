package codegen

import (
	"fmt"

	"github.com/ipld/edelweiss/plans"
)

// compilation structures

type LookupDepGoRef interface {
	LookupDepGoRef(plans.Plan) GoTypeRef
}

type PlanToGoTypeRef map[plans.BuiltinOrRefPlan]GoTypeRef

func (m PlanToGoTypeRef) LookupDepGoRef(t plans.Plan) GoTypeRef {
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
