package defs

type Inductive struct {
	Cases CaseListOrNone
}

func (Inductive) Kind() string {
	return "Inductive"
}

func MakeInductive(cases ...Case) Inductive {
	return Inductive{
		Cases: makeCases(cases),
	}
}
