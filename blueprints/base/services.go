package base

import (
	cg "github.com/ipld/edelweiss/codegen"
	"github.com/ipld/edelweiss/services"
)

var (
	SyncMutex                 = cg.GoTypeRef{PkgPath: "sync", TypeName: "Mutex"}
	Context                   = cg.GoTypeRef{PkgPath: "context", TypeName: "Context"}
	ContextWithCancel         = cg.GoTypeRef{PkgPath: "context", TypeName: "WithCancel"}
	ContextBackground         = cg.GoTypeRef{PkgPath: "context", TypeName: "Background"}
	HTTPClient                = cg.GoTypeRef{PkgPath: "net/http", TypeName: "Client"}
	HTTPDefaultClient         = cg.GoTypeRef{PkgPath: "net/http", TypeName: "DefaultClient"}
	HTTPNewRequestWithContext = cg.GoRef{PkgPath: "net/http", Name: "NewRequestWithContext"}
	HTTPHandlerFunc           = cg.GoTypeRef{PkgPath: "net/http", TypeName: "HandlerFunc"}
	HTTPRequest               = cg.GoTypeRef{PkgPath: "net/http", TypeName: "Request"}
	HTTPResponseWriter        = cg.GoTypeRef{PkgPath: "net/http", TypeName: "ResponseWriter"}
	URL                       = cg.GoTypeRef{PkgPath: "net/url", TypeName: "URL"}
	URLParse                  = cg.GoRef{PkgPath: "net/url", Name: "Parse"}
	URLValues                 = cg.GoRef{PkgPath: "net/url", Name: "Values"}
	BytesBuffer               = cg.GoTypeRef{PkgPath: "bytes", TypeName: "Buffer"}
	BytesNewReader            = cg.GoRef{PkgPath: "bytes", Name: "NewReader"}
)

var (
	IOReader           = &cg.GoRef{PkgPath: "io", Name: "Reader"}
	IOEOF              = &cg.GoRef{PkgPath: "io", Name: "EOF"}
	IOErrUnexpectedEOF = &cg.GoRef{PkgPath: "io", Name: "ErrUnexpectedEOF"}
)

const EdelweissServicesPkg = services.PkgPath

var (
	EdelweissErrContext = &cg.GoTypeRef{PkgPath: EdelweissServicesPkg, TypeName: "ErrContext"}
	EdelweissErrProto   = &cg.GoTypeRef{PkgPath: EdelweissServicesPkg, TypeName: "ErrProto"}
	EdelweissErrService = &cg.GoTypeRef{PkgPath: EdelweissServicesPkg, TypeName: "ErrService"}
	EdelweissErrSchema  = &cg.GoRef{PkgPath: EdelweissServicesPkg, Name: "ErrSchema"}
)
