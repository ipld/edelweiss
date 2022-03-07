package compile

import (
	"fmt"

	cg "github.com/ipld/edelweiss/codegen"
	"github.com/ipld/edelweiss/def"
)

type genPlan struct {
	goPkgPath string
	defToGo   map[def.Def]cg.GoTypeRef
	nameToDef map[string]def.Def
	refs      map[string]bool
	plan      []typePlan
}

type typePlan struct {
	Name  string
	Def   def.Def
	GoRef cg.GoTypeRef
}

func newGenPlan(goPkgPath string) *genPlan {
	return &genPlan{
		goPkgPath: goPkgPath,
		defToGo:   cg.DefToGoTypeRef{},
		plan:      []typePlan{},
		nameToDef: map[string]def.Def{},
		refs:      map[string]bool{},
	}
}

func (p *genPlan) DefToGo() cg.DefToGoTypeRef {
	return p.defToGo
}

func (p *genPlan) Plan() []typePlan {
	return p.plan
}

func (p *genPlan) AddNamed(name string, d def.Def) error {
	if _, ok := p.nameToDef[name]; ok {
		return fmt.Errorf("name %s already defined", name)
	}
	goTypeRef := cg.GoTypeRef{PkgPath: p.goPkgPath, TypeName: name}
	p.defToGo[def.Ref{Name: name}] = goTypeRef
	p.plan = append(p.plan, typePlan{Name: name, Def: d, GoRef: goTypeRef})
	p.nameToDef[name] = d
	return nil
}

func (p *genPlan) AddAnonymous(t def.Def) def.Ref {
	name := fmt.Sprintf("Anon%s%d", t.Kind(), len(p.plan))
	p.AddNamed(name, t)
	return def.Ref{Name: name}
}

func (p *genPlan) AddBuiltin(t def.Def, goTypeRef cg.GoTypeRef) def.Def {
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
