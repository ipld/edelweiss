package def

type List struct {
	Element Type
}

func (List) Kind() string {
	return "List"
}
