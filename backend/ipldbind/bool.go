package ipldbind

import (
	"github.com/ipld/edelweiss/def"
	"github.com/ipld/go-ipld-prime/datamodel"
)

type Bool bool

func (v Bool) IPLDNode() datamodel.Node {
	return v
}

// backend.Type implementation

func (Bool) Type() Type {
	return BoolType{}
}

// datamodel.Node implementation

func (Bool) Kind() datamodel.Kind {
	return datamodel.Kind_Bool
}

func (Bool) LookupByString(string) (datamodel.Node, error) {
	return nil, ErrNA
}

func (Bool) LookupByNode(key datamodel.Node) (datamodel.Node, error) {
	return nil, ErrNA
}

func (Bool) LookupByIndex(idx int64) (datamodel.Node, error) {
	return nil, ErrNA
}

func (Bool) LookupBySegment(seg datamodel.PathSegment) (datamodel.Node, error) {
	return nil, ErrNA
}

func (Bool) MapIterator() datamodel.MapIterator {
	return nil
}

func (Bool) ListIterator() datamodel.ListIterator {
	return nil
}

func (Bool) Length() int64 {
	return -1
}

func (Bool) IsAbsent() bool {
	return false
}

func (Bool) IsNull() bool {
	return false
}

func (v Bool) AsBool() (bool, error) {
	return bool(v), nil
}

func (Bool) AsInt() (int64, error) {
	return 0, ErrNA
}

func (Bool) AsFloat() (float64, error) {
	return 0, ErrNA
}

func (Bool) AsString() (string, error) {
	return "", ErrNA
}

func (Bool) AsBytes() ([]byte, error) {
	return nil, ErrNA
}

func (Bool) AsLink() (datamodel.Link, error) {
	return nil, ErrNA
}

func (Bool) Prototype() datamodel.NodePrototype {
	return BoolType{}
}

// datamodel.NodeAssembler implementation

func (*Bool) BeginMap(sizeHint int64) (datamodel.MapAssembler, error) {
	return nil, ErrNA
}

func (*Bool) BeginList(sizeHint int64) (datamodel.ListAssembler, error) {
	return nil, ErrNA
}

func (*Bool) AssignNull() error {
	return ErrNA
}

func (x *Bool) AssignBool(b bool) error {
	*(*bool)(x) = b
	return nil
}

func (*Bool) AssignInt(int64) error {
	return ErrNA
}

func (*Bool) AssignFloat(float64) error {
	return ErrNA
}

func (*Bool) AssignString(string) error {
	return ErrNA
}

func (*Bool) AssignBytes([]byte) error {
	return ErrNA
}

func (*Bool) AssignLink(datamodel.Link) error {
	return ErrNA
}

func (x *Bool) AssignNode(n datamodel.Node) error {
	if b, err := n.AsBool(); err != nil {
		return err
	} else {
		return x.AssignBool(b)
	}
}

// datamodel.NodeBuilder implementation

func (x *Bool) Build() datamodel.Node {
	return *x
}

func (x *Bool) Reset() {
	*x = false
}

type BoolType struct{}

func (BoolType) Def() def.Type {
	return def.Bool{}
}

func (BoolType) NewBuilder() datamodel.NodeBuilder {
	var v Bool
	return &v
}
