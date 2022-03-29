package services

// PkgPath is the fully-qualified name of this package.
const PkgPath = "github.com/ipld/edelweiss/services"

// ErrContext wraps context-related errors, like context cancellation.
type ErrContext struct {
	Cause error
}

func (e ErrContext) Error() string {
	return e.Cause.Error()
}

// ErrProto wraps protocol errors, like undecodable messages.
type ErrProto struct {
	Cause error
}

func (e ErrProto) Error() string {
	return e.Cause.Error()
}

// ErrService wraps service-level errors, produced by service implementations.
type ErrService struct {
	Cause error
}

func (e ErrService) Error() string {
	return e.Cause.Error()
}
