package values

import (
	"github.com/ipld/edelweiss/def"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/datamodel"
)

type Int int64

func (Int) Def() def.Type {
	return def.Int{}
}

func (v *Int) Parse(n datamodel.Node) error {
	if n.Kind() != ipld.Kind_Int {
		return ErrNA
	} else {
		*(*int64)(v), _ = n.AsInt()
		return nil
	}
}

func (v Int) Node() datamodel.Node {
	return v
}

// datamodel.Node implementation

func (Int) Kind() datamodel.Kind {
	return datamodel.Kind_Int
}

func (Int) LookupByString(string) (datamodel.Node, error) {
	return nil, ErrNA
}

func (Int) LookupByNode(key datamodel.Node) (datamodel.Node, error) {
	return nil, ErrNA
}

func (Int) LookupByIndex(idx int64) (datamodel.Node, error) {
	return nil, ErrNA
}

func (Int) LookupBySegment(seg datamodel.PathSegment) (datamodel.Node, error) {
	return nil, ErrNA
}

func (Int) MapIterator() datamodel.MapIterator {
	return nil
}

func (Int) ListIterator() datamodel.ListIterator {
	return nil
}

func (Int) Length() int64 {
	return -1
}

func (Int) IsAbsent() bool {
	return false
}

func (Int) IsNull() bool {
	return false
}

func (Int) AsBool() (bool, error) {
	return false, ErrNA
}

func (v Int) AsInt() (int64, error) {
	return int64(v), nil
}

func (Int) AsFloat() (float64, error) {
	return 0, ErrNA
}

func (Int) AsString() (string, error) {
	return "", ErrNA
}

func (Int) AsBytes() ([]byte, error) {
	return nil, ErrNA
}

func (Int) AsLink() (datamodel.Link, error) {
	return nil, ErrNA
}

func (Int) Prototype() datamodel.NodePrototype {
	return nil // not needed
}

func TryParseInt(n datamodel.Node) (Int, error) {
	var x Int
	return x, x.Parse(n)
}
