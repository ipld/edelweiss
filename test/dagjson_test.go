package test

import (
	"testing"

	"github.com/ipld/edelweiss/defs"
)

// TestDAGJSONStructureEmptyListAtRunTime tests that a missing list field decodes as an empty list.
func TestDAGJSONStructureEmptyListAtRunTime(t *testing.T) {
	defs := []defs.Defs{
		{defs.Named{
			Name: "UserStructure",
			Type: defs.Structure{
				Fields: defs.Fields{
					defs.Field{Name: "F", Type: defs.List{Element: defs.Int{}}},
				},
			},
		}},
	}
	testSrc := `
	var x1 UserStructure
	var x2 UserStructure
	buf := []byte("{}") // field is missing in encoding
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
