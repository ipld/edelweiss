package values

import (
	"github.com/ipld/edelweiss/blueprints/base"
	cg "github.com/ipld/edelweiss/codegen"
	"github.com/ipld/edelweiss/def"
)

func BuildCallImpl(
	lookup cg.LookupDepGoRef,
	typeDef def.Call,
	goTypeRef cg.GoTypeRef,
) cg.GoTypeImpl {
	return &GoCallImpl{
		Lookup: lookup,
		Def:    typeDef,
		Ref:    goTypeRef,
	}
}

type GoCallImpl struct {
	Lookup cg.LookupDepGoRef
	Def    def.Call
	Ref    cg.GoTypeRef
}

func (x *GoCallImpl) ProtoDef() def.Def {
	return x.Def
}

func (x *GoCallImpl) GoTypeRef() cg.GoTypeRef {
	return x.Ref
}

func (x *GoCallImpl) GoDef() cg.Blueprint {
	// build type definition
	data := cg.BlueMap{
		"Type": x.Ref,
		"TypeMapIterator": cg.GoTypeRef{
			PkgPath:  x.Ref.PkgPath,
			TypeName: x.Ref.TypeName + "_MapIterator",
		},
		"IDType":          x.Lookup.LookupDepGoRef(x.Def.ID),
		"ArgType":         x.Lookup.LookupDepGoRef(x.Def.Fn.Arg),
		"Node":            base.IPLDNodeType,
		"KindType":        base.IPLDKindType,
		"KindMap":         base.IPLDKindMap,
		"KindString":      base.IPLDKindString,
		"KindInt":         base.IPLDKindInt,
		"ErrNA":           base.EdelweissErrNA,
		"ErrBounds":       base.EdelweissErrBounds,
		"ErrNotFound":     base.EdelweissErrNotFound,
		"PathSegment":     base.IPLDPathSegment,
		"MapIterator":     base.IPLDMapIteratorType,
		"ListIterator":    base.IPLDListIteratorType,
		"Link":            base.IPLDLinkType,
		"NodePrototype":   base.IPLDNodePrototypeType,
		"EdelweissString": base.EdelweissString,
		"EdelweissInt":    base.EdelweissInt,
		"Errorf":          base.Errorf,
	}
	return cg.T{
		Data: data,
		Src: `
// -- protocol type {{.Type}} --

type {{.Type}} struct {
	ID {{.IDType}}
	Arg {{.ArgType}}
}

func (v {{.Type}}) Node() {{.Node}} {
	return v
}

func (v *{{.Type}}) Parse(n {{.Node}}) error {
	if n.Kind() != {{.KindMap}} {
		return {{.ErrNA}}
	} else {
		iter := n.MapIterator()
		idParsed, argParsed := false, false
		for !iter.Done() {
			kn, vn, _ := iter.Next()
			k, err := kn.AsString()
			if err != nil {
				continue // skip unrecognized non-string keys
			}
			switch k {
			case "ID":
				if err := v.ID.Parse(vn); err != nil {
					return {{.Errorf}}("call id (%v)", err)
				}
				idParsed = true
			case "Arg":
				if err := v.Arg.Parse(vn); err != nil {
					return {{.Errorf}}("call argument (%v)", err)
				}
				argParsed = true
			}
			if idParsed && argParsed {
				return nil
			}
		}
		if idParsed && argParsed {
			return nil
		} else {
			return {{.Errorf}}("call id or arg is missing")
		}
	}
}

func ({{.Type}}) Kind() {{.KindType}} {
	return {{.KindMap}}
}

func (v {{.Type}}) LookupByString(s string) ({{.Node}}, error) {
	switch s {
	case "ID":
		return v.ID.Node(), nil
	case "Arg":
		return v.Arg.Node(), nil
	}
	return nil, {{.ErrNA}}
}

func (v {{.Type}}) LookupByNode(key {{.Node}}) ({{.Node}}, error) {
	switch key.Kind() {
	case {{.KindString}}:
		if s, err := key.AsString(); err != nil {
			return nil, err
		} else {
			return v.LookupByString(s)
		}
	case {{.KindInt}}:
		if i, err := key.AsInt(); err != nil {
			return nil, err
		} else {
			return v.LookupByIndex(i)
		}
	}
	return nil, {{.ErrNA}}
}

func (v {{.Type}}) LookupByIndex(i int64) ({{.Node}}, error) {
	return nil, {{.ErrNA}}
}

func (v {{.Type}}) LookupBySegment(seg {{.PathSegment}}) ({{.Node}}, error) {
	return v.LookupByString(seg.String())
}

func (v {{.Type}}) MapIterator() {{.MapIterator}} {
	return &{{.TypeMapIterator}}{&v, 0}
}

func (v {{.Type}}) ListIterator() {{.ListIterator}} {
	return nil
}

func (v {{.Type}}) Length() int64 {
	return 2
}

func ({{.Type}}) IsAbsent() bool {
	return false
}

func ({{.Type}}) IsNull() bool {
	return false
}

func (v {{.Type}}) AsBool() (bool, error) {
	return false, {{.ErrNA}}
}

func ({{.Type}}) AsInt() (int64, error) {
	return 0, {{.ErrNA}}
}

func ({{.Type}}) AsFloat() (float64, error) {
	return 0, {{.ErrNA}}
}

func ({{.Type}}) AsString() (string, error) {
	return "", {{.ErrNA}}
}

func ({{.Type}}) AsBytes() ([]byte, error) {
	return nil, {{.ErrNA}}
}

func ({{.Type}}) AsLink() ({{.Link}}, error) {
	return nil, {{.ErrNA}}
}

func ({{.Type}}) Prototype() {{.Node}}Prototype {
	return nil // not needed
}

type {{.TypeMapIterator}} struct {
	m  *{{.Type}}
	at int64
}

func (iter *{{.TypeMapIterator}}) Next() ({{.Node}}, {{.Node}}, error) {
	switch iter.at {
	case 0:
		iter.at++
		return {{.EdelweissString}}("ID"), iter.m.ID.Node(), nil
	case 1:
		iter.at++
		return {{.EdelweissString}}("Arg"), iter.m.Arg.Node(), nil
	}
	return nil, nil, {{.ErrBounds}}
}

func (iter *{{.TypeMapIterator}}) Done() bool {
	return iter.at >= 2
}`,
	}
}
