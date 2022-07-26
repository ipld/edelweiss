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
			//
			"LoggerVar":           loggerVar,
			"ErrorEnvelope":       x.Lookup.LookupDepGoRef(x.Def.ErrorEnvelope),
			"ReturnEnvelope":      x.Lookup.LookupDepGoRef(x.Def.ReturnEnvelope),
			"Context":             base.Context,
			"IPLDEncodeStreaming": base.IPLDEncodeStreaming,
			"DAGJSONEncode":       base.DAGJSONEncode,
			"EdelweissString":     base.EdelweissString,
			"BytesBuffer":         base.BytesBuffer,
			"HTTPFlusher":         base.HTTPFlusher,
		}
		methodDecls = append(methodDecls, cg.T{
			Data: bmDecl,
			Src:  `{{.MethodName}}(ctx {{.Context}}, req *{{.MethodArg}}) (<-chan *{{.MethodReturnAsync}}, error)`,
		})
		methodCases = append(methodCases, cg.T{
			Data: bmDecl,
			Src: `
		case env.{{.MethodName}} != nil:
			ch, err := s.{{.MethodName}}(request.Context(), env.{{.MethodName}})
			if err != nil {
				{{.LoggerVar}}.Errorf("service rejected request (%v)", err)
				writer.Header()["Error"] = []string{err.Error()}
				writer.WriteHeader(500)
				return
			}

			writer.WriteHeader(200)
			if f, ok := writer.({{.HTTPFlusher}}); ok {
				f.Flush()
			}

			for {
				select {
				case <-request.Context().Done():
					return
				case resp, ok := <-ch:
					if !ok {
						return
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
						continue
					}
					buf.WriteByte("\n"[0])
					writer.Write(buf.Bytes())
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
		"BytesBuffer":         base.BytesBuffer,
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
		"IOUtilReadAll":      base.IOUtilReadAll,
		//
		"Interface":    x.Ref.Append("_Server"),
		"AsyncHandler": x.Ref.Append("_AsyncHandler"),
		"Logger":       cg.GoRef{PkgPath: "github.com/ipfs/go-log/v2", Name: "Logger"},
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
		env := &{{.CallEnvelope}}{}
		if err = env.Parse(n); err != nil {
			{{.LoggerVar}}.Errorf("parsing call envelope (%v)", err)
			writer.WriteHeader(400)
			return
		}

		writer.Header()["Content-Type"] = []string{
			{{.ContentType}},
		}

		// demultiplex request
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
			writer.Write(buf.Bytes())
`
)
