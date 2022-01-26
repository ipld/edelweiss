package iplddata

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
// [x] list (lib+codegen)
// [-] map (lib+codegen)
// [x] structure (lib+codegen)
// [-] tuple (lib+codegen)
// [ ] union (codegen)
// [ ] ref+named (codegen)
// [•] fn (lib+codegen)
// [•] service (codegen)
//
// [ ] specials:
//	[•] string (lib)
//	[x] any (lib)

type Value interface {
	Def() def.Type
	datamodel.Node
}

type Parser interface {
	Parse(datamodel.Node) error
}

var (
	ErrNA         = fmt.Errorf("n/a")
	ErrBounds     = fmt.Errorf("index out of bounds")
	ErrUnexpected = fmt.Errorf("unexpected")
)
