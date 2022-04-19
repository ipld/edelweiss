package main

import (
	"context"
	"fmt"

	"github.com/ipld/edelweiss/examples/greeting-service/api/proto"
	"github.com/ipld/edelweiss/values"
)

type GreetingServiceImplementation struct{}

func (GreetingServiceImplementation) Hello(ctx context.Context, req *proto.HelloRequest, respCh chan<- *proto.GreetingService_Hello_AsyncResult) error {
	name := req.Name
	resp := &proto.HelloResponse{}
	switch {
	case req.Address.US != nil:
		var state string
		switch {
		case req.Address.US.State.CA != nil:
			state = "California"
		case req.Address.US.State.NY != nil:
			state = "New York"
		case req.Address.US.State.Other != nil:
			state = string(*req.Address.US.State.Other)
		}
		greeting := values.String(fmt.Sprintf("Hello %s, from %s, US!", name, state))
		resp.English = &greeting
	case req.Address.SK != nil:
		greeting := values.String(fmt.Sprintf("Hello %s, from %s, South Korea!", name, req.Address.SK.Province))
		resp.Korean = &greeting
	case req.Address.OtherAddress != nil:
		greeting := values.String(fmt.Sprintf("Hello %s, from %s!", name, req.Address.OtherCountry))
		resp.English = &greeting
	}
	return fmt.Errorf("no valid address")
}
