package test

import (
	"testing"

	"github.com/ipld/edelweiss/defs"
)

func TestGenTest(t *testing.T) {
	RunSingleGenTest(t, defs.Defs{defs.Named{Name: "T", Type: defs.SingletonInt{Int: 23}}}, "")
}

func TestSingletonAtRunTime(t *testing.T) {
	defs := []defs.Defs{
		{defs.Named{Name: "UserSingleton", Type: defs.SingletonBool{Bool: true}}},
		{defs.Named{Name: "UserSingleton", Type: defs.SingletonInt{Int: 23}}},
		{defs.Named{Name: "UserSingleton", Type: defs.SingletonFloat{Float: 2.3}}},
		{defs.Named{Name: "UserSingleton", Type: defs.SingletonByte{Byte: 2}}},
		{defs.Named{Name: "UserSingleton", Type: defs.SingletonChar{Char: 'a'}}},
		{defs.Named{Name: "UserSingleton", Type: defs.SingletonString{String: "abc"}}},
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
	defs := []defs.Defs{
		{defs.Named{
			Name: "UserStructure",
			Type: defs.Structure{
				Fields: defs.Fields{
					defs.Field{Name: "a", GoName: "A", Type: defs.Int{}},
					defs.Field{Name: "B", GoName: "", Type: defs.String{}},
					defs.Field{Name: "C", GoName: "", Type: defs.Float{}},
					defs.Field{Name: "D", GoName: "", Type: defs.Byte{}},
					defs.Field{Name: "E", GoName: "", Type: defs.Char{}},
				},
			},
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
	defs := []defs.Defs{
		{defs.Named{
			Name: "UserInductive",
			Type: defs.Inductive{
				Cases: defs.Cases{
					defs.Case{Name: "A", Type: defs.Int{}},
					defs.Case{Name: "B", Type: defs.String{}},
					defs.Case{Name: "C", Type: defs.Float{}},
					defs.Case{Name: "D", Type: defs.Byte{}},
					defs.Case{Name: "E", Type: defs.Char{}},
				},
			},
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

func TestUnionAtRunTime(t *testing.T) {
	defs := []defs.Defs{
		{defs.Named{
			Name: "UserUnion",
			Type: defs.Union{
				Cases: defs.Cases{
					defs.Case{Name: "a", GoName: "A", Type: defs.Int{}},
					defs.Case{Name: "B", GoName: "", Type: defs.String{}},
					defs.Case{Name: "C", GoName: "", Type: defs.Float{}},
				},
			},
		}},
	}
	testSrc := `
	var x1 UserUnion
	var y values.String = "abc"
	x1.B = &y
	buf, err := ipld.Encode(x1, dagjson.Encode)
	if err != nil {
		t.Fatalf("encoding (%v)", err)
	}
	var x2 UserUnion
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
	defs := []defs.Defs{
		{defs.Named{
			Name: "UserMap",
			Type: defs.Map{Key: defs.String{}, Value: defs.Int{}},
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

// IPLD DAGJSON encoding does not support maps with non-string keys.
// This can be remedied on the Edelweiss side by using encoding non-string key maps into a list of pairs repn.
// func _TestNonStringMapAtRunTime(t *testing.T) {
// 	defs := []defs.Defs{
// 		{defs.Named{
// 			Name: "UserMap",
// 			Type: defs.Map{Key: defs.Float{}, Value: defs.Int{}},
// 		}},
// 	}
// 	testSrc := `
// 	var x1 UserMap = UserMap{
// 		{Key: 123.456, Value: 1},
// 		{Key: 456.789, Value: 2},
// 	}
// 	buf, err := ipld.Encode(x1, dagjson.Encode)
// 	if err != nil {
// 		t.Fatalf("encoding (%v)", err)
// 	}
// 	fmt.Println(string(buf))
// 	var x2 UserMap
// 	n, err := ipld.Decode(buf, dagjson.Decode)
// 	if err != nil {
// 		t.Fatalf("decoding (%v)", err)
// 	}
// 	if err = x2.Parse(n); err != nil {
// 		t.Fatalf("parsing (%v)", err)
// 	}
// 	if !ipld.DeepEqual(x1, x2) {
// 		t.Errorf("ipld values are not equal")
// 	}
// `
// 	for _, d := range defs {
// 		RunSingleGenTest(t, d, testSrc)
// 	}
// }

func TestListAtRunTime(t *testing.T) {
	defs := []defs.Defs{
		{defs.Named{
			Name: "UserList",
			Type: defs.List{Element: defs.String{}},
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
	defs := []defs.Defs{
		{defs.Named{
			Name: "UserCall",
			Type: defs.Call{
				Fn: defs.Fn{Arg: defs.Int{}, Return: defs.String{}},
				ID: defs.String{},
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
	defs := []defs.Defs{
		{defs.Named{
			Name: "UserReturn",
			Type: defs.Return{
				Fn: defs.Fn{Arg: defs.Int{}, Return: defs.String{}},
				ID: defs.String{},
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
	defs := []defs.Defs{
		{defs.Named{
			Name: "UserLink",
			Type: defs.Link{To: defs.String{}},
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
	defs := []defs.Defs{
		{defs.Named{
			Name: "UserStructure",
			Type: defs.Structure{
				Fields: defs.Fields{
					defs.Field{Name: "A", Type: defs.Int{}},
					defs.Field{Name: "B", Type: defs.Named{
						Name: "UserInductive",
						Type: defs.Inductive{
							Cases: defs.Cases{
								defs.Case{Name: "x", GoName: "X", Type: defs.String{}},
								defs.Case{Name: "Y", GoName: "", Type: defs.Int{}},
							},
						}},
					},
				},
			},
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
	defs := []defs.Defs{
		{defs.Named{
			Name: "UserList",
			Type: defs.List{
				Element: defs.Structure{
					Fields: defs.Fields{
						defs.Field{Name: "X", Type: defs.String{}},
					},
				},
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
	defs := []defs.Defs{
		{defs.Named{
			Name: "UserList",
			Type: defs.List{
				Element: defs.SingletonBool{Bool: true},
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

func TestUnionListAtRunTime(t *testing.T) {
	defs := []defs.Defs{
		{defs.Named{
			Name: "UserUnion",
			Type: defs.Union{
				Cases: defs.Cases{
					defs.Case{
						Name: "A",
						Type: defs.Named{
							Name: "UserStructure",
							Type: defs.Structure{
								Fields: defs.Fields{
									defs.Field{Name: "F", Type: defs.Int{}},
								},
							},
						},
					},
					defs.Case{
						Name: "B",
						Type: defs.Named{
							Name: "UserList",
							Type: defs.List{Element: defs.String{}},
						},
					},
				},
			},
		}},
	}
	testSrc := `
	var x1 UserUnion
	var y UserList = UserList{values.String("abc")}
	x1.B = &y
	buf, err := ipld.Encode(x1, dagjson.Encode)
	if err != nil {
		t.Fatalf("encoding (%v)", err)
	}
	var x2 UserUnion
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
