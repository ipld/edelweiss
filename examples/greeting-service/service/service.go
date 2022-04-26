package service

import (
	"context"
	"fmt"

	log "github.com/ipfs/go-log"

	"github.com/ipld/edelweiss/examples/greeting-service/api/proto"
	"github.com/ipld/edelweiss/values"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
)

var serverLogger = log.Logger("service/GreetingService")

type GreetingServiceImplementation struct{}

func (GreetingServiceImplementation) Hello(ctx context.Context, req *proto.HelloRequest) (<-chan *proto.GreetingService_Hello_AsyncResult, error) {
	respCh := make(chan *proto.GreetingService_Hello_AsyncResult)
	go func() {
		defer close(respCh)

		buf, err := ipld.Encode(req, dagjson.Encode)
		if err != nil {
			serverLogger.Fatal(err)
		}
		serverLogger.Infof("received request: %s\n", string(buf))
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
			respCh <- &proto.GreetingService_Hello_AsyncResult{Resp: resp}
		case req.Address.SK != nil:
			greeting := values.String(fmt.Sprintf("Hello %s, from %s, South Korea!", name, req.Address.SK.Province))
			resp.Korean = &greeting
			respCh <- &proto.GreetingService_Hello_AsyncResult{Resp: resp}
		case req.Address.OtherAddress != nil:
			greeting := values.String(fmt.Sprintf("Hello %s, from %s!", name, req.Address.OtherCountry))
			resp.English = &greeting
			respCh <- &proto.GreetingService_Hello_AsyncResult{Resp: resp}
		default:
			respCh <- &proto.GreetingService_Hello_AsyncResult{Err: fmt.Errorf("no valid address")}
		}
	}()
	return respCh, nil
}
