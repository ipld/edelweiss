package types

type Structure struct {
	Fields FieldListOrNone
}

type FieldListOrNone interface{}

type FieldList struct {
	Field Field
	Rest  FieldListOrNone
}

type Field struct {
	Name string
	Type Type
}
