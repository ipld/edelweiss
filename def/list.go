package def

type List struct {
	Element Type
}

func (list List) Deps() Types {
	return Types{list.Element}
}

func (List) Kind() string {
	return "List"
}
