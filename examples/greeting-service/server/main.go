package main

import (
	"net/http/httptest"

	"github.com/ipld/edelweiss/examples/greeting-service/api/proto"
)

func main() {
	httptest.NewServer(proto.GreetingService_AsyncHandler(GreetingServiceImplementation{}))
	<-(chan int)(nil)
}
