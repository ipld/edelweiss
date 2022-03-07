package compile

import (
	"fmt"
	"path"

	log "github.com/ipfs/go-log"
	blue_services "github.com/ipld/edelweiss/blueprints/services/dagjson-over-http"
	blue_values "github.com/ipld/edelweiss/blueprints/values"
	cg "github.com/ipld/edelweiss/codegen"
	"github.com/ipld/edelweiss/defs"
	"github.com/ipld/edelweiss/plans"
)

var logger = log.Logger("edelweiss")

type GoPkgCodegen struct {
	GoPkgDirPath string // local directory for the package
	GoPkgPath    string // global package name
	Defs         defs.Defs
}

func (x *GoPkgCodegen) GoPkgName() string {
	return path.Base(x.GoPkgPath)
}

func (x *GoPkgCodegen) Compile() (*cg.GoFile, error) {
	p := newGenPlan(x.GoPkgPath)
	for _, d := range x.Defs {
		if _, err := generate(p, d); err != nil {
			return nil, err
		}
	}
	if err := p.VerifyCompleteness(); err != nil {
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

func buildGoTypeImpls(typeToGen []typePlan, depToGo cg.DefToGoTypeRef) (cg.GoTypeImpls, error) {
	goTypeImpls := cg.GoTypeImpls{}
	for _, ttg := range typeToGen {
		goTypeImpls = append(goTypeImpls, buildGoTypeImpl(depToGo, ttg.Def, ttg.GoRef)...)
	}
	return goTypeImpls, nil
}

func buildGoTypeImpl(depToGo cg.DefToGoTypeRef, typeDef defs.Def, goTypeRef cg.GoTypeRef) []cg.GoTypeImpl {
	switch d := typeDef.(type) {
	case defs.SingletonBool, defs.SingletonFloat, defs.SingletonInt, defs.SingletonByte, defs.SingletonChar, defs.SingletonString:
		return []cg.GoTypeImpl{blue_values.BuildSingletonImpl(d, goTypeRef)}
	case defs.Structure:
		return []cg.GoTypeImpl{blue_values.BuildStructureImpl(depToGo, d, goTypeRef)}
	case defs.Inductive:
		return []cg.GoTypeImpl{blue_values.BuildInductiveImpl(depToGo, d, goTypeRef)}
	case defs.List:
		return []cg.GoTypeImpl{blue_values.BuildListImpl(depToGo, d, goTypeRef)}
	case defs.Link:
		return []cg.GoTypeImpl{blue_values.BuildLinkImpl(depToGo, d, goTypeRef)}
	case defs.Map:
		return []cg.GoTypeImpl{blue_values.BuildMapImpl(depToGo, d, goTypeRef)}
	case defs.Call:
		return []cg.GoTypeImpl{blue_values.BuildCallImpl(depToGo, d, goTypeRef)}
	case defs.Return:
		return []cg.GoTypeImpl{blue_values.BuildReturnImpl(depToGo, d, goTypeRef)}
	case plans.Service:
		return []cg.GoTypeImpl{
			blue_services.BuildClientImpl(depToGo, d, goTypeRef),
			blue_services.BuildServerImpl(depToGo, d, goTypeRef),
		}
	default:
		panic(fmt.Sprintf("unknown implementation plan  %#v", d))
	}
}
