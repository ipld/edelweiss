package values

import (
	"github.com/ipld/edelweiss/blueprints/base"
	cg "github.com/ipld/edelweiss/codegen"
	"github.com/ipld/edelweiss/defs"
)

func BuildMapImpl(
	lookup cg.LookupDepGoRef,
	typeDef defs.Map,
	goTypeRef cg.GoTypeRef,
) cg.GoTypeImpl {
	return &GoMapImpl{
		Lookup: lookup,
		Def:    typeDef,
		Ref:    goTypeRef,
	}
}

type GoMapImpl struct {
	Lookup cg.LookupDepGoRef
	Def    defs.Map
	Ref    cg.GoTypeRef
}

func (x *GoMapImpl) ProtoDef() defs.Def {
	return x.Def
}

func (x *GoMapImpl) GoTypeRef() cg.GoTypeRef {
	return x.Ref
}

func (x *GoMapImpl) GoDef() cg.Blueprint {
	// build type definition
	data := cg.BlueMap{
		"Type": x.Ref,
		"TypeMapIterator": cg.GoTypeRef{
			PkgPath:  x.Ref.PkgPath,
			TypeName: x.Ref.TypeName + "_MapIterator",
		},
		"TypeKeyValue": cg.GoTypeRef{
			PkgPath:  x.Ref.PkgPath,
			TypeName: x.Ref.TypeName + "_KeyValue",
		},
		"KeyType":         x.Lookup.LookupDepGoRef(x.Def.Key),
		"ValueType":       x.Lookup.LookupDepGoRef(x.Def.Value),
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
		//
		"IPLDDeepEqual": base.IPLDDeepEqual,
	}
	return cg.T{
		Data: data,
		Src: `
// -- protocol type {{.Type}} --

type {{.Type}} []{{.TypeKeyValue}}

type {{.TypeKeyValue}} struct {
	Key   {{.KeyType}}
	Value {{.ValueType}}
}

func (v {{.Type}}) Node() {{.Node}} {
	return v
}

func (v *{{.Type}}) Parse(n {{.Node}}) error {
	if n.Kind() != {{.KindMap}} {
		return {{.ErrNA}}
	} else {
		iter := n.MapIterator()
		for !iter.Done() {
			kn, vn, _ := iter.Next()
			var kv {{.TypeKeyValue}}
			if err := kv.Key.Parse(kn); err != nil {
				return err
			}
			if err := kv.Value.Parse(vn); err != nil {
				return err
			}
			*v = append(*v, kv)
		}
		return nil
	}
}

func ({{.Type}}) Kind() {{.KindType}} {
	return {{.KindMap}}
}

func (v {{.Type}}) LookupByString(s string) ({{.Node}}, error) {
	return v.LookupByNode({{.EdelweissString}}(s))
}

func (v {{.Type}}) LookupByNode(key {{.Node}}) ({{.Node}}, error) {
	for _, kv := range v {
		if {{.IPLDDeepEqual}}(kv.Key.Node(), key) {
			return kv.Value.Node(), nil
		}
	}
	return nil, {{.ErrNotFound}}
}

func (v {{.Type}}) LookupByIndex(i int64) ({{.Node}}, error) {
	return v.LookupByNode({{.EdelweissInt}}(i))
}

func (v {{.Type}}) LookupBySegment(seg {{.PathSegment}}) ({{.Node}}, error) {
	if idx, err := seg.Index(); err != nil {
		return v.LookupByString(seg.String())
	} else {
		return v.LookupByIndex(idx)
	}
}

func (v {{.Type}}) MapIterator() {{.MapIterator}} {
	return &{{.TypeMapIterator}}{v, 0}
}

func (v {{.Type}}) ListIterator() {{.ListIterator}} {
	return nil
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

func ({{.Type}}) Prototype() {{.Node}}Prototype {
	return nil // not needed
}

type {{.TypeMapIterator}} struct {
	m  {{.Type}}
	at int64
}

func (iter *{{.TypeMapIterator}}) Next() ({{.Node}}, {{.Node}}, error) {
	if iter.Done() {
		return nil, nil, {{.ErrBounds}}
	}
	v := iter.m[iter.at]
	iter.at++
	return v.Key.Node(), v.Value.Node(), nil
}

func (iter *{{.TypeMapIterator}}) Done() bool {
	return iter.at >= int64(len(iter.m))
}`,
	}
}
