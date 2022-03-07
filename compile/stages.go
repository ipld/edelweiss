package compile

import (
	"fmt"

	cg "github.com/ipld/edelweiss/codegen"
	"github.com/ipld/edelweiss/defs"
	"github.com/ipld/edelweiss/plans"
	"github.com/ipld/edelweiss/values"
)

// generate returns a resolvable (an object that is resolvable to a go type reference)
func generate(p *genPlan, s defs.Def) (defs.Def, error) {
	switch t := s.(type) {

	case defs.Named:
		plan, err := provision(p, t.Name, t.Type)
		if err != nil {
			return nil, err
		}
		if err = p.AddNamed(t.Name, plan); err != nil {
			return nil, err
		}
		return defs.Ref{Name: t.Name}, nil

	case defs.Ref:
		return t, nil

	case defs.Bool:
		return p.AddBuiltin(t, cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Bool"}), nil
	case defs.Int:
		return p.AddBuiltin(t, cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Int"}), nil
	case defs.Float:
		return p.AddBuiltin(t, cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Float"}), nil
	case defs.Byte:
		return p.AddBuiltin(t, cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Byte"}), nil
	case defs.Char:
		return p.AddBuiltin(t, cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Char"}), nil
	case defs.String:
		return p.AddBuiltin(t, cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "String"}), nil
	case defs.Bytes:
		return p.AddBuiltin(t, cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Bytes"}), nil
	case defs.Any:
		return p.AddBuiltin(t, cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Any"}), nil
	case defs.Nothing:
		return p.AddBuiltin(t, cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Nothing"}), nil

	case defs.Structure, defs.Inductive, defs.List, defs.Union, defs.Tuple, defs.Link, defs.Map,
		defs.SingletonBool, defs.SingletonByte, defs.SingletonChar, defs.SingletonFloat, defs.SingletonInt, defs.SingletonString,
		defs.Call, defs.Return:
		plan, err := provision(p, "", t)
		if err != nil {
			return nil, err
		}
		return p.AddAnonymous(plan), nil

	case defs.Fn:
		return nil, fmt.Errorf("%#v cannot be generated outside of a service context", t)

	case defs.Service:
		return nil, fmt.Errorf("anonymous service %#v cannot be generated", t)
	}

	return nil, fmt.Errorf("%#v cannot be generated", s)
}

// provision returns a generation plan
func provision(p *genPlan, named string, s defs.Def) (defs.Def, error) {
	switch t := s.(type) {

	case defs.Named:
		return generate(p, t)

	case defs.Ref:
		p.AddRef(t.Name)
		return t, nil

	case defs.Bool, defs.Int, defs.Float, defs.Byte, defs.Char, defs.String, defs.Bytes, defs.Any, defs.Nothing:
		return t, nil

	case defs.Structure:
		fields := t.Fields
		fieldPlans := make([]defs.Field, len(fields))
		for i, f := range fields {
			ftp, err := generate(p, f.Type)
			if err != nil {
				return nil, err
			}
			fieldPlans[i] = defs.Field{Name: f.Name, Type: ftp}
		}
		return defs.Structure{Fields: fieldPlans}, nil

	case defs.Inductive:
		cases := defs.FlattenCaseList(t.Cases)
		casePlans := make([]defs.Case, len(cases))
		for i, c := range cases {
			ctp, err := generate(p, c.Type)
			if err != nil {
				return nil, err
			}
			casePlans[i] = defs.Case{Name: c.Name, Type: ctp}
		}
		return defs.MakeInductive(casePlans...), nil

	case defs.Union:
		cases := defs.FlattenCaseList(t.Cases)
		casePlans := make([]defs.Case, len(cases))
		for i, c := range cases {
			ctp, err := generate(p, c.Type)
			if err != nil {
				return nil, err
			}
			casePlans[i] = defs.Case{Name: c.Name, Type: ctp}
		}
		return defs.MakeInductive(casePlans...), nil

	case defs.Tuple:
		slots := defs.FlattenSlotList(t.Slots)
		slotPlans := make([]defs.Def, len(slots))
		for i, s := range slots {
			sp, err := generate(p, s)
			if err != nil {
				return nil, err
			}
			slotPlans[i] = sp
		}
		return defs.MakeTuple(slotPlans...), nil

	case defs.List:
		ep, err := generate(p, t.Element)
		if err != nil {
			return nil, err
		}
		return defs.List{Element: ep}, nil

	case defs.Link:
		tp, err := generate(p, t.To)
		if err != nil {
			return nil, err
		}
		return defs.Link{To: tp}, nil

	case defs.Map:
		kp, err := generate(p, t.Key)
		if err != nil {
			return nil, err
		}
		vp, err := generate(p, t.Value)
		if err != nil {
			return nil, err
		}
		return defs.Map{Key: kp, Value: vp}, nil

	case defs.SingletonBool, defs.SingletonByte, defs.SingletonChar, defs.SingletonFloat, defs.SingletonInt, defs.SingletonString:
		return t, nil

	case defs.Call:
		id, err := generate(p, t.ID)
		if err != nil {
			return nil, err
		}
		arg, err := generate(p, t.Fn.Arg)
		if err != nil {
			return nil, err
		}
		r, err := generate(p, t.Fn.Return)
		if err != nil {
			return nil, err
		}
		return defs.Call{ID: id, Fn: defs.Fn{Arg: arg, Return: r}}, nil

	case defs.Return:
		id, err := generate(p, t.ID)
		if err != nil {
			return nil, err
		}
		arg, err := generate(p, t.Fn.Arg)
		if err != nil {
			return nil, err
		}
		r, err := generate(p, t.Fn.Return)
		if err != nil {
			return nil, err
		}
		return defs.Return{ID: id, Fn: defs.Fn{Arg: arg, Return: r}}, nil

	case defs.Fn:
		return nil, fmt.Errorf("%#v cannot be provisioned outside of a service context", t)

	case defs.Service:
		return provisionService(p, named, t)

	}

	return nil, fmt.Errorf("unrecognized definition %#v for provisioning", s)
}

func provisionService(p *genPlan, named string, s defs.Service) (defs.Def, error) {
	methods := defs.FlattenMethodList(s.Methods)
	plan := plans.Service{
		Methods: make([]defs.Method, len(methods)),
	}
	callCases, returnCases := make([]defs.Case, len(methods)), make([]defs.Case, len(methods))
	for i, m := range methods {
		argRef, err := generate(p, m.Type.Arg)
		if err != nil {
			return nil, fmt.Errorf("generating method argument (%v)", err)
		}
		returnRef, err := generate(p, m.Type.Return)
		if err != nil {
			return nil, fmt.Errorf("generating method return (%v)", err)
		}
		fn := defs.Fn{Arg: argRef, Return: returnRef}
		plan.Methods[i] = defs.Method{Name: m.Name, Type: fn}
		callCases[i] = defs.Case{Name: m.Name, Type: argRef}
		returnCases[i] = defs.Case{Name: m.Name, Type: returnRef}
	}
	callEnvelopeRef, err := generate(p, defs.MakeInductive(callCases...))
	if err != nil {
		return nil, fmt.Errorf("generating service call envelope (%v)", err)
	}
	returnEnvelopeRef, err := generate(p, defs.MakeInductive(returnCases...))
	if err != nil {
		return nil, fmt.Errorf("generating service call envelope (%v)", err)
	}
	plan.CallEnvelope, plan.ReturnEnvelope = callEnvelopeRef, returnEnvelopeRef
	return plan, nil
}
