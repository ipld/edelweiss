package plans

import (
	"github.com/ipld/edelweiss/defs"
)

type Service struct {
	Methods        []defs.Method
	CallEnvelope   defs.Def // ref to inductive
	ReturnEnvelope defs.Def // ref to inductive
}

func (Service) Kind() string {
	return "Service"
}
