package values

import (
	"github.com/ipld/edelweiss/def"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/datamodel"
)

type String string

func (String) Def() def.Type {
	return def.String{}
}

func (v *String) Parse(n datamodel.Node) error {
	if n.Kind() != ipld.Kind_String {
		return ErrNA
	} else {
		// TODO: verify utf8
		*(*string)(v), _ = n.AsString()
		return nil
	}
}

func (v String) Node() datamodel.Node {
	return v
}

// datamodel.Node implementation

func (String) Kind() datamodel.Kind {
	return datamodel.Kind_String
}

func (String) LookupByString(string) (datamodel.Node, error) {
	return nil, ErrNA
}

func (String) LookupByNode(key datamodel.Node) (datamodel.Node, error) {
	return nil, ErrNA
}

func (String) LookupByIndex(idx int64) (datamodel.Node, error) {
	return nil, ErrNA
}

func (String) LookupBySegment(seg datamodel.PathSegment) (datamodel.Node, error) {
	return nil, ErrNA
}

func (String) MapIterator() datamodel.MapIterator {
	return nil
}

func (String) ListIterator() datamodel.ListIterator {
	return nil
}

func (String) Length() int64 {
	return -1
}

func (String) IsAbsent() bool {
	return false
}

func (String) IsNull() bool {
	return false
}

func (String) AsBool() (bool, error) {
	return false, ErrNA
}

func (String) AsInt() (int64, error) {
	return 0, ErrNA
}

func (String) AsFloat() (float64, error) {
	return 0, ErrNA
}

func (v String) AsString() (string, error) {
	return string(v), nil
}

func (String) AsBytes() ([]byte, error) {
	return nil, ErrNA
}

func (String) AsLink() (datamodel.Link, error) {
	return nil, ErrNA
}

func (String) Prototype() datamodel.NodePrototype {
	return nil // not needed
}

func TryParseString(n datamodel.Node) (String, error) {
	var b String
	return b, b.Parse(n)
}
