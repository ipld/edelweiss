package values

import (
	"github.com/ipld/edelweiss/blueprints/base"
	cg "github.com/ipld/edelweiss/codegen"
	"github.com/ipld/edelweiss/def"
)

func BuildLinkImpl(
	lookup cg.LookupDepGoRef,
	typeDef def.Link,
	goTypeRef cg.GoTypeRef,
) cg.GoTypeImpl {
	return &GoLinkImpl{
		Lookup: lookup,
		Def:    typeDef,
		Ref:    goTypeRef,
	}
}

type GoLinkImpl struct {
	Lookup cg.LookupDepGoRef
	Def    def.Link
	Ref    cg.GoTypeRef
}

func (x *GoLinkImpl) ProtoDef() def.Type {
	return x.Def
}

func (x *GoLinkImpl) GoTypeRef() cg.GoTypeRef {
	return x.Ref
}

func (x *GoLinkImpl) GoDef() cg.Blueprint {
	// build type definition
	data := cg.BlueMap{
		"Type":            x.Ref,
		"ToType":          x.Lookup.LookupDepGoRef(x.Def.To),
		"Node":            base.IPLDNodeType,
		"KindType":        base.IPLDKindType,
		"KindLink":        base.IPLDKindLink,
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
		"Cid":         cg.GoTypeRef{PkgPath: "github.com/ipfs/go-cid", TypeName: "Cid"},
		"IPLDCidLink": cg.GoTypeRef{PkgPath: "github.com/ipld/go-ipld-prime/linking/cid", TypeName: "Link"},
	}
	return cg.T{
		Data: data,
		Src: `
// -- protocol type {{.Type}} --

type {{.Type}} {{.Cid}}

func (v *{{.Type}}) Parse(n {{.Node}}) error {
	if n.Kind() != {{.KindLink}} {
		return {{.ErrNA}}
	} else {
		ipldLink, _ := n.AsLink()
		// TODO: Is there a more general way to convert ipld.Link interface into a concrete user object?
		cidLink, ok := ipldLink.({{.IPLDCidLink}})
		if !ok {
			return {{.Errorf}}("only cid links are supported")
		} else {
			*v = {{.Type}}(cidLink.Cid)
			return nil
		}
	}
}

func (v {{.Type}}) Node() {{.Node}} {
	return v
}

func ({{.Type}}) Kind() {{.KindType}} {
	return {{.KindLink}}
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

func ({{.Type}}) LookupBySegment(seg {{.PathSegment}}) ({{.Node}}, error) {
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

func ({{.Type}}) AsBool() (bool, error) {
	return false, {{.ErrNA}}
}

func (v {{.Type}}) AsInt() (int64, error) {
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

func (v {{.Type}}) AsLink() ({{.Link}}, error) {
	return {{.IPLDCidLink}}{Cid: {{.Cid}}(v)}, nil
}

func ({{.Type}}) Prototype() {{.NodePrototype}} {
	return nil // not needed
}`,
	}
}
