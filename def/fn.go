package def

// Fn represents the type signature of a function.
type Fn struct {
	Arg    Type
	Return Type
}

func (fn Fn) Deps() Types {
	return Types{
		fn.Arg,
		fn.Return,
	}
}

func (Fn) Kind() string {
	return "Fn"
}

// Call is the type representing a function call (aka request)
type Call struct {
	Fn Fn   // type signature of the function being called
	ID Type // function instance identifier (can be a user-defined type)
}

func (call Call) Deps() Types {
	return Types{call.Fn}
}

func (Call) Kind() string {
	return "Call"
}

// Return is the type representing a function result (aka response)
type Return struct {
	Fn Fn   // type signature of the function returning a result
	ID Type // function instance identifier (can be a user-defined type)
}

func (r Return) Deps() Types {
	return Types{r.Fn}
}

func (Return) Kind() string {
	return "Return"
}
