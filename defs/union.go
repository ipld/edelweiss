package defs

type Union struct {
	Cases Cases
}

func (Union) Kind() string {
	return "Union"
}

type Case struct {
	Name string
	Type Def
}

type Cases []Case
