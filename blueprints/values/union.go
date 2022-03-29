package values

import (
	"github.com/ipld/edelweiss/blueprints/base"
	cg "github.com/ipld/edelweiss/codegen"
	"github.com/ipld/edelweiss/plans"
)

func BuildUnionImpl(
	lookup cg.LookupDepGoRef,
	plan plans.Union,
	goTypeRef cg.GoTypeRef,
) cg.GoTypeImpl {
	return &GoUnionImpl{
		Lookup: lookup,
		Def:    plan,
		Ref:    goTypeRef,
	}
}

type GoUnionImpl struct {
	Lookup cg.LookupDepGoRef
	Def    plans.Union
	Ref    cg.GoTypeRef
}

func (x *GoUnionImpl) GoTypeRef() cg.GoTypeRef {
	return x.Ref
}

func (x *GoUnionImpl) GoDef() cg.Blueprint {
	cases := x.Def.Cases
	caseData := make([]cg.BlueMap, len(cases))
	for i := range cases {
		goCaseName := cases[i].GoName
		if goCaseName == "" {
			goCaseName = cases[i].Name
		}
		caseData[i] = cg.BlueMap{
			"CaseName":        cg.V(goCaseName),
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
	// build cases for Node() method
	nodeCases := make(cg.BlueSlice, len(cases))
	for i := range cases {
		nodeCases[i] = cg.T{
			Data: caseData[i],
			Src: `if x.{{.CaseName}} != nil {
		return x.{{.CaseName}}
	}
`,
		}
	}
	// build case parse cases
	parseCases := make(cg.BlueSlice, len(cases))
	for i := range cases {
		parseCases[i] = cg.T{
			Data: caseData[i],
			Src: `
			var  {{.CaseName}} {{.CaseType}}
			if err := {{.CaseName}}.Parse(n); err == nil {
				x.{{.CaseName}} = &{{.CaseName}}
				return nil
			}
`,
		}
	}
	// build node method proxy cases
	kindCases := make(cg.BlueSlice, len(cases))
	lookupByStringCases := make(cg.BlueSlice, len(cases))
	lookupByNodeCases := make(cg.BlueSlice, len(cases))
	lookupByIndexCases := make(cg.BlueSlice, len(cases))
	lookupBySegmentCases := make(cg.BlueSlice, len(cases))
	mapIteratorCases := make(cg.BlueSlice, len(cases))
	listIteratorCases := make(cg.BlueSlice, len(cases))
	lengthCases := make(cg.BlueSlice, len(cases))
	isAbsentCases := make(cg.BlueSlice, len(cases))
	isNullCases := make(cg.BlueSlice, len(cases))
	asBoolCases := make(cg.BlueSlice, len(cases))
	asIntCases := make(cg.BlueSlice, len(cases))
	asFloatCases := make(cg.BlueSlice, len(cases))
	asStringCases := make(cg.BlueSlice, len(cases))
	asBytesCases := make(cg.BlueSlice, len(cases))
	asLinkCases := make(cg.BlueSlice, len(cases))
	for i := range cases {
		kindCases[i] = cg.T{
			Data: caseData[i],
			Src:  "if x.{{.CaseName}} != nil { return x.{{.CaseName}}.Kind() }\n",
		}
		lookupByStringCases[i] = cg.T{
			Data: caseData[i],
			Src:  "if x.{{.CaseName}} != nil { return x.{{.CaseName}}.LookupByString(key) }\n",
		}
		lookupByNodeCases[i] = cg.T{
			Data: caseData[i],
			Src:  "if x.{{.CaseName}} != nil { return x.{{.CaseName}}.LookupByNode(key) }\n",
		}
		lookupByIndexCases[i] = cg.T{
			Data: caseData[i],
			Src:  "if x.{{.CaseName}} != nil { return x.{{.CaseName}}.LookupByIndex(idx) }\n",
		}
		lookupBySegmentCases[i] = cg.T{
			Data: caseData[i],
			Src:  "if x.{{.CaseName}} != nil { return x.{{.CaseName}}.LookupBySegment(seg) }\n",
		}
		mapIteratorCases[i] = cg.T{
			Data: caseData[i],
			Src:  "if x.{{.CaseName}} != nil { return x.{{.CaseName}}.MapIterator() }\n",
		}
		listIteratorCases[i] = cg.T{
			Data: caseData[i],
			Src:  "if x.{{.CaseName}} != nil { return x.{{.CaseName}}.ListIterator() }\n",
		}
		lengthCases[i] = cg.T{
			Data: caseData[i],
			Src:  "if x.{{.CaseName}} != nil { return x.{{.CaseName}}.Length() }\n",
		}
		isAbsentCases[i] = cg.T{
			Data: caseData[i],
			Src:  "if x.{{.CaseName}} != nil { return x.{{.CaseName}}.IsAbsent() }\n",
		}
		isNullCases[i] = cg.T{
			Data: caseData[i],
			Src:  "if x.{{.CaseName}} != nil { return x.{{.CaseName}}.IsNull() }\n",
		}
		asBoolCases[i] = cg.T{
			Data: caseData[i],
			Src:  "if x.{{.CaseName}} != nil { return x.{{.CaseName}}.AsBool() }\n",
		}
		asIntCases[i] = cg.T{
			Data: caseData[i],
			Src:  "if x.{{.CaseName}} != nil { return x.{{.CaseName}}.AsInt() }\n",
		}
		asFloatCases[i] = cg.T{
			Data: caseData[i],
			Src:  "if x.{{.CaseName}} != nil { return x.{{.CaseName}}.AsFloat() }\n",
		}
		asStringCases[i] = cg.T{
			Data: caseData[i],
			Src:  "if x.{{.CaseName}} != nil { return x.{{.CaseName}}.AsString() }\n",
		}
		asBytesCases[i] = cg.T{
			Data: caseData[i],
			Src:  "if x.{{.CaseName}} != nil { return x.{{.CaseName}}.AsBytes() }\n",
		}
		asLinkCases[i] = cg.T{
			Data: caseData[i],
			Src:  "if x.{{.CaseName}} != nil { return x.{{.CaseName}}.AsLink() }\n",
		}
	}
	// build type definition
	data := cg.BlueMap{
		"Type":          cg.V(x.Ref.TypeName),
		"Node":          base.IPLDNodeType,
		"Errorf":        base.Errorf,
		"KindType":      base.IPLDKindType,
		"KindInvalid":   base.IPLDKindInvalid,
		"PathSegment":   base.IPLDPathSegment,
		"MapIterator":   base.IPLDMapIteratorType,
		"ListIterator":  base.IPLDListIteratorType,
		"Link":          base.IPLDLinkType,
		"NodePrototype": base.IPLDNodePrototypeType,
		//
		"CaseDecls":  caseDecls,
		"NodeCases":  nodeCases,
		"ParseCases": parseCases,
		//
		"KindCases":            kindCases,
		"LookupByStringCases":  lookupByStringCases,
		"LookupByNodeCases":    lookupByNodeCases,
		"LookupByIndexCases":   lookupByIndexCases,
		"LookupBySegmentCases": lookupBySegmentCases,
		"MapIteratorCases":     mapIteratorCases,
		"ListIteratorCases":    listIteratorCases,
		"LengthCases":          lengthCases,
		"IsAbsentCases":        isAbsentCases,
		"IsNullCases":          isNullCases,
		"AsBoolCases":          asBoolCases,
		"AsIntCases":           asIntCases,
		"AsFloatCases":         asFloatCases,
		"AsStringCases":        asStringCases,
		"AsBytesCases":         asBytesCases,
		"AsLinkCases":          asLinkCases,
	}
	return cg.T{
		Data: data,
		Src: `
// -- protocol type {{.Type}} --

type {{.Type}} struct {
{{range .CaseDecls}}	{{.}}{{end}}
}

func (x *{{.Type}}) Parse(n {{.Node}}) error {
	*x = {{.Type}}{}
{{range .ParseCases}}	{{.}}{{end}}
	return {{.Errorf}}("no union cases parses")
}

func (x {{.Type}}) Node() {{.Node}} {
{{range .NodeCases}}	{{.}}{{end}}
	return nil
}

// proxy Node methods to active case

func (x {{.Type}}) Kind() {{.KindType}} {
{{range .KindCases}}	{{.}}{{end}}
	return {{.KindInvalid}}
}

func (x {{.Type}}) LookupByString(key string) ({{.Node}}, error) {
{{range .LookupByStringCases}}	{{.}}{{end}}
	return nil, {{.Errorf}}("no active union case found")
}

func (x {{.Type}}) LookupByNode(key {{.Node}}) ({{.Node}}, error) {
{{range .LookupByNodeCases}}	{{.}}{{end}}
	return nil, {{.Errorf}}("no active union case found")
}

func (x {{.Type}}) LookupByIndex(idx int64) ({{.Node}}, error) {
{{range .LookupByIndexCases}}	{{.}}{{end}}
	return nil, {{.Errorf}}("no active union case found")
}

func (x {{.Type}}) LookupBySegment(seg {{.PathSegment}}) ({{.Node}}, error) {
{{range .LookupBySegmentCases}}	{{.}}{{end}}
	return nil, {{.Errorf}}("no active union case found")
}

func (x {{.Type}}) MapIterator() {{.MapIterator}} {
{{range .MapIteratorCases}}	{{.}}{{end}}
	return nil
}

func (x {{.Type}}) ListIterator() {{.ListIterator}} {
{{range .ListIteratorCases}}	{{.}}{{end}}
	return nil
}

func (x {{.Type}}) Length() int64 {
{{range .LengthCases}}	{{.}}{{end}}
	return -1
}

func (x {{.Type}}) IsAbsent() bool {
{{range .IsAbsentCases}}	{{.}}{{end}}
	return false
}

func (x {{.Type}}) IsNull() bool {
{{range .IsNullCases}}	{{.}}{{end}}
	return false
}

func (x {{.Type}}) AsBool() (bool, error) {
{{range .AsBoolCases}}	{{.}}{{end}}
	return false, {{.Errorf}}("no active union case found")
}

func (x {{.Type}}) AsInt() (int64, error) {
{{range .AsIntCases}}	{{.}}{{end}}
	return 0, {{.Errorf}}("no active union case found")
}

func (x {{.Type}}) AsFloat() (float64, error) {
{{range .AsFloatCases}}	{{.}}{{end}}
	return 0.0, {{.Errorf}}("no active union case found")
}

func (x {{.Type}}) AsString() (string, error) {
{{range .AsStringCases}}	{{.}}{{end}}
	return "", {{.Errorf}}("no active union case found")
}

func (x {{.Type}}) AsBytes() ([]byte, error) {
{{range .AsBytesCases}}	{{.}}{{end}}
	return nil, {{.Errorf}}("no active union case found")
}

func (x {{.Type}}) AsLink() ({{.Link}}, error) {
{{range .AsLinkCases}}	{{.}}{{end}}
	return nil, {{.Errorf}}("no active union case found")
}

func (x {{.Type}}) Prototype() {{.NodePrototype}} {
	return nil
}`,
	}
}
