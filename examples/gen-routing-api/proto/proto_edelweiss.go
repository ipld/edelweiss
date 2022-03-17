package proto

import(
	pd1 "github.com/ipld/go-ipld-prime/datamodel"
	pd2 "github.com/ipld/edelweiss/values"
	pd3 "fmt"
	pd4 "net/http"
	pd5 "github.com/ipld/go-ipld-prime"
	pd6 "context"
	pd7 "github.com/ipld/go-ipld-prime/codec/dagjson"
	pd8 "errors"
	pd9 "bytes"
	pd10 "io"
	pd11 "net/url"
	pd12 "github.com/ipfs/go-log"
)


// -- protocol type AnonInductive0 --

type AnonInductive0 struct {
		GetP2PProvide *GetP2PProvideRequest

}

func (x *AnonInductive0) Parse(n pd1.Node) error {
	*x = AnonInductive0{}
	if n.Kind() != pd1.Kind_Map {
		return pd2.ErrNA
	}
	iter := n.MapIterator()
	kn, vn, err := iter.Next()
	if err != nil {
		return err
	}
	k, err := kn.AsString()
	if err != nil {
		return pd3.Errorf("inductive map key is not a string")
	}
	switch k {
	case "GetP2PProvide":
		var y GetP2PProvideRequest
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.GetP2PProvide = &y
		return nil

	}
	return pd3.Errorf("inductive map has no applicable keys")
}

type AnonInductive0_MapIterator struct {
	done bool
	s    *AnonInductive0
}

func (x *AnonInductive0_MapIterator) Next() (key pd1.Node, value pd1.Node, err error) {
	if x.done {
		return nil, nil, pd2.ErrNA
	} else {
		x.done = true
		switch {
			case x.s.GetP2PProvide != nil:
			return pd2.String("GetP2PProvide"), x.s.GetP2PProvide.Node(), nil

		default:
			return nil, nil, pd3.Errorf("no inductive cases are set")
		}
	}
}

func (x *AnonInductive0_MapIterator) Done() bool {
	return x.done
}

func (x AnonInductive0) Node() pd1.Node {
	return x
}

func (x AnonInductive0) Kind() pd1.Kind {
	return pd1.Kind_Map
}

