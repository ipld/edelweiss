package compile

import (
	"fmt"
	"testing"

	"github.com/ipld/edelweiss/def"
)

func TestSingletonAtCompileTime(t *testing.T) {
	defs := def.Types{
		def.Named{Name: "T1", Type: def.SingletonBool{Bool: true}},
		def.Named{Name: "T2", Type: def.SingletonInt{Int: 23}},
	}
	x := &GoPkgCodegen{
		GoPkgDirPath: "",
		GoPkgPath:    "test",
		Defs:         defs,
	}
	goFile, err := x.Compile()
	if err != nil {
		t.Fatal(err)
	}
	fileBuf, err := goFile.Generate()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(fileBuf))
}

func TestStructureAtCompileTime(t *testing.T) {
	defs := def.Types{
		def.Named{
			Name: "S1",
			Type: def.MakeStructure(
				def.Field{Name: "Int", Type: def.Int{}},
				def.Field{Name: "Bool", Type: def.Bool{}},
			),
		},
	}
	x := &GoPkgCodegen{
		GoPkgDirPath: "",
		GoPkgPath:    "test",
		Defs:         defs,
	}
	goFile, err := x.Compile()
	if err != nil {
		t.Fatal(err)
	}
	fileBuf, err := goFile.Generate()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(fileBuf))
}

func TestInductiveAtCompileTime(t *testing.T) {
	defs := def.Types{
		def.Named{
			Name: "S1",
			Type: def.MakeInductive(
				def.Case{Name: "Int", Type: def.Int{}},
				def.Case{Name: "Bool", Type: def.Bool{}},
			),
		},
	}
	x := &GoPkgCodegen{
		GoPkgDirPath: "",
		GoPkgPath:    "test",
		Defs:         defs,
	}
	goFile, err := x.Compile()
	if err != nil {
		t.Fatal(err)
	}
	fileBuf, err := goFile.Generate()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(fileBuf))
}

func TestListAtCompileTime(t *testing.T) {
	defs := def.Types{
		def.Named{
			Name: "S1",
			Type: def.List{Element: def.Int{}},
		},
	}
	x := &GoPkgCodegen{
		GoPkgDirPath: "",
		GoPkgPath:    "test",
		Defs:         defs,
	}
	goFile, err := x.Compile()
	if err != nil {
		t.Fatal(err)
	}
	fileBuf, err := goFile.Generate()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(fileBuf))
}

func TestLinkAtCompileTime(t *testing.T) {
	defs := def.Types{
		def.Named{
			Name: "S1",
			Type: def.Link{To: def.Int{}},
		},
	}
	x := &GoPkgCodegen{
		GoPkgDirPath: "",
		GoPkgPath:    "test",
		Defs:         defs,
	}
	goFile, err := x.Compile()
	if err != nil {
		t.Fatal(err)
	}
	fileBuf, err := goFile.Generate()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(fileBuf))
}

func TestMapAtCompileTime(t *testing.T) {
	defs := def.Types{
		def.Named{
			Name: "S1",
			Type: def.Map{Key: def.Int{}, Value: def.String{}},
		},
	}
	x := &GoPkgCodegen{
		GoPkgDirPath: "",
		GoPkgPath:    "test",
		Defs:         defs,
	}
	goFile, err := x.Compile()
	if err != nil {
		t.Fatal(err)
	}
	fileBuf, err := goFile.Generate()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(fileBuf))
}

func TestCallAtCompileTime(t *testing.T) {
	defs := def.Types{
		def.Named{
			Name: "S1",
			Type: def.Call{ID: def.Int{}, Fn: def.Fn{Arg: def.Int{}, Return: def.String{}}},
		},
	}
	x := &GoPkgCodegen{
		GoPkgDirPath: "",
		GoPkgPath:    "test",
		Defs:         defs,
	}
	goFile, err := x.Compile()
	if err != nil {
		t.Fatal(err)
	}
	fileBuf, err := goFile.Generate()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(fileBuf))
}
