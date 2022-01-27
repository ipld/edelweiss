package values

import (
	"testing"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
)

func TestListRoundtrip(t *testing.T) {
	x0 := List{Any{Bool(true)}, Any{Bool(false)}}
	buf, err := ipld.Encode(x0, dagjson.Encode)
	if err != nil {
		t.Fatal(err)
	}
	n, err := ipld.Decode(buf, dagjson.Decode)
	if err != nil {
		t.Fatal(err)
	}
	var x1 List
	if err := x1.Parse(n); err != nil {
		t.Fatal(err)
	}
	if !listEqual(x0, x1) {
		t.Errorf("not equal")
	}
}
