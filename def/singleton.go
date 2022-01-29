package def

type Singleton interface{}

type SingletonBool struct {
	Bool bool
}

func (SingletonBool) Deps() Types {
	return nil
}

func (SingletonBool) Kind() string {
	return "SingletonBool"
}

type SingletonFloat struct {
	Float float64
}

func (SingletonFloat) Deps() Types {
	return nil
}

func (SingletonFloat) Kind() string {
	return "SingletonFloat"
}

type SingletonInt struct {
	Int int64
}

func (SingletonInt) Deps() Types {
	return nil
}

func (SingletonInt) Kind() string {
	return "SingletonInt"
}

type SingletonByte struct {
	Byte byte
}

func (SingletonByte) Deps() Types {
	return nil
}

func (SingletonByte) Kind() string {
	return "SingletonByte"
}

type SingletonChar struct {
	Char rune
}

func (SingletonChar) Deps() Types {
	return nil
}

func (SingletonChar) Kind() string {
	return "SingletonChar"
}

type SingletonString struct {
	String string
}

func (SingletonString) Deps() Types {
	return nil
}

func (SingletonString) Kind() string {
	return "SingletonString"
}
