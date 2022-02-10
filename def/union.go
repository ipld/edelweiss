package def

type Union struct {
	Cases CaseListOrNone
}

func (u Union) Deps() Types {
	if u.Cases == nil {
		return nil
	} else {
		return u.Cases.Deps()
	}
}

func (Union) Kind() string {
	return "Union"
}

type CaseListOrNone interface {
	Deps() Types
}

type CaseList struct {
	Case Case
	Rest CaseListOrNone
}

func (cl CaseList) Deps() Types {
	if cl.Rest == nil {
		return Types{cl.Case.Type}
	} else {
		return append(Types{cl.Case.Type}, cl.Rest.Deps()...)
	}
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

func FlattenCaseList(x CaseListOrNone) []Case {
	r, cur := []Case{}, x
	for cur != nil {
		l := cur.(CaseList)
		r = append(r, l.Case)
		cur = l.Rest
	}
	return r
}
