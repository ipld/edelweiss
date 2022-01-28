package blueprint

import (
	"github.com/ipld/edelweiss/def"
	"github.com/ipld/edelweiss/system"
)

type Compilation struct {
	PkgPath string
	Types   system.TypeMap // types to be generated
}

type compileContext struct {
	typeToGo    TypeToGo
	anonCounter int
}

type TypeToGo map[def.Type]GoType

// XXX: handle refs

func Compile(types def.Types) (Builder, error) {
	typeToGo := TypeToGo{}
	for _, typeDef := range types {
		if _, ok := typeDef.(def.Named); ok {
			XXX
		} else {
			XXX
		}
	}
	XXX
}

func disassembleTypeDef(ctx *compileContext, typeDef def.Type) {
	switch x := typeDef.(type) {
	case def.Named:
		ctx.typeToGo[typeDef] = nil
	default:
		ctx.typeToGo[def.Named{takeAutoName(ctx, typeDef), typeDef}] = nil
	}
}

func takeAutoName(ctx *compileContext, typeDef def.Type) string {
	XXX
}
