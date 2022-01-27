package def

// Service is a set of named functions.
type Service struct {
	Methods MethodListOrNone
}

type MethodListOrNone interface{}

type MethodList struct {
	Method Method
	Rest   MethodListOrNone
}

type Method struct {
	Name string
	Type Type
}

func MakeService(fields ...Method) Service {
	return Service{
		Methods: makeMethods(fields),
	}
}

func makeMethods(fields []Method) MethodListOrNone {
	if len(fields) == 0 {
		return nil
	} else {
		return MethodList{
			Method: fields[0],
			Rest:   makeMethods(fields[1:]),
		}
	}
}
