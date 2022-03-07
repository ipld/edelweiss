package defs

type Inductive struct {
	Cases Cases
}

func (Inductive) Kind() string {
	return "Inductive"
}
