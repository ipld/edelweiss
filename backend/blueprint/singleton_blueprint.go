package blueprint

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
}

var singletonTemplateCompiled = template.Must(template.New("singleton").Parse(singletonTemplateSrc))

const singletonTemplateSrc = `
// -- protocol type {{.TypeName}} --

type {{.TypeName}} struct{}

func ({{.TypeName}}) Parse(n {{.DatamodelPkgAlias}}.Node) error {
	if n.Kind() != {{.IPLDPkgAlias}}.{{.IPLDKindName}} {
		return ErrNA
	}
	if n.{{.IPLDAsMethodName}}() != {{.IPLDValueLiteral}} {
		return ErrNA
	}
	return nil
}

func (v {{.TypeName}}) Node() {{.DatamodelPkgAlias}}.Node {
	return v
}

func ({{.TypeName}}) Kind() {{.DatamodelPkgAlias}}.Kind {
	return {{.DatamodelPkgAlias}}.{{.IPLDKindName}}
}

func ({{.TypeName}}) LookupByString(string) ({{.DatamodelPkgAlias}}.Node, error) {
	return nil, ErrNA
}

func ({{.TypeName}}) LookupByNode(key {{.DatamodelPkgAlias}}.Node) ({{.DatamodelPkgAlias}}.Node, error) {
	return nil, ErrNA
}

func ({{.TypeName}}) LookupByIndex(idx int64) ({{.DatamodelPkgAlias}}.Node, error) {
	return nil, ErrNA
}

func ({{.TypeName}}) LookupBySegment(seg {{.DatamodelPkgAlias}}.PathSegment) ({{.DatamodelPkgAlias}}.Node, error) {
	return nil, ErrNA
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
	return nil, ErrNA
}

func ({{.TypeName}}) AsLink() ({{.DatamodelPkgAlias}}.Link, error) {
	return nil, ErrNA
}

func ({{.TypeName}}) Prototype() {{.DatamodelPkgAlias}}.NodePrototype {
	return nil
}
`
