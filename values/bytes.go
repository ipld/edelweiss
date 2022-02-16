package values

import (
	"github.com/ipld/edelweiss/def"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/datamodel"
)

type Bytes []byte

func (Bytes) Def() def.Type {
	return def.Bytes{}
}

func (v *Bytes) Parse(n datamodel.Node) error {
	if n.Kind() != ipld.Kind_Bytes {
		return ErrNA
	} else {
		*(*[]byte)(v), _ = n.AsBytes()
		return nil
	}
}

func (v Bytes) Node() datamodel.Node {
	return v
}

// datamodel.Node implementation

func (Bytes) Kind() datamodel.Kind {
	return datamodel.Kind_Bytes
}

func (Bytes) LookupByString(string) (datamodel.Node, error) {
	return nil, ErrNA
}

func (Bytes) LookupByNode(key datamodel.Node) (datamodel.Node, error) {
	return nil, ErrNA
}

func (Bytes) LookupByIndex(idx int64) (datamodel.Node, error) {
	return nil, ErrNA
}

func (Bytes) LookupBySegment(seg datamodel.PathSegment) (datamodel.Node, error) {
	return nil, ErrNA
}

func (Bytes) MapIterator() datamodel.MapIterator {
	return nil
}

func (Bytes) ListIterator() datamodel.ListIterator {
	return nil
}

func (Bytes) Length() int64 {
	return -1
}

func (Bytes) IsAbsent() bool {
	return false
}

func (Bytes) IsNull() bool {
	return false
}

func (Bytes) AsBool() (bool, error) {
	return false, ErrNA
}

func (Bytes) AsInt() (int64, error) {
	return 0, ErrNA
}

func (Bytes) AsFloat() (float64, error) {
	return 0, ErrNA
}

func (Bytes) AsString() (string, error) {
	return "", ErrNA
}

func (v Bytes) AsBytes() ([]byte, error) {
	return v, nil
}

func (Bytes) AsLink() (datamodel.Link, error) {
	return nil, ErrNA
}

func (Bytes) Prototype() datamodel.NodePrototype {
	return nil // not needed
}

func TryParseBytes(n datamodel.Node) (Bytes, error) {
	var b Bytes
	return b, b.Parse(n)
}
