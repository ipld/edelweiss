package parse

import "fmt"

type Source interface {
	PeekRune() (rune, error)
	TakeRune() (rune, Source, error)
	PeekN(int) (string, error)
	TakeN(int) (string, Source)
	Len() int
}

type Parser interface {
	Parse(Source) (parsed interface{}, remainder Source, err error)
}

type MatchChar struct {
	Char rune
}

func (x MatchChar) Parse(src Source) (interface{}, Source, error) {
	ch, r, err := src.TakeRune()
	if err != nil {
		return nil, nil, err
	}
	if ch != x.Char {
		return nil, nil, fmt.Errorf("expecting %v, got %v", x.Char, ch)
	}
	return ch, r, nil
}
