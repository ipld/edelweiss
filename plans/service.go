package plans

type Service struct {
	Methods []Method
}

type Method struct {
	Name string
	XXX
}
