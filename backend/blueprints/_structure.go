package blueprints

import (
	"fmt"

	"github.com/ipld/edelweiss/backend/values"
	"github.com/ipld/go-ipld-prime/datamodel"
)

type StructureXXX struct {
	F1XXX values.Int
	F2XXX values.Any
}

func (x *StructureXXX) Parse(n datamodel.Node) error {
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
				case "f1_XXX":
					if err := x.F1XXX.Parse(vn); err != nil {
						return err
					}
					nfields++
				case "f2_XXX":
					if err := x.F2XXX.Parse(vn); err != nil {
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

type StructureXXX_MapIterator struct {
	i int64
	s *StructureXXX
}

func (x *StructureXXX_MapIterator) Next() (key datamodel.Node, value datamodel.Node, err error) {
	x.i++
	switch x.i {
	case 0:
		return values.String("f1_XXX"), x.s.F1XXX.Node(), nil
	case 1:
		return values.String("f2_XXX"), x.s.F2XXX.Node(), nil
	}
	return nil, nil, values.ErrNA
}

func (x *StructureXXX_MapIterator) Done() bool {
	return x.i < 2 /*XXX*/
}

func (x StructureXXX) Kind() datamodel.Kind {
	return datamodel.Kind_Map
}

func (x StructureXXX) LookupByString(key string) (datamodel.Node, error) {
	switch key {
	case "f1_XXX":
		return x.F1XXX.Node(), nil
	case "f2_XXX":
		return x.F2XXX.Node(), nil
	}
	return nil, values.ErrNA
}

func (x StructureXXX) LookupByNode(key datamodel.Node) (datamodel.Node, error) {
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

func (x StructureXXX) LookupByIndex(idx int64) (datamodel.Node, error) {
	switch idx {
	case 0:
		return x.F1XXX.Node(), nil
	case 1:
		return x.F2XXX.Node(), nil
	}
	return nil, values.ErrNA
}

func (x StructureXXX) LookupBySegment(seg datamodel.PathSegment) (datamodel.Node, error) {
	switch seg.String() {
	case "0", "f1_XXX":
		return x.F1XXX.Node(), nil
	case "1", "f2_XXX":
		return x.F2XXX.Node(), nil
	}
	return nil, values.ErrNA
}

func (x StructureXXX) MapIterator() datamodel.MapIterator {
	return &StructureXXX_MapIterator{-1, &x}
}

func (x StructureXXX) ListIterator() datamodel.ListIterator {
	return nil
}

func (x StructureXXX) Length() int64 {
	return 2 //XXX
}

func (x StructureXXX) IsAbsent() bool {
	return false
}

func (x StructureXXX) IsNull() bool {
	return false
}

func (x StructureXXX) AsBool() (bool, error) {
	return false, values.ErrNA
}

func (x StructureXXX) AsInt() (int64, error) {
	return 0, values.ErrNA
}

func (x StructureXXX) AsFloat() (float64, error) {
	return 0, values.ErrNA
}

func (x StructureXXX) AsString() (string, error) {
	return "", values.ErrNA
}

func (x StructureXXX) AsBytes() ([]byte, error) {
	return nil, values.ErrNA
}

func (x StructureXXX) AsLink() (datamodel.Link, error) {
	return nil, values.ErrNA
}

func (x StructureXXX) Prototype() datamodel.NodePrototype {
	return nil
}
