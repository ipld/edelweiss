package dagjsonoverhttp

import (
	"fmt"

	"github.com/ipld/edelweiss/blueprints/base"
	cg "github.com/ipld/edelweiss/codegen"
	"github.com/ipld/edelweiss/def"
)

func BuildClientImpl(
	lookup cg.LookupDepGoRef,
	typeDef def.Service,
	goTypeRef cg.GoTypeRef,
) (cg.GoTypeImpl, error) {
	return &GoClientImpl{
		Lookup: lookup,
		Def:    typeDef,
		Ref:    goTypeRef,
	}, nil
}

type GoClientImpl struct {
	Lookup cg.LookupDepGoRef
	Def    def.Service
	Ref    cg.GoTypeRef
}

func (x *GoClientImpl) ProtoDef() def.Type {
	return x.Def
}

func (x *GoClientImpl) GoTypeRef() cg.GoTypeRef {
	return x.Ref
}

func (x *GoClientImpl) GoDef() cg.Blueprint {
	data := cg.BlueMap{
		"Errorf":            base.Errorf,
		"HTTPClient":        base.HTTPClient,
		"HTTPDefaultClient": base.HTTPDefaultClient,
		"URL":               base.URL,
		"URLParse":          base.URLParse,
		//
		"Interface":      x.Ref.Append("_Client"),
		"Type":           x.Ref.Prepend("client_"),
		"Option":         x.Ref.Append("_ClientOption"),
		"New":            x.Ref.Prepend("New_").Append("_Client"),
		"WithHTTPClient": x.Ref.Append("_Client").Append("_WithHTTPClient"),
		//
		"Logger":     cg.GoRef{PkgPath: "github.com/ipfs/go-log", Name: "Logger"},
		"LoggerName": cg.StringLiteral(fmt.Sprintf("service/client/%s", x.Ref.TypeName)),
		"LoggerVar":  cg.GoRef{PkgPath: x.Ref.PkgPath, Name: fmt.Sprintf("logger_client_%s", x.Ref.TypeName)},
		//
		"MethodDecls": XXX,
		"MethodImpls": XXX,
	}
	return cg.T{Data: data, Src: goClientTemplate}
}

const goClientTemplate = `
var {{.LoggerVar}} = {{.Logger}}({{.LoggerName}})

type {{.Interface}} interface {
{{.MethodDecls}} //XXX
	GetP2PProvide(ctx context.Context, req *proto.GetP2PProvideRequest) ([]*proto.GetP2PProvideResponse, error)
	GetP2PProvide_Async(ctx context.Context, req *proto.GetP2PProvideRequest) (<-chan GetP2PProvide_Async_Response, error)
}

//XXX
type GetP2PProvide_Async_Response struct {
	Resp *proto.GetP2PProvideResponse
	Err  error
}

type {{.Option}} func(*{{.Type}}) error

type {{.Type}} struct {
	httpClient       *{{.HTTPClient}}
	endpoint     *{{.URL}}
}

func {{.WithHTTPClient}}(hc *{{.HTTPClient}}) {{.Option}} {
	return func(c *{{.Type}}) error {
		c.httpClient = hc
		return nil
	}
}

func {{.New}}(endpoint string, opts ...{{.Option}}) (*{{.Type}}, error) {
	u, err := {{.URLParse}}(endpoint)
	if err != nil {
		return nil, err
	}
	c := &{{.Type}}{endpoint: u, httpClient: {{.HTTPDefaultClient}}}
	for _, o := range opts {
		if err := o(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}

{{.MethodImpls}} //XXX
`
