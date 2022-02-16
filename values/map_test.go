package values

import (
	"testing"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
)

func TestMapRoundtrip(t *testing.T) {
	x0 := Map{
		KeyValue{Key: Any{String("f1")}, Value: Any{Bool(true)}},
		KeyValue{Key: Any{String("f2")}, Value: Any{Bool(false)}},
		KeyValue{Key: Any{String("f3")}, Value: Any{String("haha")}},
		KeyValue{Key: Any{String("f4")}, Value: Any{Float(2.3)}},
		KeyValue{Key: Any{String("f5")}, Value: Any{Int(7)}},
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
	var x1 Map
	if err := x1.Parse(n); err != nil {
		t.Fatal(err)
	}
	if !ipld.DeepEqual(x0.Node(), x1.Node()) {
		t.Errorf("not equal")
	}
}
