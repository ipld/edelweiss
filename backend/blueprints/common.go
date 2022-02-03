package blueprints

import (
	cg "github.com/ipld/edelweiss/backend/codegen"
)

const (
	IPLDPkg            = "github.com/ipld/go-ipld-prime"
	IPLDDatamodelPkg   = "github.com/ipld/go-ipld-prime/datamodel"
	EdelweissValuesPkg = "github.com/ipld/edelweiss/backend/values"
)

var (
	IPLDKindType = &cg.GoTypeRef{PkgPath: IPLDDatamodelPkg, TypeName: "Kind"}
	// IPLD kind values
	IPLDKindInt = &cg.GoRef{PkgPath: IPLDDatamodelPkg, Name: "XXX"}

	IPLDNodeType         = &cg.GoTypeRef{PkgPath: IPLDDatamodelPkg, TypeName: "Node"}
	IPLDMapIteratorType  = &cg.GoTypeRef{PkgPath: IPLDDatamodelPkg, TypeName: "MapIterator"}
	IPLDListIteratorType = &cg.GoTypeRef{PkgPath: IPLDDatamodelPkg, TypeName: "ListIterator"}
)

var (
	EdelweissErrNA = &cg.GoRef{PkgPath: EdelweissValuesPkg, Name: "ErrNA"}
)
