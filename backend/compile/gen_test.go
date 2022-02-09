package compile

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"testing"

	"github.com/ipld/edelweiss/def"
)

func TestGenTest(t *testing.T) {
	RunGenTest(t, def.Types{def.Named{Name: "T", Type: def.SingletonInt{Int: 23}}}, "")
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
	github.com/ipld/edelweiss v0.0.0-20220209162310-6635ef8b8859
	github.com/ipld/go-ipld-prime v0.14.4
)
`
	if err := ioutil.WriteFile(path.Join(dir, "go.mod"), []byte(goModSrc), 0644); err != nil {
		t.Fatalf("creating go.mod (%v)", err)
	}

	// generate type code
	x := &GoPkgCodegen{
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
)

var (
	_ = fmt.Printf
	_ = os.Exit
)

func TestMain(t *testing.T) {
%s
}
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
