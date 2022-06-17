package test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"testing"

	"github.com/ipld/edelweiss/compile"
	"github.com/ipld/edelweiss/defs"
)

func RunSingleGenTest(t *testing.T, defs defs.Defs, testSrc string) {
	t.Helper()

	testFuncFmt := `func TestMain(t *testing.T) {
%s
}`
	RunGenTest(t, defs, fmt.Sprintf(testFuncFmt, testSrc))
}

func RunGenTest(t *testing.T, defs defs.Defs, testSrc string) {
	t.Helper()

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
	github.com/ipld/edelweiss 13981a091ad908bfcbce6e0db02764cdeecaaa9a
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
	"context"
	"fmt"
	"os"
	"testing"
	"net/http/httptest"
	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/edelweiss/values"
	"github.com/ipld/edelweiss/services"
	cid "github.com/ipfs/go-cid"
)

// silence pkg import errors
var (
	_ = context.Background
	_ = fmt.Printf
	_ = os.Exit
	_ = ipld.Encode
	_ = values.Any{}
	_ = dagjson.Encode
	_ = dagcbor.Encode
	_ = basicnode.Prototype
	_ = cid.NewCidV1
	_ = httptest.NewServer
	_ = services.ErrSchema
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

	var outb, errb bytes.Buffer
	goTest.Stdout = &outb
	goTest.Stderr = &errb

	if err = goTest.Run(); err != nil {
		t.Fatal("go test (", err, ")\nout:", outb.String(), "err:", errb.String())
	}
}
