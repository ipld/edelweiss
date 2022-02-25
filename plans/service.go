package plans

import (
	"github.com/ipld/edelweiss/def"
)

type Service struct {
	Methods        []def.Method
	CallEnvelope   def.Type // ref to inductive
	ReturnEnvelope def.Type // ref to inductive
}

func (Service) Kind() string {
	return "Service"
}
