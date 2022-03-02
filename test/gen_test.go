package test

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"testing"

	"github.com/ipld/edelweiss/compile"
	"github.com/ipld/edelweiss/def"
)

func TestGenTest(t *testing.T) {
	RunSingleGenTest(t, def.Types{def.Named{Name: "T", Type: def.SingletonInt{Int: 23}}}, "")
}

func TestSingletonAtRunTime(t *testing.T) {
	defs := []def.Types{
		{def.Named{Name: "UserSingleton", Type: def.SingletonBool{Bool: true}}},
		{def.Named{Name: "UserSingleton", Type: def.SingletonInt{Int: 23}}},
		{def.Named{Name: "UserSingleton", Type: def.SingletonFloat{Float: 2.3}}},
		{def.Named{Name: "UserSingleton", Type: def.SingletonByte{Byte: 2}}},
		{def.Named{Name: "UserSingleton", Type: def.SingletonChar{Char: 'a'}}},
		{def.Named{Name: "UserSingleton", Type: def.SingletonString{String: "abc"}}},
	}
	testSrc := `
	var x1 UserSingleton
	buf, err := ipld.Encode(x1, dagjson.Encode)
	if err != nil {
		t.Fatalf("encoding (%v)", err)
	}
	var x2 UserSingleton
	n, err := ipld.Decode(buf, dagjson.Decode)
	if err != nil {
		t.Fatalf("decoding (%v)", err)
	}
	if err = x2.Parse(n); err != nil {
		t.Fatalf("parsing (%v)", err)
	}
	if !ipld.DeepEqual(x1, x2) {
		t.Errorf("ipld values are not equal")
	}
`
	for _, d := range defs {
		RunSingleGenTest(t, d, testSrc)
	}
}

func TestStructureAtRunTime(t *testing.T) {
	defs := []def.Types{
		{def.Named{
			Name: "UserStructure",
			Type: def.MakeStructure(
				def.Field{Name: "A", Type: def.Int{}},
				def.Field{Name: "B", Type: def.String{}},
				def.Field{Name: "C", Type: def.Float{}},
				def.Field{Name: "D", Type: def.Byte{}},
				def.Field{Name: "E", Type: def.Char{}},
			),
		}},
	}
	testSrc := `
	var x1 UserStructure
	x1.A = 3
	x1.B = "abc"
	x1.C = 1.2
	x1.D = 123
	x1.E = 'w'
	buf, err := ipld.Encode(x1, dagjson.Encode)
	if err != nil {
		t.Fatalf("encoding (%v)", err)
	}
	var x2 UserStructure
	n, err := ipld.Decode(buf, dagjson.Decode)
	if err != nil {
		t.Fatalf("decoding (%v)", err)
	}
	if err = x2.Parse(n); err != nil {
		t.Fatalf("parsing (%v)", err)
	}
	if !ipld.DeepEqual(x1, x2) {
		t.Errorf("ipld values are not equal")
	}
`
	for _, d := range defs {
		RunSingleGenTest(t, d, testSrc)
	}
}

func TestInductiveAtRunTime(t *testing.T) {
	defs := []def.Types{
		{def.Named{
			Name: "UserInductive",
			Type: def.MakeInductive(
				def.Case{Name: "A", Type: def.Int{}},
				def.Case{Name: "B", Type: def.String{}},
				def.Case{Name: "C", Type: def.Float{}},
				def.Case{Name: "D", Type: def.Byte{}},
				def.Case{Name: "E", Type: def.Char{}},
			),
		}},
	}
	testSrc := `
	var x1 UserInductive
	var y values.String = "abc"
	x1.B = &y
	buf, err := ipld.Encode(x1, dagjson.Encode)
	if err != nil {
		t.Fatalf("encoding (%v)", err)
	}
	var x2 UserInductive
	n, err := ipld.Decode(buf, dagjson.Decode)
	if err != nil {
		t.Fatalf("decoding (%v)", err)
	}
	if err = x2.Parse(n); err != nil {
		t.Fatalf("parsing (%v)", err)
	}
	if !ipld.DeepEqual(x1, x2) {
		t.Errorf("ipld values are not equal")
	}
`
	for _, d := range defs {
		RunSingleGenTest(t, d, testSrc)
	}
}

func TestMapAtRunTime(t *testing.T) {
	defs := []def.Types{
		{def.Named{
			Name: "UserMap",
			Type: def.Map{Key: def.String{}, Value: def.Int{}},
		}},
	}
	testSrc := `
	var x1 UserMap = UserMap{
		{Key: "a", Value: 1},
		{Key: "b", Value: 2},
	}
	buf, err := ipld.Encode(x1, dagjson.Encode)
	if err != nil {
		t.Fatalf("encoding (%v)", err)
	}
	var x2 UserMap
	n, err := ipld.Decode(buf, dagjson.Decode)
	if err != nil {
		t.Fatalf("decoding (%v)", err)
	}
	if err = x2.Parse(n); err != nil {
		t.Fatalf("parsing (%v)", err)
	}
	if !ipld.DeepEqual(x1, x2) {
		t.Errorf("ipld values are not equal")
	}
`
	for _, d := range defs {
		RunSingleGenTest(t, d, testSrc)
	}
}

