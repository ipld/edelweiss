package values

import (
	"github.com/ipld/edelweiss/blueprints/base"
	cg "github.com/ipld/edelweiss/codegen"
	"github.com/ipld/edelweiss/def"
)

func BuildListImpl(
	lookup cg.LookupDepGoRef,
	typeDef def.List,
	goTypeRef cg.GoTypeRef,
) (cg.GoTypeImpl, error) {
	return &GoListImpl{
		Lookup: lookup,
		Def:    typeDef,
		Ref:    goTypeRef,
	}, nil
}

type GoListImpl struct {
	Lookup cg.LookupDepGoRef
	Def    def.List
	Ref    cg.GoTypeRef
}

func (x *GoListImpl) ProtoDef() def.Type {
	return x.Def
}

func (x *GoListImpl) GoTypeRef() cg.GoTypeRef {
	return x.Ref
}

func (x *GoListImpl) GoDef() cg.Blueprint {
	// build type definition
	data := cg.BlueMap{
		"Type": x.Ref,
		"TypeListIterator": cg.GoTypeRef{
			PkgPath:  x.Ref.PkgPath,
			TypeName: x.Ref.TypeName + "_ListIterator",
		},
		"ElemType":        x.Lookup.LookupDepGoRef(x.Def.Element),
		"Node":            base.IPLDNodeType,
		"KindType":        base.IPLDKindType,
		"KindList":        base.IPLDKindList,
		"KindString":      base.IPLDKindString,
		"KindInt":         base.IPLDKindInt,
		"ErrNA":           base.EdelweissErrNA,
		"ErrBounds":       base.EdelweissErrBounds,
		"PathSegment":     base.IPLDPathSegment,
		"MapIterator":     base.IPLDMapIteratorType,
		"ListIterator":    base.IPLDListIteratorType,
		"Link":            base.IPLDLinkType,
		"NodePrototype":   base.IPLDNodePrototypeType,
		"EdelweissString": base.EdelweissString,
		"Errorf":          base.Errorf,
		//
	}
	return cg.T{
		Data: data,
		Src: `
// -- protocol type {{.Type}} --

type {{.Type}} []{{.ElemType}}

func (v {{.Type}}) Node() {{.Node}} {
	return v
}

func (v *{{.Type}}) Parse(n {{.Node}}) error {
	if n.Kind() != {{.KindList}} {
		return {{.ErrNA}}
	} else {
		*v = make({{.Type}}, n.Length())
		iter := n.ListIterator()
		for !iter.Done() {
			if i, n, err := iter.Next(); err != nil {
				return {{.ErrNA}}
			} else if err = (*v)[i].Parse(n); err != nil {
				return err
			}
		}
		return nil
	}
}

func ({{.Type}}) Kind() {{.KindType}} {
	return {{.KindList}}
}

func ({{.Type}}) LookupByString(string) ({{.Node}}, error) {
	return nil, {{.ErrNA}}
}

func ({{.Type}}) LookupByNode(key {{.Node}}) ({{.Node}}, error) {
	return nil, {{.ErrNA}}
}

func (v {{.Type}}) LookupByIndex(i int64) ({{.Node}}, error) {
	if i < 0 || i >= v.Length() {
		return nil, {{.ErrBounds}}
	} else {
		return v[i].Node(), nil
	}
}

func (v {{.Type}}) LookupBySegment(seg {{.PathSegment}}) ({{.Node}}, error) {
	if i, err := seg.Index(); err != nil {
		return nil, {{.ErrNA}}
	} else {
		return v.LookupByIndex(i)
	}
}

func ({{.Type}}) MapIterator() {{.MapIterator}} {
	return nil
}

func (v {{.Type}}) ListIterator() {{.ListIterator}} {
	return &{{.TypeListIterator}}{v, 0}
}

func (v {{.Type}}) Length() int64 {
	return int64(len(v))
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

func ({{.Type}}) Prototype() {{.NodePrototype}} {
	return nil // not needed
}

type {{.TypeListIterator}} struct {
	list {{.Type}}
	at   int64
}

func (iter *{{.TypeListIterator}}) Next() (int64, {{.Node}}, error) {
	if iter.Done() {
		return -1, nil, {{.ErrBounds}}
	}
	v := iter.list[iter.at]
	i := int64(iter.at)
	iter.at++
	return i, v.Node(), nil
}

func (iter *{{.TypeListIterator}}) Done() bool {
	return iter.at >= iter.list.Length()
}`,
	}
}
