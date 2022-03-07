package def

type Union struct {
	Cases CaseListOrNone
}

func (Union) Kind() string {
	return "Union"
}

type CaseListOrNone interface{}

type CaseList struct {
	Case Case
	Rest CaseListOrNone
}

type Case struct {
	Name string
	Type Def
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

func FlattenCaseList(x CaseListOrNone) []Case {
	r, cur := []Case{}, x
	for cur != nil {
		l := cur.(CaseList)
		r = append(r, l.Case)
		cur = l.Rest
	}
	return r
}
