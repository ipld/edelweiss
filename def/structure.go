package def

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

func MakeStructure(fields ...Field) Structure {
	return Structure{
		Fields: makeFields(fields),
	}
}

func makeFields(fields []Field) FieldListOrNone {
	if len(fields) == 0 {
		return nil
	} else {
		return FieldList{
			Field: fields[0],
			Rest:  makeFields(fields[1:]),
		}
	}
}
