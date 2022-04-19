package main

import (
	"context"
	"flag"
	"fmt"
	"strings"

	log "github.com/ipfs/go-log"
	"github.com/ipld/edelweiss/examples/greeting-service/api/proto"
	"github.com/ipld/edelweiss/values"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
)

var flagAddress = flag.String("http", "http://localhost:8080", "http server address")
var flagName = flag.String("name", "Alice", "your name")
var flagCountry = flag.String("country", "sk", "country: us, sk or anything else")

var clientLogger = log.Logger("client/GreetingService")

func main() {
	flag.Parse()
	c, err := proto.New_GreetingService_Client(*flagAddress)
	if err != nil {
		clientLogger.Fatal(err)
	}

	req := &proto.HelloRequest{
		Name:    values.String(*flagName),
		Address: proto.Address{},
	}
	switch strings.ToLower(*flagCountry) {
	case "us":
		req.Address.US = &proto.USAddress{
			Street: "1955 Valley Dr",
			City:   "Mariposa",
			State: proto.State{
				CA: &proto.StateCA{},
			},
			ZIP: 22355,
		}
	case "sk":
		req.Address.SK = &proto.SKAddress{
			Street:     "Gangnam",
			City:       "Cheongju",
			Province:   "충청북도",
			PostalCode: 123,
		}
	default:
		req.Address.OtherCountry = *flagCountry
		req.Address.OtherAddress = &proto.AddressLines{"Other St", "Other City"}
	}

	ctx := context.Background()
	results, err := c.Hello(ctx, req)
	if err != nil {
		clientLogger.Fatal(err)
	}
	for i, r := range results {
		buf, err := ipld.Encode(r, dagjson.Encode)
		if err != nil {
			clientLogger.Fatal(err)
		}
		fmt.Printf("greeting response #%d: %s\n", i+1, string(buf))
	}
}
