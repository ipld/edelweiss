package dagjsonoverhttp

import (
	"fmt"
	"strings"

	"github.com/ipld/edelweiss/blueprints/base"
	cg "github.com/ipld/edelweiss/codegen"
	"github.com/ipld/edelweiss/plans"
)

func BuildServerImpl(
	lookup cg.LookupDepGoRef,
	typeDef plans.Service,
	goTypeRef cg.GoTypeRef,
) cg.GoTypeImpl {
	return &GoServerImpl{
		Lookup: lookup,
		Def:    typeDef,
		Ref:    goTypeRef,
	}
}

type GoServerImpl struct {
	Lookup cg.LookupDepGoRef
	Def    plans.Service
	Ref    cg.GoTypeRef
}

func (x GoServerImpl) GoTypeRef() cg.GoTypeRef {
	return x.Ref
}

func (x GoServerImpl) GoDef() cg.Blueprint {
	loggerVar := cg.GoRef{PkgPath: x.Ref.PkgPath, Name: fmt.Sprintf("logger_server_%s", x.Ref.TypeName)}
	//
	methods := x.Def.Methods
	methodDecls := cg.BlueSlice{}
	methodCases := cg.BlueSlice{}
	for _, m := range methods {
		if m.Name == x.Def.Identify.Name {
			continue
		}
		// async result type is defined by the client codegen
		asyncResultRef := &cg.GoTypeRef{PkgPath: x.Ref.PkgPath, TypeName: x.Ref.TypeName + "_" + m.Name + "_AsyncResult"}
		bmDecl := cg.BlueMap{
			"MethodName":        cg.V(m.Name),
			"MethodArg":         x.Lookup.LookupDepGoRef(m.Type.Arg),
			"MethodReturn":      x.Lookup.LookupDepGoRef(m.Type.Return),
			"MethodReturnAsync": asyncResultRef,
			"MethodCachable":    cg.BlueBool(m.Cachable),
			//
			"LoggerVar":           loggerVar,
			"ErrorEnvelope":       x.Lookup.LookupDepGoRef(x.Def.ErrorEnvelope),
			"ReturnEnvelope":      x.Lookup.LookupDepGoRef(x.Def.ReturnEnvelope),
			"Context":             base.Context,
			"IPLDEncodeStreaming": base.IPLDEncodeStreaming,
			"DAGJSONEncode":       base.DAGJSONEncode,
			"EdelweissString":     base.EdelweissString,
			"EdelweissETag":       base.EdelweissETag,
			"BytesBuffer":         base.BytesBuffer,
			"HTTPFlusher":         base.HTTPFlusher,
			"IOWriter":            base.IOWriter,
		}
		methodDecls = append(methodDecls, cg.T{
			Data: bmDecl,
			Src:  `{{.MethodName}}(ctx {{.Context}}, req *{{.MethodArg}}) (<-chan *{{.MethodReturnAsync}}, error)`,
		})
		methodCases = append(methodCases, cg.T{
			Data: bmDecl,
			Src: `
		case env.{{.MethodName}} != nil:
			{{if not .MethodCachable}}
			if isReqCachable {
				{{.LoggerVar}}.Errorf("non-cachable method called with http GET")
				writer.Header()["Error"] = []string{"non-cachable method called with GET"}
				writer.WriteHeader(500)
				return
			}
			{{end}}
			ch, err := s.{{.MethodName}}(request.Context(), env.{{.MethodName}})
			if err != nil {
				{{.LoggerVar}}.Errorf("service rejected request (%v)", err)
				writer.Header()["Error"] = []string{err.Error()}
				writer.WriteHeader(500)
				return
			}

			// if the request is cachable, collect all async results in a buffer, otherwise write them directly to http
			var resultWriter {{.IOWriter}}
			if isReqCachable {
				resultWriter = new({{.BytesBuffer}})
			} else {
				writer.WriteHeader(200)
				if f, ok := writer.({{.HTTPFlusher}}); ok {
					f.Flush()
				}
				resultWriter = writer
			}
		chanLoop_{{.MethodName}}:
			for {
				select {
				case <-request.Context().Done():
					return
				case resp, ok := <-ch:
					if !ok {
						break chanLoop_{{.MethodName}}
					}
					var env *{{.ReturnEnvelope}}
					if resp.Err != nil {
						env = &{{.ReturnEnvelope}}{ Error: &{{.ErrorEnvelope}}{Code: {{.EdelweissString}}(resp.Err.Error())} }
					} else {
						env = &{{.ReturnEnvelope}}{ {{.MethodName}}: resp.Resp }
					}
					var buf {{.BytesBuffer}}
					if err = {{.IPLDEncodeStreaming}}(&buf, env, {{.DAGJSONEncode}}); err != nil {
						{{.LoggerVar}}.Errorf("cannot encode response (%v)", err)
						continue chanLoop_{{.MethodName}}
					}
					buf.WriteByte("\n"[0])
					resultWriter.Write(buf.Bytes())
					if f, ok := resultWriter.({{.HTTPFlusher}}); ok {
						f.Flush()
					}
				}
			}
			// if the request is cachable, compute an etag and send the collected results to http
			if isReqCachable {
				result := resultWriter.(*{{.BytesBuffer}}).Bytes()
				etag, err := {{.EdelweissETag}}(result)
				if err != nil {
					{{.LoggerVar}}.Errorf("etag generation (%v)", err)
					writer.Header()["Error"] = []string{err.Error()}
					writer.WriteHeader(500)
					return
				}
				// if the request has an If-None-Match header, respond appropriately
				ifNoneMatchValue := request.Header["If-None-Match"]
				if len(ifNoneMatchValue) == 1 && ifNoneMatchValue[0] == etag {
					writer.WriteHeader(304)
				} else {
					writer.Header()["ETag"] = []string{etag}
					writer.Write(result)
					if f, ok := writer.({{.HTTPFlusher}}); ok {
						f.Flush()
					}
				}
			}
`,
		})
	}
	//
	methodNameStrings := cg.BlueSlice{}
	for _, m := range methods {
		if m.Name == x.Def.Identify.Name {
			continue
		}
		methodNameStrings = append(methodNameStrings, cg.StringLiteral(m.Name))
	}
	identifyData := cg.BlueMap{
		"LoggerVar":           loggerVar,
		"ReturnEnvelope":      x.Lookup.LookupDepGoRef(x.Def.ReturnEnvelope),
		"IPLDEncodeStreaming": base.IPLDEncodeStreaming,
		"DAGJSONEncode":       base.DAGJSONEncode,
		"EdelweissString":     base.EdelweissString,
		"EdelweissETag":       base.EdelweissETag,
		"BytesBuffer":         base.BytesBuffer,
		"HTTPFlusher":         base.HTTPFlusher,
		//
		"IdentifyMethodName":   cg.V(x.Def.Identify.Name),
		"IdentifyMethodArg":    x.Lookup.LookupDepGoRef(x.Def.Identify.Type.Arg),
		"IdentifyMethodReturn": x.Lookup.LookupDepGoRef(x.Def.Identify.Type.Return),
		"MethodNameStrings":    methodNameStrings,
	}
	//
	data := cg.BlueMap{
		"Errorf":             base.Errorf,
		"HTTPHandlerFunc":    base.HTTPHandlerFunc,
		"HTTPRequest":        base.HTTPRequest,
		"HTTPResponseWriter": base.HTTPResponseWriter,
		"IPLDDecode":         base.IPLDDecode,
		"DAGJSONDecode":      base.DAGJSONDecode,
		"DAGCBORDecode":      base.DAGCBORDecode,
		"IOUtilReadAll":      base.IOUtilReadAll,
		//
		"Interface":    x.Ref.Append("_Server"),
		"AsyncHandler": x.Ref.Append("_AsyncHandler"),
		"Logger":       cg.GoRef{PkgPath: "github.com/ipfs/go-log", Name: "Logger"},
		"LoggerName":   cg.StringLiteral(fmt.Sprintf("service/server/%s", strings.ToLower(x.Ref.TypeName))),
		"LoggerVar":    loggerVar,
		//
		"CallEnvelope": x.Lookup.LookupDepGoRef(x.Def.CallEnvelope),
		"MethodDecls":  methodDecls,
		"MethodCases":  methodCases,
		"IdentifyCase": cg.T{Data: identifyData, Src: goIdentifyCaseTemplate},
		//
		"ContentType": cg.StringLiteral(ContentTypeV1),
	}
	return cg.T{Data: data, Src: goServerTemplate}
}

