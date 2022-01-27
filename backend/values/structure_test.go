package values

import (
	"testing"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
)

func TestStructureRoundtrip(t *testing.T) {
	x0 := Structure{
		Field{Name: "f1", Value: Any{Bool(true)}},
		Field{Name: "f2", Value: Any{Bool(false)}},
		Field{Name: "f3", Value: Any{String("haha")}},
		Field{Name: "f4", Value: Any{Float(2.3)}},
		Field{Name: "f5", Value: Any{Int(7)}},
	}
	buf, err := ipld.Encode(x0.Node(), dagjson.Encode)
	if err != nil {
		t.Fatal(err)
	}
	// fmt.Println(string(buf))
	n, err := ipld.Decode(buf, dagjson.Decode)
	if err != nil {
		t.Fatal(err)
	}
	var x1 Structure
	if err := x1.Parse(n); err != nil {
		t.Fatal(err)
	}
	if !structureEqual(x0, x1) {
		t.Errorf("not equal")
	}
}
