package test

import (
	"testing"

	"github.com/ipld/edelweiss/defs"
)

func TestService(t *testing.T) {
	defs := defs.Defs{
		defs.Named{Name: "TestService1",
			Type: defs.Service{
				Methods: defs.Methods{
					defs.Method{Name: "Method1", Type: defs.Fn{Arg: defs.Int{}, Return: defs.Bool{}}},
					defs.Method{Name: "Method2", Type: defs.Fn{Arg: defs.String{}, Return: defs.Float{}}},
				},
			},
		},
		defs.Named{Name: "TestService2",
			Type: defs.Service{
				Methods: defs.Methods{
					defs.Method{Name: "Method1", Type: defs.Fn{Arg: defs.Int{}, Return: defs.Bool{}}},
					defs.Method{Name: "Method2", Type: defs.Fn{Arg: defs.String{}, Return: defs.Float{}}},
					defs.Method{Name: "Method3", Type: defs.Fn{Arg: defs.String{}, Return: defs.String{}}},
				},
			},
		},
	}
	testSrc := `

type TestService1_ServerImpl struct{}

func (TestService1_ServerImpl) Method1(ctx context.Context, req *values.Int) (<-chan *TestService1_Method1_AsyncResult, error) {
	respCh := make(chan *TestService1_Method1_AsyncResult)
	go func() {
		defer close(respCh)
		var r1 values.Bool = true
		respCh <- &TestService1_Method1_AsyncResult{ Resp: &r1 }
	}()
	return respCh, nil
}

func (TestService1_ServerImpl) Method2(ctx context.Context, req *values.String) (<-chan *TestService1_Method2_AsyncResult, error) {
	respCh := make(chan *TestService1_Method2_AsyncResult)
	go func() {
		defer close(respCh)
		var r1 values.Float = 1.23
		respCh <- &TestService1_Method2_AsyncResult{ Resp: &r1 }
		// TODO: dagjson.Decode does not support multiple streaming values
		// var r2 values.Float = 4.56
		// respCh <- &TestService1_Method2_AsyncResult{ Resp: &r2 }
	}()
	return respCh, nil
}

var testServiceIdentifyResult = &TestService1_IdentifyResult{
	Methods: []values.String{"Method1", "Method2"},
}

func TestRoundtrip(t *testing.T) {

	s := httptest.NewServer(TestService1_AsyncHandler(TestService1_ServerImpl{}))
	defer s.Close()

	c1, err := New_TestService1_Client(s.URL, TestService1_Client_WithHTTPClient(s.Client()))
	if err != nil {
		t.Fatal(err)
	}

	c2, err := New_TestService2_Client(s.URL, TestService2_Client_WithHTTPClient(s.Client()))
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()

	r1, err := c1.Method1(ctx, values.NewInt(5))
	if err != nil {
		t.Fatal(err)
	}
	if *r1[0] != values.Bool(true) {
		t.Errorf("expecting true, fot false")
	}

	r2, err := c1.Method2(ctx, values.NewString("5"))
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

	// test Identify method
	r3, err := c1.Identify(ctx, &TestService1_IdentifyArg{})
	if err != nil {
		t.Fatal(err)
	}
	if len(r3) != 1 {
		t.Fatalf("expecting 1 results, got %d", len(r3))
	}
	if !ipld.DeepEqual(r3[0], testServiceIdentifyResult) {
		t.Fatalf("expecting #%v, got %v", testServiceIdentifyResult, r3[0])
	}

	// test error handling when unsupported method is called
	_, err = c2.Method3(ctx, values.NewString("X"))
	if err != services.ErrSchema {
		t.Errorf("expecting error %v, got %v", services.ErrSchema, err)
	}

}`
	RunGenTest(t, defs, testSrc)
}
