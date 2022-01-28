package def

type Type interface {
	Kind() string
}

type Types []Type
