package values

import (
	"github.com/ipld/edelweiss/def"
	"github.com/ipld/go-ipld-prime/datamodel"
)

type Structure []Field

type Field struct {
	Name  string
	Value Any
}

func (v Structure) Def() def.Type {
	fs := make([]def.Field, len(v))
	for i := range v {
		fs[i] = def.Field{Name: v[i].Name, Type: v[i].Value.Def()}
	}
	return def.MakeStructure(fs...)
}

func (v Structure) Node() datamodel.Node {
	return v.Envelope().Node()
}

func (v Structure) Envelope() Map {
	m := make(Map, len(v)+1)
	m[0] = makeEnvelopeTag(envelopeStructureTagValue)
	for i := range v {
		m[i+1] = KeyValue{Key: Any{String(v[i].Name)}, Value: v[i].Value}
	}
	return m
}

func (x *Structure) Parse(n datamodel.Node) error {
	*x = Structure{}
	var m Map
	if err := m.Parse(n); err != nil {
		return err
	}
	tagFound := false
	for _, kv := range m {
		if tag, ok := extractEnvelopeTag(kv); ok && !tagFound {
			if tag != envelopeStructureTagValue {
				return ErrInvalid
			}
			tagFound = true
		} else {
			if n, ok := kv.Key.Value.(String); ok {
				*x = append(*x, Field{Name: string(n), Value: kv.Value})
			}
		}
	}
	if tagFound {
		return nil
	} else {
		return ErrInvalid
	}
}

func structureEqual(x, y Structure) bool {
	if len(x) != len(y) {
		return false
	} else {
		for i := range x {
			if x[i] != y[i] {
				return false
			}
		}
		return true
	}
}

func TryParseStructure(n datamodel.Node) (Structure, error) {
	var s Structure
	return s, s.Parse(n)
}
