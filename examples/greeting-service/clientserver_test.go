package helloservice

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/ipld/edelweiss/examples/greeting-service/api/proto"
	"github.com/ipld/edelweiss/examples/greeting-service/service"
	"github.com/ipld/edelweiss/values"
	"github.com/ipld/go-ipld-prime"
)

var testGreetingResponse = &proto.HelloResponse{
	English: values.NewString("Hello TestName, from California, US!"),
}

var testGreetingServiceIdentifyResult = &proto.GreetingService_IdentifyResult{
	Methods: []values.String{"Hello"},
}

func TestRoundtrip(t *testing.T) {

	s := httptest.NewServer(proto.GreetingService_AsyncHandler(service.GreetingServiceImplementation{}))
	defer s.Close()

	c, err := proto.New_GreetingService_Client(s.URL, proto.GreetingService_Client_WithHTTPClient(s.Client()))
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()

	// test calling Hello
	req := &proto.HelloRequest{
		Name: "TestName",
		Address: proto.Address{
			US: &proto.USAddress{
				Street: "1955 Valley Dr",
				City:   "Mariposa",
				State: proto.State{
					CA: &proto.StateCA{},
				},
				ZIP: values.Int(22355),
			},
		},
	}

	r1, err := c.Hello(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	if len(r1) != 1 {
		t.Fatalf("expecting 1 result, got %d", len(r1))
	}
	if !ipld.DeepEqual(r1[0], testGreetingResponse) {
		t.Errorf("expecting %v, got %v", testGreetingResponse, r1[0])
	}

	// test calling Identify
	r2, err := c.Identify(ctx, &proto.GreetingService_IdentifyArg{})
	if err != nil {
		t.Fatal(err)
	}
	if len(r2) != 1 {
		t.Fatalf("expecting 1 results, got %d", len(r2))
	}
	if !ipld.DeepEqual(r2[0], testGreetingServiceIdentifyResult) {
		t.Fatalf("expecting %v, got %v", testGreetingServiceIdentifyResult, r2[0])
	}

}
