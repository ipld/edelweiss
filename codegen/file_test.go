package codegen

import (
	"testing"

	"github.com/ipld/edelweiss/def"
)

func TestFile(t *testing.T) {
	f := &GoFile{
		FilePath: "test.go",
		PkgPath:  "test",
		Types:    GoTypeImpls{testTypeImpl{}},
	}
	buf, err := f.Generate()
	if err != nil {
		t.Fatal(err)
	}
	expect := `package test

import(
	pd1 "fmt"
)

var _ = pd1.Errorf`
	if string(buf) != expect {
		t.Errorf("got %q, expecting %q", string(buf), expect)
	}
}

type testTypeImpl struct{}

func (testTypeImpl) ProtoDef() def.Type {
	return nil
}

func (testTypeImpl) GoTypeRef() GoTypeRef {
	return GoTypeRef{}
}

func (testTypeImpl) GoDef() Blueprint {
	return BlueSlice{
		V("var _ = "),
		GoRef{PkgPath: "fmt", Name: "Errorf"},
	}
}
