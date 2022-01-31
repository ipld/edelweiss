package blueprint

import (
	"fmt"

	"github.com/ipld/edelweiss/def"
)

// XXX: handle refs

func Compile(defs def.Types) error {
	nameToDef, err := ComputeNameToDef(defs)
	if err != nil {
		return err
	}
	defToGoTypeRef, refs, err := AssignDefToGoTypeRef(defs) // these definitions are codegen implementation targets
	if err != nil {
		return err
	}
	refToGoTypeRef, err := LinkRefToGoTypeRef(refs, nameToDef, defToGoTypeRef)
	if err != nil {
		return err
	}
	_ = refToGoTypeRef
	// defToGoTypeImpl, err := BuildDefToGoTypeImpl(defToGoTypeRef, refToGoTypeRef)
	// if err != nil {
	// 	return err
	// }
	panic("xxx")
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

func AssignDefToGoTypeRef(defs def.Types) (DefToGoTypeRef, def.Refs, error) {
	defToGo := DefToGoTypeRef{} // all defs that must be named and implemented in go
	refs := def.Refs{}          // references found throughout type definitions
	for _, typeDef := range defs {
		switch t := typeDef.(type) {
		case def.Named:
			if err := assignDefToGoTypeRef(defToGo, refs, t.Type, &GoTypeRef{
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

func assignDefToGoTypeRef(defToGo DefToGoTypeRef, refs def.Refs, typeDef def.Type, goTypeRef *GoTypeRef) error {
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
		if err := assignDefToGoTypeRef(defToGo, refs, d, nil); err != nil {
			return err
		}
	}
	return nil
}

func makeTypeName(defToGo DefToGoTypeRef, typeDef def.Type) string {
	return fmt.Sprintf("Anon%s%d", typeDef.Kind(), len(defToGo))
}

// link refs to go type refs

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
