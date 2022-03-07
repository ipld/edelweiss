package routingapi

import (
	"github.com/ipld/edelweiss/defs"
)

var proto = defs.Defs{
	defs.Named{
		Name: "DelegatedRouting",
		Type: defs.Service{},
	},
}
