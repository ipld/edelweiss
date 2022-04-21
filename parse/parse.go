package parse

import "fmt"

// XXX: substitute Source with Tape
type Source interface {
	PeekRune() (rune, error)
	TakeRune() (rune, Source, error)
	PeekN(int) (string, error)
	TakeN(int) (string, Source, error)
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

type MatchCharGroup []rune

func (x MatchCharGroup) Parse(src Source) (interface{}, Source, error) {
	ch, r, err := src.TakeRune()
	if err != nil {
		return nil, nil, err
	}
	for _, g := range x {
		if g == ch {
			return ch, r, nil
		}
	}
	return nil, nil, fmt.Errorf("expecting a character in %v, got %v", x, ch)
}

type MatchString struct {
	String string
}

func (x MatchString) Parse(src Source) (interface{}, Source, error) {
	s, r, err := src.TakeN(len(x.String))
	if err != nil {
		return nil, nil, err
	}
	if s != x.String {
		return nil, nil, fmt.Errorf("expecting %v, got %v", x.String, s)
	}
	return s, r, nil
}

type MatchEOF struct{}

func (x MatchEOF) Parse(src Source) (interface{}, Source, error) {
	if src.Len() != 0 {
		return nil, nil, fmt.Errorf("expecting end of file")
	}
	return nil, src, nil
}
