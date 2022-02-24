package compile

import (
	"fmt"
	"path"

	blue_services "github.com/ipld/edelweiss/blueprints/services/dagjson-over-http"
	blue_values "github.com/ipld/edelweiss/blueprints/values"
	cg "github.com/ipld/edelweiss/codegen"
	"github.com/ipld/edelweiss/def"
	"github.com/ipld/edelweiss/values"
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
	p, err := processDefs(x.GoPkgPath, x.Defs)
	if err != nil {
		return nil, err
	}
	goTypeImpls, err := buildGoTypeImpls(p.Plan(), p.DefToGo())
	if err != nil {
		return nil, err
	}
	file := &cg.GoFile{
		FilePath: path.Join(x.GoPkgDirPath, fmt.Sprintf("%s_edelweiss.go", x.GoPkgName())),
		PkgPath:  x.GoPkgPath,
		Types:    goTypeImpls,
	}
	return file, nil
}

func processDefs(goPkgPath string, defs def.Types) (*genPlan, error) {
	p := newGenPlan(goPkgPath)
	for _, d := range defs {
		switch t := d.(type) {
		case def.Named:
			p.AddNamed(t.Name, t.Type)
			if err := processDeps(p, t.Type); err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("anonymous type at top level")
		}
	}
	if err := p.VerifyCompleteness(); err != nil {
		return nil, err
	}
	return p, nil
}

func processDeps(p *genPlan, t def.Type) error {
	for _, dep := range t.Deps() {
		if p.IsKnown(dep) {
			continue
		}
		switch t := dep.(type) {
		case def.Named:
			return fmt.Errorf("named types must be at the top level")
		case def.Ref:
			p.AddRef(t.Name)
		// non-parametric types have static/non-codegen implementation
		// whenever we encounter a non-parametric type, we refer to its static implementation
		case def.Bool:
			p.AddBuiltin(t, cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Bool"})
		case def.Int:
			p.AddBuiltin(t, cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Int"})
		case def.Float:
			p.AddBuiltin(t, cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Float"})
		case def.Byte:
			p.AddBuiltin(t, cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Byte"})
		case def.Char:
			p.AddBuiltin(t, cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Char"})
		case def.String:
			p.AddBuiltin(t, cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "String"})
		case def.Bytes:
			p.AddBuiltin(t, cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Bytes"})
		case def.Any:
			p.AddBuiltin(t, cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Any"})
		case def.Nothing:
			p.AddBuiltin(t, cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Nothing"})
		// all other types are anonymous inline parametric types
		default:
			p.AddAnonymous(t)
			// process the dependencies of the dependency
			if err := processDeps(p, dep); err != nil {
				return err
			}
		}
	}
	return nil
}

func buildGoTypeImpls(typeToGen []typePlan, depToGo cg.DefToGoTypeRef) (cg.GoTypeImpls, error) {
	goTypeImpls := cg.GoTypeImpls{}
	for _, ttg := range typeToGen {
		if goTypeImpl, err := buildGoTypeImpl(depToGo, ttg.Def, ttg.GoRef); err != nil {
			return nil, err
		} else {
			if goTypeImpl != nil {
				goTypeImpls = append(goTypeImpls, goTypeImpl)
			}
		}
	}
	return goTypeImpls, nil
}

func buildGoTypeImpl(depToGo cg.DefToGoTypeRef, typeDef def.Type, goTypeRef cg.GoTypeRef) (cg.GoTypeImpl, error) {
	switch d := typeDef.(type) {
	case def.SingletonBool, def.SingletonFloat, def.SingletonInt, def.SingletonByte, def.SingletonChar, def.SingletonString:
		return blue_values.BuildSingletonImpl(d, goTypeRef)
	case def.Structure:
		return blue_values.BuildStructureImpl(depToGo, d, goTypeRef)
	case def.Inductive:
		return blue_values.BuildInductiveImpl(depToGo, d, goTypeRef)
	case def.List:
		return blue_values.BuildListImpl(depToGo, d, goTypeRef)
	case def.Link:
		return blue_values.BuildLinkImpl(depToGo, d, goTypeRef)
	case def.Map:
		return blue_values.BuildMapImpl(depToGo, d, goTypeRef)
	case def.Fn:
		// fn types define functional signatures. they don't have a corresponding value type.
		return nil, nil
	case def.Call:
		return blue_values.BuildCallImpl(depToGo, d, goTypeRef)
	case def.Return:
		return blue_values.BuildReturnImpl(depToGo, d, goTypeRef)
	case def.Service:
		return blue_services.BuildClientImpl(depToGo, d, goTypeRef)
	default:
		return nil, fmt.Errorf("unsupported user type definition %#v", typeDef)
	}
}
