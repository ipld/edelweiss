package ipldbind

import (
	"fmt"

	"github.com/ipld/edelweiss/def"
	"github.com/ipld/go-ipld-prime/datamodel"
)

// [ ] primitives:
//	[x] bool (lib)
//	[ ] int (lib)
//	[ ] float (lib)
//	[ ] byte (lib)
//	[ ] char (lib)
//
// [ ] singletons:
//	[ ] singleton bool (codegen)
//	[ ] singleton int (codegen)
//	[ ] singleton float (codegen)
//	[ ] singleton byte (codegen)
//	[ ] singleton char (codegen)
//
// [ ] link (lib)
// [ ] list (lib+codegen)
// [ ] map (lib+codegen)
// [ ] structure (lib+codegen)
// [ ] tuple (lib+codegen)
// [ ] union (codegen)
// [ ] ref+named (codegen)
// [ ] fn (lib+codegen)
// [ ] service (codegen)
//
// [ ] specials:
//	[ ] string (lib)
//	[ ] any (lib)

type Value interface {
	IPLDNode() datamodel.Node
	Type() Type
}

type Type interface {
	Def() def.Type
	datamodel.NodePrototype
}

var ErrNA = fmt.Errorf("n/a")
