package def

type Fn struct {
	Arg    Type
	Return Type
}

func (fn Fn) Deps() Types {
	return Types{fn.Arg, fn.Return}
}

func (Fn) Kind() string {
	return "Fn"
}
