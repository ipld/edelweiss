package main

import (
	"flag"
	"net/http"

	log "github.com/ipfs/go-log"
	"github.com/ipld/edelweiss/examples/greeting-service/api/proto"
)

var serverLogger = log.Logger("server/GreetingService")

var flagAddress = flag.String("bind", ":8080", "http bind address")

func main() {
	flag.Parse()
	serverLogger.Infof("Starting GreetingService on %s", *flagAddress)
	s := &http.Server{
		Addr:    *flagAddress,
		Handler: proto.GreetingService_AsyncHandler(GreetingServiceImplementation{}),
	}
	serverLogger.Fatal(s.ListenAndServe())
	<-(chan int)(nil)
}
