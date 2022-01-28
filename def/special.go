package def

type String struct{}

func (String) Kind() string {
	return "String"
}

type Any struct{}

func (Any) Kind() string {
	return "Any"
}

type Nothing struct{}

func (Nothing) Kind() string {
	return "Nothing"
}
