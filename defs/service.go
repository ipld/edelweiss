package defs

// Service is a set of named functions.
type Service struct {
	Methods MethodListOrNone
}

func (Service) Kind() string {
	return "Service"
}

type MethodListOrNone interface{}

type MethodList struct {
	Method Method
	Rest   MethodListOrNone
}

type Method struct {
	Name string
	Type Fn
}

func (m Method) Call() Call {
	return Call{ID: String{}, Fn: m.Type}
}

func (m Method) Return() Return {
	return Return{ID: String{}, Fn: m.Type}
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

func FlattenMethodList(x MethodListOrNone) []Method {
	r, cur := []Method{}, x
	for cur != nil {
		l := cur.(MethodList)
		r = append(r, l.Method)
		cur = l.Rest
	}
	return r
}
