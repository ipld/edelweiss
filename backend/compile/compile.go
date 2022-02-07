package compile

import (
	"fmt"
	"path"

	blue "github.com/ipld/edelweiss/backend/blueprints"
	cg "github.com/ipld/edelweiss/backend/codegen"
	"github.com/ipld/edelweiss/backend/values"
	"github.com/ipld/edelweiss/def"
)

type GoPkgCodegen struct {
	GoPkgDirPath string // local directory for the package
	GoPkgPath    string // global package name
	Defs         def.Types
}

func (x *GoPkgCodegen) GoPkgName() string {
	return path.Base(x.GoPkgPath)
}

func (x *GoPkgCodegen) Compile() (*cg.GoFile, error) {
	nameToDef, err := ComputeNameToDef(x.Defs)
	if err != nil {
		return nil, err
	}
	defToGoTypeRef, refs, err := AssignGoTypeRefToDef(x.GoPkgPath, x.Defs)
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
		FilePath: path.Join(x.GoPkgDirPath, fmt.Sprintf("%s_edelweiss.go", x.GoPkgName())),
		PkgPath:  x.GoPkgPath,
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

func AssignGoTypeRefToDef(goPkgPath string, defs def.Types) (cg.DefToGoTypeRef, def.Refs, error) {
	defToGo := cg.DefToGoTypeRef{} // all defs that must be named and implemented in go
	refs := def.Refs{}             // references found throughout type definitions
	for _, typeDef := range defs {
		switch t := typeDef.(type) {
		case def.Named:
			if err := assignGoTypeRefToDef(goPkgPath, defToGo, refs, t.Type, &cg.GoTypeRef{
				PkgPath:  goPkgPath,
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

func assignGoTypeRefToDef(
	goPkgPath string,
	defToGo cg.DefToGoTypeRef,
	refs def.Refs,
	typeDef def.Type,
	goTypeRef *cg.GoTypeRef,
) error {
	switch t := typeDef.(type) {
	case def.Named:
		return fmt.Errorf("named types must be at the top level")
	case def.Ref:
		refs = append(refs, t)
	}
	if goTypeRef != nil {
		defToGo[typeDef] = *goTypeRef
	} else {
		switch t := typeDef.(type) {
		case def.Ref: // don't name anonymous references

		case def.Bool: // non-codegen types refer to static implementations
			defToGo[t] = cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Bool"}
		case def.Int: // non-codegen types refer to static implementations
			defToGo[t] = cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Int"}
		case def.Float: // non-codegen types refer to static implementations
			defToGo[t] = cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Float"}
		case def.String: // non-codegen types refer to static implementations
			defToGo[t] = cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "String"}
		case def.Byte: // non-codegen types refer to static implementations
			defToGo[t] = cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Byte"}
		case def.Char: // non-codegen types refer to static implementations
			defToGo[t] = cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Char"}
		case def.Any: // non-codegen types refer to static implementations
			defToGo[t] = cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Any"}
		case def.Nothing: // non-codegen types refer to static implementations
			defToGo[t] = cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Nothing"}

		default:
			defToGo[typeDef] = cg.GoTypeRef{
				PkgPath:  goPkgPath,
				TypeName: makeTypeName(defToGo, typeDef),
			}
		}
	}
	for _, d := range typeDef.Deps() {
		if err := assignGoTypeRefToDef(goPkgPath, defToGo, refs, d, nil); err != nil {
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
	case def.SingletonBool, def.SingletonFloat, def.SingletonInt, def.SingletonByte, def.SingletonChar, def.SingletonString:
		return blue.BuildSingletonImpl(d, goTypeRef)
	case def.Structure:
		return blue.BuildStructureImpl(plan.DefToGoTypeRef, d, goTypeRef)
	default:
		return nil, fmt.Errorf("unsupported user type definition %#v", typeDef)
	}
}
