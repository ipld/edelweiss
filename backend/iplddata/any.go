package iplddata

import (
	"github.com/ipld/edelweiss/def"
	"github.com/ipld/go-ipld-prime/datamodel"
)

type Any struct {
	Value
}

func (Any) Def() def.Type {
	return def.Any{}
}

func (v *Any) Parse(n datamodel.Node) error {
	if x, err := TryParseBool(n); err == nil {
		v.Value = x
		return nil
	}
	if x, err := TryParseList(n); err == nil {
		v.Value = x
		return nil
	}
	if x, err := TryParseStructure(n); err == nil {
		v.Value = x
		return nil
	}
	// etc
	return ErrUnexpected
}
