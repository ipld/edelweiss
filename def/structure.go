package def

type Structure struct {
	Fields FieldListOrNone
}

func (Structure) Kind() string {
	return "Structure"
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

func FlattenFieldList(x FieldListOrNone) []Field {
	r, cur := []Field{}, x
	for cur != nil {
		l := cur.(FieldList)
		r = append(r, l.Field)
		cur = l.Rest
	}
	return r
}
