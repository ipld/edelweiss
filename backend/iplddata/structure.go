package iplddata

import (
	"github.com/ipld/edelweiss/def"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/node/basicnode"
)

type Structure []Field

type Field struct {
	Name  string
	Value Any
}

func (v Structure) Def() def.Type {
	fs := make([]def.Field, len(v))
	for i := range v {
		fs[i] = def.Field{v[i].Name, v[i].Value.Def()}
	}
	return def.MakeStructure(fs...)
}

func (v *Structure) Parse(n datamodel.Node) error {
	if n.Kind() != ipld.Kind_Map {
		return ErrNA
	} else {
		structureTagFound := false
		iter := n.MapIterator()
		for !iter.Done() {
			kn, vn, _ := iter.Next()
			kns, err := kn.AsString()
			if err != nil {
				return ErrNA
			}
			vns, _ := vn.AsString()
			if kns == repnTypeKey { // parse type key
				if vns == repnStructure {
					structureTagFound = true
					continue
				} else {
					return ErrNA
				}
			} else { // parse user keys
				var f Field
				f.Name = kns
				if err = f.Value.Parse(vn); err != nil {
					return ErrNA
				} else {
					*v = append(*v, f)
				}
			}
		}
		if structureTagFound {
			return nil
		} else {
			return ErrNA
		}
	}
}

// datamodel.Node implementation

func (Structure) Kind() datamodel.Kind {
	return datamodel.Kind_Map
}

const repnStructure = "structure"

var repnStructureNode = basicnode.NewString(repnStructure)

func (v Structure) LookupByString(s string) (datamodel.Node, error) {
	if s == repnTypeKey {
		return repnStructureNode, nil
	}
	for _, f := range v {
		if f.Name == s {
			return f.Value, nil
		}
	}
	return nil, ErrNA
}

func (v Structure) LookupByNode(key datamodel.Node) (datamodel.Node, error) {
	ks, err := key.AsString()
	if err != nil {
		return nil, ErrNA
	}
	return v.LookupByString(ks)
}

func (v Structure) LookupByIndex(i int64) (datamodel.Node, error) {
	return nil, ErrNA
}

func (v Structure) LookupBySegment(seg datamodel.PathSegment) (datamodel.Node, error) {
	return v.LookupByString(seg.String())
}

func (v Structure) MapIterator() datamodel.MapIterator {
	return &structureIterator{v, -1}
}

func (v Structure) ListIterator() datamodel.ListIterator {
	return nil
}

func (v Structure) Length() int64 {
	return int64(len(v) + 1) // account for the demux field
}

func (Structure) IsAbsent() bool {
	return false
}

func (Structure) IsNull() bool {
	return false
}

func (v Structure) AsBool() (bool, error) {
	return false, ErrNA
}

func (Structure) AsInt() (int64, error) {
	return 0, ErrNA
}

func (Structure) AsFloat() (float64, error) {
	return 0, ErrNA
}

func (Structure) AsString() (string, error) {
	return "", ErrNA
}

func (Structure) AsBytes() ([]byte, error) {
	return nil, ErrNA
}

func (Structure) AsLink() (datamodel.Link, error) {
	return nil, ErrNA
}

func (Structure) Prototype() datamodel.NodePrototype {
	return nil // not needed
}

type structureIterator struct {
	s  Structure
	at int64
}

func (iter *structureIterator) Next() (datamodel.Node, datamodel.Node, error) {
	if iter.Done() {
		return nil, nil, ErrBounds
	}
	if iter.at == -1 {
		iter.at++
		return repnTypeKeyNode, repnStructureNode, nil
	} else {
		v := iter.s[iter.at]
		iter.at++
		return basicnode.NewString(v.Name), v.Value, nil
	}
}

func (iter *structureIterator) Done() bool {
	return iter.at >= int64(len(iter.s))
}

func TryParseStructure(n datamodel.Node) (Structure, error) {
	var s Structure
	return s, s.Parse(n)
}
