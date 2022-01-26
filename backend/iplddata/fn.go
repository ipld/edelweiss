package iplddata

import (
	"github.com/ipld/edelweiss/def"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/node/basicnode"
)

// type Return struct {
// 	Value Any
// }

type Call struct {
	Name string
	Arg  Any
}

func (x Call) Def() def.Type {
	return def.Fn{Arg: x.Arg.Def()} //XXX
}

func (x Call) Node() datamodel.Node {
	XXX
}

const envelopeCallTagValue = "call"

var envelopeCallTagValueNode = basicnode.NewString(envelopeCallTagValue)

func (x *Call) Parse(n datamodel.Node) error {
	parseTag := func(tag string) error {
		if tag != envelopeCallTagValue {
			return ErrInvalid
		}
		return nil
	}
	foundName, foundArg := false, false
	parseKeyValue := func(k datamodel.Node, v datamodel.Node) error {
		ks, err := k.AsString()
		if err != nil { // ignore unknown fields
			return nil
		}
		switch ks {
		case "name":
			vs, err := v.AsString()
			if err != nil {
				return ErrInvalid
			}
			x.Name = vs
			foundName = true
		case "arg":
			if err := x.Arg.Parse(v); err != nil {
				return err
			}
			foundArg = true
		}
		return nil
	}
	if err := ParseEnvelope(n, parseTag, parseKeyValue); err != nil {
		return err
	}
	if !foundName || !foundArg {
		return ErrInvalid
	}
	return nil
}

// datamodel.Node implementation
