package compile

import (
	"fmt"

	cg "github.com/ipld/edelweiss/codegen"
	"github.com/ipld/edelweiss/defs"
)

type genPlan struct {
	goPkgPath string
	defToGo   map[defs.Def]cg.GoTypeRef
	nameToDef map[string]defs.Def
	refs      map[string]bool
	plan      []typePlan
}

type typePlan struct {
	Name  string
	Def   defs.Def
	GoRef cg.GoTypeRef
}

func newGenPlan(goPkgPath string) *genPlan {
	return &genPlan{
		goPkgPath: goPkgPath,
		defToGo:   cg.DefToGoTypeRef{},
		plan:      []typePlan{},
		nameToDef: map[string]defs.Def{},
		refs:      map[string]bool{},
	}
}

func (p *genPlan) DefToGo() cg.DefToGoTypeRef {
	return p.defToGo
}

func (p *genPlan) Plan() []typePlan {
	return p.plan
}

func (p *genPlan) AddNamed(name string, d defs.Def) error {
	if _, ok := p.nameToDef[name]; ok {
		return fmt.Errorf("name %s already defined", name)
	}
	goTypeRef := cg.GoTypeRef{PkgPath: p.goPkgPath, TypeName: name}
	p.defToGo[defs.Ref{Name: name}] = goTypeRef
	p.plan = append(p.plan, typePlan{Name: name, Def: d, GoRef: goTypeRef})
	p.nameToDef[name] = d
	return nil
}

func (p *genPlan) AddAnonymous(t defs.Def) defs.Ref {
	name := fmt.Sprintf("Anon%s%d", t.Kind(), len(p.plan))
	p.AddNamed(name, t)
	return defs.Ref{Name: name}
}

func (p *genPlan) AddBuiltin(t defs.Def, goTypeRef cg.GoTypeRef) defs.Def {
	p.defToGo[t] = goTypeRef
	return t
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
