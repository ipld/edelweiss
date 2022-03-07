package defs

type Structure struct {
	Fields []Field
}

func (Structure) Kind() string {
	return "Structure"
}

type Field struct {
	Name string
	Type Def
}

type Fields []Field
