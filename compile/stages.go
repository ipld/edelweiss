package compile

import (
	"fmt"

	cg "github.com/ipld/edelweiss/codegen"
	"github.com/ipld/edelweiss/def"
	"github.com/ipld/edelweiss/plans"
	"github.com/ipld/edelweiss/values"
)

// generate returns a resolvable (an object that is resolvable to a go type reference)
func generate(p *genPlan, s def.Def) (def.Def, error) {
	switch t := s.(type) {

	case def.Named:
		plan, err := provision(p, t.Name, t.Type)
		if err != nil {
			return nil, err
		}
		if err = p.AddNamed(t.Name, plan); err != nil {
			return nil, err
		}
		return def.Ref{Name: t.Name}, nil

	case def.Ref:
		return t, nil

	case def.Bool:
		return p.AddBuiltin(t, cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Bool"}), nil
	case def.Int:
		return p.AddBuiltin(t, cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Int"}), nil
	case def.Float:
		return p.AddBuiltin(t, cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Float"}), nil
	case def.Byte:
		return p.AddBuiltin(t, cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Byte"}), nil
	case def.Char:
		return p.AddBuiltin(t, cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Char"}), nil
	case def.String:
		return p.AddBuiltin(t, cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "String"}), nil
	case def.Bytes:
		return p.AddBuiltin(t, cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Bytes"}), nil
	case def.Any:
		return p.AddBuiltin(t, cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Any"}), nil
	case def.Nothing:
		return p.AddBuiltin(t, cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Nothing"}), nil

	case def.Structure, def.Inductive, def.List, def.Union, def.Tuple, def.Link, def.Map,
		def.SingletonBool, def.SingletonByte, def.SingletonChar, def.SingletonFloat, def.SingletonInt, def.SingletonString,
		def.Call, def.Return:
		plan, err := provision(p, "", t)
		if err != nil {
			return nil, err
		}
		return p.AddAnonymous(plan), nil

	case def.Fn:
		return nil, fmt.Errorf("%#v cannot be generated outside of a service context", t)

	case def.Service:
		return nil, fmt.Errorf("anonymous service %#v cannot be generated", t)
	}

	return nil, fmt.Errorf("%#v cannot be generated", s)
}

// provision returns a generation plan
func provision(p *genPlan, named string, s def.Def) (def.Def, error) {
	switch t := s.(type) {

	case def.Named:
		return generate(p, t)

	case def.Ref:
		p.AddRef(t.Name)
		return t, nil

	case def.Bool, def.Int, def.Float, def.Byte, def.Char, def.String, def.Bytes, def.Any, def.Nothing:
		return t, nil

	case def.Structure:
		fields := def.FlattenFieldList(t.Fields)
		fieldPlans := make([]def.Field, len(fields))
		for i, f := range fields {
			ftp, err := generate(p, f.Type)
			if err != nil {
				return nil, err
			}
			fieldPlans[i] = def.Field{Name: f.Name, Type: ftp}
		}
		return def.MakeStructure(fieldPlans...), nil

	case def.Inductive:
		cases := def.FlattenCaseList(t.Cases)
		casePlans := make([]def.Case, len(cases))
		for i, c := range cases {
			ctp, err := generate(p, c.Type)
			if err != nil {
				return nil, err
			}
			casePlans[i] = def.Case{Name: c.Name, Type: ctp}
		}
		return def.MakeInductive(casePlans...), nil

	case def.Union:
		cases := def.FlattenCaseList(t.Cases)
		casePlans := make([]def.Case, len(cases))
		for i, c := range cases {
			ctp, err := generate(p, c.Type)
			if err != nil {
				return nil, err
			}
			casePlans[i] = def.Case{Name: c.Name, Type: ctp}
		}
		return def.MakeInductive(casePlans...), nil

	case def.Tuple:
		slots := def.FlattenSlotList(t.Slots)
		slotPlans := make([]def.Def, len(slots))
		for i, s := range slots {
			sp, err := generate(p, s)
			if err != nil {
				return nil, err
			}
			slotPlans[i] = sp
		}
		return def.MakeTuple(slotPlans...), nil

	case def.List:
		ep, err := generate(p, t.Element)
		if err != nil {
			return nil, err
		}
		return def.List{Element: ep}, nil

	case def.Link:
		tp, err := generate(p, t.To)
		if err != nil {
			return nil, err
		}
		return def.Link{To: tp}, nil

	case def.Map:
		kp, err := generate(p, t.Key)
		if err != nil {
			return nil, err
		}
		vp, err := generate(p, t.Value)
		if err != nil {
			return nil, err
		}
		return def.Map{Key: kp, Value: vp}, nil

	case def.SingletonBool, def.SingletonByte, def.SingletonChar, def.SingletonFloat, def.SingletonInt, def.SingletonString:
		return t, nil

	case def.Call:
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
		return def.Call{ID: id, Fn: def.Fn{Arg: arg, Return: r}}, nil

	case def.Return:
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
		return def.Return{ID: id, Fn: def.Fn{Arg: arg, Return: r}}, nil

	case def.Fn:
		return nil, fmt.Errorf("%#v cannot be provisioned outside of a service context", t)

	case def.Service:
		return provisionService(p, named, t)

	}

	return nil, fmt.Errorf("unrecognized definition %#v for provisioning", s)
}

func provisionService(p *genPlan, named string, s def.Service) (def.Def, error) {
	methods := def.FlattenMethodList(s.Methods)
	plan := plans.Service{
		Methods: make([]def.Method, len(methods)),
	}
	callCases, returnCases := make([]def.Case, len(methods)), make([]def.Case, len(methods))
	for i, m := range methods {
		argRef, err := generate(p, m.Type.Arg)
		if err != nil {
			return nil, fmt.Errorf("generating method argument (%v)", err)
		}
		returnRef, err := generate(p, m.Type.Return)
		if err != nil {
			return nil, fmt.Errorf("generating method return (%v)", err)
		}
		fn := def.Fn{Arg: argRef, Return: returnRef}
		plan.Methods[i] = def.Method{Name: m.Name, Type: fn}
		callCases[i] = def.Case{Name: m.Name, Type: argRef}
		returnCases[i] = def.Case{Name: m.Name, Type: returnRef}
	}
	callEnvelopeRef, err := generate(p, def.MakeInductive(callCases...))
	if err != nil {
		return nil, fmt.Errorf("generating service call envelope (%v)", err)
	}
	returnEnvelopeRef, err := generate(p, def.MakeInductive(returnCases...))
	if err != nil {
		return nil, fmt.Errorf("generating service call envelope (%v)", err)
	}
	plan.CallEnvelope, plan.ReturnEnvelope = callEnvelopeRef, returnEnvelopeRef
	return plan, nil
}
