package base

import (
	cg "github.com/ipld/edelweiss/codegen"
)

var (
	Context                   = cg.GoTypeRef{PkgPath: "context", TypeName: "Context"}
	ContextWithCancel         = cg.GoTypeRef{PkgPath: "context", TypeName: "WithCancel"}
	HTTPClient                = cg.GoTypeRef{PkgPath: "net/http", TypeName: "Client"}
	HTTPDefaultClient         = cg.GoTypeRef{PkgPath: "net/http", TypeName: "DefaultClient"}
	HTTPNewRequestWithContext = cg.GoRef{PkgPath: "net/http", Name: "NewRequestWithContext"}
	URL                       = cg.GoTypeRef{PkgPath: "net/url", TypeName: "URL"}
	URLParse                  = cg.GoRef{PkgPath: "net/url", Name: "Parse"}
	URLValues                 = cg.GoRef{PkgPath: "net/url", Name: "Values"}
	BytesNewReader            = cg.GoRef{PkgPath: "bytes", Name: "NewReader"}
)

var (
	IOReader           = &cg.GoRef{PkgPath: "io", Name: "Reader"}
	IOEOF              = &cg.GoRef{PkgPath: "io", Name: "EOF"}
	IOErrUnexpectedEOF = &cg.GoRef{PkgPath: "io", Name: "ErrUnexpectedEOF"}
)
