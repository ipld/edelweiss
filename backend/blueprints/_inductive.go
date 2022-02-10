package blueprints

import (
	"fmt"

	"github.com/ipld/edelweiss/backend/values"
	"github.com/ipld/go-ipld-prime/datamodel"
)

type InductiveXXX struct {
	C1XXX values.Int
	C2XXX values.Any
}

func (x *InductiveXXX) Parse(n datamodel.Node) error {
	if n.Kind() != datamodel.Kind_Map {
		return values.ErrNA
	}
	iter := n.MapIterator()
	nfields := 0
	for !iter.Done() {
		if kn, vn, err := iter.Next(); err != nil {
			return err
		} else {
			if k, err := kn.AsString(); err != nil {
				return fmt.Errorf("structure map key is not a string")
			} else {
				switch k {
				case "c1_XXX":
					if err := x.C1XXX.Parse(vn); err != nil {
						return err
					}
					nfields++
				case "c2_XXX":
					if err := x.C2XXX.Parse(vn); err != nil {
						return err
					}
					nfields++
				}
			}
		}
	}
	if nfields != 2 /*XXX*/ {
		return values.ErrNA
	} else {
		return nil
	}
}

type InductiveXXX_MapIterator struct {
	i int64
	s *InductiveXXX
}

func (x *InductiveXXX_MapIterator) Next() (key datamodel.Node, value datamodel.Node, err error) {
	x.i++
	switch x.i {
	case 0:
		return values.String("c1_XXX"), x.s.C1XXX.Node(), nil
	case 1:
		return values.String("c2_XXX"), x.s.C2XXX.Node(), nil
	}
	return nil, nil, values.ErrNA
}

func (x *InductiveXXX_MapIterator) Done() bool {
	return x.i < 2 /*XXX*/
}

func (x InductiveXXX) Kind() datamodel.Kind {
	return datamodel.Kind_Map
}

func (x InductiveXXX) LookupByString(key string) (datamodel.Node, error) {
	switch key {
	case "c1_XXX":
		return x.C1XXX.Node(), nil
	case "c2_XXX":
		return x.C2XXX.Node(), nil
	}
	return nil, values.ErrNA
}

func (x InductiveXXX) LookupByNode(key datamodel.Node) (datamodel.Node, error) {
	switch key.Kind() {
	case datamodel.Kind_String:
		if s, err := key.AsString(); err != nil {
			return nil, err
		} else {
			return x.LookupByString(s)
		}
	case datamodel.Kind_Int:
		if i, err := key.AsInt(); err != nil {
			return nil, err
		} else {
			return x.LookupByIndex(i)
		}
	}
	return nil, values.ErrNA
}

func (x InductiveXXX) LookupByIndex(idx int64) (datamodel.Node, error) {
	switch idx {
	case 0:
		return x.C1XXX.Node(), nil
	case 1:
		return x.C2XXX.Node(), nil
	}
	return nil, values.ErrNA
}

func (x InductiveXXX) LookupBySegment(seg datamodel.PathSegment) (datamodel.Node, error) {
	switch seg.String() {
	case "0", "c1_XXX":
		return x.C1XXX.Node(), nil
	case "1", "c2_XXX":
		return x.C2XXX.Node(), nil
	}
	return nil, values.ErrNA
}

func (x InductiveXXX) MapIterator() datamodel.MapIterator {
	return &InductiveXXX_MapIterator{-1, &x}
}

func (x InductiveXXX) ListIterator() datamodel.ListIterator {
	return nil
}

func (x InductiveXXX) Length() int64 {
	return 2 //XXX
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
