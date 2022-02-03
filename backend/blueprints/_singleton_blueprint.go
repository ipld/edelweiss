package blueprints

import "html/template"

type singletonBlueprint struct {
	// type dependent
	TypeName         string
	IPLDKindName     string
	IPLDAsMethodName string
	IPLDValueLiteral string
	AsBoolReturns    string
	AsIntReturns     string
	AsFloatReturns   string
	AsStringReturns  string
	// file context dependent
	IPLDPkgAlias      string
	DatamodelPkgAlias string
	ValuesPkgAlias    string
}

var singletonTemplateCompiled = template.Must(template.New("singleton").Parse(singletonTemplateSrc))

type IPLDRefs struct {
	KindType                  string
	DatamodelKindType         string
	DatamodelNodeType         string
	DatamodelMapIteratorType  string
	DatamodelListIteratorType string
}

type EdelweissRefs struct {
	ErrNA string
}

type singletonBlueprint2 struct {
	IPLD             IPLDRefs
	Edelweiss        EdelweissRefs
	IPLDKindValue    string
	IPLDAsMethodName string
	IPLDValueLiteral string
	AsBoolReturns    [2]string
	AsIntReturns     [2]string
	AsFloatReturns   [2]string
	AsStringReturns  [2]string
}

const singletonTemplateSrc = `
// -- protocol type {{.TypeName}} --

type {{.TypeName}} struct{}

func ({{.TypeName}}) Parse(n {{.IPLD.DatamodelNodeType}}) error {
	if n.Kind() != {{.IPLDKindValue}} {
		return {{.Edelweiss.ErrNA}}
	}
	if n.{{.IPLDAsMethodName}}() != {{.IPLDValueLiteral}} {
		return {{.Edelweiss.ErrNA}}
	}
	return nil
}

func (v {{.TypeName}}) Node() {{.IPLD.DatamodelNodeType}} {
	return v
}

func ({{.TypeName}}) Kind() {{.IPLD.DatamodelKindType}} {
	return {{.IPLDKindValue}}
}

func ({{.TypeName}}) LookupByString(string) ({{.IPLD.DatamodelNodeType}}, error) {
	return nil, {{.Edelweiss.ErrNA}}
}

func ({{.TypeName}}) LookupByNode(key {{.IPLD.DatamodelNodeType}}) ({{.IPLD.DatamodelNodeType}}, error) {
	return nil, {{.Edelweiss.ErrNA}}
}

func ({{.TypeName}}) LookupByIndex(idx int64) ({{.IPLD.DatamodelNodeType}}, error) {
	return nil, {{.Edelweiss.ErrNA}}
}

func ({{.TypeName}}) LookupBySegment(seg {{.DatamodelPkgAlias}}.PathSegment) ({{.IPLD.DatamodelNodeType}}, error) {
	return nil, {{.Edelweiss.ErrNA}}
}

func ({{.TypeName}}) MapIterator() {{.DatamodelPkgAlias}}.MapIterator {
	return nil
}

func ({{.TypeName}}) ListIterator() {{.DatamodelPkgAlias}}.ListIterator {
	return nil
}

func ({{.TypeName}}) Length() int64 {
	return -1
}

func ({{.TypeName}}) IsAbsent() bool {
	return false
}

func ({{.TypeName}}) IsNull() bool {
	return false
}

func (v {{.TypeName}}) AsBool() (bool, error) {
	return {{.AsBoolReturns}}
}

func ({{.TypeName}}) AsInt() (int64, error) {
	return {{.AsIntReturns}}
}

func ({{.TypeName}}) AsFloat() (float64, error) {
	return {{.AsFloatReturns}}
}

func ({{.TypeName}}) AsString() (string, error) {
	return {{.AsStringReturns}}
}

func ({{.TypeName}}) AsBytes() ([]byte, error) {
	return nil, {{.Edelweiss.ErrNA}}
}

func ({{.TypeName}}) AsLink() ({{.DatamodelPkgAlias}}.Link, error) {
	return nil, {{.Edelweiss.ErrNA}}
}

func ({{.TypeName}}) Prototype() {{.IPLD.DatamodelNodeType}}Prototype {
	return nil
}
`
