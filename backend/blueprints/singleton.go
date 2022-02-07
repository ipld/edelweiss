package blueprints

import (
	cg "github.com/ipld/edelweiss/backend/codegen"
	"github.com/ipld/edelweiss/def"
)

func BuildSingletonImpl(typeDef def.Type, goTypeRef cg.GoTypeRef) (cg.GoTypeImpl, error) {
	return &GoSingletonImpl{Def: typeDef, Ref: goTypeRef}, nil
}

type GoSingletonImpl struct {
	Def def.Type
	Ref cg.GoTypeRef
}

func (x *GoSingletonImpl) ProtoDef() def.Type {
	return x.Def
}

func (x *GoSingletonImpl) GoTypeRef() cg.GoTypeRef {
	return x.Ref
}

func (x *GoSingletonImpl) GoDef() cg.Blueprint {
	common := cg.BlueMap{
		"Type":          x.Ref,
		"Node":          IPLDNodeType,
		"KindType":      IPLDKindType,
		"ErrNA":         EdelweissErrNA,
		"PathSegment":   IPLDPathSegment,
		"MapIterator":   IPLDMapIteratorType,
		"ListIterator":  IPLDListIteratorType,
		"Link":          IPLDLinkType,
		"NodePrototype": IPLDNodePrototypeType,
	}
	var specific cg.BlueMap
	switch t := x.Def.(type) {
	case def.SingletonBool:
		specific = cg.BlueMap{
			"KindValue":             IPLDKindBool,
			"AsMethod":              cg.V("AsBool"),
			"ValueLiteral":          cg.BoolLiteral(t.Bool),
			"AsBoolReturnsResult":   cg.BoolLiteral(t.Bool),
			"AsBoolReturnsError":    Nil,
			"AsIntReturnsResult":    cg.IntLiteral(0),
			"AsIntReturnsError":     EdelweissErrNA,
			"AsFloatReturnsResult":  cg.FloatLiteral(0),
			"AsFloatReturnsError":   EdelweissErrNA,
			"AsStringReturnsResult": cg.StringLiteral(""),
			"AsStringReturnsError":  EdelweissErrNA,
		}
	case def.SingletonInt:
		specific = cg.BlueMap{
			"KindValue":             IPLDKindInt,
			"AsMethod":              cg.V("AsInt"),
			"ValueLiteral":          cg.IntLiteral(t.Int),
			"AsBoolReturnsResult":   cg.BoolLiteral(false),
			"AsBoolReturnsError":    EdelweissErrNA,
			"AsIntReturnsResult":    cg.IntLiteral(t.Int),
			"AsIntReturnsError":     Nil,
			"AsFloatReturnsResult":  cg.FloatLiteral(0),
			"AsFloatReturnsError":   EdelweissErrNA,
			"AsStringReturnsResult": cg.StringLiteral(""),
			"AsStringReturnsError":  EdelweissErrNA,
		}
	case def.SingletonFloat:
		specific = cg.BlueMap{
			"KindValue":             IPLDKindFloat,
			"AsMethod":              cg.V("AsFloat"),
			"ValueLiteral":          cg.FloatLiteral(t.Float),
			"AsBoolReturnsResult":   cg.BoolLiteral(false),
			"AsBoolReturnsError":    EdelweissErrNA,
			"AsIntReturnsResult":    cg.IntLiteral(0),
			"AsIntReturnsError":     EdelweissErrNA,
			"AsFloatReturnsResult":  cg.FloatLiteral(t.Float),
			"AsFloatReturnsError":   Nil,
			"AsStringReturnsResult": cg.StringLiteral(""),
			"AsStringReturnsError":  EdelweissErrNA,
		}
	case def.SingletonString:
		specific = cg.BlueMap{
			"KindValue":             IPLDKindString,
			"AsMethod":              cg.V("AsString"),
			"ValueLiteral":          cg.StringLiteral(t.String),
			"AsBoolReturnsResult":   cg.BoolLiteral(false),
			"AsBoolReturnsError":    EdelweissErrNA,
			"AsIntReturnsResult":    cg.IntLiteral(0),
			"AsIntReturnsError":     EdelweissErrNA,
			"AsFloatReturnsResult":  cg.FloatLiteral(0),
			"AsFloatReturnsError":   EdelweissErrNA,
			"AsStringReturnsResult": cg.StringLiteral(t.String),
			"AsStringReturnsError":  Nil,
		}
	case def.SingletonChar:
		specific = cg.BlueMap{
			"KindValue":             IPLDKindInt,
			"AsMethod":              cg.V("AsInt"),
			"ValueLiteral":          cg.IntLiteral(t.Char),
			"AsBoolReturnsResult":   cg.BoolLiteral(false),
			"AsBoolReturnsError":    EdelweissErrNA,
			"AsIntReturnsResult":    cg.IntLiteral(t.Char),
			"AsIntReturnsError":     Nil,
			"AsFloatReturnsResult":  cg.FloatLiteral(0),
			"AsFloatReturnsError":   EdelweissErrNA,
			"AsStringReturnsResult": cg.StringLiteral(""),
			"AsStringReturnsError":  EdelweissErrNA,
		}
	case def.SingletonByte:
		specific = cg.BlueMap{
			"KindValue":             IPLDKindInt,
			"AsMethod":              cg.V("AsInt"),
			"ValueLiteral":          cg.IntLiteral(t.Byte),
			"AsBoolReturnsResult":   cg.BoolLiteral(false),
			"AsBoolReturnsError":    EdelweissErrNA,
			"AsIntReturnsResult":    cg.IntLiteral(t.Byte),
			"AsIntReturnsError":     Nil,
			"AsFloatReturnsResult":  cg.FloatLiteral(0),
			"AsFloatReturnsError":   EdelweissErrNA,
			"AsStringReturnsResult": cg.StringLiteral(""),
			"AsStringReturnsError":  EdelweissErrNA,
		}
	default:
		panic("unrecognized singleton type")
	}
	return cg.T{
		Data: cg.MergeBlueMaps(common, specific),
		Src: `
// -- protocol type {{.Type}} --

type {{.Type}} struct{}

func ({{.Type}}) Parse(n {{.Node}}) error {
	if n.Kind() != {{.KindValue}} {
		return {{.ErrNA}}
	}
	v, err := n.{{.AsMethod}}()
	if err != nil {
		return err
	}
	if v != {{.ValueLiteral}} {
		return {{.ErrNA}}
	}
	return nil
}

func (v {{.Type}}) Node() {{.Node}} {
	return v
}

func ({{.Type}}) Kind() {{.KindType}} {
	return {{.KindValue}}
}

func ({{.Type}}) LookupByString(string) ({{.Node}}, error) {
	return nil, {{.ErrNA}}
}

func ({{.Type}}) LookupByNode(key {{.Node}}) ({{.Node}}, error) {
	return nil, {{.ErrNA}}
}

func ({{.Type}}) LookupByIndex(idx int64) ({{.Node}}, error) {
	return nil, {{.ErrNA}}
}

func ({{.Type}}) LookupBySegment(_ {{.PathSegment}}) ({{.Node}}, error) {
	return nil, {{.ErrNA}}
}

func ({{.Type}}) MapIterator() {{.MapIterator}} {
	return nil
}

func ({{.Type}}) ListIterator() {{.ListIterator}} {
	return nil
}

func ({{.Type}}) Length() int64 {
	return -1
}

func ({{.Type}}) IsAbsent() bool {
	return false
}

func ({{.Type}}) IsNull() bool {
	return false
}

func (v {{.Type}}) AsBool() (bool, error) {
	return {{.AsBoolReturnsResult}}, {{.AsBoolReturnsError}}
}

func ({{.Type}}) AsInt() (int64, error) {
	return {{.AsIntReturnsResult}}, {{.AsIntReturnsError}}
}

func ({{.Type}}) AsFloat() (float64, error) {
	return {{.AsFloatReturnsResult}}, {{.AsFloatReturnsError}}
}

func ({{.Type}}) AsString() (string, error) {
	return {{.AsStringReturnsResult}}, {{.AsStringReturnsError}}
}

func ({{.Type}}) AsBytes() ([]byte, error) {
	return nil, {{.ErrNA}}
}

func ({{.Type}}) AsLink() ({{.Link}}, error) {
	return nil, {{.ErrNA}}
}

func ({{.Type}}) Prototype() {{.NodePrototype}} {
	return nil
}
`,
	}
}
