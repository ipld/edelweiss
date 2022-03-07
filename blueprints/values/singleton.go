package values

import (
	"github.com/ipld/edelweiss/blueprints/base"
	cg "github.com/ipld/edelweiss/codegen"
	"github.com/ipld/edelweiss/defs"
)

func BuildSingletonImpl(typeDef defs.Def, goTypeRef cg.GoTypeRef) cg.GoTypeImpl {
	return &GoSingletonImpl{Def: typeDef, Ref: goTypeRef}
}

type GoSingletonImpl struct {
	Def defs.Def
	Ref cg.GoTypeRef
}

func (x *GoSingletonImpl) ProtoDef() defs.Def {
	return x.Def
}

func (x *GoSingletonImpl) GoTypeRef() cg.GoTypeRef {
	return x.Ref
}

func (x *GoSingletonImpl) GoDef() cg.Blueprint {
	common := cg.BlueMap{
		"Type":          x.Ref,
		"Node":          base.IPLDNodeType,
		"KindType":      base.IPLDKindType,
		"ErrNA":         base.EdelweissErrNA,
		"PathSegment":   base.IPLDPathSegment,
		"MapIterator":   base.IPLDMapIteratorType,
		"ListIterator":  base.IPLDListIteratorType,
		"Link":          base.IPLDLinkType,
		"NodePrototype": base.IPLDNodePrototypeType,
	}
	var specific cg.BlueMap
	switch t := x.Def.(type) {
	case defs.SingletonBool:
		specific = cg.BlueMap{
			"KindValue":             base.IPLDKindBool,
			"AsMethod":              cg.V("AsBool"),
			"ValueLiteral":          cg.BoolLiteral(t.Bool),
			"AsBoolReturnsResult":   cg.BoolLiteral(t.Bool),
			"AsBoolReturnsError":    base.Nil,
			"AsIntReturnsResult":    cg.IntLiteral(0),
			"AsIntReturnsError":     base.EdelweissErrNA,
			"AsFloatReturnsResult":  cg.FloatLiteral(0),
			"AsFloatReturnsError":   base.EdelweissErrNA,
			"AsStringReturnsResult": cg.StringLiteral(""),
			"AsStringReturnsError":  base.EdelweissErrNA,
		}
	case defs.SingletonInt:
		specific = cg.BlueMap{
			"KindValue":             base.IPLDKindInt,
			"AsMethod":              cg.V("AsInt"),
			"ValueLiteral":          cg.IntLiteral(t.Int),
			"AsBoolReturnsResult":   cg.BoolLiteral(false),
			"AsBoolReturnsError":    base.EdelweissErrNA,
			"AsIntReturnsResult":    cg.IntLiteral(t.Int),
			"AsIntReturnsError":     base.Nil,
			"AsFloatReturnsResult":  cg.FloatLiteral(0),
			"AsFloatReturnsError":   base.EdelweissErrNA,
			"AsStringReturnsResult": cg.StringLiteral(""),
			"AsStringReturnsError":  base.EdelweissErrNA,
		}
	case defs.SingletonFloat:
		specific = cg.BlueMap{
			"KindValue":             base.IPLDKindFloat,
			"AsMethod":              cg.V("AsFloat"),
			"ValueLiteral":          cg.FloatLiteral(t.Float),
			"AsBoolReturnsResult":   cg.BoolLiteral(false),
			"AsBoolReturnsError":    base.EdelweissErrNA,
			"AsIntReturnsResult":    cg.IntLiteral(0),
			"AsIntReturnsError":     base.EdelweissErrNA,
			"AsFloatReturnsResult":  cg.FloatLiteral(t.Float),
			"AsFloatReturnsError":   base.Nil,
			"AsStringReturnsResult": cg.StringLiteral(""),
			"AsStringReturnsError":  base.EdelweissErrNA,
		}
	case defs.SingletonString:
		specific = cg.BlueMap{
			"KindValue":             base.IPLDKindString,
			"AsMethod":              cg.V("AsString"),
			"ValueLiteral":          cg.StringLiteral(t.String),
			"AsBoolReturnsResult":   cg.BoolLiteral(false),
			"AsBoolReturnsError":    base.EdelweissErrNA,
			"AsIntReturnsResult":    cg.IntLiteral(0),
			"AsIntReturnsError":     base.EdelweissErrNA,
			"AsFloatReturnsResult":  cg.FloatLiteral(0),
			"AsFloatReturnsError":   base.EdelweissErrNA,
			"AsStringReturnsResult": cg.StringLiteral(t.String),
			"AsStringReturnsError":  base.Nil,
		}
	case defs.SingletonChar:
		specific = cg.BlueMap{
			"KindValue":             base.IPLDKindInt,
			"AsMethod":              cg.V("AsInt"),
			"ValueLiteral":          cg.IntLiteral(t.Char),
			"AsBoolReturnsResult":   cg.BoolLiteral(false),
			"AsBoolReturnsError":    base.EdelweissErrNA,
			"AsIntReturnsResult":    cg.IntLiteral(t.Char),
			"AsIntReturnsError":     base.Nil,
			"AsFloatReturnsResult":  cg.FloatLiteral(0),
			"AsFloatReturnsError":   base.EdelweissErrNA,
			"AsStringReturnsResult": cg.StringLiteral(""),
			"AsStringReturnsError":  base.EdelweissErrNA,
		}
	case defs.SingletonByte:
		specific = cg.BlueMap{
			"KindValue":             base.IPLDKindInt,
			"AsMethod":              cg.V("AsInt"),
			"ValueLiteral":          cg.IntLiteral(t.Byte),
			"AsBoolReturnsResult":   cg.BoolLiteral(false),
			"AsBoolReturnsError":    base.EdelweissErrNA,
			"AsIntReturnsResult":    cg.IntLiteral(t.Byte),
			"AsIntReturnsError":     base.Nil,
			"AsFloatReturnsResult":  cg.FloatLiteral(0),
			"AsFloatReturnsError":   base.EdelweissErrNA,
			"AsStringReturnsResult": cg.StringLiteral(""),
			"AsStringReturnsError":  base.EdelweissErrNA,
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
