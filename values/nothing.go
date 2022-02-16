package values

import (
	"github.com/ipld/edelweiss/def"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/datamodel"
)

type Nothing struct{}

func (Nothing) Def() def.Type {
	return def.Nothing{}
}

func (v *Nothing) Parse(n datamodel.Node) error {
	if n.Kind() != ipld.Kind_Null {
		return ErrNA
	} else {
		return nil
	}
}

func (v Nothing) Node() datamodel.Node {
	return v
}

// datamodel.Node implementation

func (Nothing) Kind() datamodel.Kind {
	return datamodel.Kind_Null
}

func (Nothing) LookupByString(string) (datamodel.Node, error) {
	return nil, ErrNA
}

func (Nothing) LookupByNode(key datamodel.Node) (datamodel.Node, error) {
	return nil, ErrNA
}

func (Nothing) LookupByIndex(idx int64) (datamodel.Node, error) {
	return nil, ErrNA
}

func (Nothing) LookupBySegment(seg datamodel.PathSegment) (datamodel.Node, error) {
	return nil, ErrNA
}

func (Nothing) MapIterator() datamodel.MapIterator {
	return nil
}

func (Nothing) ListIterator() datamodel.ListIterator {
	return nil
}

func (Nothing) Length() int64 {
	return -1
}

func (Nothing) IsAbsent() bool {
	return false
}

func (Nothing) IsNull() bool {
	return true
}

func (v Nothing) AsBool() (bool, error) {
	return false, ErrNA
}

func (Nothing) AsInt() (int64, error) {
	return 0, ErrNA
}

func (Nothing) AsFloat() (float64, error) {
	return 0, ErrNA
}

func (Nothing) AsString() (string, error) {
	return "", ErrNA
}

func (Nothing) AsBytes() ([]byte, error) {
	return nil, ErrNA
}

func (Nothing) AsLink() (datamodel.Link, error) {
	return nil, ErrNA
}

func (Nothing) Prototype() datamodel.NodePrototype {
	return nil // not needed
}

func TryParseNothing(n datamodel.Node) (Nothing, error) {
	var nth Nothing
	return nth, nth.Parse(n)
}
