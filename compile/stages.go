package compile

import (
	"fmt"

	cg "github.com/ipld/edelweiss/codegen"
	"github.com/ipld/edelweiss/def"
	"github.com/ipld/edelweiss/values"
)

// generate returns a resolvable (an object that is resolvable to a go type reference)
func generate(p *genPlan, s def.Type) (def.Type, error) {
	switch t := s.(type) {

	case def.Named:
		plan, err := provision(p, t.Name, t.Type)
		if err != nil {
			return nil, err
		}
		p.AddNamed(t.Name, plan)
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
		def.SingletonBool, def.SingletonByte, def.SingletonChar, def.SingletonFloat, def.SingletonInt, def.SingletonString:
		plan, err := provision(p, "", t)
		if err != nil {
			return nil, err
		}
		return p.AddAnonymous(plan), nil

	case def.Fn, def.Call, def.Return:
		return nil, fmt.Errorf("%#v cannot be generated outside of a service context", t)

	case def.Service:
		return nil, fmt.Errorf("anonymous service %#v cannot be generated", t)
	}

	return nil, fmt.Errorf("%#v cannot be generated", s)
}

// provision returns a generation plan
func provision(p *genPlan, named string, s def.Type) (def.Type, error) {
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
		slotPlans := make([]def.Type, len(slots))
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

	case def.Fn, def.Call, def.Return:
		return nil, fmt.Errorf("%#v cannot be provisioned outside of a service context", t)

	case def.Service:
		XXX
		return nil, fmt.Errorf("anonymous service %#v cannot be generated", t)

	}

	return nil, fmt.Errorf("XXX")
}
