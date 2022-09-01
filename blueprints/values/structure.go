package values

import (
	"strconv"

	"github.com/ipld/edelweiss/blueprints/base"
	cg "github.com/ipld/edelweiss/codegen"
	"github.com/ipld/edelweiss/plans"
)

func BuildStructureImpl(
	lookup cg.LookupDepGoRef,
	plan plans.Structure,
	goTypeRef cg.GoTypeRef,
) cg.GoTypeImpl {
	return &GoStructureImpl{
		Lookup: lookup,
		Def:    plan,
		Ref:    goTypeRef,
	}
}

type GoStructureImpl struct {
	Lookup cg.LookupDepGoRef
	Def    plans.Structure
	Ref    cg.GoTypeRef
}

func (x *GoStructureImpl) GoTypeRef() cg.GoTypeRef {
	return x.Ref
}

func (x *GoStructureImpl) GoDef() cg.Blueprint {
	fields := x.Def.Fields
	fieldData := make([]cg.BlueMap, len(fields))
	for i := range fields {
		goFieldName := fields[i].GoName
		if goFieldName == "" {
			goFieldName = fields[i].Name
		}
		fieldData[i] = cg.BlueMap{
			"FieldIndex":       cg.IntLiteral(i),
			"FieldIndexString": cg.StringLiteral(strconv.Itoa(i)),
			"FieldName":        cg.V(goFieldName),
			"FieldNameString":  cg.StringLiteral(fields[i].Name),
			"FieldType":        x.Lookup.LookupDepGoRef(fields[i].Type),
			"EdelweissString":  base.EdelweissString,
			"Errorf":           base.Errorf,
		}
	}
	// build field declarations
	fieldDecls := make(cg.BlueSlice, len(fields))
	for i := range fields {
		fieldDecls[i] = cg.T{
			Data: fieldData[i],
			Src: "	{{.FieldName}} {{.FieldType}}\n",
		}
	}
	// build field parse cases
	fieldParseCases := make(cg.BlueSlice, len(fields))
	fieldParseMapCases := make(cg.BlueSlice, len(fields))
	for i := range fields {
		fieldParseMapCases[i] = cg.T{
			Data: fieldData[i],
			Src: "		{{.FieldNameString}}: x.{{.FieldName}}.Parse,\n",
		}
		fieldParseCases[i] = cg.T{
			Data: fieldData[i],
			Src: `		case {{.FieldNameString}}:
			if _, notParsed := fieldMap[{{.FieldNameString}}]; !notParsed {
				return {{.Errorf}}("field %s already parsed", {{.FieldNameString}})
			}
			if err := x.{{.FieldName}}.Parse(vn); err != nil {
				return err
			}
			delete(fieldMap, {{.FieldNameString}})
`,
		}
	}
	// build field next cases
	fieldNextCases := make(cg.BlueSlice, len(fields))
	for i := range fields {
		fieldNextCases[i] = cg.T{
			Data: fieldData[i],
			Src: `		case {{.FieldIndex}}:
			return {{.EdelweissString}}({{.FieldNameString}}), x.s.{{.FieldName}}.Node(), nil
`,
		}
	}
	// build field lookup by string cases
	fieldLookupByStringCases := make(cg.BlueSlice, len(fields))
	for i := range fields {
		fieldLookupByStringCases[i] = cg.T{
			Data: fieldData[i],
			Src: `	case {{.FieldNameString}}:
		return x.{{.FieldName}}.Node(), nil
`,
		}
	}
	// build field lookup by index cases
	fieldLookupByIndexCases := make(cg.BlueSlice, len(fields))
	for i := range fields {
		fieldLookupByIndexCases[i] = cg.T{
			Data: fieldData[i],
			Src: `	case {{.FieldIndex}}:
		return x.{{.FieldName}}.Node(), nil
`,
		}
	}
	// build field lookup by segment cases
	fieldLookupBySegmentCases := make(cg.BlueSlice, len(fields))
	for i := range fields {
		fieldLookupBySegmentCases[i] = cg.T{
			Data: fieldData[i],
			Src: `	case {{.FieldIndexString}}, {{.FieldNameString}}:
		return x.{{.FieldName}}.Node(), nil
`,
		}
	}
	// build type definition
	data := cg.BlueMap{
		"Type":               x.Ref,
		"Node":               base.IPLDNodeType,
		"Null":               base.IPLDNull,
		"KindType":           base.IPLDKindType,
		"KindMap":            base.IPLDKindMap,
		"KindString":         base.IPLDKindString,
		"KindInt":            base.IPLDKindInt,
		"ErrNA":              base.EdelweissErrNA,
		"PathSegment":        base.IPLDPathSegment,
		"MapIterator":        base.IPLDMapIteratorType,
		"ListIterator":       base.IPLDListIteratorType,
		"Link":               base.IPLDLinkType,
		"NodePrototype":      base.IPLDNodePrototypeType,
		"Length":             cg.IntLiteral(len(fields)),
		"EdelweissString":    base.EdelweissString,
		"EdelweissParseFunc": base.EdelweissParseFunc,
		"Errorf":             base.Errorf,
		//
		"FieldDecls":                fieldDecls,
		"FieldParseMapCases":        fieldParseMapCases,
		"FieldParseCases":           fieldParseCases,
		"FieldNextCases":            fieldNextCases,
		"FieldLookupByStringCases":  fieldLookupByStringCases,
		"FieldLookupByIndexCases":   fieldLookupByIndexCases,
		"FieldLookupBySegmentCases": fieldLookupBySegmentCases,
	}
	return cg.T{
		Data: data,
		Src: `
// -- protocol type {{.Type}} --

type {{.Type}} struct {
{{range .FieldDecls}}	{{.}}{{end}}
}

func (x {{.Type}}) Node() {{.Node}} {
	return x
}

func (x *{{.Type}}) Parse(n {{.Node}}) error {
	if n.Kind() != {{.KindMap}} {
		return {{.ErrNA}}
	}
	iter := n.MapIterator()
	fieldMap := map[string]{{.EdelweissParseFunc}}{
		{{range .FieldParseMapCases}}{{.}}{{end}}
	}
	for !iter.Done() {
		if kn, vn, err := iter.Next(); err != nil {
			return err
		} else {
			if k, err := kn.AsString(); err != nil {
				return {{.Errorf}}("structure map key is not a string")
			} else {
				_ = vn
				switch k {
{{range .FieldParseCases}}	{{.}}{{end}}
				}
			}
		}
	}
	for _, fieldParse := range fieldMap {
		if err := fieldParse({{.Null}}); err != nil {
			return err
		}
	}
	return nil
}

type {{.Type}}_MapIterator struct {
	i int64
	s *{{.Type}}
}

func (x *{{.Type}}_MapIterator) Next() (key {{.Node}}, value {{.Node}}, err error) {
	x.i++
	switch x.i {
{{range .FieldNextCases}}	{{.}}{{end}}
	}
	return nil, nil, {{.ErrNA}}
}

func (x *{{.Type}}_MapIterator) Done() bool {
	return x.i + 1 >= {{.Length}}
}

func (x {{.Type}}) Kind() {{.KindType}} {
	return {{.KindMap}}
}

func (x {{.Type}}) LookupByString(key string) ({{.Node}}, error) {
	switch key {
{{range .FieldLookupByStringCases}}	{{.}}{{end}}
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
{{range .FieldLookupByIndexCases}}	{{.}}{{end}}
	}
	return nil, {{.ErrNA}}
}

func (x {{.Type}}) LookupBySegment(seg {{.PathSegment}}) ({{.Node}}, error) {
	switch seg.String() {
{{range .FieldLookupBySegmentCases}}	{{.}}{{end}}
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
