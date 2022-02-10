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

type genPlan struct {
	depToRef  cg.DefToGoTypeRef // deps = (builtin) non-parametric types + anonymous/inline types + references
	typeToGen typesToGen        // types to generate = named types
	names     map[string]bool
	refs      map[string]bool
}

type typeToGen struct {
	Name  string
	Def   def.Type
	GoRef cg.GoTypeRef
}

type typesToGen []typeToGen

func (x *GoPkgCodegen) Compile() (*cg.GoFile, error) {
	p, err := processDefs(x.GoPkgPath, x.Defs)
	if err != nil {
		return nil, err
	}
	goTypeImpls, err := buildGoTypeImpls(p.typeToGen, p.depToRef)
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
	p := &genPlan{
		depToRef:  cg.DefToGoTypeRef{},
		typeToGen: typesToGen{},
		names:     map[string]bool{},
		refs:      map[string]bool{},
	}
	for _, d := range defs {
		switch t := d.(type) {
		case def.Named:
			goRef := cg.GoTypeRef{PkgPath: goPkgPath, TypeName: t.Name}
			p.depToRef[def.Ref{Name: t.Name}] = goRef
			p.typeToGen = append(p.typeToGen, typeToGen{Name: t.Name, Def: t.Type, GoRef: goRef})
			p.names[t.Name] = true
			if err := processDeps(goPkgPath, p, t.Type); err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("anonymous type at top level")
		}
	}
	for r := range p.refs {
		if !p.names[r] {
			return nil, fmt.Errorf("reference %s cannot be resolved", r)
		}
	}
	return p, nil
}

func processDeps(goPkgPath string, p *genPlan, t def.Type) error {
	for _, dep := range t.Deps() {
		switch t := dep.(type) {
		case def.Named:
			return fmt.Errorf("named types must be at the top level")
		case def.Ref:
			p.refs[t.Name] = true
		// non-parametric types have static/non-codegen implementation
		// whenever we encounter a non-parametric type, we refer to its static implementation
		case def.Bool:
			p.depToRef[t] = cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Bool"}
		case def.Int:
			p.depToRef[t] = cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Int"}
		case def.Float:
			p.depToRef[t] = cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Float"}
		case def.String:
			p.depToRef[t] = cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "String"}
		case def.Byte:
			p.depToRef[t] = cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Byte"}
		case def.Char:
			p.depToRef[t] = cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Char"}
		case def.Any:
			p.depToRef[t] = cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Any"}
		case def.Nothing:
			p.depToRef[t] = cg.GoTypeRef{PkgPath: values.PkgPath, TypeName: "Nothing"}
		// all other types are anonymous inline parametric types
		default:
			name := fmt.Sprintf("Anon%s%d", t.Kind(), len(p.typeToGen))
			goRef := cg.GoTypeRef{PkgPath: goPkgPath, TypeName: name}
			p.depToRef[def.Ref{Name: name}] = goRef
			p.typeToGen = append(p.typeToGen, typeToGen{Name: name, Def: t, GoRef: goRef})
			p.names[name] = true
			// process the dependencies of the dependency
			if err := processDeps(goPkgPath, p, dep); err != nil {
				return err
			}
		}
	}
	return nil
}

func buildGoTypeImpls(typeToGen typesToGen, depToGo cg.DefToGoTypeRef) (cg.GoTypeImpls, error) {
	goTypeImpls := cg.GoTypeImpls{}
	for _, ttg := range typeToGen {
		if goTypeImpl, err := buildGoTypeImpl(depToGo, ttg.Def, ttg.GoRef); err != nil {
			return nil, err
		} else {
			goTypeImpls = append(goTypeImpls, goTypeImpl)
		}
	}
	return goTypeImpls, nil
}

func buildGoTypeImpl(depToGo cg.DefToGoTypeRef, typeDef def.Type, goTypeRef cg.GoTypeRef) (cg.GoTypeImpl, error) {
	switch d := typeDef.(type) {
	case def.SingletonBool, def.SingletonFloat, def.SingletonInt, def.SingletonByte, def.SingletonChar, def.SingletonString:
		return blue.BuildSingletonImpl(d, goTypeRef)
	case def.Structure:
		return blue.BuildStructureImpl(depToGo, d, goTypeRef)
	case def.Inductive:
		return blue.BuildInductiveImpl(depToGo, d, goTypeRef)
	default:
		return nil, fmt.Errorf("unsupported user type definition %#v", typeDef)
	}
}
