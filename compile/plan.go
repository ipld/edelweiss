package compile

import (
	cg "github.com/ipld/edelweiss/codegen"
	"github.com/ipld/edelweiss/def"
)

type genPlan struct {
	depToRef  cg.DefToGoTypeRef // deps = (builtin) non-parametric types + anonymous/inline types + references
	typeToGen typesToGen        // types to generate = named types
	names     map[string]bool
	refs      map[string]bool
}

type typeToGen struct {
	Name  string
	Def   def.Type
	GoRef cg.GoTypeRef
}

type typesToGen []typeToGen

func newGenPlan() *genPlan {
	return &genPlan{
		depToRef:  cg.DefToGoTypeRef{},
		typeToGen: typesToGen{},
		names:     map[string]bool{},
		refs:      map[string]bool{},
	}
}

func (p *genPlan) AddNamed(goPkgPath string, name string, d def.Type) {
	goTypeRef := cg.GoTypeRef{PkgPath: goPkgPath, TypeName: name}
	p.depToRef[def.Ref{Name: name}] = goTypeRef
	p.typeToGen = append(p.typeToGen, typeToGen{Name: name, Def: d, GoRef: goTypeRef})
	p.names[name] = true
}

func (p *genPlan) AddBuiltin(t def.Type, goTypeRef cg.GoTypeRef) {
	p.depToRef[t] = goTypeRef
}

func (p *genPlan) IsKnown(t def.Type) bool {
	_, known := p.depToRef[t]
	return known
}

func (p *genPlan) AddRef(to string) {
	p.refs[to] = true
}
