package compile

import (
	"fmt"
	"path"

	blue "github.com/ipld/edelweiss/backend/blueprint"
	cg "github.com/ipld/edelweiss/backend/codegen"
	"github.com/ipld/edelweiss/def"
)

type GoPkgCodegen struct {
	GoPkgDirPath string
	GoPkgName    string
	Defs         def.Types
}

func (x *GoPkgCodegen) Compile() (*cg.GoFile, error) {
	nameToDef, err := ComputeNameToDef(x.Defs)
	if err != nil {
		return nil, err
	}
	defToGoTypeRef, refs, err := AssignGoTypeRefToDef(x.Defs)
	if err != nil {
		return nil, err
	}
	refToGoTypeRef, err := LinkRefToGoTypeRef(refs, nameToDef, defToGoTypeRef)
	if err != nil {
		return nil, err
	}
	plan := cg.GoTypeImplPlan{
		DefToGoTypeRef: defToGoTypeRef,
		RefToGoTypeRef: refToGoTypeRef,
	}
	defToGoTypeImpl, err := BuildGoTypeImpl(plan)
	if err != nil {
		return nil, err
	}
	file := &cg.GoFile{
		FilePath: path.Join(x.GoPkgDirPath, fmt.Sprintf("%s_edelweiss.go", x.GoPkgName)),
		PkgName:  x.GoPkgName,
	}
	for _, goTypeImpl := range defToGoTypeImpl {
		file.Types = append(file.Types, goTypeImpl)
	}
	return file, nil
}

// def name -> def

func ComputeNameToDef(defs def.Types) (cg.NameToDef, error) {
	nameToDef := cg.NameToDef{}
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

func AssignGoTypeRefToDef(defs def.Types) (cg.DefToGoTypeRef, def.Refs, error) {
	defToGo := cg.DefToGoTypeRef{} // all defs that must be named and implemented in go
	refs := def.Refs{}             // references found throughout type definitions
	for _, typeDef := range defs {
		switch t := typeDef.(type) {
		case def.Named:
			if err := assignGoTypeRefToDef(defToGo, refs, t.Type, &cg.GoTypeRef{
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

func assignGoTypeRefToDef(defToGo cg.DefToGoTypeRef, refs def.Refs, typeDef def.Type, goTypeRef *cg.GoTypeRef) error {
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
			defToGo[typeDef] = cg.GoTypeRef{
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

func makeTypeName(defToGo cg.DefToGoTypeRef, typeDef def.Type) string {
	return fmt.Sprintf("Anon%s%d", typeDef.Kind(), len(defToGo))
}

// link refs to go type refs: ref -> go type ref

func LinkRefToGoTypeRef(refs def.Refs, nameToDef cg.NameToDef, defToGoTypeRef cg.DefToGoTypeRef) (cg.RefToGoTypeRef, error) {
	refToGoTypeRef := cg.RefToGoTypeRef{}
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

func BuildGoTypeImpl(plan cg.GoTypeImplPlan) (cg.DefToGoTypeImpl, error) {
	defToGoTypeImpl := cg.DefToGoTypeImpl{}
	for typeDef, goTypeRef := range plan.DefToGoTypeRef {
		if goTypeImpl, err := buildGoTypeImpl(plan, typeDef, goTypeRef); err != nil {
			return nil, err
		} else {
			defToGoTypeImpl[typeDef] = goTypeImpl
		}
	}
	return defToGoTypeImpl, nil
}

func buildGoTypeImpl(plan cg.GoTypeImplPlan, typeDef def.Type, goTypeRef cg.GoTypeRef) (cg.GoTypeImpl, error) {
	switch d := typeDef.(type) {
	case def.SingletonBool:
		return blue.BuildSingletonBoolGoImpl(d, goTypeRef)
	case def.SingletonFloat:
		return blue.BuildSingletonFloatGoImpl(d, goTypeRef)
	case def.SingletonInt:
		return blue.BuildSingletonIntGoImpl(d, goTypeRef)
	case def.SingletonByte:
		return blue.BuildSingletonByteGoImpl(d, goTypeRef)
	case def.SingletonChar:
		return blue.BuildSingletonCharGoImpl(d, goTypeRef)
	case def.SingletonString:
		return blue.BuildSingletonStringGoImpl(d, goTypeRef)
	default:
		return nil, fmt.Errorf("unsupported user type definition %#v", typeDef)
	}
}