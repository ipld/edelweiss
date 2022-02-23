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

func (x GoClientImpl) ProtoDef() def.Type {
	return x.Def
}

func (x GoClientImpl) GoTypeRef() cg.GoTypeRef {
	return x.Ref
}

func (x GoClientImpl) GoDef() cg.Blueprint {
	methods := def.FlattenMethodList(x.Def.Methods)
	methodSyncDecls, methodAsyncDecls := make(cg.BlueSlice, len(methods)), make(cg.BlueSlice, len(methods))
	methodAsyncResultDefs := make(cg.BlueSlice, len(methods))
	for i, m := range methods {
		asyncResultRef := &cg.GoTypeRef{PkgPath: x.Ref.PkgPath, TypeName: x.Ref.TypeName + "_" + m.Name + "_AsyncResult"}
		bm := cg.BlueMap{
			"Context":           base.Context,
			"MethodName":        cg.V(m.Name),
			"MethodArg":         x.Lookup.LookupDepGoRef(m.Type.Arg),
			"MethodReturn":      x.Lookup.LookupDepGoRef(m.Type.Arg),
			"MethodReturnAsync": asyncResultRef,
		}
		methodAsyncResultDefs[i] = cg.T{
			Data: bm,
			Src: `type {{.MethodReturnAsync}} struct {
	Resp *{{.MethodReturn}}
	Err  error
}`,
		}
		methodSyncDecls[i] = cg.T{
			Data: bm,
			Src:  `{{.MethodName}}(ctx {{.Context}}, req *{{.MethodArg}}) ([]{{.MethodReturn}}, error)`,
		}
		methodAsyncDecls[i] = cg.T{
			Data: bm,
			Src:  `{{.MethodName}}_Async(ctx {{.Context}}, req *{{.MethodArg}}) (<-chan {{.MethodReturnAsync}}, error)`,
		}
	}
	//
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
		"MethodSyncDecls":       methodSyncDecls,
		"MethodAsyncDecls":      methodAsyncDecls,
		"MethodAsyncResultDefs": methodAsyncResultDefs,
		// "MethodImpls":      XXX,
	}
	return cg.T{Data: data, Src: goClientTemplate}
}

const goClientTemplate = `
var {{.LoggerVar}} = {{.Logger}}({{.LoggerName}})

type {{.Interface}} interface {
{{range .MethodSyncDecls}}
{{.}}
{{end}}
{{range .MethodAsyncDecls}}
{{.}}
{{end}}
}

{{range .MethodAsyncResultDefs}}
{{.}}
{{end}}

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
