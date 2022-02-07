package blueprints

import (
	cg "github.com/ipld/edelweiss/backend/codegen"
	"github.com/ipld/edelweiss/def"
)

func BuildStructureImpl(typeDef def.Structure, goTypeRef cg.GoTypeRef) (cg.GoTypeImpl, error) {
	return &GoStructureImpl{Def: typeDef, Ref: goTypeRef}, nil
}

type GoStructureImpl struct {
	Def def.Structure
	Ref cg.GoTypeRef
}

func (x *GoStructureImpl) ProtoDef() def.Type {
	return x.Def
}

func (x *GoStructureImpl) GoTypeRef() cg.GoTypeRef {
	return x.Ref
}

func (x *GoStructureImpl) GoDef() cg.Blueprint {
	fields := def.FlattenFieldList(x.Def.Fields)
	fieldDecls := make(cg.Blueprints, len(fields))
	for i := range fields {
		fieldDecls[i] = cg.T{
			Data: cg.BlueMap{
				"FieldName": cg.V(fields[i].Name),
				"FieldType": XXX,
			},
			Src: "	{{.FieldName}} {{.FieldType}}\n",
		}
	}
	data := cg.BlueMap{
		"Type":            x.Ref,
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
		"FieldDecls":      fieldDecls,
		"EdelweissString": EdelweissString,
	}
	return cg.T{
		Data: data,
		Src: `
// -- protocol type {{.Type}} --

type {{.Type}} struct {
{{.FieldDecls}}
}

func (x *{{.Type}}) Parse(n {{.Node}}) error {
	if n.Kind() != {{.KindMap}} {
		return {{.ErrNA}}
	}
	iter := n.MapIterator()
	nfields := 0
	for !iter.Done() {
		if kn, vn, err := iter.Next(); err != nil {
			return err
		} else {
			if k, err := kn.AsString(); err != nil {
				return fmt.Errorf("structure map key is not a string")
			} else {
				switch k {
				case "f1_XXX":
					if err := x.F1XXX.Parse(vn); err != nil {
						return err
					}
					nfields++
				case "f2_XXX":
					if err := x.F2XXX.Parse(vn); err != nil {
						return err
					}
					nfields++
				}
			}
		}
	}
	if nfields != 2 /*XXX*/ {
		return {{.ErrNA}}
	} else {
		return nil
	}
}

type {{.Type}}_MapIterator struct {
	i int64
	s *{{.Type}}
}

func (x *{{.Type}}_MapIterator) Next() (key {{.Node}}, value {{.Node}}, err error) {
	x.i++
	switch x.i {
	case 0:
		return {{.EdelweissString}}("f1_XXX"), x.s.F1XXX.Node(), nil
	case 1:
		return {{.EdelweissString}}("f2_XXX"), x.s.F2XXX.Node(), nil
	}
	return nil, nil, {{.ErrNA}}
}

func (x *{{.Type}}_MapIterator) Done() bool {
	return x.i < {{.Length}}
}

func (x {{.Type}}) Kind() {{.KindType}} {
	return {{.KindMap}}
}

func (x {{.Type}}) LookupByString(key string) ({{.Node}}, error) {
	switch key {
	case "f1_XXX":
		return x.F1XXX.Node(), nil
	case "f2_XXX":
		return x.F2XXX.Node(), nil
	}
	return nil, {{.ErrNA}}
}

func (x {{.Type}}) LookupByNode(key {{.Node}}) ({{.Node}}, error) {
	switch key.Kind() {
	case {{.KindString}}:
		if s, err := key.AsString(); err != nil {
			return nil, err
		} else {
			return x.LookupByString(s)
		}
	case {{.KindInt}}:
		if i, err := key.AsInt(); err != nil {
			return nil, err
		} else {
			return x.LookupByIndex(i)
		}
	}
	return nil, {{.ErrNA}}
}

func (x {{.Type}}) LookupByIndex(idx int64) ({{.Node}}, error) {
	switch idx {
	case 0:
		return x.F1XXX.Node(), nil
	case 1:
		return x.F2XXX.Node(), nil
	}
	return nil, {{.ErrNA}}
}

func (x {{.Type}}) LookupBySegment(seg datamodel.PathSegment) ({{.Node}}, error) {
	switch seg.String() {
	case "0", "f1_XXX":
		return x.F1XXX.Node(), nil
	case "1", "f2_XXX":
		return x.F2XXX.Node(), nil
	}
	return nil, {{.ErrNA}}
}

func (x {{.Type}}) MapIterator() {{.MapIterator}} {
	return &{{.Type}}_MapIterator{-1, &x}
}

func (x {{.Type}}) ListIterator() {{.ListIterator}} {
	return nil
}

func (x {{.Type}}) Length() int64 {
	return {{.Length}}
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
}
`,
	}
}
