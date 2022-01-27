package values

import (
	"github.com/ipld/edelweiss/def"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/datamodel"
)

type Char rune

func (Char) Def() def.Type {
	return def.Char{}
}

func (v *Char) Parse(n datamodel.Node) error {
	if n.Kind() != ipld.Kind_Int {
		return ErrNA
	} else {
		if i, _ := n.AsInt(); int64(rune(i)) != i {
			return ErrNA
		} else {
			*(*rune)(v) = rune(i)
			return nil
		}
	}
}

func (v Char) Node() datamodel.Node {
	return v
}

// datamodel.Node implementation

func (Char) Kind() datamodel.Kind {
	return datamodel.Kind_Int
}

func (Char) LookupByString(string) (datamodel.Node, error) {
	return nil, ErrNA
}

func (Char) LookupByNode(key datamodel.Node) (datamodel.Node, error) {
	return nil, ErrNA
}

func (Char) LookupByIndex(idx int64) (datamodel.Node, error) {
	return nil, ErrNA
}

func (Char) LookupBySegment(seg datamodel.PathSegment) (datamodel.Node, error) {
	return nil, ErrNA
}

func (Char) MapIterator() datamodel.MapIterator {
	return nil
}

func (Char) ListIterator() datamodel.ListIterator {
	return nil
}

func (Char) Length() int64 {
	return -1
}

func (Char) IsAbsent() bool {
	return false
}

func (Char) IsNull() bool {
	return false
}

func (Char) AsBool() (bool, error) {
	return false, ErrNA
}

func (v Char) AsInt() (int64, error) {
	return int64(v), nil
}

func (Char) AsFloat() (float64, error) {
	return 0, ErrNA
}

func (Char) AsString() (string, error) {
	return "", ErrNA
}

func (Char) AsBytes() ([]byte, error) {
	return nil, ErrNA
}

func (Char) AsLink() (datamodel.Link, error) {
	return nil, ErrNA
}

func (Char) Prototype() datamodel.NodePrototype {
	return nil // not needed
}

func TryParseChar(n datamodel.Node) (Char, error) {
	var x Char
	return x, x.Parse(n)
}
