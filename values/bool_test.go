package values

import (
	"testing"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
)

func TestBoolRoundtrip(t *testing.T) {
	b0 := Bool(true)
	buf, err := ipld.Encode(b0, dagjson.Encode)
	if err != nil {
		t.Fatal(err)
	}
	// fmt.Println(string(buf))
	n, err := ipld.Decode(buf, dagjson.Decode)
	if err != nil {
		t.Fatal(err)
	}
	var b1 Bool
	if err := b1.Parse(n); err != nil {
		t.Fatal(err)
	}
	if b1 != b0 {
		t.Errorf("bools not equal")
	}
}
