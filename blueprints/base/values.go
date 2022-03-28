package base

import (
	cg "github.com/ipld/edelweiss/codegen"
	"github.com/ipld/edelweiss/values"
)

const (
	IPLDPkg            = "github.com/ipld/go-ipld-prime"
	IPLDDatamodelPkg   = "github.com/ipld/go-ipld-prime/datamodel"
	EdelweissValuesPkg = values.PkgPath
)

var (
	Nil          = cg.V("nil")
	IPLDKindType = &cg.GoTypeRef{PkgPath: IPLDDatamodelPkg, TypeName: "Kind"}
	IPLDNull     = &cg.GoTypeRef{PkgPath: IPLDDatamodelPkg, TypeName: "Null"}
	// IPLD kind values
	IPLDKindInvalid = &cg.GoRef{PkgPath: IPLDDatamodelPkg, Name: "Kind_Invalid"}
	IPLDKindNull    = &cg.GoRef{PkgPath: IPLDDatamodelPkg, Name: "Kind_Null"}
	IPLDKindBool    = &cg.GoRef{PkgPath: IPLDDatamodelPkg, Name: "Kind_Bool"}
	IPLDKindInt     = &cg.GoRef{PkgPath: IPLDDatamodelPkg, Name: "Kind_Int"}
	IPLDKindFloat   = &cg.GoRef{PkgPath: IPLDDatamodelPkg, Name: "Kind_Float"}
	IPLDKindLink    = &cg.GoRef{PkgPath: IPLDDatamodelPkg, Name: "Kind_Link"}
	IPLDKindString  = &cg.GoRef{PkgPath: IPLDDatamodelPkg, Name: "Kind_String"}
	IPLDKindList    = &cg.GoRef{PkgPath: IPLDDatamodelPkg, Name: "Kind_List"}
	IPLDKindMap     = &cg.GoRef{PkgPath: IPLDDatamodelPkg, Name: "Kind_Map"}

	IPLDNodeType          = &cg.GoTypeRef{PkgPath: IPLDDatamodelPkg, TypeName: "Node"}
	IPLDMapIteratorType   = &cg.GoTypeRef{PkgPath: IPLDDatamodelPkg, TypeName: "MapIterator"}
	IPLDListIteratorType  = &cg.GoTypeRef{PkgPath: IPLDDatamodelPkg, TypeName: "ListIterator"}
	IPLDPathSegment       = &cg.GoTypeRef{PkgPath: IPLDDatamodelPkg, TypeName: "PathSegment"}
	IPLDLinkType          = &cg.GoTypeRef{PkgPath: IPLDDatamodelPkg, TypeName: "Link"}
	IPLDNodePrototypeType = &cg.GoTypeRef{PkgPath: IPLDDatamodelPkg, TypeName: "NodePrototype"}

	IPLDDeepEqual       = &cg.GoRef{PkgPath: IPLDPkg, Name: "DeepEqual"}
	IPLDEncode          = &cg.GoRef{PkgPath: IPLDPkg, Name: "Encode"}
	IPLDDecode          = &cg.GoRef{PkgPath: IPLDPkg, Name: "Decode"}
	IPLDDecodeStreaming = &cg.GoRef{PkgPath: IPLDPkg, Name: "DecodeStreaming"}
	DAGJSONEncode       = cg.GoRef{PkgPath: "github.com/ipld/go-ipld-prime/codec/dagjson", Name: "Encode"}
	DAGJSONDecode       = cg.GoRef{PkgPath: "github.com/ipld/go-ipld-prime/codec/dagjson", Name: "Decode"}
)

var (
	EdelweissErrNA       = &cg.GoRef{PkgPath: EdelweissValuesPkg, Name: "ErrNA"}
	EdelweissErrBounds   = &cg.GoRef{PkgPath: EdelweissValuesPkg, Name: "ErrBounds"}
	EdelweissErrNotFound = &cg.GoRef{PkgPath: EdelweissValuesPkg, Name: "ErrNotFound"}
	EdelweissString      = &cg.GoRef{PkgPath: EdelweissValuesPkg, Name: "String"}
	EdelweissInt         = &cg.GoRef{PkgPath: EdelweissValuesPkg, Name: "Int"}
	EdelweissParseFunc   = &cg.GoRef{PkgPath: EdelweissValuesPkg, Name: "ParseFunc"}
)

var (
	Errorf   = &cg.GoRef{PkgPath: "fmt", Name: "Errorf"}
	ErrorsIs = &cg.GoRef{PkgPath: "errors", Name: "Is"}
)
