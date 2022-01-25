package def

type Union struct {
	Cases CaseListOrNone
}

type CaseListOrNone interface{}

type CaseList struct {
	Case Case
	Rest CaseListOrNone
}

type Case struct {
	Name string
	Type Type
}

func MakeUnion(cases ...Case) Union {
	return Union{
		Cases: makeCases(cases),
	}
}

func makeCases(cases []Case) FieldListOrNone {
	if len(cases) == 0 {
		return nil
	} else {
		return CaseList{
			Case: cases[0],
			Rest: makeCases(cases[1:]),
		}
	}
}
