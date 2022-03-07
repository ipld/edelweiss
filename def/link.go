package def

type Link struct {
	To Type
}

func (Link) Kind() string {
	return "Link"
}
