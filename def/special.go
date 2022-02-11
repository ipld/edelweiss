package def

type String struct{}

func (String) Deps() Types {
	return nil
}

func (String) Kind() string {
	return "String"
}

type Bytes struct{}

func (Bytes) Deps() Types {
	return nil
}

func (Bytes) Kind() string {
	return "Bytes"
}

type Any struct{}

func (Any) Deps() Types {
	return nil
}

func (Any) Kind() string {
	return "Any"
}

type Nothing struct{}

func (Nothing) Deps() Types {
	return nil
}

func (Nothing) Kind() string {
	return "Nothing"
}
