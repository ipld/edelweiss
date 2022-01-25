package ipld

import (
	"fmt"

	"github.com/ipld/edelweiss/def"
	"github.com/ipld/go-ipld-prime/datamodel"
)

// [ ] fn
// [ ] link
// [ ] list
// [ ] map
// [ ] named
// [ ] primitives:
//	[x] bool
//	[ ] int
//	[ ] float
//	[ ] byte
//	[ ] char
// [ ] ref
// [ ] service
// [ ] singletons:
//	[x] singleton bool
//	[ ] singleton int
//	[ ] singleton float
//	[ ] singleton byte
//	[ ] singleton char
// [ ] specials:
//	[ ] string
//	[ ] any
//	[ ] nothing
// [ ] structure
// [ ] tuple
// [ ] union

type Value interface {
	IPLDNode() datamodel.Node
	Type() Type
}

type Type interface {
	Def() def.Type
	datamodel.NodePrototype
}

var ErrNA = fmt.Errorf("n/a")
