package def

type Map struct {
	Key   Type
	Value Type
}

func (Map) Kind() string {
	return "Map"
}
