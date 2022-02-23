package base

import (
	cg "github.com/ipld/edelweiss/codegen"
)

var (
	Context           = cg.GoTypeRef{PkgPath: "context", TypeName: "Context"}
	ContextWithCancel = cg.GoTypeRef{PkgPath: "context", TypeName: "WithCancel"}
	HTTPClient        = cg.GoTypeRef{PkgPath: "http", TypeName: "Client"}
	HTTPDefaultClient = cg.GoTypeRef{PkgPath: "http", TypeName: "DefaultClient"}
	URL               = cg.GoTypeRef{PkgPath: "url", TypeName: "URL"}
	URLParse          = cg.GoRef{PkgPath: "url", Name: "Parse"}
)
