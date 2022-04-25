package dagjsonoverhttp

import (
	"fmt"
	"strings"

	"github.com/ipld/edelweiss/blueprints/base"
	cg "github.com/ipld/edelweiss/codegen"
	"github.com/ipld/edelweiss/plans"
)

func BuildClientImpl(
	lookup cg.LookupDepGoRef,
	plan plans.Service,
	goTypeRef cg.GoTypeRef,
) cg.GoTypeImpl {
	return &GoClientImpl{
		Lookup: lookup,
		Def:    plan,
		Ref:    goTypeRef,
	}
}

type GoClientImpl struct {
	Lookup cg.LookupDepGoRef
	Def    plans.Service
	Ref    cg.GoTypeRef
}

func (x GoClientImpl) GoTypeRef() cg.GoTypeRef {
	return x.Ref
}

func (x GoClientImpl) GoDef() cg.Blueprint {
	methods := x.Def.Methods
	methodSyncDecls, methodAsyncDecls := make(cg.BlueSlice, len(methods)), make(cg.BlueSlice, len(methods))
	methodAsyncResultDefs := make(cg.BlueSlice, len(methods))
	methodImpls := make(cg.BlueSlice, len(methods))
	typ := x.Ref.Prepend("client_")
	loggerVar := cg.GoRef{PkgPath: x.Ref.PkgPath, Name: fmt.Sprintf("logger_client_%s", x.Ref.TypeName)}
	for i, m := range methods {
		// async result type is also used by the server code
		asyncResultRef := &cg.GoTypeRef{PkgPath: x.Ref.PkgPath, TypeName: x.Ref.TypeName + "_" + m.Name + "_AsyncResult"}
		processAsyncResultRef := &cg.GoRef{PkgPath: x.Ref.PkgPath, Name: "process_" + x.Ref.TypeName + "_" + m.Name + "_AsyncResult"}
		bmDecl := cg.BlueMap{
			"Type":               typ,
			"MethodName":         cg.V(m.Name),
			"MethodArg":          x.Lookup.LookupDepGoRef(m.Type.Arg),
			"MethodReturn":       x.Lookup.LookupDepGoRef(m.Type.Return),
			"MethodReturnAsync":  asyncResultRef,
			"ProcessReturnAsync": processAsyncResultRef,
			//
			"DAGJSONEncode":             base.DAGJSONEncode,
			"Context":                   base.Context,
			"ContextWithCancel":         base.ContextWithCancel,
			"LoggerVar":                 loggerVar,
			"IOReader":                  base.IOReader,
			"IPLDEncode":                base.IPLDEncode,
			"IPLDDecodeStreaming":       base.IPLDDecodeStreaming,
			"Errorf":                    base.Errorf,
			"URLValues":                 base.URLValues,
			"HTTPNewRequestWithContext": base.HTTPNewRequestWithContext,
			"BytesNewReader":            base.BytesNewReader,
			"DAGJSONDecode":             base.DAGJSONDecode,
			"ErrorsIs":                  base.ErrorsIs,
			"IOEOF":                     base.IOEOF,
			"IOErrUnexpectedEOF":        base.IOErrUnexpectedEOF,
			//
			"AcceptContentType": cg.StringLiteral(ContentTypeV1),
		}
		methodAsyncResultDefs[i] = cg.T{
			Data: bmDecl,
			Src: `type {{.MethodReturnAsync}} struct {
	Resp *{{.MethodReturn}}
	Err  error
}`,
		}
		syncMethodDecl := cg.T{
			Data: bmDecl,
			Src:  `{{.MethodName}}(ctx {{.Context}}, req *{{.MethodArg}}) ([]*{{.MethodReturn}}, error)`,
		}
		methodSyncDecls[i] = syncMethodDecl
		asyncMethodDecl := cg.T{
			Data: bmDecl,
			Src:  `{{.MethodName}}_Async(ctx {{.Context}}, req *{{.MethodArg}}) (<-chan {{.MethodReturnAsync}}, error)`,
		}
		methodAsyncDecls[i] = asyncMethodDecl
		bmImpl := cg.MergeBlueMaps(bmDecl,
			cg.BlueMap{
				"ErrorEnvelope":   x.Lookup.LookupDepGoRef(x.Def.ErrorEnvelope),
				"CallEnvelope":    x.Lookup.LookupDepGoRef(x.Def.CallEnvelope),
				"ReturnEnvelope":  x.Lookup.LookupDepGoRef(x.Def.ReturnEnvelope),
				"SyncMethodDecl":  syncMethodDecl,
				"AsyncMethodDecl": asyncMethodDecl,
				"ErrContext":      base.EdelweissErrContext,
				"ErrProto":        base.EdelweissErrProto,
				"ErrService":      base.EdelweissErrService,
				"ErrorsNew":       base.ErrorsNew,
			},
		)
		methodImpls[i] = cg.T{Data: bmImpl, Src: goClientMethodImplTemplate}
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
		"Type":           typ,
		"Option":         x.Ref.Append("_ClientOption"),
		"New":            x.Ref.Prepend("New_").Append("_Client"),
		"WithHTTPClient": x.Ref.Append("_Client").Append("_WithHTTPClient"),
		//
		"Logger":     cg.GoRef{PkgPath: "github.com/ipfs/go-log", Name: "Logger"},
		"LoggerName": cg.StringLiteral(fmt.Sprintf("service/client/%s", strings.ToLower(x.Ref.TypeName))),
		"LoggerVar":  loggerVar,
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
					{{.LoggerVar}}.Errorf("client received error response (%v)", r.Err)
					cancel()
					return resps, r.Err
				}
			}
		case <-ctx.Done():
			return resps, ctx.Err()
		}
	}
}

