package dagjsonoverhttp

import (
	"fmt"

	"github.com/ipld/edelweiss/blueprints/base"
	cg "github.com/ipld/edelweiss/codegen"
	"github.com/ipld/edelweiss/def"
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

func (x GoServerImpl) ProtoDef() def.Def {
	return x.Def
}

func (x GoServerImpl) GoTypeRef() cg.GoTypeRef {
	return x.Ref
}

func (x GoServerImpl) GoDef() cg.Blueprint {
	loggerVar := cg.GoRef{PkgPath: x.Ref.PkgPath, Name: fmt.Sprintf("logger_server_%s", x.Ref.TypeName)}
	//
	methods := x.Def.Methods
	methodDecls := make(cg.BlueSlice, len(methods))
	methodCases := make(cg.BlueSlice, len(methods))
	for i, m := range methods {
		bmDecl := cg.BlueMap{
			"MethodName":   cg.V(m.Name),
			"MethodArg":    x.Lookup.LookupDepGoRef(m.Type.Arg),
			"MethodReturn": x.Lookup.LookupDepGoRef(m.Type.Return),
			//
			"LoggerVar":         loggerVar,
			"ReturnEnvelope":    x.Lookup.LookupDepGoRef(x.Def.ReturnEnvelope),
			"Context":           base.Context,
			"ContextBackground": base.ContextBackground,
			"IPLDEncode":        base.IPLDEncode,
			"DAGJSONEncode":     base.DAGJSONEncode,
		}
		methodDecls[i] = cg.T{
			Data: bmDecl,
			Src:  `{{.MethodName}}(ctx {{.Context}}, req *{{.MethodArg}}, respCh chan<- *{{.MethodReturn}}) error`,
		}
		methodCases[i] = cg.T{
			Data: bmDecl,
			Src: `
		case env.{{.MethodName}} != nil:
			ch := make(chan *{{.MethodReturn}})
			if err = s.{{.MethodName}}({{.ContextBackground}}(), env.{{.MethodName}}, ch); err != nil {
				{{.LoggerVar}}.Errorf("get p2p provider rejected request (%v)", err)
				writer.WriteHeader(500)
				return
			}
			for resp := range ch {
				env := &{{.ReturnEnvelope}}{ {{.MethodName}}: resp }
				buf, err := {{.IPLDEncode}}(env, {{.DAGJSONEncode}})
				if err != nil {
					{{.LoggerVar}}.Errorf("cannot encode response (%v)", err)
					continue
				}
				writer.Write(buf)
			}
`,
		}
	}
	//
	data := cg.BlueMap{
		"Errorf":             base.Errorf,
		"HTTPHandlerFunc":    base.HTTPHandlerFunc,
		"HTTPRequest":        base.HTTPRequest,
		"HTTPResponseWriter": base.HTTPResponseWriter,
		"IPLDDecode":         base.IPLDDecode,
		"DAGJSONDecode":      base.DAGJSONDecode,
		//
		"Interface":    x.Ref.Append("_Server"),
		"AsyncHandler": x.Ref.Append("_AsyncHandler"),
		"Logger":       cg.GoRef{PkgPath: "github.com/ipfs/go-log", Name: "Logger"},
		"LoggerName":   cg.StringLiteral(fmt.Sprintf("service/server/%s", x.Ref.TypeName)),
		"LoggerVar":    loggerVar,
		//
		"CallEnvelope": x.Lookup.LookupDepGoRef(x.Def.CallEnvelope),
		"MethodDecls":  methodDecls,
		"MethodCases":  methodCases,
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
		msg := request.URL.Query().Get("q")
		n, err := {{.IPLDDecode}}([]byte(msg), {{.DAGJSONDecode}})
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

		// demultiplex request
		switch {
{{range .MethodCases}}{{.}}{{end}}
		default:
			{{.LoggerVar}}.Errorf("missing or unknown request")
			writer.WriteHeader(404)
		}
	}
}
`
)
