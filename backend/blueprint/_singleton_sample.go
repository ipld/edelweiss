package XXX_PkgName

type Singleton_XXX_TypeName struct{}

func (Singleton_XXX_TypeName) Def() def.Type {
	return nil
}

func (Singleton_XXX_TypeName) Parse(n datamodel.Node) error {
	if n.Kind() != ipld.XXX_IPLDKind {
		return ErrNA
	}
	if n.XXX_IPLDAs() != XXX_IPLDLiteral {
		return ErrNA
	}
	return nil
}

func (v Singleton_XXX_TypeName) Node() datamodel.Node {
	return v
}

func (Singleton_XXX_TypeName) Kind() datamodel.Kind {
	return datamodel.XXX_IPLDKind
}

func (Singleton_XXX_TypeName) LookupByString(string) (datamodel.Node, error) {
	return nil, ErrNA
}

func (Singleton_XXX_TypeName) LookupByNode(key datamodel.Node) (datamodel.Node, error) {
	return nil, ErrNA
}

func (Singleton_XXX_TypeName) LookupByIndex(idx int64) (datamodel.Node, error) {
	return nil, ErrNA
}

func (Singleton_XXX_TypeName) LookupBySegment(seg datamodel.PathSegment) (datamodel.Node, error) {
	return nil, ErrNA
}

func (Singleton_XXX_TypeName) MapIterator() datamodel.MapIterator {
	return nil
}

func (Singleton_XXX_TypeName) ListIterator() datamodel.ListIterator {
	return nil
}

func (Singleton_XXX_TypeName) Length() int64 {
	return -1
}

func (Singleton_XXX_TypeName) IsAbsent() bool {
	return false
}

func (Singleton_XXX_TypeName) IsNull() bool {
	return false
}

func (v Singleton_XXX_TypeName) AsBool() (bool, error) {
	return XXX_AsBoolReturn
}

func (Singleton_XXX_TypeName) AsInt() (int64, error) {
	return XXX_AsIntReturn
}

func (Singleton_XXX_TypeName) AsFloat() (float64, error) {
	return XXX_AsFloatReturn
}

func (Singleton_XXX_TypeName) AsString() (string, error) {
	return XXX_AsStringReturn
}

func (Singleton_XXX_TypeName) AsBytes() ([]byte, error) {
	return nil, ErrNA
}

func (Singleton_XXX_TypeName) AsLink() (datamodel.Link, error) {
	return nil, ErrNA
}

func (Singleton_XXX_TypeName) Prototype() datamodel.NodePrototype {
	return nil // not needed
}
