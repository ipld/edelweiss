package def

// Ref is a reference to a type by name.
type Ref struct {
	Name string
}

func (Ref) Deps() Types {
	return nil
}

func (Ref) Kind() string {
	return "Ref"
}