func TestListAtRunTime(t *testing.T) {
	defs := []def.Types{
		{def.Named{
			Name: "UserList",
			Type: def.List{Element: def.String{}},
		}},
	}
	testSrc := `
	var x1 UserList = UserList{
		"abc",
		"def",
	}
	buf, err := ipld.Encode(x1, dagjson.Encode)
	if err != nil {
		t.Fatalf("encoding (%v)", err)
	}
	var x2 UserList
	n, err := ipld.Decode(buf, dagjson.Decode)
	if err != nil {
		t.Fatalf("decoding (%v)", err)
	}
	if err = x2.Parse(n); err != nil {
		t.Fatalf("parsing (%v)", err)
	}
	if !ipld.DeepEqual(x1, x2) {
		t.Errorf("ipld values are not equal")
	}
`
	for _, d := range defs {
		RunSingleGenTest(t, d, testSrc)
	}
}

func TestCallAtRunTime(t *testing.T) {
	defs := []def.Types{
		{def.Named{
			Name: "UserCall",
			Type: def.Call{
				Fn: def.Fn{Arg: def.Int{}, Return: def.String{}},
				ID: def.String{},
			},
		}},
	}
	testSrc := `
	var x1 UserCall = UserCall{
		ID: "abc",
		Arg: 3,
	}
	buf, err := ipld.Encode(x1, dagjson.Encode)
	if err != nil {
		t.Fatalf("encoding (%v)", err)
	}
	var x2 UserCall
	n, err := ipld.Decode(buf, dagjson.Decode)
	if err != nil {
		t.Fatalf("decoding (%v)", err)
	}
	if err = x2.Parse(n); err != nil {
		t.Fatalf("parsing (%v)", err)
	}
	if !ipld.DeepEqual(x1, x2) {
		t.Errorf("ipld values are not equal")
	}
`
	for _, d := range defs {
		RunSingleGenTest(t, d, testSrc)
	}
}

func TestReturnAtRunTime(t *testing.T) {
	defs := []def.Types{
		{def.Named{
			Name: "UserReturn",
			Type: def.Return{
				Fn: def.Fn{Arg: def.Int{}, Return: def.String{}},
				ID: def.String{},
			},
		}},
	}
	testSrc := `
	var x1 UserReturn = UserReturn{
		ID: "abc",
		Return: "def",
	}
	buf, err := ipld.Encode(x1, dagjson.Encode)
	if err != nil {
		t.Fatalf("encoding (%v)", err)
	}
	var x2 UserReturn
	n, err := ipld.Decode(buf, dagjson.Decode)
	if err != nil {
		t.Fatalf("decoding (%v)", err)
	}
	if err = x2.Parse(n); err != nil {
		t.Fatalf("parsing (%v)", err)
	}
	if !ipld.DeepEqual(x1, x2) {
		t.Errorf("ipld values are not equal")
	}
`
	for _, d := range defs {
		RunSingleGenTest(t, d, testSrc)
	}
}

func TestLinkAtRunTime(t *testing.T) {
	defs := []def.Types{
		{def.Named{
			Name: "UserLink",
			Type: def.Link{To: def.String{}},
		}},
	}
	testSrc := `
	c, err := cid.Decode("QmYyQSo1c1Ym7orWxLYvCrM2EmxFTANf8wXmmE7DWjhx5N")
	if err != nil {
		t.Fatal(err)
	}
	var x1 UserLink = UserLink(c)
	buf, err := ipld.Encode(x1, dagjson.Encode)
	if err != nil {
		t.Fatalf("encoding (%v)", err)
	}
	var x2 UserLink
	n, err := ipld.Decode(buf, dagjson.Decode)
	if err != nil {
		t.Fatalf("decoding (%v)", err)
	}
	if err = x2.Parse(n); err != nil {
		t.Fatalf("parsing (%v)", err)
	}
	if !ipld.DeepEqual(x1, x2) {
		t.Errorf("ipld values are not equal")
	}
`
	for _, d := range defs {
		RunSingleGenTest(t, d, testSrc)
	}
}

func TestStructureInductiveAtRunTime(t *testing.T) {
	defs := []def.Types{
		{def.Named{
			Name: "UserStructure",
			Type: def.MakeStructure(
				def.Field{Name: "A", Type: def.Int{}},
				def.Field{Name: "B", Type: def.Named{
					Name: "UserInductive",
					Type: def.MakeInductive(
						def.Case{Name: "X", Type: def.String{}},
						def.Case{Name: "Y", Type: def.Int{}},
					)},
				},
			),
		}},
	}
	testSrc := `
	var x1 UserStructure
	x1.A = 3
	var x1b UserInductive
	y := values.Int(5)
	x1b.Y = &y
	x1.B = x1b
	buf, err := ipld.Encode(x1, dagjson.Encode)
	if err != nil {
		t.Fatalf("encoding (%v)", err)
	}
	var x2 UserStructure
	n, err := ipld.Decode(buf, dagjson.Decode)
	if err != nil {
		t.Fatalf("decoding (%v)", err)
	}
	if err = x2.Parse(n); err != nil {
		t.Fatalf("parsing (%v)", err)
	}
	if !ipld.DeepEqual(x1, x2) {
		t.Errorf("ipld values are not equal")
	}
`
	for _, d := range defs {
		RunSingleGenTest(t, d, testSrc)
	}
}

