package test

import (
	"testing"

	"github.com/ipld/edelweiss/defs"
)

func TestServiceWithoutErrors(t *testing.T) {
	defs := defs.Defs{
		defs.Named{Name: "TestService1",
			Type: defs.Service{
				Methods: defs.Methods{
					defs.Method{Name: "Method1", Type: defs.Fn{Arg: defs.Int{}, Return: defs.Bool{}}},
					defs.Method{Name: "Method2", Type: defs.Fn{Arg: defs.String{}, Return: defs.Float{}}},
					defs.Method{Name: "Method4", Type: defs.Fn{Arg: defs.String{}, Return: defs.Float{}}},
				},
			},
		},
		defs.Named{Name: "TestService2",
			Type: defs.Service{
				Methods: defs.Methods{
					defs.Method{Name: "Method1", Type: defs.Fn{Arg: defs.Int{}, Return: defs.Bool{}}},
					defs.Method{Name: "Method2", Type: defs.Fn{Arg: defs.String{}, Return: defs.Float{}}},
					defs.Method{Name: "Method3", Type: defs.Fn{Arg: defs.String{}, Return: defs.String{}}},
					defs.Method{Name: "Method4", Type: defs.Fn{Arg: defs.String{}, Return: defs.Float{}}},
				},
			},
		},
	}
	testSrc := `

type TestService1_ServerImpl struct{
	Release4 chan struct{}
}

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
		var r2 values.Float = 4.56
		respCh <- &TestService1_Method2_AsyncResult{ Resp: &r2 }
	}()
	return respCh, nil
}

// This method returns one response and stalls. The test verifies that the first response is received.
func (s TestService1_ServerImpl) Method4(ctx context.Context, req *values.String) (<-chan *TestService1_Method4_AsyncResult, error) {
	respCh := make(chan *TestService1_Method4_AsyncResult)
	go func() {
		defer close(respCh)
		var r1 values.Float = 1.23
		respCh <- &TestService1_Method4_AsyncResult{ Resp: &r1 }
		<-s.Release4
	}()
	return respCh, nil
}

var testServiceIdentifyResult = &TestService1_IdentifyResult{
	Methods: []values.String{"Method1", "Method2", "Method4"},
}

func TestRoundtrip(t *testing.T) {

	svc := TestService1_ServerImpl{Release4: make(chan struct{})}
	s := httptest.NewServer(TestService1_AsyncHandler(svc))
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

	// test method returning one response
	r1, err := c1.Method1(ctx, values.NewInt(5))
	if err != nil {
		t.Fatal(err)
	}
	if *r1[0] != values.Bool(true) {
		t.Errorf("expecting true, fot false")
	}

	// test method returning two responses
	r2, err := c1.Method2(ctx, values.NewString("5"))
	if err != nil {
		t.Fatal(err)
	}
	if len(r2) != 2 {
		t.Fatalf("expecting 2 results, got %d", len(r2))
	}
	if *r2[0] != values.Float(1.23) {
		t.Fatalf("expecting 1.23, got %v", *r2[0])
	}
	if *r2[1] != values.Float(4.56) {
		t.Fatalf("expecting 4.56, got %v", *r2[1])
	}

	// test method that returns one response and stalls
	rch, err := c1.Method4_Async(ctx, values.NewString("5"))
	if err != nil {
		t.Fatal(err)
	}
	resp := <-rch
	if resp.Resp == nil {
		t.Fatalf("expecting response")
	}
	if *resp.Resp != values.Float(1.23) {
		t.Fatalf("expecting 1.23, got %v", *r2[0])
	}
	close(svc.Release4)

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

func TestServiceWithErrors(t *testing.T) {
	defs := defs.Defs{
		defs.Named{Name: "TestService1",
			Type: defs.Service{
				Methods: defs.Methods{
					defs.Method{Name: "Method1", Type: defs.Fn{Arg: defs.Int{}, Return: defs.Bool{}}},
					defs.Method{Name: "Method2", Type: defs.Fn{Arg: defs.String{}, Return: defs.Float{}}},
				},
			},
		},
	}
	testSrc := `

var testAsyncError = "async error"
var testSyncError = "sync error"

type TestService1_ServerImpl struct{}

func (TestService1_ServerImpl) Method1(ctx context.Context, req *values.Int) (<-chan *TestService1_Method1_AsyncResult, error) {
	respCh := make(chan *TestService1_Method1_AsyncResult)
	go func() {
		defer close(respCh)
		respCh <- &TestService1_Method1_AsyncResult{ Err: fmt.Errorf(testAsyncError) }
	}()
	return respCh, nil
}

func (TestService1_ServerImpl) Method2(ctx context.Context, req *values.String) (<-chan *TestService1_Method2_AsyncResult, error) {
	return nil, fmt.Errorf(testSyncError)
}

func TestRoundtrip(t *testing.T) {

	s := httptest.NewServer(TestService1_AsyncHandler(TestService1_ServerImpl{}))
	defer s.Close()

	c1, err := New_TestService1_Client(s.URL, TestService1_Client_WithHTTPClient(s.Client()))
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()

	ch1, err := c1.Method1_Async(ctx, values.NewInt(5))
	if err != nil {
		t.Fatalf("sync error not expected (%v)", err)
	}
	for asyncResult := range ch1 {
		if asyncResult.Err == nil {
			t.Errorf("async error expectded")
		}
		if asyncResult.Err.Error() != testAsyncError {
			t.Errorf("expected %v, got %v", testAsyncError, asyncResult.Err.Error())
		}
	}

	_, err = c1.Method2_Async(ctx, values.NewString("5"))
	if err == nil {
		t.Fatalf("sync error expected")
	}
	if err.Error() != testSyncError {
		t.Fatalf("expected %v, got %v", testSyncError, err.Error())
	}
}`
	RunGenTest(t, defs, testSrc)
}
