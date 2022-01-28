package kernel

import (
	"fmt"

	"github.com/ipld/edelweiss/def"
	"github.com/ipld/edelweiss/system"
)

func CompileTypeMap(defs def.Types) (system.TypeMap, error) {
	m := system.TypeMap{}
	for _, d := range defs {
		switch x := d.(type) {
		case def.Named:
			if _, ok := m[x.Name]; ok {
				return nil, fmt.Errorf("type %s already defined", x.Name)
			} else {
				m[x.Name] = x.Type
			}
		default:
			return nil, fmt.Errorf("anonymous top-level type")
		}
	}
	return m, nil
}
