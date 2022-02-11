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

type Call struct {
	Fn Fn
	ID Type // function instance identifier can be user-defined
}

func (call Call) Deps() Types {
	return Types{call.Fn}
}

func (Call) Kind() string {
	return "Call"
}

type Return struct {
	Fn Fn
	ID Type // function instance identifier can be user-defined
}

func (r Return) Deps() Types {
	return Types{r.Fn}
}

func (Return) Kind() string {
	return "Return"
}
