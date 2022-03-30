package plans

type Service struct {
	Methods        Methods
	ErrorEnvelope  BuiltinOrRefPlan
	CallEnvelope   BuiltinOrRefPlan
	ReturnEnvelope BuiltinOrRefPlan
	Identify       Method
}

const IdentifyName = "Identify" // auto-generated

func (Service) IAmPlan()     {}
func (Service) Kind() string { return "Service" }

type Methods []Method

type Method struct {
	Name string
	Type Fn
}
