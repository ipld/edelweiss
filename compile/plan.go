package compile

import (
	"fmt"

	cg "github.com/ipld/edelweiss/codegen"
	"github.com/ipld/edelweiss/def"
)

type genPlan struct {
	goPkgPath string
	defToGo   map[def.Type]cg.GoTypeRef
	nameToDef map[string]def.Type
	refs      map[string]bool
	plan      []typePlan
}

type typePlan struct {
	Name  string
	Def   def.Type
	GoRef cg.GoTypeRef
}

func newGenPlan(goPkgPath string) *genPlan {
	return &genPlan{
		goPkgPath: goPkgPath,
		defToGo:   cg.DefToGoTypeRef{},
		plan:      []typePlan{},
		nameToDef: map[string]def.Type{},
		refs:      map[string]bool{},
	}
}

func (p *genPlan) DefToGo() cg.DefToGoTypeRef {
	return p.defToGo
}

func (p *genPlan) Plan() []typePlan {
	return p.plan
}

func (p *genPlan) AddNamed(name string, d def.Type) {
	goTypeRef := cg.GoTypeRef{PkgPath: p.goPkgPath, TypeName: name}
	// p.defToGo[d] = goTypeRef
	p.defToGo[def.Ref{Name: name}] = goTypeRef
	p.plan = append(p.plan, typePlan{Name: name, Def: d, GoRef: goTypeRef})
	p.nameToDef[name] = d
}

func (p *genPlan) AddAnonymous(t def.Type) def.Ref {
	name := fmt.Sprintf("Anon%s%d", t.Kind(), len(p.plan))
	p.AddNamed(name, t)
	return def.Ref{Name: name}
}

func (p *genPlan) AddBuiltin(t def.Type, goTypeRef cg.GoTypeRef) {
	p.defToGo[t] = goTypeRef
}

func (p *genPlan) AddRef(to string) {
	p.refs[to] = true
}

func (p *genPlan) VerifyCompleteness() error {
	for r := range p.refs {
		if _, ok := p.nameToDef[r]; !ok {
			return fmt.Errorf("reference %s cannot be resolved", r)
		}
	}
	return nil
}