func TestListStructureAtRunTime(t *testing.T) {
	defs := []def.Types{
		{def.Named{
			Name: "UserList",
			Type: def.List{
				Element: def.MakeStructure(
					def.Field{Name: "X", Type: def.String{}},
				),
			},
		}},
	}
	testSrc := `
	var x1 UserList = UserList{{X: "abc"}}
	buf, err := ipld.Encode(x1, dagjson.Encode)
	if err != nil {
		t.Fatalf("encoding (%v)", err)
	}
	var x2 UserList
	n, err := ipld.Decode(buf, dagjson.Decode)
	if err != nil {
		t.Fatalf("decoding (%v)", err)
	}
	if err = x2.Parse(n); err != nil {
		t.Fatalf("parsing (%v)", err)
	}
	if !ipld.DeepEqual(x1, x2) {
		t.Errorf("ipld values are not equal")
	}
`
	for _, d := range defs {
		RunSingleGenTest(t, d, testSrc)
	}
}

func TestListSingletonAtRunTime(t *testing.T) {
	defs := []def.Types{
		{def.Named{
			Name: "UserList",
			Type: def.List{
				Element: def.SingletonBool{Bool: true},
			},
		}},
	}
	testSrc := `
	var x1 UserList = UserList{{}}
	buf, err := ipld.Encode(x1, dagjson.Encode)
	if err != nil {
		t.Fatalf("encoding (%v)", err)
	}
	var x2 UserList
	n, err := ipld.Decode(buf, dagjson.Decode)
	if err != nil {
		t.Fatalf("decoding (%v)", err)
	}
	if err = x2.Parse(n); err != nil {
		t.Fatalf("parsing (%v)", err)
	}
	if !ipld.DeepEqual(x1, x2) {
		t.Errorf("ipld values are not equal")
	}
`
	for _, d := range defs {
		RunSingleGenTest(t, d, testSrc)
	}
}

func RunSingleGenTest(t *testing.T, defs def.Types, testSrc string) {
	testFuncFmt := `func TestMain(t *testing.T) {
%s
}`
	RunGenTest(t, defs, fmt.Sprintf(testFuncFmt, testSrc))
}

func RunGenTest(t *testing.T, defs def.Types, testSrc string) {

	// create tmp dir
	dir, err := os.MkdirTemp("", "edelweiss_test")
	if err != nil {
		t.Fatal(err)
	}
	// defer os.RemoveAll(dir)
	fmt.Printf("generating test in %s\n", dir)

	// create go.mod
	goModSrc := `
module test

go 1.16

require (
	github.com/ipld/edelweiss c5150ca51dd47435b63b10b0a72743e53d476e6a
	github.com/ipld/go-ipld-prime v0.14.4
	github.com/ipfs/go-cid v0.0.4
)
`
	if err := ioutil.WriteFile(path.Join(dir, "go.mod"), []byte(goModSrc), 0644); err != nil {
		t.Fatalf("creating go.mod (%v)", err)
	}

	// generate type code
	x := &compile.GoPkgCodegen{
		GoPkgDirPath: dir,
		GoPkgPath:    "test",
		Defs:         defs,
	}
	goFile, err := x.Compile()
	if err != nil {
		t.Fatal(err)
	}
	if err = goFile.Build(); err != nil {
		t.Fatal(err)
	}

	// generate test code
	testGoSrc := `
package test

import (
	"fmt"
	"os"
	"testing"
	"net/http/httptest"
	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/edelweiss/values"
	cid "github.com/ipfs/go-cid"
)

// silence pkg import errors
var (
	_ = fmt.Printf
	_ = os.Exit
	_ = ipld.Encode
	_ = values.Any{}
	_ = dagjson.Encode
	_ = dagcbor.Encode
	_ = basicnode.Prototype
	_ = cid.NewCidV1
	_ = httptest.NewServer
)

%s
`
	if err := ioutil.WriteFile(path.Join(dir, "edelweiss_test.go"), []byte(fmt.Sprintf(testGoSrc, testSrc)), 0644); err != nil {
		t.Fatalf("creating test.go (%v)", err)
	}

	// run go mod tidy
	goModTidy := exec.Command("go", "mod", "tidy")
	goModTidy.Dir = dir
	if err = goModTidy.Run(); err != nil {
		t.Fatalf("go mod tidy (%v)", err)
	}

	// run test
	goTest := exec.Command("go", "test")
	goTest.Dir = dir
	if err = goTest.Run(); err != nil {
		t.Fatalf("go test (%v)", err)
	}
}
