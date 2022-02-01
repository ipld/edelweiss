package blueprint

import (
	"fmt"
	"io"

	"github.com/ipld/edelweiss/def"
)

func BuildSingletonBoolGoImpl(plan GoTypeImplPlan, typeDef def.SingletonBool, goTypeRef GoTypeRef) (GoTypeImpl, error) {
	return &GoSingletonImpl{
		def: typeDef,
		ref: goTypeRef,
		data: singletonBlueprintData{
			TypeName:         goTypeRef.TypeName,
			IPLDKindName:     "Kind_Bool",
			IPLDAsMethodName: "AsBool",
			IPLDValueLiteral: fmt.Sprintf("%v", typeDef.Bool),
			AsBoolReturns:    fmt.Sprintf("%v, nil", typeDef.Bool),
			AsIntReturns:     "0, ErrNA", //XXX: move error defs in dedicated pkg
			AsFloatReturns:   "0.0, ErrNA",
			AsStringReturns:  "\"\", ErrNA",
		},
	}, nil
}

func BuildSingletonFloatGoImpl(plan GoTypeImplPlan, typeDef def.SingletonFloat, goTypeRef GoTypeRef) (GoTypeImpl, error) {
	return &GoSingletonImpl{
		def: typeDef,
		ref: goTypeRef,
		data: singletonBlueprintData{
			TypeName:         goTypeRef.TypeName,
			IPLDKindName:     "Kind_Float",
			IPLDAsMethodName: "AsFloat",
			IPLDValueLiteral: fmt.Sprintf("%v", typeDef.Float),
			AsBoolReturns:    "false, ErrNA",
			AsIntReturns:     "0, ErrNA",
			AsFloatReturns:   fmt.Sprintf("%v, nil", typeDef.Float),
			AsStringReturns:  "\"\", ErrNA",
		},
	}, nil
}

func BuildSingletonIntGoImpl(plan GoTypeImplPlan, typeDef def.SingletonInt, goTypeRef GoTypeRef) (GoTypeImpl, error) {
	return &GoSingletonImpl{
		def: typeDef,
		ref: goTypeRef,
		data: singletonBlueprintData{
			TypeName:         goTypeRef.TypeName,
			IPLDKindName:     "Kind_Int",
			IPLDAsMethodName: "AsInt",
			IPLDValueLiteral: fmt.Sprintf("%d", typeDef.Int),
			AsBoolReturns:    "false, ErrNA",
			AsIntReturns:     fmt.Sprintf("%d, nil", typeDef.Int),
			AsFloatReturns:   "0.0, ErrNA",
			AsStringReturns:  "\"\", ErrNA",
		},
	}, nil
}

func BuildSingletonByteGoImpl(plan GoTypeImplPlan, typeDef def.SingletonByte, goTypeRef GoTypeRef) (GoTypeImpl, error) {
	return &GoSingletonImpl{
		def: typeDef,
		ref: goTypeRef,
		data: singletonBlueprintData{
			TypeName:         goTypeRef.TypeName,
			IPLDKindName:     "Kind_Int",
			IPLDAsMethodName: "AsInt",
			IPLDValueLiteral: fmt.Sprintf("%d", typeDef.Byte),
			AsBoolReturns:    "false, ErrNA",
			AsIntReturns:     fmt.Sprintf("%d, nil", typeDef.Byte),
			AsFloatReturns:   "0.0, ErrNA",
			AsStringReturns:  "\"\", ErrNA",
		},
	}, nil
}

func BuildSingletonCharGoImpl(plan GoTypeImplPlan, typeDef def.SingletonChar, goTypeRef GoTypeRef) (GoTypeImpl, error) {
	return &GoSingletonImpl{
		def: typeDef,
		ref: goTypeRef,
		data: singletonBlueprintData{
			TypeName:         goTypeRef.TypeName,
			IPLDKindName:     "Kind_Int",
			IPLDAsMethodName: "AsInt",
			IPLDValueLiteral: fmt.Sprintf("%d", int(typeDef.Char)),
			AsBoolReturns:    "false, ErrNA",
			AsIntReturns:     fmt.Sprintf("%d, nil", int(typeDef.Char)),
			AsFloatReturns:   "0.0, ErrNA",
			AsStringReturns:  "\"\", ErrNA",
		},
	}, nil
}

func BuildSingletonStringGoImpl(plan GoTypeImplPlan, typeDef def.SingletonString, goTypeRef GoTypeRef) (GoTypeImpl, error) {
	return &GoSingletonImpl{
		def: typeDef,
		ref: goTypeRef,
		data: singletonBlueprintData{
			TypeName:         goTypeRef.TypeName,
			IPLDKindName:     "Kind_String",
			IPLDAsMethodName: "AsString",
			IPLDValueLiteral: fmt.Sprintf("%q", typeDef.String),
			AsBoolReturns:    "false, ErrNA",
			AsIntReturns:     "0, ErrNA",
			AsFloatReturns:   "0.0, ErrNA",
			AsStringReturns:  fmt.Sprintf("%q, nil", typeDef.String),
		},
	}, nil
}

type GoSingletonImpl struct {
	def  def.Type
	ref  GoTypeRef
	data singletonBlueprintData
}

func (g GoSingletonImpl) Def() def.Type {
	return g.def
}

func (g GoSingletonImpl) GoTypeRef() GoTypeRef {
	return g.ref
}

func (g GoSingletonImpl) WriteDef(w io.Writer) (int, error) {
	panic("XXX")
}

func (g GoSingletonImpl) WriteRef(w io.Writer) (int, error) {
	return g.ref.WriteRef(w)
}
