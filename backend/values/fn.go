package values

import (
	"github.com/ipld/edelweiss/def"
	"github.com/ipld/go-ipld-prime/datamodel"
)

type Call struct {
	Name string
	Arg  Any
}

func (x Call) Def() def.Type {
	return def.Fn{Arg: x.Arg.Def(), Return: def.Nothing{}}
}

func (x Call) Node() datamodel.Node {
	return x.Envelope().Node()
}

func (x Call) Envelope() Map {
	m := make(Map, 2)
	m[0] = makeEnvelopeTag(envelopeCallTagValue)
	m[1] = KeyValue{Key: Any{String("arg")}, Value: x.Arg}
	return m
}

func (x *Call) Parse(n datamodel.Node) error {
	*x = Call{}
	var m Map
	if err := m.Parse(n); err != nil {
		return err
	}
	tagFound, argFound := false, false
	for _, kv := range m {
		if tag, ok := extractEnvelopeTag(kv); ok {
			if tag != envelopeCallTagValue {
				return ErrInvalid
			}
			tagFound = true
		} else {
			if n, ok := kv.Key.Value.(String); ok && n == "arg" {
				x.Arg = kv.Value
				argFound = true
			}
		}
	}
	if tagFound && argFound {
		return nil
	}
	return ErrInvalid
}

func TryParseCall(n datamodel.Node) (Call, error) {
	var s Call
	return s, s.Parse(n)
}
