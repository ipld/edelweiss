package test

import (
	"testing"

	"github.com/ipld/edelweiss/defs"
)

func TestService(t *testing.T) {
	defs := defs.Defs{
		defs.Named{Name: "TestService",
			Type: defs.Service{
				Methods: defs.Methods{
					defs.Method{Name: "Method1", Type: defs.Fn{Arg: defs.Int{}, Return: defs.Bool{}}},
					defs.Method{Name: "Method2", Type: defs.Fn{Arg: defs.String{}, Return: defs.Float{}}},
				},
			},
		},
	}
	testSrc := `

type TestService_ServerImpl struct{}

func (TestService_ServerImpl) Method1(ctx context.Context, req *values.Int, respCh chan<- *TestService_Method1_AsyncResult) error {
	go func() {
		defer close(respCh)
		var r1 values.Bool = true
		respCh <- &TestService_Method1_AsyncResult{ Resp: &r1 }
	}()
	return nil
}

func (TestService_ServerImpl) Method2(ctx context.Context, req *values.String, respCh chan<- *TestService_Method2_AsyncResult) error {
	go func() {
		defer close(respCh)
		var r1 values.Float = 1.23
		respCh <- &TestService_Method2_AsyncResult{ Resp: &r1 }
		// TODO: dagjson.Decode does not support multiple streaming values
		// var r2 values.Float = 4.56
		// respCh <- &TestService_Method2_AsyncResult{ Resp: &r2 }
	}()
	return nil
}

var testServiceIdentifyResult = &TestService_IdentifyResult{
	Methods: []values.String{"Method1", "Method2"},
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

	r3, err := c.Identify(ctx, &TestService_IdentifyArg{})
	if err != nil {
		t.Fatal(err)
	}
	if len(r3) != 1 {
		t.Fatalf("expecting 1 results, got %d", len(r3))
	}
	if !ipld.DeepEqual(r3[0], testServiceIdentifyResult) {
		t.Fatalf("expecting #%v, got %v", testServiceIdentifyResult, r3[0])
	}

}`
	RunGenTest(t, defs, testSrc)
}
