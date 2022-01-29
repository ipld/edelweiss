package def

type Map struct {
	Key   Type
	Value Type
}

func (m Map) Deps() Types {
	return Types{m.Key, m.Value}
}

func (Map) Kind() string {
	return "Map"
}
