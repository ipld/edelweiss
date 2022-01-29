package def

type Named struct {
	Name string
	Type Type
}

func (n Named) Deps() Types {
	return Types{n.Type}
}

func (Named) Kind() string {
	return "Named"
}
