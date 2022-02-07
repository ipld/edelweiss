package values

import (
	"fmt"

	"github.com/ipld/edelweiss/def"
	"github.com/ipld/go-ipld-prime/datamodel"
)

// PkgPath is the fully-qualified name of this package.
const PkgPath = "github.com/ipld/edelweiss/backend/values"

// [x] primitives:
//	[x] bool (lib)
//	[x] int (lib)
//	[x] float (lib)
//	[x] byte (lib)
//	[x] char (lib)
//
// [ ] singletons:
//	[ ] singleton bool (codegen)
//	[ ] singleton int (codegen)
//	[ ] singleton float (codegen)
//	[ ] singleton byte (codegen)
//	[ ] singleton char (codegen)
//
// [•] link (lib)
// [x] list (lib+codegen)
// [x] map (lib+codegen)
// [x] structure (lib+codegen)
// [-] tuple (lib+codegen)
// [ ] union (codegen)
// [ ] ref+named (codegen)
// [x] fn: call/return (lib+codegen)
// [ ] service (codegen)
//
// [x] specials:
//	[x] string (lib)
//	[x] any (lib)

type Value interface {
	Def() def.Type
	Node() datamodel.Node
}

type Parser interface {
	Parse(datamodel.Node) error
}

var (
	ErrNA           = fmt.Errorf("n/a")
	ErrBounds       = fmt.Errorf("index out of bounds")
	ErrUnexpected   = fmt.Errorf("unexpected")
	ErrInvalid      = fmt.Errorf("invalid format")
	ErrNotSupported = fmt.Errorf("not supported")
)
