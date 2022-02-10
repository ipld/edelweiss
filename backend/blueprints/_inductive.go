package blueprints

import (
	"fmt"

	"github.com/ipld/edelweiss/backend/values"
	"github.com/ipld/go-ipld-prime/datamodel"
)

type InductiveXXX struct {
	C1XXX *values.Int
	C2XXX *values.Any
}

func (x *InductiveXXX) Parse(n datamodel.Node) error {
	*x = InductiveXXX{}
	if n.Kind() != datamodel.Kind_Map {
		return values.ErrNA
	}
	iter := n.MapIterator()
	kn, vn, err := iter.Next()
	if err != nil {
		return err
	}
	k, err := kn.AsString()
	if err != nil {
		return fmt.Errorf("inductive map key is not a string")
	}
	switch k {
	case "c1_XXX":
		var y T1XXX
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.C1XXX = &y
	}
	return nil
}

type InductiveXXX_MapIterator struct {
	done bool
	s    *InductiveXXX
}

func (x *InductiveXXX_MapIterator) Next() (key datamodel.Node, value datamodel.Node, err error) {
	if x.done {
		return nil, nil, values.ErrNA
	} else {
		x.done = true
		switch {
		case x.C1XXX != nil:
			return values.String("c1_XXX"), x.s.C1XXX.Node(), nil
		default:
			return nil, nil, fmt.Errorf("no inductive cases are set")
		}
	}
}

func (x *InductiveXXX_MapIterator) Done() bool {
	return x.done
}

func (x InductiveXXX) Kind() datamodel.Kind {
	return datamodel.Kind_Map
}

func (x InductiveXXX) LookupByString(key string) (datamodel.Node, error) {
	switch key {
	case x.C1XXX != nil && key == "c1_XXX":
		return x.C1XXX.Node(), nil
	}
	return nil, values.ErrNA
}

func (x InductiveXXX) LookupByNode(key datamodel.Node) (datamodel.Node, error) {
	if key.Kind() != datamodel.Kind_String {
		return nil, values.ErrNA
	}
	if s, err := key.AsString(); err != nil {
		return nil, err
	} else {
		return x.LookupByString(s)
	}
}

func (x InductiveXXX) LookupByIndex(idx int64) (datamodel.Node, error) {
	return nil, values.ErrNA
}

func (x InductiveXXX) LookupBySegment(seg datamodel.PathSegment) (datamodel.Node, error) {
	switch seg.String() {
	case "c1_XXX":
		return x.C1XXX.Node(), nil
	}
	return nil, values.ErrNA
}

func (x InductiveXXX) MapIterator() datamodel.MapIterator {
	return &InductiveXXX_MapIterator{false, &x}
}

func (x InductiveXXX) ListIterator() datamodel.ListIterator {
	return nil
}

func (x InductiveXXX) Length() int64 {
	return 1
}

func (x InductiveXXX) IsAbsent() bool {
	return false
}

func (x InductiveXXX) IsNull() bool {
	return false
}

func (x InductiveXXX) AsBool() (bool, error) {
	return false, values.ErrNA
}

func (x InductiveXXX) AsInt() (int64, error) {
	return 0, values.ErrNA
}

func (x InductiveXXX) AsFloat() (float64, error) {
	return 0, values.ErrNA
}

func (x InductiveXXX) AsString() (string, error) {
	return "", values.ErrNA
}

func (x InductiveXXX) AsBytes() ([]byte, error) {
	return nil, values.ErrNA
}

func (x InductiveXXX) AsLink() (datamodel.Link, error) {
	return nil, values.ErrNA
}

func (x InductiveXXX) Prototype() datamodel.NodePrototype {
	return nil
}
