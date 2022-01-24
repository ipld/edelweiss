package types

type Union struct {
	Cases CaseListOrNone
}

type CaseListOrNone interface{}

type CaseList struct {
	Case Case
	Rest CaseListOrNone
}

type Case struct {
	Name string
	Type Type
}
