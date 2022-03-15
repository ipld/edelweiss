package compile

import (
	"testing"

	"github.com/ipld/edelweiss/defs"
)

func TestSingletonAtCompileTime(t *testing.T) {
	defs := defs.Defs{
		defs.Named{Name: "T1", Type: defs.SingletonBool{Bool: true}},
		defs.Named{Name: "T2", Type: defs.SingletonInt{Int: 23}},
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
	_, err = goFile.Generate()
	if err != nil {
		t.Fatal(err)
	}
}

func TestStructureAtCompileTime(t *testing.T) {
	defs := defs.Defs{
		defs.Named{
			Name: "S1",
			Type: defs.Structure{
				Fields: defs.Fields{
					defs.Field{Name: "Int", Type: defs.Int{}},
					defs.Field{Name: "Bool", Type: defs.Bool{}},
				},
			},
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
	_, err = goFile.Generate()
	if err != nil {
		t.Fatal(err)
	}
}

func TestInductiveAtCompileTime(t *testing.T) {
	defs := defs.Defs{
		defs.Named{
			Name: "S1",
			Type: defs.Inductive{
				Cases: defs.Cases{
					defs.Case{Name: "Int", Type: defs.Int{}},
					defs.Case{Name: "Bool", Type: defs.Bool{}},
				},
			},
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
	_, err = goFile.Generate()
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnionAtCompileTime(t *testing.T) {
	defs := defs.Defs{
		defs.Named{
			Name: "S1",
			Type: defs.Union{
				Cases: defs.Cases{
					defs.Case{Name: "Int", Type: defs.Int{}},
					defs.Case{Name: "Bool", Type: defs.Bool{}},
				},
			},
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
	_, err = goFile.Generate()
	if err != nil {
		t.Fatal(err)
	}
}

func TestListAtCompileTime(t *testing.T) {
	defs := defs.Defs{
		defs.Named{
			Name: "S1",
			Type: defs.List{Element: defs.Int{}},
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
	_, err = goFile.Generate()
	if err != nil {
		t.Fatal(err)
	}
}

func TestLinkAtCompileTime(t *testing.T) {
	defs := defs.Defs{
		defs.Named{
			Name: "S1",
			Type: defs.Link{To: defs.Int{}},
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
	_, err = goFile.Generate()
	if err != nil {
		t.Fatal(err)
	}
}

func TestMapAtCompileTime(t *testing.T) {
	defs := defs.Defs{
		defs.Named{
			Name: "S1",
			Type: defs.Map{Key: defs.Int{}, Value: defs.String{}},
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
	_, err = goFile.Generate()
	if err != nil {
		t.Fatal(err)
	}
}

func TestCallAtCompileTime(t *testing.T) {
	defs := defs.Defs{
		defs.Named{
			Name: "S1",
			Type: defs.Call{ID: defs.Int{}, Fn: defs.Fn{Arg: defs.Int{}, Return: defs.String{}}},
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
	_, err = goFile.Generate()
	if err != nil {
		t.Fatal(err)
	}
}

func TestReturnAtCompileTime(t *testing.T) {
	defs := defs.Defs{
		defs.Named{
			Name: "S1",
			Type: defs.Return{ID: defs.Int{}, Fn: defs.Fn{Arg: defs.Int{}, Return: defs.String{}}},
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
	_, err = goFile.Generate()
	if err != nil {
		t.Fatal(err)
	}
}

func TestServiceAtCompileTime(t *testing.T) {
	defs := defs.Defs{
		defs.Named{Name: "TestService",
			Type: defs.Service{
				Methods: defs.Methods{
					defs.Method{Name: "Method1", Type: defs.Fn{Arg: defs.Int{}, Return: defs.Bool{}}},
					defs.Method{Name: "Method2", Type: defs.Fn{Arg: defs.String{}, Return: defs.Float{}}},
				},
			},
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
	_, err = goFile.Generate()
	if err != nil {
		t.Fatal(err)
	}
}
