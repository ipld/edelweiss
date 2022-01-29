package def

// Bool, Float, Int, Byte, Char

type Bool struct{}

func (Bool) Deps() Types {
	return nil
}

func (Bool) Kind() string {
	return "Bool"
}

type Float struct{}

func (Float) Deps() Types {
	return nil
}

func (Float) Kind() string {
	return "Float"
}

type Int struct{}

func (Int) Deps() Types {
	return nil
}

func (Int) Kind() string {
	return "Int"
}

type Byte struct{}

func (Byte) Deps() Types {
	return nil
}

func (Byte) Kind() string {
	return "Byte"
}

type Char struct{}

func (Char) Deps() Types {
	return nil
}

func (Char) Kind() string {
	return "Char"
}
