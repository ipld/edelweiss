package compile

import (
	"fmt"

	cg "github.com/ipld/edelweiss/codegen"
	"github.com/ipld/edelweiss/plans"
)

type genPlan struct {
	goPkgPath  string
	planToGo   cg.PlanToGoTypeRef
	nameToPlan map[string]plans.Plan
	refs       map[string]bool
	plan       []typePlan
}

type typePlan struct {
	Name  string
	Plan  plans.Plan
	GoRef cg.GoTypeRef
}

func newGenPlan(goPkgPath string) *genPlan {
	return &genPlan{
		goPkgPath:  goPkgPath,
		planToGo:   cg.PlanToGoTypeRef{},
		plan:       []typePlan{},
		nameToPlan: map[string]plans.Plan{},
		refs:       map[string]bool{},
	}
}

func (p *genPlan) PlanToGo() cg.PlanToGoTypeRef {
	return p.planToGo
}

func (p *genPlan) Plan() []typePlan {
	return p.plan
}

func (p *genPlan) AddNamed(name string, d plans.Plan) (plans.Ref, error) {
	if _, ok := p.nameToPlan[name]; ok {
		return plans.Ref{}, fmt.Errorf("name %s already defined", name)
	}
	goTypeRef := cg.GoTypeRef{PkgPath: p.goPkgPath, TypeName: name}
	ref := plans.Ref{Name: name}
	p.planToGo[ref] = goTypeRef
	p.plan = append(p.plan, typePlan{Name: name, Plan: d, GoRef: goTypeRef})
	p.nameToPlan[name] = d
	return ref, nil
}

func (p *genPlan) AddAnonymous(t plans.Plan) plans.Ref {
	name := fmt.Sprintf("Anon%s%d", t.Kind(), len(p.plan))
	p.AddNamed(name, t)
	return plans.Ref{Name: name}
}

func (p *genPlan) AddBuiltin(t plans.Plan, goTypeRef cg.GoTypeRef) plans.Plan {
	p.planToGo[t] = goTypeRef
	return t
}

func (p *genPlan) AddRef(to string) plans.Ref {
	p.refs[to] = true
	return plans.Ref{Name: to}
}

func (p *genPlan) VerifyCompleteness() error {
	for r := range p.refs {
		if _, ok := p.nameToPlan[r]; !ok {
			return fmt.Errorf("reference %s cannot be resolved", r)
		}
	}
	return nil
}
