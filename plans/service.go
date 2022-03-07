package plans

import (
	"github.com/ipld/edelweiss/def"
)

type Service struct {
	Methods        []def.Method
	CallEnvelope   def.Def // ref to inductive
	ReturnEnvelope def.Def // ref to inductive
}

func (Service) Kind() string {
	return "Service"
}