func (x AnonInductive0) LookupByString(key string) (pd1.Node, error) {
	switch {
		case x.GetP2PProvide != nil && key == "GetP2PProvide":
		return x.GetP2PProvide.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x AnonInductive0) LookupByNode(key pd1.Node) (pd1.Node, error) {
	if key.Kind() != pd1.Kind_String {
		return nil, pd2.ErrNA
	}
	if s, err := key.AsString(); err != nil {
		return nil, err
	} else {
		return x.LookupByString(s)
	}
}

func (x AnonInductive0) LookupByIndex(idx int64) (pd1.Node, error) {
	return nil, pd2.ErrNA
}

func (x AnonInductive0) LookupBySegment(seg pd1.PathSegment) (pd1.Node, error) {
	switch seg.String() {
		case "GetP2PProvide":
		return x.GetP2PProvide.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x AnonInductive0) MapIterator() pd1.MapIterator {
	return &AnonInductive0_MapIterator{false, &x}
}

func (x AnonInductive0) ListIterator() pd1.ListIterator {
	return nil
}

func (x AnonInductive0) Length() int64 {
	return 1
}

func (x AnonInductive0) IsAbsent() bool {
	return false
}

func (x AnonInductive0) IsNull() bool {
	return false
}

func (x AnonInductive0) AsBool() (bool, error) {
	return false, pd2.ErrNA
}

func (x AnonInductive0) AsInt() (int64, error) {
	return 0, pd2.ErrNA
}

func (x AnonInductive0) AsFloat() (float64, error) {
	return 0, pd2.ErrNA
}

func (x AnonInductive0) AsString() (string, error) {
	return "", pd2.ErrNA
}

func (x AnonInductive0) AsBytes() ([]byte, error) {
	return nil, pd2.ErrNA
}

func (x AnonInductive0) AsLink() (pd1.Link, error) {
	return nil, pd2.ErrNA
}

func (x AnonInductive0) Prototype() pd1.NodePrototype {
	return nil
}
// -- protocol type AnonInductive1 --

type AnonInductive1 struct {
		GetP2PProvide *GetP2PProvideResponse

}

func (x *AnonInductive1) Parse(n pd1.Node) error {
	*x = AnonInductive1{}
	if n.Kind() != pd1.Kind_Map {
		return pd2.ErrNA
	}
	iter := n.MapIterator()
	kn, vn, err := iter.Next()
	if err != nil {
		return err
	}
	k, err := kn.AsString()
	if err != nil {
		return pd3.Errorf("inductive map key is not a string")
	}
	switch k {
	case "GetP2PProvide":
		var y GetP2PProvideResponse
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.GetP2PProvide = &y
		return nil

	}
	return pd3.Errorf("inductive map has no applicable keys")
}

type AnonInductive1_MapIterator struct {
	done bool
	s    *AnonInductive1
}

func (x *AnonInductive1_MapIterator) Next() (key pd1.Node, value pd1.Node, err error) {
	if x.done {
		return nil, nil, pd2.ErrNA
	} else {
		x.done = true
		switch {
			case x.s.GetP2PProvide != nil:
			return pd2.String("GetP2PProvide"), x.s.GetP2PProvide.Node(), nil

		default:
			return nil, nil, pd3.Errorf("no inductive cases are set")
		}
	}
}

func (x *AnonInductive1_MapIterator) Done() bool {
	return x.done
}

func (x AnonInductive1) Node() pd1.Node {
	return x
}

func (x AnonInductive1) Kind() pd1.Kind {
	return pd1.Kind_Map
}

func (x AnonInductive1) LookupByString(key string) (pd1.Node, error) {
	switch {
		case x.GetP2PProvide != nil && key == "GetP2PProvide":
		return x.GetP2PProvide.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x AnonInductive1) LookupByNode(key pd1.Node) (pd1.Node, error) {
	if key.Kind() != pd1.Kind_String {
		return nil, pd2.ErrNA
	}
	if s, err := key.AsString(); err != nil {
		return nil, err
	} else {
		return x.LookupByString(s)
	}
}

func (x AnonInductive1) LookupByIndex(idx int64) (pd1.Node, error) {
	return nil, pd2.ErrNA
}

func (x AnonInductive1) LookupBySegment(seg pd1.PathSegment) (pd1.Node, error) {
	switch seg.String() {
		case "GetP2PProvide":
		return x.GetP2PProvide.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x AnonInductive1) MapIterator() pd1.MapIterator {
	return &AnonInductive1_MapIterator{false, &x}
}

func (x AnonInductive1) ListIterator() pd1.ListIterator {
	return nil
}

func (x AnonInductive1) Length() int64 {
	return 1
}

func (x AnonInductive1) IsAbsent() bool {
	return false
}

func (x AnonInductive1) IsNull() bool {
	return false
}

func (x AnonInductive1) AsBool() (bool, error) {
	return false, pd2.ErrNA
}

func (x AnonInductive1) AsInt() (int64, error) {
	return 0, pd2.ErrNA
}

func (x AnonInductive1) AsFloat() (float64, error) {
	return 0, pd2.ErrNA
}

func (x AnonInductive1) AsString() (string, error) {
	return "", pd2.ErrNA
}

func (x AnonInductive1) AsBytes() ([]byte, error) {
	return nil, pd2.ErrNA
}

func (x AnonInductive1) AsLink() (pd1.Link, error) {
	return nil, pd2.ErrNA
}

func (x AnonInductive1) Prototype() pd1.NodePrototype {
	return nil
}
var logger_client_DelegatedRouting = pd12.Logger("service/client/DelegatedRouting")

type DelegatedRouting_Client interface {

GetP2PProvide(ctx pd6.Context, req *GetP2PProvideRequest) ([]*GetP2PProvideResponse, error)


GetP2PProvide_Async(ctx pd6.Context, req *GetP2PProvideRequest) (<-chan DelegatedRouting_GetP2PProvide_AsyncResult, error)

}


type DelegatedRouting_GetP2PProvide_AsyncResult struct {
	Resp *GetP2PProvideResponse
	Err  error
}


type DelegatedRouting_ClientOption func(*client_DelegatedRouting) error

type client_DelegatedRouting struct {
	httpClient       *pd4.Client
	endpoint     *pd11.URL
}

func DelegatedRouting_Client_WithHTTPClient(hc *pd4.Client) DelegatedRouting_ClientOption {
	return func(c *client_DelegatedRouting) error {
		c.httpClient = hc
		return nil
	}
}

func New_DelegatedRouting_Client(endpoint string, opts ...DelegatedRouting_ClientOption) (*client_DelegatedRouting, error) {
	u, err := pd11.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	c := &client_DelegatedRouting{endpoint: u, httpClient: pd4.DefaultClient}
	for _, o := range opts {
		if err := o(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}



func (c *client_DelegatedRouting) GetP2PProvide(ctx pd6.Context, req *GetP2PProvideRequest) ([]*GetP2PProvideResponse, error) {
	ctx, cancel := pd6.WithCancel(ctx)
	defer cancel()
	ch, err := c.GetP2PProvide_Async(ctx, req)
	if err != nil {
		return nil, err
	}
	var resps []*GetP2PProvideResponse
	for {
		select {
		case r, ok := <-ch:
			if !ok {
				cancel()
				return resps, nil
			} else {
				if r.Err == nil {
					resps = append(resps, r.Resp)
				} else {
					logger_client_DelegatedRouting.Errorf("client received invalid response (%v)", r.Err)
				}
			}
		case <-ctx.Done():
			return resps, ctx.Err()
		}
	}
}

func (c *client_DelegatedRouting) GetP2PProvide_Async(ctx pd6.Context, req *GetP2PProvideRequest) (<-chan DelegatedRouting_GetP2PProvide_AsyncResult, error) {
	envelope := &AnonInductive0{
		GetP2PProvide: req,
	}

	buf, err := pd5.Encode(envelope, pd7.Encode)
	if err != nil {
		return nil, pd3.Errorf("unexpected serialization error (%v)", err)
	}

	// encode request in URL
	u := *c.endpoint
	q := pd11.Values{}
	q.Set("q", string(buf))
	u.RawQuery = q.Encode()
	httpReq, err := pd4.NewRequestWithContext(ctx, "GET", u.String(), pd9.NewReader(buf))
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, pd3.Errorf("sending HTTP request (%v)", err)
	}

	ch := make(chan DelegatedRouting_GetP2PProvide_AsyncResult, 1)
	go process_DelegatedRouting_GetP2PProvide_AsyncResult(ctx, ch, resp.Body)
	return ch, nil
}

func process_DelegatedRouting_GetP2PProvide_AsyncResult(ctx pd6.Context, ch chan<- DelegatedRouting_GetP2PProvide_AsyncResult, r pd10.Reader) {
	defer close(ch)
	for {
		if ctx.Err() != nil {
			return
		}

		n, err := pd5.DecodeStreaming(r, pd7.Decode)
		if pd8.Is(err, pd10.EOF) || pd8.Is(err, pd10.ErrUnexpectedEOF) {
			return
		}
		if err != nil {
			ch <- DelegatedRouting_GetP2PProvide_AsyncResult{Err: err}
			return
		}
		env := &AnonInductive1{}
		if err = env.Parse(n); err != nil {
			ch <- DelegatedRouting_GetP2PProvide_AsyncResult{Err: err}
			return
		}

		if env.GetP2PProvide == nil {
			continue
		}
		ch <- DelegatedRouting_GetP2PProvide_AsyncResult{Resp: env.GetP2PProvide}
	}
}


var logger_server_DelegatedRouting = pd12.Logger("service/server/DelegatedRouting")

type DelegatedRouting_Server interface {

	GetP2PProvide(ctx pd6.Context, req *GetP2PProvideRequest, respCh chan<- *GetP2PProvideResponse) error
}

func DelegatedRouting_AsyncHandler(s DelegatedRouting_Server) pd4.HandlerFunc {
	return func(writer pd4.ResponseWriter, request *pd4.Request) {
		// parse request
		msg := request.URL.Query().Get("q")
		n, err := pd5.Decode([]byte(msg), pd7.Decode)
		if err != nil {
			logger_server_DelegatedRouting.Errorf("received request not decodeable (%v)", err)
			writer.WriteHeader(400)
			return
		}
		env := &AnonInductive0{}
		if err = env.Parse(n); err != nil {
			logger_server_DelegatedRouting.Errorf("parsing call envelope (%v)", err)
			writer.WriteHeader(400)
			return
		}

		// demultiplex request
		switch {

		case env.GetP2PProvide != nil:
			ch := make(chan *GetP2PProvideResponse)
			if err = s.GetP2PProvide(pd6.Background(), env.GetP2PProvide, ch); err != nil {
				logger_server_DelegatedRouting.Errorf("get p2p provider rejected request (%v)", err)
				writer.WriteHeader(500)
				return
			}
			for resp := range ch {
				env := &AnonInductive1{ GetP2PProvide: resp }
				buf, err := pd5.Encode(env, pd7.Encode)
				if err != nil {
					logger_server_DelegatedRouting.Errorf("cannot encode response (%v)", err)
					continue
				}
				writer.Write(buf)
			}

		default:
			logger_server_DelegatedRouting.Errorf("missing or unknown request")
			writer.WriteHeader(404)
		}
	}
}

// -- protocol type AnonList3 --

type AnonList3 []Multihash

func (v AnonList3) Node() pd1.Node {
	return v
}

func (v *AnonList3) Parse(n pd1.Node) error {
	if n.Kind() != pd1.Kind_List {
		return pd2.ErrNA
	} else {
		*v = make(AnonList3, n.Length())
		iter := n.ListIterator()
		for !iter.Done() {
			if i, n, err := iter.Next(); err != nil {
				return pd2.ErrNA
			} else if err = (*v)[i].Parse(n); err != nil {
				return err
			}
		}
		return nil
	}
}

func (AnonList3) Kind() pd1.Kind {
	return pd1.Kind_List
}

func (AnonList3) LookupByString(string) (pd1.Node, error) {
	return nil, pd2.ErrNA
}

func (AnonList3) LookupByNode(key pd1.Node) (pd1.Node, error) {
	return nil, pd2.ErrNA
}

func (v AnonList3) LookupByIndex(i int64) (pd1.Node, error) {
	if i < 0 || i >= v.Length() {
		return nil, pd2.ErrBounds
	} else {
		return v[i].Node(), nil
	}
}

func (v AnonList3) LookupBySegment(seg pd1.PathSegment) (pd1.Node, error) {
	if i, err := seg.Index(); err != nil {
		return nil, pd2.ErrNA
	} else {
		return v.LookupByIndex(i)
	}
}

func (AnonList3) MapIterator() pd1.MapIterator {
	return nil
}

func (v AnonList3) ListIterator() pd1.ListIterator {
	return &AnonList3_ListIterator{v, 0}
}

func (v AnonList3) Length() int64 {
	return int64(len(v))
}

func (AnonList3) IsAbsent() bool {
	return false
}

func (AnonList3) IsNull() bool {
	return false
}

func (v AnonList3) AsBool() (bool, error) {
	return false, pd2.ErrNA
}

func (AnonList3) AsInt() (int64, error) {
	return 0, pd2.ErrNA
}

func (AnonList3) AsFloat() (float64, error) {
	return 0, pd2.ErrNA
}

func (AnonList3) AsString() (string, error) {
	return "", pd2.ErrNA
}

func (AnonList3) AsBytes() ([]byte, error) {
	return nil, pd2.ErrNA
}

func (AnonList3) AsLink() (pd1.Link, error) {
	return nil, pd2.ErrNA
}

func (AnonList3) Prototype() pd1.NodePrototype {
	return nil // not needed
}

type AnonList3_ListIterator struct {
	list AnonList3
	at   int64
}

func (iter *AnonList3_ListIterator) Next() (int64, pd1.Node, error) {
	if iter.Done() {
		return -1, nil, pd2.ErrBounds
	}
	v := iter.list[iter.at]
	i := int64(iter.at)
	iter.at++
	return i, v.Node(), nil
}

func (iter *AnonList3_ListIterator) Done() bool {
	return iter.at >= iter.list.Length()
}
// -- protocol type GetP2PProvideRequest --

type GetP2PProvideRequest struct {
		Keys AnonList3

}

func (x GetP2PProvideRequest) Node() pd1.Node {
	return x
}

func (x *GetP2PProvideRequest) Parse(n pd1.Node) error {
	if n.Kind() != pd1.Kind_Map {
		return pd2.ErrNA
	}
	iter := n.MapIterator()
	nfields := 0
	for !iter.Done() {
		if kn, vn, err := iter.Next(); err != nil {
			return err
		} else {
			if k, err := kn.AsString(); err != nil {
				return pd3.Errorf("structure map key is not a string")
			} else {
				_ = vn
				switch k {
			case "Keys":
			if err := x.Keys.Parse(vn); err != nil {
				return err
			}
			nfields++

				}
			}
		}
	}
	if nfields != 1 {
		return pd2.ErrNA
	} else {
		return nil
	}
}

type GetP2PProvideRequest_MapIterator struct {
	i int64
	s *GetP2PProvideRequest
}

func (x *GetP2PProvideRequest_MapIterator) Next() (key pd1.Node, value pd1.Node, err error) {
	x.i++
	switch x.i {
			case 0:
			return pd2.String("Keys"), x.s.Keys.Node(), nil

	}
	return nil, nil, pd2.ErrNA
}

func (x *GetP2PProvideRequest_MapIterator) Done() bool {
	return x.i + 1 >= 1
}

func (x GetP2PProvideRequest) Kind() pd1.Kind {
	return pd1.Kind_Map
}

func (x GetP2PProvideRequest) LookupByString(key string) (pd1.Node, error) {
	switch key {
		case "Keys":
		return x.Keys.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x GetP2PProvideRequest) LookupByNode(key pd1.Node) (pd1.Node, error) {
	switch key.Kind() {
	case pd1.Kind_String:
		if s, err := key.AsString(); err != nil {
			return nil, err
		} else {
			return x.LookupByString(s)
		}
	case pd1.Kind_Int:
		if i, err := key.AsInt(); err != nil {
			return nil, err
		} else {
			return x.LookupByIndex(i)
		}
	}
	return nil, pd2.ErrNA
}

func (x GetP2PProvideRequest) LookupByIndex(idx int64) (pd1.Node, error) {
	switch idx {
		case 0:
		return x.Keys.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x GetP2PProvideRequest) LookupBySegment(seg pd1.PathSegment) (pd1.Node, error) {
	switch seg.String() {
		case "0", "Keys":
		return x.Keys.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x GetP2PProvideRequest) MapIterator() pd1.MapIterator {
	return &GetP2PProvideRequest_MapIterator{-1, &x}
}

func (x GetP2PProvideRequest) ListIterator() pd1.ListIterator {
	return nil
}

func (x GetP2PProvideRequest) Length() int64 {
	return 1
}

func (x GetP2PProvideRequest) IsAbsent() bool {
	return false
}

func (x GetP2PProvideRequest) IsNull() bool {
	return false
}

func (x GetP2PProvideRequest) AsBool() (bool, error) {
	return false, pd2.ErrNA
}

func (x GetP2PProvideRequest) AsInt() (int64, error) {
	return 0, pd2.ErrNA
}

func (x GetP2PProvideRequest) AsFloat() (float64, error) {
	return 0, pd2.ErrNA
}

func (x GetP2PProvideRequest) AsString() (string, error) {
	return "", pd2.ErrNA
}

func (x GetP2PProvideRequest) AsBytes() ([]byte, error) {
	return nil, pd2.ErrNA
}

func (x GetP2PProvideRequest) AsLink() (pd1.Link, error) {
	return nil, pd2.ErrNA
}

func (x GetP2PProvideRequest) Prototype() pd1.NodePrototype {
	return nil
}

// -- protocol type AnonList5 --

type AnonList5 []ProvidersByKey

func (v AnonList5) Node() pd1.Node {
	return v
}

func (v *AnonList5) Parse(n pd1.Node) error {
	if n.Kind() != pd1.Kind_List {
		return pd2.ErrNA
	} else {
		*v = make(AnonList5, n.Length())
		iter := n.ListIterator()
		for !iter.Done() {
			if i, n, err := iter.Next(); err != nil {
				return pd2.ErrNA
			} else if err = (*v)[i].Parse(n); err != nil {
				return err
			}
		}
		return nil
	}
}

func (AnonList5) Kind() pd1.Kind {
	return pd1.Kind_List
}

func (AnonList5) LookupByString(string) (pd1.Node, error) {
	return nil, pd2.ErrNA
}

func (AnonList5) LookupByNode(key pd1.Node) (pd1.Node, error) {
	return nil, pd2.ErrNA
}

func (v AnonList5) LookupByIndex(i int64) (pd1.Node, error) {
	if i < 0 || i >= v.Length() {
		return nil, pd2.ErrBounds
	} else {
		return v[i].Node(), nil
	}
}

func (v AnonList5) LookupBySegment(seg pd1.PathSegment) (pd1.Node, error) {
	if i, err := seg.Index(); err != nil {
		return nil, pd2.ErrNA
	} else {
		return v.LookupByIndex(i)
	}
}

func (AnonList5) MapIterator() pd1.MapIterator {
	return nil
}

func (v AnonList5) ListIterator() pd1.ListIterator {
	return &AnonList5_ListIterator{v, 0}
}

func (v AnonList5) Length() int64 {
	return int64(len(v))
}

func (AnonList5) IsAbsent() bool {
	return false
}

func (AnonList5) IsNull() bool {
	return false
}

func (v AnonList5) AsBool() (bool, error) {
	return false, pd2.ErrNA
}

func (AnonList5) AsInt() (int64, error) {
	return 0, pd2.ErrNA
}

func (AnonList5) AsFloat() (float64, error) {
	return 0, pd2.ErrNA
}

func (AnonList5) AsString() (string, error) {
	return "", pd2.ErrNA
}

func (AnonList5) AsBytes() ([]byte, error) {
	return nil, pd2.ErrNA
}

func (AnonList5) AsLink() (pd1.Link, error) {
	return nil, pd2.ErrNA
}

func (AnonList5) Prototype() pd1.NodePrototype {
	return nil // not needed
}

type AnonList5_ListIterator struct {
	list AnonList5
	at   int64
}

func (iter *AnonList5_ListIterator) Next() (int64, pd1.Node, error) {
	if iter.Done() {
		return -1, nil, pd2.ErrBounds
	}
	v := iter.list[iter.at]
	i := int64(iter.at)
	iter.at++
	return i, v.Node(), nil
}

func (iter *AnonList5_ListIterator) Done() bool {
	return iter.at >= iter.list.Length()
}
// -- protocol type GetP2PProvideResponse --

type GetP2PProvideResponse struct {
		ProvidersByKey AnonList5

}

func (x GetP2PProvideResponse) Node() pd1.Node {
	return x
}

func (x *GetP2PProvideResponse) Parse(n pd1.Node) error {
	if n.Kind() != pd1.Kind_Map {
		return pd2.ErrNA
	}
	iter := n.MapIterator()
	nfields := 0
	for !iter.Done() {
		if kn, vn, err := iter.Next(); err != nil {
			return err
		} else {
			if k, err := kn.AsString(); err != nil {
				return pd3.Errorf("structure map key is not a string")
			} else {
				_ = vn
				switch k {
			case "ProvidersByKey":
			if err := x.ProvidersByKey.Parse(vn); err != nil {
				return err
			}
			nfields++

				}
			}
		}
	}
	if nfields != 1 {
		return pd2.ErrNA
	} else {
		return nil
	}
}

type GetP2PProvideResponse_MapIterator struct {
	i int64
	s *GetP2PProvideResponse
}

func (x *GetP2PProvideResponse_MapIterator) Next() (key pd1.Node, value pd1.Node, err error) {
	x.i++
	switch x.i {
			case 0:
			return pd2.String("ProvidersByKey"), x.s.ProvidersByKey.Node(), nil

	}
	return nil, nil, pd2.ErrNA
}

func (x *GetP2PProvideResponse_MapIterator) Done() bool {
	return x.i + 1 >= 1
}

func (x GetP2PProvideResponse) Kind() pd1.Kind {
	return pd1.Kind_Map
}

func (x GetP2PProvideResponse) LookupByString(key string) (pd1.Node, error) {
	switch key {
		case "ProvidersByKey":
		return x.ProvidersByKey.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x GetP2PProvideResponse) LookupByNode(key pd1.Node) (pd1.Node, error) {
	switch key.Kind() {
	case pd1.Kind_String:
		if s, err := key.AsString(); err != nil {
			return nil, err
		} else {
			return x.LookupByString(s)
		}
	case pd1.Kind_Int:
		if i, err := key.AsInt(); err != nil {
			return nil, err
		} else {
			return x.LookupByIndex(i)
		}
	}
	return nil, pd2.ErrNA
}

func (x GetP2PProvideResponse) LookupByIndex(idx int64) (pd1.Node, error) {
	switch idx {
		case 0:
		return x.ProvidersByKey.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x GetP2PProvideResponse) LookupBySegment(seg pd1.PathSegment) (pd1.Node, error) {
	switch seg.String() {
		case "0", "ProvidersByKey":
		return x.ProvidersByKey.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x GetP2PProvideResponse) MapIterator() pd1.MapIterator {
	return &GetP2PProvideResponse_MapIterator{-1, &x}
}

func (x GetP2PProvideResponse) ListIterator() pd1.ListIterator {
	return nil
}

func (x GetP2PProvideResponse) Length() int64 {
	return 1
}

func (x GetP2PProvideResponse) IsAbsent() bool {
	return false
}

func (x GetP2PProvideResponse) IsNull() bool {
	return false
}

func (x GetP2PProvideResponse) AsBool() (bool, error) {
	return false, pd2.ErrNA
}

func (x GetP2PProvideResponse) AsInt() (int64, error) {
	return 0, pd2.ErrNA
}

func (x GetP2PProvideResponse) AsFloat() (float64, error) {
	return 0, pd2.ErrNA
}

func (x GetP2PProvideResponse) AsString() (string, error) {
	return "", pd2.ErrNA
}

func (x GetP2PProvideResponse) AsBytes() ([]byte, error) {
	return nil, pd2.ErrNA
}

func (x GetP2PProvideResponse) AsLink() (pd1.Link, error) {
	return nil, pd2.ErrNA
}

func (x GetP2PProvideResponse) Prototype() pd1.NodePrototype {
	return nil
}

// -- protocol type ProvidersByKey --

type ProvidersByKey struct {
		Key Multihash
		Provider Provider

}

func (x ProvidersByKey) Node() pd1.Node {
	return x
}

func (x *ProvidersByKey) Parse(n pd1.Node) error {
	if n.Kind() != pd1.Kind_Map {
		return pd2.ErrNA
	}
	iter := n.MapIterator()
	nfields := 0
	for !iter.Done() {
		if kn, vn, err := iter.Next(); err != nil {
			return err
		} else {
			if k, err := kn.AsString(); err != nil {
				return pd3.Errorf("structure map key is not a string")
			} else {
				_ = vn
				switch k {
			case "Key":
			if err := x.Key.Parse(vn); err != nil {
				return err
			}
			nfields++
			case "Provider":
			if err := x.Provider.Parse(vn); err != nil {
				return err
			}
			nfields++

				}
			}
		}
	}
	if nfields != 2 {
		return pd2.ErrNA
	} else {
		return nil
	}
}

type ProvidersByKey_MapIterator struct {
	i int64
	s *ProvidersByKey
}

func (x *ProvidersByKey_MapIterator) Next() (key pd1.Node, value pd1.Node, err error) {
	x.i++
	switch x.i {
			case 0:
			return pd2.String("Key"), x.s.Key.Node(), nil
			case 1:
			return pd2.String("Provider"), x.s.Provider.Node(), nil

	}
	return nil, nil, pd2.ErrNA
}

func (x *ProvidersByKey_MapIterator) Done() bool {
	return x.i + 1 >= 2
}

func (x ProvidersByKey) Kind() pd1.Kind {
	return pd1.Kind_Map
}

func (x ProvidersByKey) LookupByString(key string) (pd1.Node, error) {
	switch key {
		case "Key":
		return x.Key.Node(), nil
		case "Provider":
		return x.Provider.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x ProvidersByKey) LookupByNode(key pd1.Node) (pd1.Node, error) {
	switch key.Kind() {
	case pd1.Kind_String:
		if s, err := key.AsString(); err != nil {
			return nil, err
		} else {
			return x.LookupByString(s)
		}
	case pd1.Kind_Int:
		if i, err := key.AsInt(); err != nil {
			return nil, err
		} else {
			return x.LookupByIndex(i)
		}
	}
	return nil, pd2.ErrNA
}

func (x ProvidersByKey) LookupByIndex(idx int64) (pd1.Node, error) {
	switch idx {
		case 0:
		return x.Key.Node(), nil
		case 1:
		return x.Provider.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x ProvidersByKey) LookupBySegment(seg pd1.PathSegment) (pd1.Node, error) {
	switch seg.String() {
		case "0", "Key":
		return x.Key.Node(), nil
		case "1", "Provider":
		return x.Provider.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x ProvidersByKey) MapIterator() pd1.MapIterator {
	return &ProvidersByKey_MapIterator{-1, &x}
}

func (x ProvidersByKey) ListIterator() pd1.ListIterator {
	return nil
}

func (x ProvidersByKey) Length() int64 {
	return 2
}

func (x ProvidersByKey) IsAbsent() bool {
	return false
}

func (x ProvidersByKey) IsNull() bool {
	return false
}

func (x ProvidersByKey) AsBool() (bool, error) {
	return false, pd2.ErrNA
}

func (x ProvidersByKey) AsInt() (int64, error) {
	return 0, pd2.ErrNA
}

func (x ProvidersByKey) AsFloat() (float64, error) {
	return 0, pd2.ErrNA
}

func (x ProvidersByKey) AsString() (string, error) {
	return "", pd2.ErrNA
}

func (x ProvidersByKey) AsBytes() ([]byte, error) {
	return nil, pd2.ErrNA
}

func (x ProvidersByKey) AsLink() (pd1.Link, error) {
	return nil, pd2.ErrNA
}

func (x ProvidersByKey) Prototype() pd1.NodePrototype {
	return nil
}

// -- protocol type Multihash --

type Multihash struct {
		Bytes pd2.Bytes

}

func (x Multihash) Node() pd1.Node {
	return x
}

func (x *Multihash) Parse(n pd1.Node) error {
	if n.Kind() != pd1.Kind_Map {
		return pd2.ErrNA
	}
	iter := n.MapIterator()
	nfields := 0
	for !iter.Done() {
		if kn, vn, err := iter.Next(); err != nil {
			return err
		} else {
			if k, err := kn.AsString(); err != nil {
				return pd3.Errorf("structure map key is not a string")
			} else {
				_ = vn
				switch k {
			case "Bytes":
			if err := x.Bytes.Parse(vn); err != nil {
				return err
			}
			nfields++

				}
			}
		}
	}
	if nfields != 1 {
		return pd2.ErrNA
	} else {
		return nil
	}
}

type Multihash_MapIterator struct {
	i int64
	s *Multihash
}

func (x *Multihash_MapIterator) Next() (key pd1.Node, value pd1.Node, err error) {
	x.i++
	switch x.i {
			case 0:
			return pd2.String("Bytes"), x.s.Bytes.Node(), nil

	}
	return nil, nil, pd2.ErrNA
}

func (x *Multihash_MapIterator) Done() bool {
	return x.i + 1 >= 1
}

func (x Multihash) Kind() pd1.Kind {
	return pd1.Kind_Map
}

func (x Multihash) LookupByString(key string) (pd1.Node, error) {
	switch key {
		case "Bytes":
		return x.Bytes.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x Multihash) LookupByNode(key pd1.Node) (pd1.Node, error) {
	switch key.Kind() {
	case pd1.Kind_String:
		if s, err := key.AsString(); err != nil {
			return nil, err
		} else {
			return x.LookupByString(s)
		}
	case pd1.Kind_Int:
		if i, err := key.AsInt(); err != nil {
			return nil, err
		} else {
			return x.LookupByIndex(i)
		}
	}
	return nil, pd2.ErrNA
}

func (x Multihash) LookupByIndex(idx int64) (pd1.Node, error) {
	switch idx {
		case 0:
		return x.Bytes.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x Multihash) LookupBySegment(seg pd1.PathSegment) (pd1.Node, error) {
	switch seg.String() {
		case "0", "Bytes":
		return x.Bytes.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x Multihash) MapIterator() pd1.MapIterator {
	return &Multihash_MapIterator{-1, &x}
}

func (x Multihash) ListIterator() pd1.ListIterator {
	return nil
}

func (x Multihash) Length() int64 {
	return 1
}

func (x Multihash) IsAbsent() bool {
	return false
}

func (x Multihash) IsNull() bool {
	return false
}

func (x Multihash) AsBool() (bool, error) {
	return false, pd2.ErrNA
}

func (x Multihash) AsInt() (int64, error) {
	return 0, pd2.ErrNA
}

func (x Multihash) AsFloat() (float64, error) {
	return 0, pd2.ErrNA
}

func (x Multihash) AsString() (string, error) {
	return "", pd2.ErrNA
}

func (x Multihash) AsBytes() ([]byte, error) {
	return nil, pd2.ErrNA
}

func (x Multihash) AsLink() (pd1.Link, error) {
	return nil, pd2.ErrNA
}

func (x Multihash) Prototype() pd1.NodePrototype {
	return nil
}

// -- protocol type AnonList9 --

type AnonList9 []Node

func (v AnonList9) Node() pd1.Node {
	return v
}

func (v *AnonList9) Parse(n pd1.Node) error {
	if n.Kind() != pd1.Kind_List {
		return pd2.ErrNA
	} else {
		*v = make(AnonList9, n.Length())
		iter := n.ListIterator()
		for !iter.Done() {
			if i, n, err := iter.Next(); err != nil {
				return pd2.ErrNA
			} else if err = (*v)[i].Parse(n); err != nil {
				return err
			}
		}
		return nil
	}
}

func (AnonList9) Kind() pd1.Kind {
	return pd1.Kind_List
}

func (AnonList9) LookupByString(string) (pd1.Node, error) {
	return nil, pd2.ErrNA
}

func (AnonList9) LookupByNode(key pd1.Node) (pd1.Node, error) {
	return nil, pd2.ErrNA
}

func (v AnonList9) LookupByIndex(i int64) (pd1.Node, error) {
	if i < 0 || i >= v.Length() {
		return nil, pd2.ErrBounds
	} else {
		return v[i].Node(), nil
	}
}

func (v AnonList9) LookupBySegment(seg pd1.PathSegment) (pd1.Node, error) {
	if i, err := seg.Index(); err != nil {
		return nil, pd2.ErrNA
	} else {
		return v.LookupByIndex(i)
	}
}

func (AnonList9) MapIterator() pd1.MapIterator {
	return nil
}

func (v AnonList9) ListIterator() pd1.ListIterator {
	return &AnonList9_ListIterator{v, 0}
}

func (v AnonList9) Length() int64 {
	return int64(len(v))
}

func (AnonList9) IsAbsent() bool {
	return false
}

func (AnonList9) IsNull() bool {
	return false
}

func (v AnonList9) AsBool() (bool, error) {
	return false, pd2.ErrNA
}

func (AnonList9) AsInt() (int64, error) {
	return 0, pd2.ErrNA
}

func (AnonList9) AsFloat() (float64, error) {
	return 0, pd2.ErrNA
}

func (AnonList9) AsString() (string, error) {
	return "", pd2.ErrNA
}

func (AnonList9) AsBytes() ([]byte, error) {
	return nil, pd2.ErrNA
}

func (AnonList9) AsLink() (pd1.Link, error) {
	return nil, pd2.ErrNA
}

func (AnonList9) Prototype() pd1.NodePrototype {
	return nil // not needed
}

type AnonList9_ListIterator struct {
	list AnonList9
	at   int64
}

func (iter *AnonList9_ListIterator) Next() (int64, pd1.Node, error) {
	if iter.Done() {
		return -1, nil, pd2.ErrBounds
	}
	v := iter.list[iter.at]
	i := int64(iter.at)
	iter.at++
	return i, v.Node(), nil
}

func (iter *AnonList9_ListIterator) Done() bool {
	return iter.at >= iter.list.Length()
}
// -- protocol type AnonList10 --

type AnonList10 []TransferProto

func (v AnonList10) Node() pd1.Node {
	return v
}

func (v *AnonList10) Parse(n pd1.Node) error {
	if n.Kind() != pd1.Kind_List {
		return pd2.ErrNA
	} else {
		*v = make(AnonList10, n.Length())
		iter := n.ListIterator()
		for !iter.Done() {
			if i, n, err := iter.Next(); err != nil {
				return pd2.ErrNA
			} else if err = (*v)[i].Parse(n); err != nil {
				return err
			}
		}
		return nil
	}
}

func (AnonList10) Kind() pd1.Kind {
	return pd1.Kind_List
}

func (AnonList10) LookupByString(string) (pd1.Node, error) {
	return nil, pd2.ErrNA
}

func (AnonList10) LookupByNode(key pd1.Node) (pd1.Node, error) {
	return nil, pd2.ErrNA
}

func (v AnonList10) LookupByIndex(i int64) (pd1.Node, error) {
	if i < 0 || i >= v.Length() {
		return nil, pd2.ErrBounds
	} else {
		return v[i].Node(), nil
	}
}

func (v AnonList10) LookupBySegment(seg pd1.PathSegment) (pd1.Node, error) {
	if i, err := seg.Index(); err != nil {
		return nil, pd2.ErrNA
	} else {
		return v.LookupByIndex(i)
	}
}

func (AnonList10) MapIterator() pd1.MapIterator {
	return nil
}

func (v AnonList10) ListIterator() pd1.ListIterator {
	return &AnonList10_ListIterator{v, 0}
}

func (v AnonList10) Length() int64 {
	return int64(len(v))
}

func (AnonList10) IsAbsent() bool {
	return false
}

func (AnonList10) IsNull() bool {
	return false
}

func (v AnonList10) AsBool() (bool, error) {
	return false, pd2.ErrNA
}

func (AnonList10) AsInt() (int64, error) {
	return 0, pd2.ErrNA
}

func (AnonList10) AsFloat() (float64, error) {
	return 0, pd2.ErrNA
}

func (AnonList10) AsString() (string, error) {
	return "", pd2.ErrNA
}

func (AnonList10) AsBytes() ([]byte, error) {
	return nil, pd2.ErrNA
}

func (AnonList10) AsLink() (pd1.Link, error) {
	return nil, pd2.ErrNA
}

func (AnonList10) Prototype() pd1.NodePrototype {
	return nil // not needed
}

type AnonList10_ListIterator struct {
	list AnonList10
	at   int64
}

func (iter *AnonList10_ListIterator) Next() (int64, pd1.Node, error) {
	if iter.Done() {
		return -1, nil, pd2.ErrBounds
	}
	v := iter.list[iter.at]
	i := int64(iter.at)
	iter.at++
	return i, v.Node(), nil
}

func (iter *AnonList10_ListIterator) Done() bool {
	return iter.at >= iter.list.Length()
}
// -- protocol type Provider --

type Provider struct {
		Nodes AnonList9
		Proto AnonList10

}

func (x Provider) Node() pd1.Node {
	return x
}

func (x *Provider) Parse(n pd1.Node) error {
	if n.Kind() != pd1.Kind_Map {
		return pd2.ErrNA
	}
	iter := n.MapIterator()
	nfields := 0
	for !iter.Done() {
		if kn, vn, err := iter.Next(); err != nil {
			return err
		} else {
			if k, err := kn.AsString(); err != nil {
				return pd3.Errorf("structure map key is not a string")
			} else {
				_ = vn
				switch k {
			case "Nodes":
			if err := x.Nodes.Parse(vn); err != nil {
				return err
			}
			nfields++
			case "Proto":
			if err := x.Proto.Parse(vn); err != nil {
				return err
			}
			nfields++

				}
			}
		}
	}
	if nfields != 2 {
		return pd2.ErrNA
	} else {
		return nil
	}
}

type Provider_MapIterator struct {
	i int64
	s *Provider
}

func (x *Provider_MapIterator) Next() (key pd1.Node, value pd1.Node, err error) {
	x.i++
	switch x.i {
			case 0:
			return pd2.String("Nodes"), x.s.Nodes.Node(), nil
			case 1:
			return pd2.String("Proto"), x.s.Proto.Node(), nil

	}
	return nil, nil, pd2.ErrNA
}

func (x *Provider_MapIterator) Done() bool {
	return x.i + 1 >= 2
}

func (x Provider) Kind() pd1.Kind {
	return pd1.Kind_Map
}

func (x Provider) LookupByString(key string) (pd1.Node, error) {
	switch key {
		case "Nodes":
		return x.Nodes.Node(), nil
		case "Proto":
		return x.Proto.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x Provider) LookupByNode(key pd1.Node) (pd1.Node, error) {
	switch key.Kind() {
	case pd1.Kind_String:
		if s, err := key.AsString(); err != nil {
			return nil, err
		} else {
			return x.LookupByString(s)
		}
	case pd1.Kind_Int:
		if i, err := key.AsInt(); err != nil {
			return nil, err
		} else {
			return x.LookupByIndex(i)
		}
	}
	return nil, pd2.ErrNA
}

func (x Provider) LookupByIndex(idx int64) (pd1.Node, error) {
	switch idx {
		case 0:
		return x.Nodes.Node(), nil
		case 1:
		return x.Proto.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x Provider) LookupBySegment(seg pd1.PathSegment) (pd1.Node, error) {
	switch seg.String() {
		case "0", "Nodes":
		return x.Nodes.Node(), nil
		case "1", "Proto":
		return x.Proto.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x Provider) MapIterator() pd1.MapIterator {
	return &Provider_MapIterator{-1, &x}
}

func (x Provider) ListIterator() pd1.ListIterator {
	return nil
}

func (x Provider) Length() int64 {
	return 2
}

func (x Provider) IsAbsent() bool {
	return false
}

func (x Provider) IsNull() bool {
	return false
}

func (x Provider) AsBool() (bool, error) {
	return false, pd2.ErrNA
}

func (x Provider) AsInt() (int64, error) {
	return 0, pd2.ErrNA
}

func (x Provider) AsFloat() (float64, error) {
	return 0, pd2.ErrNA
}

func (x Provider) AsString() (string, error) {
	return "", pd2.ErrNA
}

func (x Provider) AsBytes() ([]byte, error) {
	return nil, pd2.ErrNA
}

func (x Provider) AsLink() (pd1.Link, error) {
	return nil, pd2.ErrNA
}

func (x Provider) Prototype() pd1.NodePrototype {
	return nil
}

// -- protocol type Node --

type Node struct {
		Peer *Peer

}

func (x *Node) Parse(n pd1.Node) error {
	*x = Node{}
	if n.Kind() != pd1.Kind_Map {
		return pd2.ErrNA
	}
	iter := n.MapIterator()
	kn, vn, err := iter.Next()
	if err != nil {
		return err
	}
	k, err := kn.AsString()
	if err != nil {
		return pd3.Errorf("inductive map key is not a string")
	}
	switch k {
	case "Peer":
		var y Peer
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.Peer = &y
		return nil

	}
	return pd3.Errorf("inductive map has no applicable keys")
}

type Node_MapIterator struct {
	done bool
	s    *Node
}

func (x *Node_MapIterator) Next() (key pd1.Node, value pd1.Node, err error) {
	if x.done {
		return nil, nil, pd2.ErrNA
	} else {
		x.done = true
		switch {
			case x.s.Peer != nil:
			return pd2.String("Peer"), x.s.Peer.Node(), nil

		default:
			return nil, nil, pd3.Errorf("no inductive cases are set")
		}
	}
}

func (x *Node_MapIterator) Done() bool {
	return x.done
}

func (x Node) Node() pd1.Node {
	return x
}

func (x Node) Kind() pd1.Kind {
	return pd1.Kind_Map
}

func (x Node) LookupByString(key string) (pd1.Node, error) {
	switch {
		case x.Peer != nil && key == "Peer":
		return x.Peer.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x Node) LookupByNode(key pd1.Node) (pd1.Node, error) {
	if key.Kind() != pd1.Kind_String {
		return nil, pd2.ErrNA
	}
	if s, err := key.AsString(); err != nil {
		return nil, err
	} else {
		return x.LookupByString(s)
	}
}

func (x Node) LookupByIndex(idx int64) (pd1.Node, error) {
	return nil, pd2.ErrNA
}

func (x Node) LookupBySegment(seg pd1.PathSegment) (pd1.Node, error) {
	switch seg.String() {
		case "Peer":
		return x.Peer.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x Node) MapIterator() pd1.MapIterator {
	return &Node_MapIterator{false, &x}
}

func (x Node) ListIterator() pd1.ListIterator {
	return nil
}

func (x Node) Length() int64 {
	return 1
}

func (x Node) IsAbsent() bool {
	return false
}

func (x Node) IsNull() bool {
	return false
}

func (x Node) AsBool() (bool, error) {
	return false, pd2.ErrNA
}

func (x Node) AsInt() (int64, error) {
	return 0, pd2.ErrNA
}

func (x Node) AsFloat() (float64, error) {
	return 0, pd2.ErrNA
}

func (x Node) AsString() (string, error) {
	return "", pd2.ErrNA
}

func (x Node) AsBytes() ([]byte, error) {
	return nil, pd2.ErrNA
}

func (x Node) AsLink() (pd1.Link, error) {
	return nil, pd2.ErrNA
}

func (x Node) Prototype() pd1.NodePrototype {
	return nil
}
// -- protocol type AnonList13 --

type AnonList13 []pd2.Bytes

func (v AnonList13) Node() pd1.Node {
	return v
}

func (v *AnonList13) Parse(n pd1.Node) error {
	if n.Kind() != pd1.Kind_List {
		return pd2.ErrNA
	} else {
		*v = make(AnonList13, n.Length())
		iter := n.ListIterator()
		for !iter.Done() {
			if i, n, err := iter.Next(); err != nil {
				return pd2.ErrNA
			} else if err = (*v)[i].Parse(n); err != nil {
				return err
			}
		}
		return nil
	}
}

func (AnonList13) Kind() pd1.Kind {
	return pd1.Kind_List
}

func (AnonList13) LookupByString(string) (pd1.Node, error) {
	return nil, pd2.ErrNA
}

func (AnonList13) LookupByNode(key pd1.Node) (pd1.Node, error) {
	return nil, pd2.ErrNA
}

func (v AnonList13) LookupByIndex(i int64) (pd1.Node, error) {
	if i < 0 || i >= v.Length() {
		return nil, pd2.ErrBounds
	} else {
		return v[i].Node(), nil
	}
}

func (v AnonList13) LookupBySegment(seg pd1.PathSegment) (pd1.Node, error) {
	if i, err := seg.Index(); err != nil {
		return nil, pd2.ErrNA
	} else {
		return v.LookupByIndex(i)
	}
}

func (AnonList13) MapIterator() pd1.MapIterator {
	return nil
}

func (v AnonList13) ListIterator() pd1.ListIterator {
	return &AnonList13_ListIterator{v, 0}
}

func (v AnonList13) Length() int64 {
	return int64(len(v))
}

func (AnonList13) IsAbsent() bool {
	return false
}

func (AnonList13) IsNull() bool {
	return false
}

func (v AnonList13) AsBool() (bool, error) {
	return false, pd2.ErrNA
}

func (AnonList13) AsInt() (int64, error) {
	return 0, pd2.ErrNA
}

func (AnonList13) AsFloat() (float64, error) {
	return 0, pd2.ErrNA
}

func (AnonList13) AsString() (string, error) {
	return "", pd2.ErrNA
}

func (AnonList13) AsBytes() ([]byte, error) {
	return nil, pd2.ErrNA
}

func (AnonList13) AsLink() (pd1.Link, error) {
	return nil, pd2.ErrNA
}

func (AnonList13) Prototype() pd1.NodePrototype {
	return nil // not needed
}

type AnonList13_ListIterator struct {
	list AnonList13
	at   int64
}

func (iter *AnonList13_ListIterator) Next() (int64, pd1.Node, error) {
	if iter.Done() {
		return -1, nil, pd2.ErrBounds
	}
	v := iter.list[iter.at]
	i := int64(iter.at)
	iter.at++
	return i, v.Node(), nil
}

func (iter *AnonList13_ListIterator) Done() bool {
	return iter.at >= iter.list.Length()
}
// -- protocol type Peer --

type Peer struct {
		ID pd2.Bytes
		Multiaddresses AnonList13

}

func (x Peer) Node() pd1.Node {
	return x
}

func (x *Peer) Parse(n pd1.Node) error {
	if n.Kind() != pd1.Kind_Map {
		return pd2.ErrNA
	}
	iter := n.MapIterator()
	nfields := 0
	for !iter.Done() {
		if kn, vn, err := iter.Next(); err != nil {
			return err
		} else {
			if k, err := kn.AsString(); err != nil {
				return pd3.Errorf("structure map key is not a string")
			} else {
				_ = vn
				switch k {
			case "ID":
			if err := x.ID.Parse(vn); err != nil {
				return err
			}
			nfields++
			case "Multiaddresses":
			if err := x.Multiaddresses.Parse(vn); err != nil {
				return err
			}
			nfields++

				}
			}
		}
	}
	if nfields != 2 {
		return pd2.ErrNA
	} else {
		return nil
	}
}

type Peer_MapIterator struct {
	i int64
	s *Peer
}

func (x *Peer_MapIterator) Next() (key pd1.Node, value pd1.Node, err error) {
	x.i++
	switch x.i {
			case 0:
			return pd2.String("ID"), x.s.ID.Node(), nil
			case 1:
			return pd2.String("Multiaddresses"), x.s.Multiaddresses.Node(), nil

	}
	return nil, nil, pd2.ErrNA
}

func (x *Peer_MapIterator) Done() bool {
	return x.i + 1 >= 2
}

func (x Peer) Kind() pd1.Kind {
	return pd1.Kind_Map
}

func (x Peer) LookupByString(key string) (pd1.Node, error) {
	switch key {
		case "ID":
		return x.ID.Node(), nil
		case "Multiaddresses":
		return x.Multiaddresses.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x Peer) LookupByNode(key pd1.Node) (pd1.Node, error) {
	switch key.Kind() {
	case pd1.Kind_String:
		if s, err := key.AsString(); err != nil {
			return nil, err
		} else {
			return x.LookupByString(s)
		}
	case pd1.Kind_Int:
		if i, err := key.AsInt(); err != nil {
			return nil, err
		} else {
			return x.LookupByIndex(i)
		}
	}
	return nil, pd2.ErrNA
}

func (x Peer) LookupByIndex(idx int64) (pd1.Node, error) {
	switch idx {
		case 0:
		return x.ID.Node(), nil
		case 1:
		return x.Multiaddresses.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x Peer) LookupBySegment(seg pd1.PathSegment) (pd1.Node, error) {
	switch seg.String() {
		case "0", "ID":
		return x.ID.Node(), nil
		case "1", "Multiaddresses":
		return x.Multiaddresses.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x Peer) MapIterator() pd1.MapIterator {
	return &Peer_MapIterator{-1, &x}
}

func (x Peer) ListIterator() pd1.ListIterator {
	return nil
}

func (x Peer) Length() int64 {
	return 2
}

func (x Peer) IsAbsent() bool {
	return false
}

func (x Peer) IsNull() bool {
	return false
}

func (x Peer) AsBool() (bool, error) {
	return false, pd2.ErrNA
}

func (x Peer) AsInt() (int64, error) {
	return 0, pd2.ErrNA
}

func (x Peer) AsFloat() (float64, error) {
	return 0, pd2.ErrNA
}

func (x Peer) AsString() (string, error) {
	return "", pd2.ErrNA
}

func (x Peer) AsBytes() ([]byte, error) {
	return nil, pd2.ErrNA
}

func (x Peer) AsLink() (pd1.Link, error) {
	return nil, pd2.ErrNA
}

func (x Peer) Prototype() pd1.NodePrototype {
	return nil
}

// -- protocol type TransferProto --

type TransferProto struct {
		Bitswap *BitswapTransfer

}

func (x *TransferProto) Parse(n pd1.Node) error {
	*x = TransferProto{}
	if n.Kind() != pd1.Kind_Map {
		return pd2.ErrNA
	}
	iter := n.MapIterator()
	kn, vn, err := iter.Next()
	if err != nil {
		return err
	}
	k, err := kn.AsString()
	if err != nil {
		return pd3.Errorf("inductive map key is not a string")
	}
	switch k {
	case "Bitswap":
		var y BitswapTransfer
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.Bitswap = &y
		return nil

	}
	return pd3.Errorf("inductive map has no applicable keys")
}

type TransferProto_MapIterator struct {
	done bool
	s    *TransferProto
}

func (x *TransferProto_MapIterator) Next() (key pd1.Node, value pd1.Node, err error) {
	if x.done {
		return nil, nil, pd2.ErrNA
	} else {
		x.done = true
		switch {
			case x.s.Bitswap != nil:
			return pd2.String("Bitswap"), x.s.Bitswap.Node(), nil

		default:
			return nil, nil, pd3.Errorf("no inductive cases are set")
		}
	}
}

func (x *TransferProto_MapIterator) Done() bool {
	return x.done
}

func (x TransferProto) Node() pd1.Node {
	return x
}

func (x TransferProto) Kind() pd1.Kind {
	return pd1.Kind_Map
}

func (x TransferProto) LookupByString(key string) (pd1.Node, error) {
	switch {
		case x.Bitswap != nil && key == "Bitswap":
		return x.Bitswap.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x TransferProto) LookupByNode(key pd1.Node) (pd1.Node, error) {
	if key.Kind() != pd1.Kind_String {
		return nil, pd2.ErrNA
	}
	if s, err := key.AsString(); err != nil {
		return nil, err
	} else {
		return x.LookupByString(s)
	}
}

func (x TransferProto) LookupByIndex(idx int64) (pd1.Node, error) {
	return nil, pd2.ErrNA
}

func (x TransferProto) LookupBySegment(seg pd1.PathSegment) (pd1.Node, error) {
	switch seg.String() {
		case "Bitswap":
		return x.Bitswap.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x TransferProto) MapIterator() pd1.MapIterator {
	return &TransferProto_MapIterator{false, &x}
}

func (x TransferProto) ListIterator() pd1.ListIterator {
	return nil
}

func (x TransferProto) Length() int64 {
	return 1
}

func (x TransferProto) IsAbsent() bool {
	return false
}

func (x TransferProto) IsNull() bool {
	return false
}

func (x TransferProto) AsBool() (bool, error) {
	return false, pd2.ErrNA
}

func (x TransferProto) AsInt() (int64, error) {
	return 0, pd2.ErrNA
}

func (x TransferProto) AsFloat() (float64, error) {
	return 0, pd2.ErrNA
}

func (x TransferProto) AsString() (string, error) {
	return "", pd2.ErrNA
}

func (x TransferProto) AsBytes() ([]byte, error) {
	return nil, pd2.ErrNA
}

func (x TransferProto) AsLink() (pd1.Link, error) {
	return nil, pd2.ErrNA
}

func (x TransferProto) Prototype() pd1.NodePrototype {
	return nil
}
// -- protocol type BitswapTransfer --

type BitswapTransfer struct {

}

func (x BitswapTransfer) Node() pd1.Node {
	return x
}

func (x *BitswapTransfer) Parse(n pd1.Node) error {
	if n.Kind() != pd1.Kind_Map {
		return pd2.ErrNA
	}
	iter := n.MapIterator()
	nfields := 0
	for !iter.Done() {
		if kn, vn, err := iter.Next(); err != nil {
			return err
		} else {
			if k, err := kn.AsString(); err != nil {
				return pd3.Errorf("structure map key is not a string")
			} else {
				_ = vn
				switch k {

				}
			}
		}
	}
	if nfields != 0 {
		return pd2.ErrNA
	} else {
		return nil
	}
}

type BitswapTransfer_MapIterator struct {
	i int64
	s *BitswapTransfer
}

func (x *BitswapTransfer_MapIterator) Next() (key pd1.Node, value pd1.Node, err error) {
	x.i++
	switch x.i {

	}
	return nil, nil, pd2.ErrNA
}

func (x *BitswapTransfer_MapIterator) Done() bool {
	return x.i + 1 >= 0
}

func (x BitswapTransfer) Kind() pd1.Kind {
	return pd1.Kind_Map
}

func (x BitswapTransfer) LookupByString(key string) (pd1.Node, error) {
	switch key {

	}
	return nil, pd2.ErrNA
}

func (x BitswapTransfer) LookupByNode(key pd1.Node) (pd1.Node, error) {
	switch key.Kind() {
	case pd1.Kind_String:
		if s, err := key.AsString(); err != nil {
			return nil, err
		} else {
			return x.LookupByString(s)
		}
	case pd1.Kind_Int:
		if i, err := key.AsInt(); err != nil {
			return nil, err
		} else {
			return x.LookupByIndex(i)
		}
	}
	return nil, pd2.ErrNA
}

func (x BitswapTransfer) LookupByIndex(idx int64) (pd1.Node, error) {
	switch idx {

	}
	return nil, pd2.ErrNA
}

func (x BitswapTransfer) LookupBySegment(seg pd1.PathSegment) (pd1.Node, error) {
	switch seg.String() {

	}
	return nil, pd2.ErrNA
}

func (x BitswapTransfer) MapIterator() pd1.MapIterator {
	return &BitswapTransfer_MapIterator{-1, &x}
}

func (x BitswapTransfer) ListIterator() pd1.ListIterator {
	return nil
}

func (x BitswapTransfer) Length() int64 {
	return 0
}

func (x BitswapTransfer) IsAbsent() bool {
	return false
}

func (x BitswapTransfer) IsNull() bool {
	return false
}

func (x BitswapTransfer) AsBool() (bool, error) {
	return false, pd2.ErrNA
}

func (x BitswapTransfer) AsInt() (int64, error) {
	return 0, pd2.ErrNA
}

func (x BitswapTransfer) AsFloat() (float64, error) {
	return 0, pd2.ErrNA
}

func (x BitswapTransfer) AsString() (string, error) {
	return "", pd2.ErrNA
}

func (x BitswapTransfer) AsBytes() ([]byte, error) {
	return nil, pd2.ErrNA
}

func (x BitswapTransfer) AsLink() (pd1.Link, error) {
	return nil, pd2.ErrNA
}

func (x BitswapTransfer) Prototype() pd1.NodePrototype {
	return nil
}
