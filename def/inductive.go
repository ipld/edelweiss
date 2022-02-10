package def

type Inductive struct {
	Cases CaseListOrNone
}

func (u Inductive) Deps() Types {
	if u.Cases == nil {
		return nil
	} else {
		return u.Cases.Deps()
	}
}

func (Inductive) Kind() string {
	return "Inductive"
}

func MakeInductive(cases ...Case) Inductive {
	return Inductive{
		Cases: makeCases(cases),
	}
}
