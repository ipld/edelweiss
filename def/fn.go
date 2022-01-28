package def

type Fn struct {
	Arg    Type
	Return Type
}

func (Fn) Kind() string {
	return "Fn"
}
