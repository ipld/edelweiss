package values

import (
	"github.com/ipld/edelweiss/def"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/datamodel"
)

type Float float64

func (Float) Def() def.Type {
	return def.Int{}
}

func (v *Float) Parse(n datamodel.Node) error {
	if n.Kind() != ipld.Kind_Float {
		return ErrNA
	} else {
		*(*float64)(v), _ = n.AsFloat()
		return nil
	}
}

func (v Float) Node() datamodel.Node {
	return v
}

// datamodel.Node implementation

func (Float) Kind() datamodel.Kind {
	return datamodel.Kind_Float
}

func (Float) LookupByString(string) (datamodel.Node, error) {
	return nil, ErrNA
}

func (Float) LookupByNode(key datamodel.Node) (datamodel.Node, error) {
	return nil, ErrNA
}

func (Float) LookupByIndex(idx int64) (datamodel.Node, error) {
	return nil, ErrNA
}

func (Float) LookupBySegment(seg datamodel.PathSegment) (datamodel.Node, error) {
	return nil, ErrNA
}

func (Float) MapIterator() datamodel.MapIterator {
	return nil
}

func (Float) ListIterator() datamodel.ListIterator {
	return nil
}

func (Float) Length() int64 {
	return -1
}

func (Float) IsAbsent() bool {
	return false
}

func (Float) IsNull() bool {
	return false
}

func (Float) AsBool() (bool, error) {
	return false, ErrNA
}

func (Float) AsInt() (int64, error) {
	return 0, ErrNA
}

func (v Float) AsFloat() (float64, error) {
	return float64(v), nil
}

func (Float) AsString() (string, error) {
	return "", ErrNA
}

func (Float) AsBytes() ([]byte, error) {
	return nil, ErrNA
}

func (Float) AsLink() (datamodel.Link, error) {
	return nil, ErrNA
}

func (Float) Prototype() datamodel.NodePrototype {
	return nil // not needed
}

func TryParseFloat(n datamodel.Node) (Float, error) {
	var x Float
	return x, x.Parse(n)
}
