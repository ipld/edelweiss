package compile

import (
	"fmt"
	"testing"

	"github.com/ipld/edelweiss/def"
)

func TestSingleton(t *testing.T) {
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
