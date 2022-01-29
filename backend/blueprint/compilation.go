package blueprint

import (
	"fmt"

	"github.com/ipld/edelweiss/def"
)

type compileContext struct {
	nameToDef   NameToDef
	defToGo     TypeToGo
	anonCounter int
}

type NameToDef map[string]def.Type
type TypeToGo map[def.Type]GoType

// XXX: handle refs

func Compile(defs def.Types) (Builder, error) {
	ctx := &compileContext{}
	if err := compileNameToDef(ctx, defs); err != nil {
		return nil, err
	}
	ctx.defToGo = TypeToGo{}
	for _, typeDef := range defs {
		if _, ok := typeDef.(def.Named); ok {
			XXX
		} else {
			XXX
		}
	}
	XXX
}

func compileNameToDef(ctx *compileContext, defs def.Types) error {
	ctx.nameToDef = NameToDef{}
	for _, d := range defs {
		switch x := d.(type) {
		case def.Named:
			if _, ok := ctx.nameToDef[x.Name]; ok {
				return fmt.Errorf("type %s already defined", x.Name)
			} else {
				ctx.nameToDef[x.Name] = x
			}
		default:
			return fmt.Errorf("anonymous top-level type")
		}
	}
	return nil
}

func disassembleTypeDef(ctx *compileContext, typeDef def.Type) {
	switch x := typeDef.(type) {
	case def.Named:
		ctx.defToGo[typeDef] = nil
	default:
		ctx.defToGo[def.Named{takeAutoName(ctx, typeDef), typeDef}] = nil
	}
}

func takeAutoName(ctx *compileContext, typeDef def.Type) string {
	XXX
}