func (c *{{.Type}}) {{.AsyncMethodDecl}} {
	envelope := &{{.CallEnvelope}}{
		{{.MethodName}}: req,
	}

	buf, err := {{.IPLDEncode}}(envelope, {{.DAGJSONEncode}})
	if err != nil {
		return nil, {{.Errorf}}("unexpected serialization error (%v)", err)
	}

	// encode request in URL
	u := *c.endpoint
	q := {{.URLValues}}{}
	q.Set("q", string(buf))
	u.RawQuery = q.Encode()
	httpReq, err := {{.HTTPNewRequestWithContext}}(ctx, "POST", u.String(), {{.BytesNewReader}}(buf))
	if err != nil {
		return nil, err
	}
	httpReq.Header = map[string][]string{
		"Accept": {
			{{.AcceptContentType}},
		},
	}

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, {{.Errorf}}("sending HTTP request (%v)", err)
	}

	ch := make(chan {{.MethodReturnAsync}}, 1)
	go {{.ProcessReturnAsync}}(ctx, ch, resp.Body)
	return ch, nil
}

func {{.ProcessReturnAsync}}(ctx {{.Context}}, ch chan<- {{.MethodReturnAsync}}, r {{.IOReader}}) {
	defer close(ch)
	for {
		if ctx.Err() != nil {
			ch <- {{.MethodReturnAsync}}{Err: {{.ErrContext}}{Cause: ctx.Err()}} // context cancelled
			return
		}

		n, err := {{.IPLDDecodeStreaming}}(r, {{.DAGJSONDecode}})
		if {{.ErrorsIs}}(err, {{.IOEOF}}) || {{.ErrorsIs}}(err, {{.IOErrUnexpectedEOF}}) {
			return
		}
		if err != nil {
			ch <- {{.MethodReturnAsync}}{Err: {{.ErrProto}}{Cause: err}} // IPLD decode error
			return
		}
		env := &{{.ReturnEnvelope}}{}
		if err = env.Parse(n); err != nil {
			ch <- {{.MethodReturnAsync}}{Err: {{.ErrProto}}{Cause: err}} // schema decode error
			return
		}

		if env.Error != nil {
			ch <- {{.MethodReturnAsync}}{Err: {{.ErrService}}{Cause: {{.ErrorsNew}}(string(env.Error.Code))}} // service-level error
			return
		}
		if env.{{.MethodName}} == nil {
			continue
		}
		ch <- {{.MethodReturnAsync}}{Resp: env.{{.MethodName}}}
	}
}`
)
