package values

import (
	"testing"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
)

func TestAnyRoundtrip(t *testing.T) {
	var x0 Any
	x0.Value = Bool(true)
	buf, err := ipld.Encode(x0.Node(), dagjson.Encode)
	if err != nil {
		t.Fatal(err)
	}
	n, err := ipld.Decode(buf, dagjson.Decode)
	if err != nil {
		t.Fatal(err)
	}
	var x1 Any
	if err := x1.Parse(n); err != nil {
		t.Fatal(err)
	}
	if x1 != x0 {
		t.Errorf("not equal")
	}
}
