package values

import (
	"github.com/ipld/edelweiss/blueprints/base"
	cg "github.com/ipld/edelweiss/codegen"
	"github.com/ipld/edelweiss/plans"
)

func BuildInductiveImpl(
	lookup cg.LookupDepGoRef,
	plan plans.Inductive,
	goTypeRef cg.GoTypeRef,
) cg.GoTypeImpl {
	return &GoInductiveImpl{
		Lookup: lookup,
		Def:    plan,
		Ref:    goTypeRef,
	}
}

type GoInductiveImpl struct {
	Lookup cg.LookupDepGoRef
	Def    plans.Inductive
	Ref    cg.GoTypeRef
}

func (x *GoInductiveImpl) GoTypeRef() cg.GoTypeRef {
	return x.Ref
}

func (x *GoInductiveImpl) GoDef() cg.Blueprint {
	cases := x.Def.Cases
	caseData := make([]cg.BlueMap, len(cases))
	for i := range cases {
		goCaseName := cases[i].GoName
		if goCaseName == "" {
			goCaseName = cases[i].Name
		}
		caseData[i] = cg.BlueMap{
			"CaseName":        cg.V(goCaseName),
			"CaseNameString":  cg.StringLiteral(cases[i].Name),
			"CaseType":        x.Lookup.LookupDepGoRef(cases[i].Type),
			"EdelweissString": base.EdelweissString,
		}
	}
	// build case declarations
	caseDecls := make(cg.BlueSlice, len(cases))
	for i := range cases {
		caseDecls[i] = cg.T{
			Data: caseData[i],
			Src: "	{{.CaseName}} *{{.CaseType}}\n",
		}
	}
	// build case parse cases
	caseParseCases := make(cg.BlueSlice, len(cases))
	for i := range cases {
		caseParseCases[i] = cg.T{
			Data: caseData[i],
			Src: `case {{.CaseNameString}}:
		var y {{.CaseType}}
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.{{.CaseName}} = &y
		return nil
`,
		}
	}
	// build case next cases
	caseNextCases := make(cg.BlueSlice, len(cases))
	for i := range cases {
		caseNextCases[i] = cg.T{
			Data: caseData[i],
			Src: `		case x.s.{{.CaseName}} != nil:
			return {{.EdelweissString}}({{.CaseNameString}}), x.s.{{.CaseName}}.Node(), nil
`,
		}
	}
	// build case lookup by string cases
	caseLookupByStringCases := make(cg.BlueSlice, len(cases))
	for i := range cases {
		caseLookupByStringCases[i] = cg.T{
			Data: caseData[i],
			Src: `	case x.{{.CaseName}} != nil && key == {{.CaseNameString}}:
		return x.{{.CaseName}}.Node(), nil
`,
		}
	}
	// build case lookup by segment cases
	caseLookupBySegmentCases := make(cg.BlueSlice, len(cases))
	for i := range cases {
		caseLookupBySegmentCases[i] = cg.T{
			Data: caseData[i],
			Src: `	case {{.CaseNameString}}:
		return x.{{.CaseName}}.Node(), nil
`,
		}
	}

	// default case
	var defaultData cg.BlueMap
	if x.Def.Default.Type != nil {
		defaultData = cg.BlueMap{
			"DefaultKeyName":   cg.V(x.Def.Default.GoKeyName),
			"DefaultValueName": cg.V(x.Def.Default.GoValueName),
			"DefaultType":      x.Lookup.LookupDepGoRef(x.Def.Default.Type),
			"EdelweissString":  base.EdelweissString,
		}
	}

	// build type definition
	data := cg.BlueMap{
		"Type":            cg.V(x.Ref.TypeName),
		"Node":            base.IPLDNodeType,
		"KindType":        base.IPLDKindType,
		"KindMap":         base.IPLDKindMap,
		"KindString":      base.IPLDKindString,
		"KindInt":         base.IPLDKindInt,
		"ErrNA":           base.EdelweissErrNA,
		"PathSegment":     base.IPLDPathSegment,
		"MapIterator":     base.IPLDMapIteratorType,
		"ListIterator":    base.IPLDListIteratorType,
		"Link":            base.IPLDLinkType,
		"NodePrototype":   base.IPLDNodePrototypeType,
		"EdelweissString": base.EdelweissString,
		"Errorf":          base.Errorf,
		// cases
		"CaseDecls":                caseDecls,
		"CaseParseCases":           caseParseCases,
		"CaseNextCases":            caseNextCases,
		"CaseLookupByStringCases":  caseLookupByStringCases,
		"CaseLookupBySegmentCases": caseLookupBySegmentCases,
		// default case, if present
		"Default": defaultData,
	}
	return cg.T{
		Data: data,
		Src: `
// -- protocol type {{.Type}} --

type {{.Type}} struct {
{{range .CaseDecls}}	{{.}}{{end}}
{{with .Default}}
		{{.DefaultKeyName}} string
		{{.DefaultValueName}} *{{.DefaultType}}
{{end}}
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
{{range .CaseParseCases}}	{{.}}{{end}}
{{with .Default}}
	default:
		var y {{.DefaultType}}
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.{{.DefaultKeyName}} = k
		x.{{.DefaultValueName}} = &y
		return nil
{{end}}
	}
	return {{.Errorf}}("inductive map has no applicable keys")
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
{{range .CaseNextCases}}	{{.}}{{end}}
{{with .Default}}
	case x.s.{{.DefaultValueName}} != nil:
		return {{.EdelweissString}}(x.s.{{.DefaultKeyName}}), x.s.{{.DefaultValueName}}.Node(), nil
{{end}}
		default:
			return nil, nil, {{.Errorf}}("no inductive cases are set")
		}
	}
}

func (x *{{.Type}}_MapIterator) Done() bool {
	return x.done
}

func (x {{.Type}}) Node() {{.Node}} {
	return x
}

func (x {{.Type}}) Kind() {{.KindType}} {
	return {{.KindMap}}
}

func (x {{.Type}}) LookupByString(key string) ({{.Node}}, error) {
	switch {
{{range .CaseLookupByStringCases}}	{{.}}{{end}}
{{with .Default}}
	case x.{{.DefaultValueName}} != nil && key == x.{{.DefaultKeyName}}:
		return x.{{.DefaultValueName}}.Node(), nil
{{end}}
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
{{range .CaseLookupBySegmentCases}}	{{.}}{{end}}
{{with .Default}}
	case x.{{.DefaultKeyName}}:
		return x.{{.DefaultValueName}}.Node(), nil
{{end}}
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
