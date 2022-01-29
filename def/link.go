package def

type Link struct {
	To Type
}

func (link Link) Deps() Types {
	return Types{link.To}
}

func (Link) Kind() string {
	return "Link"
}
