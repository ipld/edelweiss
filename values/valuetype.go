package values

import (
	"fmt"

	"github.com/ipld/edelweiss/def"
	"github.com/ipld/go-ipld-prime/datamodel"
)

// PkgPath is the fully-qualified name of this package.
const PkgPath = "github.com/ipld/edelweiss/values"

type Value interface {
	Def() def.Def
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
	ErrNotFound     = fmt.Errorf("not found")
)
