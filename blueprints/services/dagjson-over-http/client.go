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
	methodImpls := make(cg.BlueSlice, len(methods))
	for i, m := range methods {
		asyncResultRef := &cg.GoTypeRef{PkgPath: x.Ref.PkgPath, TypeName: x.Ref.TypeName + "_" + m.Name + "_AsyncResult"}
		processAsyncResultRef := &cg.GoRef{PkgPath: x.Ref.PkgPath, Name: "process_" + x.Ref.TypeName + "_" + m.Name + "_AsyncResult"}
		bm := cg.BlueMap{
			"Context":            base.Context,
			"MethodName":         cg.V(m.Name),
			"MethodArg":          x.Lookup.LookupDepGoRef(m.Type.Arg),
			"MethodReturn":       x.Lookup.LookupDepGoRef(m.Type.Arg),
			"MethodReturnAsync":  asyncResultRef,
			"ProcessReturnAsync": processAsyncResultRef,
		}
		methodAsyncResultDefs[i] = cg.T{
			Data: bm,
			Src: `type {{.MethodReturnAsync}} struct {
	Resp *{{.MethodReturn}}
	Err  error
}`,
		}
		syncMethodDecl := cg.T{
			Data: bm,
			Src:  `{{.MethodName}}(ctx {{.Context}}, req *{{.MethodArg}}) ([]*{{.MethodReturn}}, error)`,
		}
		methodSyncDecls[i] = syncMethodDecl
		asyncMethodDecl := cg.T{
			Data: bm,
			Src:  `{{.MethodName}}_Async(ctx {{.Context}}, req *{{.MethodArg}}) (<-chan {{.MethodReturnAsync}}, error)`,
		}
		methodAsyncDecls[i] = asyncMethodDecl
		bm["SyncMethodDecl"] = syncMethodDecl
		bm["AsyncMethodDecl"] = asyncMethodDecl
		methodImpls[i] = cg.T{Data: bm, Src: goClientMethodImplTemplate}
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
		"MethodImpls":           methodImpls,
	}
	return cg.T{Data: data, Src: goClientTemplate}
}

const (
	goClientTemplate = `
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

{{range .MethodImpls}}
{{.}}
{{end}}
`
	//.Type
	//.ContextWithCancel
	//.LoggerVar
	//.Context
	//.IOReader
	goClientMethodImplTemplate = `
func (c *{{.Type}}) {{.SyncMethodDecl}} {
	ctx, cancel := {{.ContextWithCancel}}(ctx)
	defer cancel()
	ch, err := c.{{.MethodName}}_Async(ctx, req)
	if err != nil {
		return nil, err
	}
	var resps []*{{.MethodReturn}}
	for {
		select {
		case r, ok := <-ch:
			if !ok {
				cancel()
				return resps, nil
			} else {
				if r.Err == nil {
					resps = append(resps, r.Resp)
				} else {
					{{.LoggerVar}}.Errorf("client received invalid response (%v)", r.Err)
				}
			}
		case <-ctx.Done():
			return resps, ctx.Err()
		}
	}
}

func (c *{{.Type}}) {{.AsyncMethodDecl}} {
	//XXX
	envelope := &proto.ServiceEnvelope{
		GetP2PProvideRequest: req,
	}

	buf, err := ipld.Marshal(dagjson.Encode, envelope, proto.Prototypes.ServiceEnvelope.Type())
	if err != nil {
		return nil, fmt.Errorf("unexpected serialization error (%w)", err)
	}

	// encode request in URL
	u := *c.endpoint
	q := url.Values{}
	q.Set("q", string(buf))
	u.RawQuery = q.Encode()
	httpReq, err := http.NewRequestWithContext(ctx, "GET", u.String(), bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("sending HTTP request (%v)", err)
	}

	ch := make(chan {{.MethodReturnAsync}}, 1)
	go {{.ProcessReturnAsync}}(ctx, ch, resp.Body)
	return ch, nil
}

func {{.ProcessReturnAsync}}(ctx {{.Context}}, ch chan<- {{.MethodReturnAsync}}, r {{.IOReader}}) {
	defer close(ch)
	for {
		if ctx.Err() != nil {
			return
		}

		// XXX
		env := &proto.ServiceEnvelope{}
		_, err := ipld.UnmarshalStreaming(r, dagjson.Decode, env, proto.Prototypes.ServiceEnvelope.Type())
		if errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF) {
			return
		}
		if err != nil {
			ch <- {{.MethodReturnAsync}}{Err: err}
			return
		}

		if env.GetP2PProvideResponse == nil {
			continue
		}
		ch <- {{.MethodReturnAsync}}{Resp: env.GetP2PProvideResponse}
	}
}`
)
