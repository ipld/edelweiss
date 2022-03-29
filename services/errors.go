package services

// PkgPath is the fully-qualified name of this package.
const PkgPath = "github.com/ipld/edelweiss/services"

// ErrContext wraps context-related errors, like context cancellation.
type ErrContext struct {
	error
}

// ErrProto wraps protocol errors, like undecodable messages.
type ErrProto struct {
	error
}

// ErrService wraps service-level errors, produced by service implementations.
type ErrService struct {
	error
}
