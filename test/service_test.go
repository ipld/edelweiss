package test

import (
	"testing"

	"github.com/ipld/edelweiss/def"
)

func TestService(t *testing.T) {
	defs := def.Types{
		def.Named{Name: "TestService",
			Type: def.MakeService(
				def.Method{Name: "Method1", Type: def.Fn{Arg: def.Int{}, Return: def.Bool{}}},
				def.Method{Name: "Method2", Type: def.Fn{Arg: def.String{}, Return: def.Float{}}},
			),
		},
	}
	testSrc := `

type TestService_ServerImpl struct{}

func (TestService_ServerImpl) Method1(ctx context.Context, req *values.Int, respCh chan<- *values.Bool) error {
	go func() {
		defer close(respCh)
		var r1 values.Bool = true
		respCh <- &r1
	}()
	return nil
}

func (TestService_ServerImpl) Method2(ctx context.Context, req *values.String, respCh chan<- *values.Float) error {
	go func() {
		defer close(respCh)
		var r1 values.Float = 1.23
		respCh <- &r1
		// TODO: dagjson.Decode does not support multiple streaming values
		// var r2 values.Float = 4.56
		// respCh <- &r2
	}()
	return nil
}

func TestRoundtrip(t *testing.T) {

	s := httptest.NewServer(TestService_AsyncHandler(TestService_ServerImpl{}))
	defer s.Close()

	c, err := New_TestService_Client(s.URL, TestService_Client_WithHTTPClient(s.Client()))
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()

	r1, err := c.Method1(ctx, values.NewInt(5))
	if err != nil {
		t.Fatal(err)
	}
	if *r1[0] != values.Bool(true) {
		t.Errorf("expecting true, fot false")
	}

	r2, err := c.Method2(ctx, values.NewString("5"))
	if err != nil {
		t.Fatal(err)
	}
	if len(r2) != 1 {
		t.Fatalf("expecting 1 results, got %d", len(r2))
	}
	if *r2[0] != values.Float(1.23) {
		t.Fatalf("expecting 1.23, got %v", *r2[0])
	}
	// TODO: dagjson.Decode does not support multiple streaming values
	// if *r2[1] != values.Float(4.56) {
	// 	t.Fatalf("expecting 4.56, got %v", *r2[1])
	// }

}`
	RunGenTest(t, defs, testSrc)
}
