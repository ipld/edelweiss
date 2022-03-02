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

func (TestService_ServerImpl) Method1(ctx pd9.Context, req *pd2.Int, respCh chan<- *pd2.Bool) error {
	defer close(respCh)
	var r1 pd2.Bool = true
	respCh <- &r1
}

// XXX: pkg aliases
func (TestService_ServerImpl) Method2(ctx pd9.Context, req *pd2.String, respCh chan<- *pd2.Float) error {
	defer close(respCh)
	var r1 pd2.Float = 1.23
	respCh <- &r1
	var r2 pd2.Float = 4.56
	respCh <- &r2
}

func TestRoundtrip(t *testing.T) {

	s := httptest.NewServer(TestService_AsyncHandler(TestService_ServerImpl{}))
	defer s.Close()

	c, err := New_TestService_Client(s.URL, TestService_Client_WithHTTPClient(s.Client()))
	if err != nil {
		t.Fatal(err)
	}

	XXX
}`
	RunGenTest(t, defs, testSrc)
}
