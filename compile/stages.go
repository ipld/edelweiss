package compile

import (
	"fmt"

	cg "github.com/ipld/edelweiss/codegen"
	"github.com/ipld/edelweiss/defs"
	"github.com/ipld/edelweiss/plans"
	"github.com/ipld/edelweiss/values"
)

// generate returns a resolvable (an object that is resolvable to a go type reference)
func generate(p *genPlan, s defs.Def) (plans.BuiltinOrRefPlan, error) {
	switch t := s.(type) {

	case defs.Named:
		plan, err := provision(p, t.Name, t.Type)
		if err != nil {
			return nil, err
		}
		if _, err = p.AddNamed(t.Name, plan); err != nil {
			return nil, err
		}
		return plans.Ref{Name: t.Name}, nil

	case defs.Ref:
		return plans.Ref{Name: t.Name}, nil

	case defs.Bool:
		return p.AddBuiltin(plans.Bool{}, cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Bool"}), nil
	case defs.Int:
		return p.AddBuiltin(plans.Int{}, cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Int"}), nil
	case defs.Float:
		return p.AddBuiltin(plans.Float{}, cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Float"}), nil
	case defs.Byte:
		return p.AddBuiltin(plans.Byte{}, cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Byte"}), nil
	case defs.Char:
		return p.AddBuiltin(plans.Char{}, cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Char"}), nil
	case defs.String:
		return p.AddBuiltin(plans.String{}, cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "String"}), nil
	case defs.Bytes:
		return p.AddBuiltin(plans.Bytes{}, cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Bytes"}), nil
	case defs.Any:
		return p.AddBuiltin(plans.Any{}, cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Any"}), nil
	case defs.Nothing:
		return p.AddBuiltin(plans.Nothing{}, cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Nothing"}), nil

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
func provision(p *genPlan, named string, s defs.Def) (plans.Plan, error) {
	switch t := s.(type) {

	case defs.Named:
		return generate(p, t)

	case defs.Ref:
		return p.AddRef(t.Name), nil

	case defs.Bool:
		return plans.Bool{}, nil
	case defs.Int:
		return plans.Int{}, nil
	case defs.Float:
		return plans.Float{}, nil
	case defs.Byte:
		return plans.Byte{}, nil
	case defs.Char:
		return plans.Char{}, nil
	case defs.String:
		return plans.String{}, nil
	case defs.Bytes:
		return plans.Byte{}, nil
	case defs.Any:
		return plans.Any{}, nil
	case defs.Nothing:
		return plans.Nothing{}, nil

	case defs.Structure:
		fields := t.Fields
		fieldPlans := make([]plans.Field, len(fields))
		for i, f := range fields {
			ftp, err := generate(p, f.Type)
			if err != nil {
				return nil, err
			}
			fieldPlans[i] = plans.Field{Name: f.Name, GoName: f.GoName, Type: ftp}
		}
		return plans.Structure{Fields: fieldPlans}, nil

	case defs.Inductive:
		cases := t.Cases
		casePlans := make([]plans.Case, len(cases))
		for i, c := range cases {
			ctp, err := generate(p, c.Type)
			if err != nil {
				return nil, err
			}
			casePlans[i] = plans.Case{Name: c.Name, GoName: c.GoName, Type: ctp}
		}
		var dp plans.BuiltinOrRefPlan
		switch {
		case t.Default.GoKeyName == "" && t.Default.GoValueName == "" && t.Default.Type == nil:
		case t.Default.GoKeyName != "" && t.Default.GoValueName != "" && t.Default.Type != nil:
			var err error
			if dp, err = generate(p, t.Default.Type); err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("invalid inductive default case definition")
		}
		return plans.Inductive{
			Cases: casePlans,
			Default: plans.DefaultCase{
				GoKeyName:   t.Default.GoKeyName,
				GoValueName: t.Default.GoValueName,
				Type:        dp,
			},
		}, nil

	case defs.Union:
		cases := t.Cases
		casePlans := make([]plans.Case, len(cases))
		for i, c := range cases {
			ctp, err := generate(p, c.Type)
			if err != nil {
				return nil, err
			}
			casePlans[i] = plans.Case{Name: c.Name, GoName: c.GoName, Type: ctp}
		}
		return plans.Union{Cases: casePlans}, nil

	case defs.Tuple:
		slots := t.Slots
		slotPlans := make(plans.Slots, len(slots))
		for i, s := range slots {
			sp, err := generate(p, s)
			if err != nil {
				return nil, err
			}
			slotPlans[i] = sp
		}
		return plans.Tuple{Slots: slotPlans}, nil

	case defs.List:
		ep, err := generate(p, t.Element)
		if err != nil {
			return nil, err
		}
		return plans.List{Element: ep}, nil

	case defs.Link:
		tp, err := generate(p, t.To)
		if err != nil {
			return nil, err
		}
		return plans.Link{To: tp}, nil

	case defs.Map:
		kp, err := generate(p, t.Key)
		if err != nil {
			return nil, err
		}
		vp, err := generate(p, t.Value)
		if err != nil {
			return nil, err
		}
		return plans.Map{Key: kp, Value: vp}, nil

	case defs.SingletonBool:
		return plans.SingletonBool(t), nil
	case defs.SingletonByte:
		return plans.SingletonByte(t), nil
	case defs.SingletonChar:
		return plans.SingletonChar(t), nil
	case defs.SingletonFloat:
		return plans.SingletonFloat(t), nil
	case defs.SingletonInt:
		return plans.SingletonInt(t), nil
	case defs.SingletonString:
		return plans.SingletonString(t), nil

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
		return plans.Call{ID: id, Fn: plans.Fn{Arg: arg, Return: r}}, nil

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
		return plans.Return{ID: id, Fn: plans.Fn{Arg: arg, Return: r}}, nil

	case defs.Fn:
		return nil, fmt.Errorf("%#v cannot be provisioned outside of a service context", t)

	case defs.Service:
		return provisionService(p, named, t)

	}

	return nil, fmt.Errorf("unrecognized definition %#v for provisioning", s)
}

func provisionService(p *genPlan, named string, s defs.Service) (plans.Plan, error) {
	// inject identify method
	for _, m := range s.Methods {
		if m.Name == plans.IdentifyName {
			return nil, fmt.Errorf("method Identify is reserved")
		}
	}
	methods := append(defs.Methods{
		defs.Method{
			Name: plans.IdentifyName,
			Type: defs.Fn{
				Arg: defs.Named{
					Name: fmt.Sprintf("%s_IdentifyArg", named),
					Type: defs.Structure{},
				},
				Return: defs.Named{
					Name: fmt.Sprintf("%s_IdentifyResult", named),
					Type: defs.Structure{
						Fields: defs.Fields{
							defs.Field{Name: "Methods", GoName: "Methods", Type: defs.List{Element: defs.String{}}},
						},
					},
				},
			},
			Cachable: true,
		},
	}, s.Methods...)

	// prepare service plan
	plan := plans.Service{
		Methods: make(plans.Methods, len(methods)),
	}
	callCases, returnCases := make(plans.Cases, len(methods)), make(plans.Cases, len(methods)+1)
	for i, m := range methods {
		argRef, err := generate(p, m.Type.Arg)
		if err != nil {
			return nil, fmt.Errorf("generating method argument (%v)", err)
		}
		returnRef, err := generate(p, m.Type.Return)
		if err != nil {
			return nil, fmt.Errorf("generating method return (%v)", err)
		}
		fn := plans.Fn{Arg: argRef, Return: returnRef}
		if m.Name == plans.IdentifyName {
			plan.Identify = plans.Method{Name: plans.IdentifyName, Type: fn}
		}
		plan.Methods[i] = plans.Method{Name: m.Name, Type: fn, Cachable: m.Cachable}
		callCases[i] = plans.Case{Name: m.Name + "Request", GoName: m.Name, Type: argRef}
		returnCases[i] = plans.Case{Name: m.Name + "Response", GoName: m.Name, Type: returnRef}
	}
	errorEnvelope := plans.Structure{
		Fields: plans.Fields{
			plans.Field{Name: "Code", GoName: "Code", Type: plans.String{}},
		},
	}
	var err error
	errEnvName := fmt.Sprintf("%s_Error", named)
	if plan.ErrorEnvelope, err = p.AddNamed(errEnvName, errorEnvelope); err != nil {
		return nil, fmt.Errorf("service envelope name %s already in use", errEnvName)
	}
	returnCases[len(returnCases)-1] = plans.Case{Name: "Error", Type: plan.ErrorEnvelope}
	//
	// callEnvelopeRef, err := generate(p, plans.Inductive{Cases: callCases})
	// if err != nil {
	// 	return nil, fmt.Errorf("generating service call envelope (%v)", err)
	// }
	// returnEnvelopeRef, err := generate(p, plans.Inductive{Cases: returnCases})
	// if err != nil {
	// 	return nil, fmt.Errorf("generating service return envelope (%v)", err)
	// }
	//
	// plan.CallEnvelope, plan.ReturnEnvelope = callEnvelopeRef, returnEnvelopeRef
	plan.CallEnvelope = p.AddAnonymous(plans.Inductive{Cases: callCases})
	plan.ReturnEnvelope = p.AddAnonymous(plans.Inductive{Cases: returnCases})
	return plan, nil
}
