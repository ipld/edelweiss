package parse

type Tape interface {
	Peek() (interface{}, error)
	Take() (interface{}, Tape, error)
	PeekN(int) (string, error)
	TakeN(int) (string, Tape, error)
	Len() int
}

type StringTape interface {
	PeekRune() (rune, error)
	TakeRune() (rune, Source, error)
	PeekStringN(int) (string, error)
	TakeStringN(int) (string, Source, error)
	Len() int
}
