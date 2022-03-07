package def

type Def interface {
	Kind() string
}

type Defs []Def
