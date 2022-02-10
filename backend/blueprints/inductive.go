package blueprints

import (
	cg "github.com/ipld/edelweiss/backend/codegen"
	"github.com/ipld/edelweiss/def"
)

func BuildInductiveImpl(
	lookup cg.LookupDepGoRef,
	typeDef def.Inductive,
	goTypeRef cg.GoTypeRef,
) (cg.GoTypeImpl, error) {
	return &GoInductiveImpl{
		Lookup: lookup,
		Def:    typeDef,
		Ref:    goTypeRef,
	}, nil
}

type GoInductiveImpl struct {
	Lookup cg.LookupDepGoRef
	Def    def.Inductive
	Ref    cg.GoTypeRef
}

func (x *GoInductiveImpl) ProtoDef() def.Type {
	return x.Def
}

func (x *GoInductiveImpl) GoTypeRef() cg.GoTypeRef {
	return x.Ref
}

func (x *GoInductiveImpl) GoDef() cg.Blueprint {
	XXX
	// build type definition
	data := cg.BlueMap{
		"Type":            cg.V(x.Ref.TypeName),
		"Node":            IPLDNodeType,
		"KindType":        IPLDKindType,
		"KindMap":         IPLDKindMap,
		"KindString":      IPLDKindString,
		"KindInt":         IPLDKindInt,
		"ErrNA":           EdelweissErrNA,
		"PathSegment":     IPLDPathSegment,
		"MapIterator":     IPLDMapIteratorType,
		"ListIterator":    IPLDListIteratorType,
		"Link":            IPLDLinkType,
		"NodePrototype":   IPLDNodePrototypeType,
		"Length":          cg.IntLiteral(len(fields)),
		"EdelweissString": EdelweissString,
		"Errorf":          Errorf,
		//
		"CaseDecls": caseDecls,
	}
	return cg.T{
		Data: data,
		Src: `
// -- protocol type {{.Type}} --

type {{.Type}} struct {
	C1XXX *values.Int
	C2XXX *values.Any
}

func (x *{{.Type}}) Parse(n {{.Node}}) error {
	*x = {{.Type}}{}
	if n.Kind() != {{.KindMap}} {
		return {{.ErrNA}}
	}
	iter := n.MapIterator()
	kn, vn, err := iter.Next()
	if err != nil {
		return err
	}
	k, err := kn.AsString()
	if err != nil {
		return {{.Errorf}}("inductive map key is not a string")
	}
	switch k {
	case "c1_XXX":
		var y T1XXX
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.C1XXX = &y
	}
	return nil
}

type {{.Type}}_MapIterator struct {
	done bool
	s    *{{.Type}}
}

func (x *{{.Type}}_MapIterator) Next() (key {{.Node}}, value {{.Node}}, err error) {
	if x.done {
		return nil, nil, {{.ErrNA}}
	} else {
		x.done = true
		switch {
		case x.C1XXX != nil:
			return values.String("c1_XXX"), x.s.C1XXX.Node(), nil
		default:
			return nil, nil, fmt.Errorf("no inductive cases are set")
		}
	}
}

func (x *{{.Type}}_MapIterator) Done() bool {
	return x.done
}

func (x {{.Type}}) Kind() {{.KindType}} {
	return {{.KindMap}}
}

func (x {{.Type}}) LookupByString(key string) ({{.Node}}, error) {
	switch key {
	case x.C1XXX != nil && key == "c1_XXX":
		return x.C1XXX.Node(), nil
	}
	return nil, {{.ErrNA}}
}

func (x {{.Type}}) LookupByNode(key {{.Node}}) ({{.Node}}, error) {
	if key.Kind() != {{.KindString}} {
		return nil, {{.ErrNA}}
	}
	if s, err := key.AsString(); err != nil {
		return nil, err
	} else {
		return x.LookupByString(s)
	}
}

func (x {{.Type}}) LookupByIndex(idx int64) ({{.Node}}, error) {
	return nil, {{.ErrNA}}
}

func (x {{.Type}}) LookupBySegment(seg {{.PathSegment}}) ({{.Node}}, error) {
	switch seg.String() {
	case "c1_XXX":
		return x.C1XXX.Node(), nil
	}
	return nil, {{.ErrNA}}
}

func (x {{.Type}}) MapIterator() {{.MapIterator}} {
	return &{{.Type}}_MapIterator{false, &x}
}

func (x {{.Type}}) ListIterator() {{.ListIterator}} {
	return nil
}

func (x {{.Type}}) Length() int64 {
	return 1
}

func (x {{.Type}}) IsAbsent() bool {
	return false
}

func (x {{.Type}}) IsNull() bool {
	return false
}

func (x {{.Type}}) AsBool() (bool, error) {
	return false, {{.ErrNA}}
}

func (x {{.Type}}) AsInt() (int64, error) {
	return 0, {{.ErrNA}}
}

func (x {{.Type}}) AsFloat() (float64, error) {
	return 0, {{.ErrNA}}
}

func (x {{.Type}}) AsString() (string, error) {
	return "", {{.ErrNA}}
}

func (x {{.Type}}) AsBytes() ([]byte, error) {
	return nil, {{.ErrNA}}
}

func (x {{.Type}}) AsLink() ({{.Link}}, error) {
	return nil, {{.ErrNA}}
}

func (x {{.Type}}) Prototype() {{.NodePrototype}} {
	return nil
}`,
	}
}
