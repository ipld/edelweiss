package blueprint

type singletonBlueprintData struct {
	TypeName         string
	IPLDKindName     string
	IPLDAsMethodName string
	IPLDValueLiteral string
	AsBoolReturns    string
	AsIntReturns     string
	AsFloatReturns   string
	AsStringReturns  string
}

const singletonBlueprint = `
type {{.TypeName}} struct{}

func ({{.TypeName}}) Parse(n datamodel.Node) error {
	if n.Kind() != ipld.{{.IPLDKindName}} {
		return ErrNA
	}
	if n.{{.IPLDAsMethodName}}() != {{.IPLDValueLiteral}} {
		return ErrNA
	}
	return nil
}

func (v {{.TypeName}}) Node() datamodel.Node {
	return v
}

func ({{.TypeName}}) Kind() datamodel.Kind {
	return datamodel.{{.IPLDKindName}}
}

func ({{.TypeName}}) LookupByString(string) (datamodel.Node, error) {
	return nil, ErrNA
}

func ({{.TypeName}}) LookupByNode(key datamodel.Node) (datamodel.Node, error) {
	return nil, ErrNA
}

func ({{.TypeName}}) LookupByIndex(idx int64) (datamodel.Node, error) {
	return nil, ErrNA
}

func ({{.TypeName}}) LookupBySegment(seg datamodel.PathSegment) (datamodel.Node, error) {
	return nil, ErrNA
}

func ({{.TypeName}}) MapIterator() datamodel.MapIterator {
	return nil
}

func ({{.TypeName}}) ListIterator() datamodel.ListIterator {
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

func ({{.TypeName}}) AsLink() (datamodel.Link, error) {
	return nil, ErrNA
}

func ({{.TypeName}}) Prototype() datamodel.NodePrototype {
	return nil
}
`
