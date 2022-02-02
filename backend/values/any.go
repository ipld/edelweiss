package values

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
	// primitives
	if x, err := TryParseBool(n); err == nil {
		v.Value = x
		return nil
	}
	if x, err := TryParseInt(n); err == nil {
		v.Value = x
		return nil
	}
	if x, err := TryParseFloat(n); err == nil {
		v.Value = x
		return nil
	}
	if x, err := TryParseString(n); err == nil {
		v.Value = x
		return nil
	}
	if x, err := TryParseLink(n); err == nil {
		v.Value = x
		return nil
	}
	// composite
	if x, err := TryParseList(n); err == nil {
		v.Value = x
		return nil
	}
	// enveloped composite
	if x, err := TryParseStructure(n); err == nil {
		v.Value = x
		return nil
	}
	if x, err := TryParseCall(n); err == nil {
		v.Value = x
		return nil
	}
	if x, err := TryParseMap(n); err == nil {
		v.Value = x
		return nil
	}
	// etc
	return ErrUnexpected
}

func TryParseAny(n datamodel.Node) (Any, error) {
	var x Any
	return x, x.Parse(n)
}
