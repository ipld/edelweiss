package blueprint

import (
	"fmt"

	"github.com/ipld/edelweiss/def"
)

func Compile(defs def.Types) error {
	nameToDef, err := ComputeNameToDef(defs)
	if err != nil {
		return err
	}
	defToGoTypeRef, refs, err := AssignGoTypeRefToDef(defs)
	if err != nil {
		return err
	}
	refToGoTypeRef, err := LinkRefToGoTypeRef(refs, nameToDef, defToGoTypeRef)
	if err != nil {
		return err
	}
	plan := GoTypeImplPlan{
		DefToGoTypeRef: defToGoTypeRef,
		RefToGoTypeRef: refToGoTypeRef,
	}
	defToGoTypeImpl, err := BuildGoTypeImpl(plan)
	if err != nil {
		return err
	}
	_ = defToGoTypeImpl
	panic("XXX")
}

// def name -> def

type NameToDef map[string]def.Type

func ComputeNameToDef(defs def.Types) (NameToDef, error) {
	nameToDef := NameToDef{}
	for _, d := range defs {
		switch x := d.(type) {
		case def.Named:
			if _, ok := nameToDef[x.Name]; ok {
				return nil, fmt.Errorf("type %s already defined", x.Name)
			} else {
				nameToDef[x.Name] = x
			}
		default:
			return nil, fmt.Errorf("anonymous top-level type")
		}
	}
	return nameToDef, nil
}

// assign go names to defs: def -> go type ref

type DefToGoTypeRef map[def.Type]GoTypeRef

func AssignGoTypeRefToDef(defs def.Types) (DefToGoTypeRef, def.Refs, error) {
	defToGo := DefToGoTypeRef{} // all defs that must be named and implemented in go
	refs := def.Refs{}          // references found throughout type definitions
	for _, typeDef := range defs {
		switch t := typeDef.(type) {
		case def.Named:
			if err := assignGoTypeRefToDef(defToGo, refs, t.Type, &GoTypeRef{
				PkgPath:  "", // for now everything lives in one package
				TypeName: t.Name,
			}); err != nil {
				return nil, nil, err
			}
		default:
			return nil, nil, fmt.Errorf("anonymous top-level type")
		}
	}
	return defToGo, refs, nil
}

func assignGoTypeRefToDef(defToGo DefToGoTypeRef, refs def.Refs, typeDef def.Type, goTypeRef *GoTypeRef) error {
	switch t := typeDef.(type) {
	case def.Named:
		return fmt.Errorf("named types must be at the top level")
	case def.Ref:
		refs = append(refs, t)
	}
	if goTypeRef != nil {
		defToGo[typeDef] = *goTypeRef
	} else {
		switch typeDef.(type) {
		case def.Ref: // don't name anonymous references
		default:
			defToGo[typeDef] = GoTypeRef{
				PkgPath:  "", // for now everything lives in one package
				TypeName: makeTypeName(defToGo, typeDef),
			}
		}
	}
	for _, d := range typeDef.Deps() {
		if err := assignGoTypeRefToDef(defToGo, refs, d, nil); err != nil {
			return err
		}
	}
	return nil
}

func makeTypeName(defToGo DefToGoTypeRef, typeDef def.Type) string {
	return fmt.Sprintf("Anon%s%d", typeDef.Kind(), len(defToGo))
}

// link refs to go type refs: ref -> go type ref

type RefToGoTypeRef map[def.Ref]GoTypeRef

func LinkRefToGoTypeRef(refs def.Refs, nameToDef NameToDef, defToGoTypeRef DefToGoTypeRef) (RefToGoTypeRef, error) {
	refToGoTypeRef := RefToGoTypeRef{}
	for _, ref := range refs {
		refDef, ok := nameToDef[ref.Name]
		if !ok {
			return nil, fmt.Errorf("reference to undefined user type %s", ref.Name)
		}
		goRef, ok := defToGoTypeRef[refDef]
		if !ok {
			return nil, fmt.Errorf("missing go reference for definition %v", refDef)
		}
		refToGoTypeRef[ref] = goRef
	}
	return refToGoTypeRef, nil
}

// build go implementations for each def: def -> go type impl

type GoTypeImplPlan struct {
	DefToGoTypeRef // definitions that must be code-generated
	RefToGoTypeRef // references used throughout definitions
}

type DefToGoTypeImpl map[def.Type]GoTypeImpl

func BuildGoTypeImpl(plan GoTypeImplPlan) (DefToGoTypeImpl, error) {
	defToGoTypeImpl := DefToGoTypeImpl{}
	for typeDef, goTypeRef := range plan.DefToGoTypeRef {
		if goTypeImpl, err := buildGoTypeImpl(plan, typeDef, goTypeRef); err != nil {
			return nil, err
		} else {
			defToGoTypeImpl[typeDef] = goTypeImpl
		}
	}
	return defToGoTypeImpl, nil
}

func buildGoTypeImpl(plan GoTypeImplPlan, typeDef def.Type, goTypeRef GoTypeRef) (GoTypeImpl, error) {
	switch d := typeDef.(type) {
	case def.SingletonBool:
		return BuildSingletonBoolGoImpl(plan, d, goTypeRef)
	case def.SingletonFloat:
		return BuildSingletonFloatGoImpl(plan, d, goTypeRef)
	case def.SingletonInt:
		return BuildSingletonIntGoImpl(plan, d, goTypeRef)
	case def.SingletonByte:
		return BuildSingletonByteGoImpl(plan, d, goTypeRef)
	case def.SingletonChar:
		return BuildSingletonCharGoImpl(plan, d, goTypeRef)
	case def.SingletonString:
		return BuildSingletonStringGoImpl(plan, d, goTypeRef)
	default:
		return nil, fmt.Errorf("unsupported user type definition %#v", typeDef)
	}
}
