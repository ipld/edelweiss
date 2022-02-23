package def

// Service is a set of named functions.
type Service struct {
	Methods MethodListOrNone
}

func (s Service) Deps() Types {
	if s.Methods != nil {
		return s.Methods.Deps()
	} else {
		return nil
	}
}

func (Service) Kind() string {
	return "Service"
}

type MethodListOrNone interface {
	Deps() Types
}

type MethodList struct {
	Method Method
	Rest   MethodListOrNone
}

func (ml MethodList) Deps() Types {
	if ml.Rest == nil {
		return Types{ml.Method.Type}
	} else {
		return append(Types{ml.Method.Type}, ml.Rest.Deps()...)
	}
}

type Method struct {
	Name string
	Type Fn
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