const (
	goServerTemplate = `
var {{.LoggerVar}} = {{.Logger}}({{.LoggerName}})

type {{.Interface}} interface {
{{range .MethodDecls}}
	{{.}}{{end}}
}

func {{.AsyncHandler}}(s {{.Interface}}) {{.HTTPHandlerFunc}} {
	return func(writer {{.HTTPResponseWriter}}, request *{{.HTTPRequest}}) {
		// parse request
		env := &{{.CallEnvelope}}{}
		isReqCachable := false
		switch request.Method {
		case "POST":
			isReqCachable = false
			msg, err := {{.IOUtilReadAll}}(request.Body)
			if err != nil {
				{{.LoggerVar}}.Errorf("reading request body (%v)", err)
				writer.WriteHeader(400)
				return
			}
			n, err := {{.IPLDDecode}}(msg, {{.DAGJSONDecode}})
			if err != nil {
				{{.LoggerVar}}.Errorf("received request not decodeable (%v)", err)
				writer.WriteHeader(400)
				return
			}
			if err = env.Parse(n); err != nil {
				{{.LoggerVar}}.Errorf("parsing call envelope (%v)", err)
				writer.WriteHeader(400)
				return
			}
		case "GET":
			isReqCachable = true
			msg := request.URL.Query().Get("q")
			n, err := {{.IPLDDecode}}([]byte(msg), {{.DAGCBORDecode}})
			if err != nil {
				{{.LoggerVar}}.Errorf("received url not decodeable (%v)", err)
				writer.WriteHeader(400)
				return
			}
			if err = env.Parse(n); err != nil {
				{{.LoggerVar}}.Errorf("parsing call envelope (%v)", err)
				writer.WriteHeader(400)
				return
			}
		default:
			{{.LoggerVar}}.Errorf("http method not supported")
			writer.WriteHeader(400)
			return
		}
		_ = isReqCachable

		writer.Header()["Content-Type"] = []string{
			{{.ContentType}},
		}

		// demultiplex request
		var err error
		switch {
{{range .MethodCases}}{{.}}{{end}}
{{.IdentifyCase}}
		default:
			{{.LoggerVar}}.Errorf("missing or unknown request")
			writer.WriteHeader(404)
		}
	}
}
`
	goIdentifyCaseTemplate = `
		case env.{{.IdentifyMethodName}} != nil:
			var env *{{.ReturnEnvelope}}
			env = &{{.ReturnEnvelope}}{
				{{.IdentifyMethodName}}: &{{.IdentifyMethodReturn}}{
					Methods: []{{.EdelweissString}}{
{{range .MethodNameStrings}}						{{.}},
{{end}}
					},
				},
			}
			var buf {{.BytesBuffer}}
			if err = {{.IPLDEncodeStreaming}}(&buf, env, {{.DAGJSONEncode}}); err != nil {
				{{.LoggerVar}}.Errorf("cannot encode identify response (%v)", err)
				writer.WriteHeader(500)
				return
			}
			buf.WriteByte("\n"[0])

			// compute etag, since Identify is cachable
			result := buf.Bytes()
			etag, err := {{.EdelweissETag}}(result)
			if err != nil {
				{{.LoggerVar}}.Errorf("etag generation (%v)", err)
				writer.Header()["Error"] = []string{err.Error()}
				writer.WriteHeader(500)
				return
			}
			// if the request has an If-None-Match header, respond appropriately
			ifNoneMatchValue := request.Header["If-None-Match"]
			if len(ifNoneMatchValue) == 1 && ifNoneMatchValue[0] == etag {
				writer.WriteHeader(304)
			} else {
				writer.Header()["ETag"] = []string{etag}
				writer.Write(result)
				if f, ok := writer.({{.HTTPFlusher}}); ok {
					f.Flush()
				}
			}
`
)
