package def

type Type interface {
	Kind() string
	Deps() Types
}

type Types []Type
