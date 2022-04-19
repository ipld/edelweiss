package proto

import(
	pd1 "github.com/ipld/go-ipld-prime/datamodel"
	pd2 "github.com/ipld/edelweiss/values"
	pd3 "fmt"
	pd4 "net/http"
	pd5 "net/url"
	pd6 "github.com/ipfs/go-log"
	pd7 "context"
	pd8 "github.com/ipld/go-ipld-prime"
	pd9 "io"
	pd10 "bytes"
	pd11 "github.com/ipld/go-ipld-prime/codec/dagjson"
	pd12 "errors"
	pd13 "github.com/ipld/edelweiss/services"
)


// -- protocol type Greeting_IdentifyArg --

type Greeting_IdentifyArg struct {

}

func (x Greeting_IdentifyArg) Node() pd1.Node {
	return x
}

func (x *Greeting_IdentifyArg) Parse(n pd1.Node) error {
	if n.Kind() != pd1.Kind_Map {
		return pd2.ErrNA
	}
	iter := n.MapIterator()
	fieldMap := map[string]pd2.ParseFunc{
		
	}
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
	for _, fieldParse := range fieldMap {
		if err := fieldParse(pd1.Null); err != nil {
			return err
		}
	}
	return nil
}

type Greeting_IdentifyArg_MapIterator struct {
	i int64
	s *Greeting_IdentifyArg
}

func (x *Greeting_IdentifyArg_MapIterator) Next() (key pd1.Node, value pd1.Node, err error) {
	x.i++
	switch x.i {

	}
	return nil, nil, pd2.ErrNA
}

func (x *Greeting_IdentifyArg_MapIterator) Done() bool {
	return x.i + 1 >= 0
}

func (x Greeting_IdentifyArg) Kind() pd1.Kind {
	return pd1.Kind_Map
}

func (x Greeting_IdentifyArg) LookupByString(key string) (pd1.Node, error) {
	switch key {

	}
	return nil, pd2.ErrNA
}

func (x Greeting_IdentifyArg) LookupByNode(key pd1.Node) (pd1.Node, error) {
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

func (x Greeting_IdentifyArg) LookupByIndex(idx int64) (pd1.Node, error) {
	switch idx {

	}
	return nil, pd2.ErrNA
}

func (x Greeting_IdentifyArg) LookupBySegment(seg pd1.PathSegment) (pd1.Node, error) {
	switch seg.String() {

	}
	return nil, pd2.ErrNA
}

func (x Greeting_IdentifyArg) MapIterator() pd1.MapIterator {
	return &Greeting_IdentifyArg_MapIterator{-1, &x}
}

func (x Greeting_IdentifyArg) ListIterator() pd1.ListIterator {
	return nil
}

func (x Greeting_IdentifyArg) Length() int64 {
	return 0
}

func (x Greeting_IdentifyArg) IsAbsent() bool {
	return false
}

func (x Greeting_IdentifyArg) IsNull() bool {
	return false
}

func (x Greeting_IdentifyArg) AsBool() (bool, error) {
	return false, pd2.ErrNA
}

func (x Greeting_IdentifyArg) AsInt() (int64, error) {
	return 0, pd2.ErrNA
}

func (x Greeting_IdentifyArg) AsFloat() (float64, error) {
	return 0, pd2.ErrNA
}

func (x Greeting_IdentifyArg) AsString() (string, error) {
	return "", pd2.ErrNA
}

func (x Greeting_IdentifyArg) AsBytes() ([]byte, error) {
	return nil, pd2.ErrNA
}

func (x Greeting_IdentifyArg) AsLink() (pd1.Link, error) {
	return nil, pd2.ErrNA
}

func (x Greeting_IdentifyArg) Prototype() pd1.NodePrototype {
	return nil
}

// -- protocol type AnonList1 --

type AnonList1 []pd2.String

func (v AnonList1) Node() pd1.Node {
	return v
}

func (v *AnonList1) Parse(n pd1.Node) error {
	if n.Kind() == pd1.Kind_Null {
		*v = nil
		return nil
	}
	if n.Kind() != pd1.Kind_List {
		return pd2.ErrNA
	} else {
		*v = make(AnonList1, n.Length())
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

func (AnonList1) Kind() pd1.Kind {
	return pd1.Kind_List
}

func (AnonList1) LookupByString(string) (pd1.Node, error) {
	return nil, pd2.ErrNA
}

func (AnonList1) LookupByNode(key pd1.Node) (pd1.Node, error) {
	return nil, pd2.ErrNA
}

func (v AnonList1) LookupByIndex(i int64) (pd1.Node, error) {
	if i < 0 || i >= v.Length() {
		return nil, pd2.ErrBounds
	} else {
		return v[i].Node(), nil
	}
}

func (v AnonList1) LookupBySegment(seg pd1.PathSegment) (pd1.Node, error) {
	if i, err := seg.Index(); err != nil {
		return nil, pd2.ErrNA
	} else {
		return v.LookupByIndex(i)
	}
}

func (AnonList1) MapIterator() pd1.MapIterator {
	return nil
}

func (v AnonList1) ListIterator() pd1.ListIterator {
	return &AnonList1_ListIterator{v, 0}
}

func (v AnonList1) Length() int64 {
	return int64(len(v))
}

func (AnonList1) IsAbsent() bool {
	return false
}

func (AnonList1) IsNull() bool {
	return false
}

func (v AnonList1) AsBool() (bool, error) {
	return false, pd2.ErrNA
}

func (AnonList1) AsInt() (int64, error) {
	return 0, pd2.ErrNA
}

func (AnonList1) AsFloat() (float64, error) {
	return 0, pd2.ErrNA
}

func (AnonList1) AsString() (string, error) {
	return "", pd2.ErrNA
}

func (AnonList1) AsBytes() ([]byte, error) {
	return nil, pd2.ErrNA
}

func (AnonList1) AsLink() (pd1.Link, error) {
	return nil, pd2.ErrNA
}

func (AnonList1) Prototype() pd1.NodePrototype {
	return nil // not needed
}

type AnonList1_ListIterator struct {
	list AnonList1
	at   int64
}

func (iter *AnonList1_ListIterator) Next() (int64, pd1.Node, error) {
	if iter.Done() {
		return -1, nil, pd2.ErrBounds
	}
	v := iter.list[iter.at]
	i := int64(iter.at)
	iter.at++
	return i, v.Node(), nil
}

func (iter *AnonList1_ListIterator) Done() bool {
	return iter.at >= iter.list.Length()
}
// -- protocol type Greeting_IdentifyResult --

type Greeting_IdentifyResult struct {
		Methods AnonList1

}

func (x Greeting_IdentifyResult) Node() pd1.Node {
	return x
}

func (x *Greeting_IdentifyResult) Parse(n pd1.Node) error {
	if n.Kind() != pd1.Kind_Map {
		return pd2.ErrNA
	}
	iter := n.MapIterator()
	fieldMap := map[string]pd2.ParseFunc{
				"Methods": x.Methods.Parse,

	}
	for !iter.Done() {
		if kn, vn, err := iter.Next(); err != nil {
			return err
		} else {
			if k, err := kn.AsString(); err != nil {
				return pd3.Errorf("structure map key is not a string")
			} else {
				_ = vn
				switch k {
			case "Methods":
			if _, notParsed := fieldMap["Methods"]; !notParsed {
				return pd3.Errorf("field %s already parsed", "Methods")
			}
			if err := x.Methods.Parse(vn); err != nil {
				return err
			}
			delete(fieldMap, "Methods")

				}
			}
		}
	}
	for _, fieldParse := range fieldMap {
		if err := fieldParse(pd1.Null); err != nil {
			return err
		}
	}
	return nil
}

type Greeting_IdentifyResult_MapIterator struct {
	i int64
	s *Greeting_IdentifyResult
}

func (x *Greeting_IdentifyResult_MapIterator) Next() (key pd1.Node, value pd1.Node, err error) {
	x.i++
	switch x.i {
			case 0:
			return pd2.String("Methods"), x.s.Methods.Node(), nil

	}
	return nil, nil, pd2.ErrNA
}

func (x *Greeting_IdentifyResult_MapIterator) Done() bool {
	return x.i + 1 >= 1
}

func (x Greeting_IdentifyResult) Kind() pd1.Kind {
	return pd1.Kind_Map
}

func (x Greeting_IdentifyResult) LookupByString(key string) (pd1.Node, error) {
	switch key {
		case "Methods":
		return x.Methods.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x Greeting_IdentifyResult) LookupByNode(key pd1.Node) (pd1.Node, error) {
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

func (x Greeting_IdentifyResult) LookupByIndex(idx int64) (pd1.Node, error) {
	switch idx {
		case 0:
		return x.Methods.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x Greeting_IdentifyResult) LookupBySegment(seg pd1.PathSegment) (pd1.Node, error) {
	switch seg.String() {
		case "0", "Methods":
		return x.Methods.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x Greeting_IdentifyResult) MapIterator() pd1.MapIterator {
	return &Greeting_IdentifyResult_MapIterator{-1, &x}
}

func (x Greeting_IdentifyResult) ListIterator() pd1.ListIterator {
	return nil
}

func (x Greeting_IdentifyResult) Length() int64 {
	return 1
}

func (x Greeting_IdentifyResult) IsAbsent() bool {
	return false
}

func (x Greeting_IdentifyResult) IsNull() bool {
	return false
}

func (x Greeting_IdentifyResult) AsBool() (bool, error) {
	return false, pd2.ErrNA
}

func (x Greeting_IdentifyResult) AsInt() (int64, error) {
	return 0, pd2.ErrNA
}

func (x Greeting_IdentifyResult) AsFloat() (float64, error) {
	return 0, pd2.ErrNA
}

func (x Greeting_IdentifyResult) AsString() (string, error) {
	return "", pd2.ErrNA
}

func (x Greeting_IdentifyResult) AsBytes() ([]byte, error) {
	return nil, pd2.ErrNA
}

func (x Greeting_IdentifyResult) AsLink() (pd1.Link, error) {
	return nil, pd2.ErrNA
}

func (x Greeting_IdentifyResult) Prototype() pd1.NodePrototype {
	return nil
}

// -- protocol type Greeting_Error --

type Greeting_Error struct {
		Code pd2.String

}

func (x Greeting_Error) Node() pd1.Node {
	return x
}

func (x *Greeting_Error) Parse(n pd1.Node) error {
	if n.Kind() != pd1.Kind_Map {
		return pd2.ErrNA
	}
	iter := n.MapIterator()
	fieldMap := map[string]pd2.ParseFunc{
				"Code": x.Code.Parse,

	}
	for !iter.Done() {
		if kn, vn, err := iter.Next(); err != nil {
			return err
		} else {
			if k, err := kn.AsString(); err != nil {
				return pd3.Errorf("structure map key is not a string")
			} else {
				_ = vn
				switch k {
			case "Code":
			if _, notParsed := fieldMap["Code"]; !notParsed {
				return pd3.Errorf("field %s already parsed", "Code")
			}
			if err := x.Code.Parse(vn); err != nil {
				return err
			}
			delete(fieldMap, "Code")

				}
			}
		}
	}
	for _, fieldParse := range fieldMap {
		if err := fieldParse(pd1.Null); err != nil {
			return err
		}
	}
	return nil
}

type Greeting_Error_MapIterator struct {
	i int64
	s *Greeting_Error
}

func (x *Greeting_Error_MapIterator) Next() (key pd1.Node, value pd1.Node, err error) {
	x.i++
	switch x.i {
			case 0:
			return pd2.String("Code"), x.s.Code.Node(), nil

	}
	return nil, nil, pd2.ErrNA
}

func (x *Greeting_Error_MapIterator) Done() bool {
	return x.i + 1 >= 1
}

func (x Greeting_Error) Kind() pd1.Kind {
	return pd1.Kind_Map
}

func (x Greeting_Error) LookupByString(key string) (pd1.Node, error) {
	switch key {
		case "Code":
		return x.Code.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x Greeting_Error) LookupByNode(key pd1.Node) (pd1.Node, error) {
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

func (x Greeting_Error) LookupByIndex(idx int64) (pd1.Node, error) {
	switch idx {
		case 0:
		return x.Code.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x Greeting_Error) LookupBySegment(seg pd1.PathSegment) (pd1.Node, error) {
	switch seg.String() {
		case "0", "Code":
		return x.Code.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x Greeting_Error) MapIterator() pd1.MapIterator {
	return &Greeting_Error_MapIterator{-1, &x}
}

func (x Greeting_Error) ListIterator() pd1.ListIterator {
	return nil
}

func (x Greeting_Error) Length() int64 {
	return 1
}

func (x Greeting_Error) IsAbsent() bool {
	return false
}

func (x Greeting_Error) IsNull() bool {
	return false
}

func (x Greeting_Error) AsBool() (bool, error) {
	return false, pd2.ErrNA
}

func (x Greeting_Error) AsInt() (int64, error) {
	return 0, pd2.ErrNA
}

func (x Greeting_Error) AsFloat() (float64, error) {
	return 0, pd2.ErrNA
}

func (x Greeting_Error) AsString() (string, error) {
	return "", pd2.ErrNA
}

func (x Greeting_Error) AsBytes() ([]byte, error) {
	return nil, pd2.ErrNA
}

func (x Greeting_Error) AsLink() (pd1.Link, error) {
	return nil, pd2.ErrNA
}

func (x Greeting_Error) Prototype() pd1.NodePrototype {
	return nil
}

// -- protocol type AnonInductive4 --

type AnonInductive4 struct {
		Identify *Greeting_IdentifyArg
		Hello *HelloRequest


}

func (x *AnonInductive4) Parse(n pd1.Node) error {
	*x = AnonInductive4{}
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
	case "IdentifyRequest":
		var y Greeting_IdentifyArg
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.Identify = &y
		return nil
	case "HelloRequest":
		var y HelloRequest
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.Hello = &y
		return nil


	}

	return pd3.Errorf("inductive map has no applicable keys")

}

type AnonInductive4_MapIterator struct {
	done bool
	s    *AnonInductive4
}

func (x *AnonInductive4_MapIterator) Next() (key pd1.Node, value pd1.Node, err error) {
	if x.done {
		return nil, nil, pd2.ErrNA
	} else {
		x.done = true
		switch {
			case x.s.Identify != nil:
			return pd2.String("IdentifyRequest"), x.s.Identify.Node(), nil
			case x.s.Hello != nil:
			return pd2.String("HelloRequest"), x.s.Hello.Node(), nil


		default:
			return nil, nil, pd3.Errorf("no inductive cases are set")
		}
	}
}

func (x *AnonInductive4_MapIterator) Done() bool {
	return x.done
}

func (x AnonInductive4) Node() pd1.Node {
	return x
}

func (x AnonInductive4) Kind() pd1.Kind {
	return pd1.Kind_Map
}

func (x AnonInductive4) LookupByString(key string) (pd1.Node, error) {
	switch {
		case x.Identify != nil && key == "IdentifyRequest":
		return x.Identify.Node(), nil
		case x.Hello != nil && key == "HelloRequest":
		return x.Hello.Node(), nil


	}
	return nil, pd2.ErrNA
}

func (x AnonInductive4) LookupByNode(key pd1.Node) (pd1.Node, error) {
	if key.Kind() != pd1.Kind_String {
		return nil, pd2.ErrNA
	}
	if s, err := key.AsString(); err != nil {
		return nil, err
	} else {
		return x.LookupByString(s)
	}
}

func (x AnonInductive4) LookupByIndex(idx int64) (pd1.Node, error) {
	return nil, pd2.ErrNA
}

func (x AnonInductive4) LookupBySegment(seg pd1.PathSegment) (pd1.Node, error) {
	switch seg.String() {
		case "IdentifyRequest":
		return x.Identify.Node(), nil
		case "HelloRequest":
		return x.Hello.Node(), nil


	}
	return nil, pd2.ErrNA
}

func (x AnonInductive4) MapIterator() pd1.MapIterator {
	return &AnonInductive4_MapIterator{false, &x}
}

func (x AnonInductive4) ListIterator() pd1.ListIterator {
	return nil
}

func (x AnonInductive4) Length() int64 {
	return 1
}

func (x AnonInductive4) IsAbsent() bool {
	return false
}

func (x AnonInductive4) IsNull() bool {
	return false
}

func (x AnonInductive4) AsBool() (bool, error) {
	return false, pd2.ErrNA
}

func (x AnonInductive4) AsInt() (int64, error) {
	return 0, pd2.ErrNA
}

func (x AnonInductive4) AsFloat() (float64, error) {
	return 0, pd2.ErrNA
}

func (x AnonInductive4) AsString() (string, error) {
	return "", pd2.ErrNA
}

func (x AnonInductive4) AsBytes() ([]byte, error) {
	return nil, pd2.ErrNA
}

func (x AnonInductive4) AsLink() (pd1.Link, error) {
	return nil, pd2.ErrNA
}

func (x AnonInductive4) Prototype() pd1.NodePrototype {
	return nil
}
// -- protocol type AnonInductive5 --

type AnonInductive5 struct {
		Identify *Greeting_IdentifyResult
		Hello *HelloResponse
		Error *Greeting_Error


}

func (x *AnonInductive5) Parse(n pd1.Node) error {
	*x = AnonInductive5{}
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
	case "IdentifyResponse":
		var y Greeting_IdentifyResult
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.Identify = &y
		return nil
	case "HelloResponse":
		var y HelloResponse
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.Hello = &y
		return nil
	case "Error":
		var y Greeting_Error
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.Error = &y
		return nil


	}

	return pd3.Errorf("inductive map has no applicable keys")

}

type AnonInductive5_MapIterator struct {
	done bool
	s    *AnonInductive5
}

func (x *AnonInductive5_MapIterator) Next() (key pd1.Node, value pd1.Node, err error) {
	if x.done {
		return nil, nil, pd2.ErrNA
	} else {
		x.done = true
		switch {
			case x.s.Identify != nil:
			return pd2.String("IdentifyResponse"), x.s.Identify.Node(), nil
			case x.s.Hello != nil:
			return pd2.String("HelloResponse"), x.s.Hello.Node(), nil
			case x.s.Error != nil:
			return pd2.String("Error"), x.s.Error.Node(), nil


		default:
			return nil, nil, pd3.Errorf("no inductive cases are set")
		}
	}
}

func (x *AnonInductive5_MapIterator) Done() bool {
	return x.done
}

func (x AnonInductive5) Node() pd1.Node {
	return x
}

func (x AnonInductive5) Kind() pd1.Kind {
	return pd1.Kind_Map
}

func (x AnonInductive5) LookupByString(key string) (pd1.Node, error) {
	switch {
		case x.Identify != nil && key == "IdentifyResponse":
		return x.Identify.Node(), nil
		case x.Hello != nil && key == "HelloResponse":
		return x.Hello.Node(), nil
		case x.Error != nil && key == "Error":
		return x.Error.Node(), nil


	}
	return nil, pd2.ErrNA
}

func (x AnonInductive5) LookupByNode(key pd1.Node) (pd1.Node, error) {
	if key.Kind() != pd1.Kind_String {
		return nil, pd2.ErrNA
	}
	if s, err := key.AsString(); err != nil {
		return nil, err
	} else {
		return x.LookupByString(s)
	}
}

func (x AnonInductive5) LookupByIndex(idx int64) (pd1.Node, error) {
	return nil, pd2.ErrNA
}

func (x AnonInductive5) LookupBySegment(seg pd1.PathSegment) (pd1.Node, error) {
	switch seg.String() {
		case "IdentifyResponse":
		return x.Identify.Node(), nil
		case "HelloResponse":
		return x.Hello.Node(), nil
		case "Error":
		return x.Error.Node(), nil


	}
	return nil, pd2.ErrNA
}

func (x AnonInductive5) MapIterator() pd1.MapIterator {
	return &AnonInductive5_MapIterator{false, &x}
}

func (x AnonInductive5) ListIterator() pd1.ListIterator {
	return nil
}

func (x AnonInductive5) Length() int64 {
	return 1
}

func (x AnonInductive5) IsAbsent() bool {
	return false
}

func (x AnonInductive5) IsNull() bool {
	return false
}

func (x AnonInductive5) AsBool() (bool, error) {
	return false, pd2.ErrNA
}

func (x AnonInductive5) AsInt() (int64, error) {
	return 0, pd2.ErrNA
}

func (x AnonInductive5) AsFloat() (float64, error) {
	return 0, pd2.ErrNA
}

func (x AnonInductive5) AsString() (string, error) {
	return "", pd2.ErrNA
}

func (x AnonInductive5) AsBytes() ([]byte, error) {
	return nil, pd2.ErrNA
}

func (x AnonInductive5) AsLink() (pd1.Link, error) {
	return nil, pd2.ErrNA
}

func (x AnonInductive5) Prototype() pd1.NodePrototype {
	return nil
}
var logger_client_Greeting = pd6.Logger("service/client/Greeting")

type Greeting_Client interface {

Identify(ctx pd7.Context, req *Greeting_IdentifyArg) ([]*Greeting_IdentifyResult, error)

Hello(ctx pd7.Context, req *HelloRequest) ([]*HelloResponse, error)


Identify_Async(ctx pd7.Context, req *Greeting_IdentifyArg) (<-chan Greeting_Identify_AsyncResult, error)

Hello_Async(ctx pd7.Context, req *HelloRequest) (<-chan Greeting_Hello_AsyncResult, error)

}


type Greeting_Identify_AsyncResult struct {
	Resp *Greeting_IdentifyResult
	Err  error
}

type Greeting_Hello_AsyncResult struct {
	Resp *HelloResponse
	Err  error
}


type Greeting_ClientOption func(*client_Greeting) error

type client_Greeting struct {
	httpClient       *pd4.Client
	endpoint     *pd5.URL
}

func Greeting_Client_WithHTTPClient(hc *pd4.Client) Greeting_ClientOption {
	return func(c *client_Greeting) error {
		c.httpClient = hc
		return nil
	}
}

func New_Greeting_Client(endpoint string, opts ...Greeting_ClientOption) (*client_Greeting, error) {
	u, err := pd5.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	c := &client_Greeting{endpoint: u, httpClient: pd4.DefaultClient}
	for _, o := range opts {
		if err := o(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}



func (c *client_Greeting) Identify(ctx pd7.Context, req *Greeting_IdentifyArg) ([]*Greeting_IdentifyResult, error) {
	ctx, cancel := pd7.WithCancel(ctx)
	defer cancel()
	ch, err := c.Identify_Async(ctx, req)
	if err != nil {
		return nil, err
	}
	var resps []*Greeting_IdentifyResult
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
					logger_client_Greeting.Errorf("client received error response (%v)", r.Err)
					cancel()
					return resps, r.Err
				}
			}
		case <-ctx.Done():
			return resps, ctx.Err()
		}
	}
}

func (c *client_Greeting) Identify_Async(ctx pd7.Context, req *Greeting_IdentifyArg) (<-chan Greeting_Identify_AsyncResult, error) {
	envelope := &AnonInductive4{
		Identify: req,
	}

	buf, err := pd8.Encode(envelope, pd11.Encode)
	if err != nil {
		return nil, pd3.Errorf("unexpected serialization error (%v)", err)
	}

	// encode request in URL
	u := *c.endpoint
	q := pd5.Values{}
	q.Set("q", string(buf))
	u.RawQuery = q.Encode()
	httpReq, err := pd4.NewRequestWithContext(ctx, "POST", u.String(), pd10.NewReader(buf))
	if err != nil {
		return nil, err
	}
	httpReq.Header = map[string][]string{
		"Accept": {
			"application/vnd.ipfs.rpc+dag-json; version=1",
		},
	}

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, pd3.Errorf("sending HTTP request (%v)", err)
	}

	ch := make(chan Greeting_Identify_AsyncResult, 1)
	go process_Greeting_Identify_AsyncResult(ctx, ch, resp.Body)
	return ch, nil
}

func process_Greeting_Identify_AsyncResult(ctx pd7.Context, ch chan<- Greeting_Identify_AsyncResult, r pd9.Reader) {
	defer close(ch)
	for {
		if ctx.Err() != nil {
			ch <- Greeting_Identify_AsyncResult{Err: pd13.ErrContext{Cause: ctx.Err()}} // context cancelled
			return
		}

		n, err := pd8.DecodeStreaming(r, pd11.Decode)
		if pd12.Is(err, pd9.EOF) || pd12.Is(err, pd9.ErrUnexpectedEOF) {
			return
		}
		if err != nil {
			ch <- Greeting_Identify_AsyncResult{Err: pd13.ErrProto{Cause: err}} // IPLD decode error
			return
		}
		env := &AnonInductive5{}
		if err = env.Parse(n); err != nil {
			ch <- Greeting_Identify_AsyncResult{Err: pd13.ErrProto{Cause: err}} // schema decode error
			return
		}

		if env.Error != nil {
			ch <- Greeting_Identify_AsyncResult{Err: pd13.ErrService{Cause: pd12.New(string(env.Error.Code))}} // service-level error
			return
		}
		if env.Identify == nil {
			continue
		}
		ch <- Greeting_Identify_AsyncResult{Resp: env.Identify}
	}
}


func (c *client_Greeting) Hello(ctx pd7.Context, req *HelloRequest) ([]*HelloResponse, error) {
	ctx, cancel := pd7.WithCancel(ctx)
	defer cancel()
	ch, err := c.Hello_Async(ctx, req)
	if err != nil {
		return nil, err
	}
	var resps []*HelloResponse
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
					logger_client_Greeting.Errorf("client received error response (%v)", r.Err)
					cancel()
					return resps, r.Err
				}
			}
		case <-ctx.Done():
			return resps, ctx.Err()
		}
	}
}

func (c *client_Greeting) Hello_Async(ctx pd7.Context, req *HelloRequest) (<-chan Greeting_Hello_AsyncResult, error) {
	envelope := &AnonInductive4{
		Hello: req,
	}

	buf, err := pd8.Encode(envelope, pd11.Encode)
	if err != nil {
		return nil, pd3.Errorf("unexpected serialization error (%v)", err)
	}

	// encode request in URL
	u := *c.endpoint
	q := pd5.Values{}
	q.Set("q", string(buf))
	u.RawQuery = q.Encode()
	httpReq, err := pd4.NewRequestWithContext(ctx, "POST", u.String(), pd10.NewReader(buf))
	if err != nil {
		return nil, err
	}
	httpReq.Header = map[string][]string{
		"Accept": {
			"application/vnd.ipfs.rpc+dag-json; version=1",
		},
	}

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, pd3.Errorf("sending HTTP request (%v)", err)
	}

	ch := make(chan Greeting_Hello_AsyncResult, 1)
	go process_Greeting_Hello_AsyncResult(ctx, ch, resp.Body)
	return ch, nil
}

func process_Greeting_Hello_AsyncResult(ctx pd7.Context, ch chan<- Greeting_Hello_AsyncResult, r pd9.Reader) {
	defer close(ch)
	for {
		if ctx.Err() != nil {
			ch <- Greeting_Hello_AsyncResult{Err: pd13.ErrContext{Cause: ctx.Err()}} // context cancelled
			return
		}

		n, err := pd8.DecodeStreaming(r, pd11.Decode)
		if pd12.Is(err, pd9.EOF) || pd12.Is(err, pd9.ErrUnexpectedEOF) {
			return
		}
		if err != nil {
			ch <- Greeting_Hello_AsyncResult{Err: pd13.ErrProto{Cause: err}} // IPLD decode error
			return
		}
		env := &AnonInductive5{}
		if err = env.Parse(n); err != nil {
			ch <- Greeting_Hello_AsyncResult{Err: pd13.ErrProto{Cause: err}} // schema decode error
			return
		}

		if env.Error != nil {
			ch <- Greeting_Hello_AsyncResult{Err: pd13.ErrService{Cause: pd12.New(string(env.Error.Code))}} // service-level error
			return
		}
		if env.Hello == nil {
			continue
		}
		ch <- Greeting_Hello_AsyncResult{Resp: env.Hello}
	}
}


var logger_server_Greeting = pd6.Logger("service/server/Greeting")

type Greeting_Server interface {

	Hello(ctx pd7.Context, req *HelloRequest, respCh chan<- *Greeting_Hello_AsyncResult) error
}

func Greeting_AsyncHandler(s Greeting_Server) pd4.HandlerFunc {
	return func(writer pd4.ResponseWriter, request *pd4.Request) {
		// parse request
		msg := request.URL.Query().Get("q")
		n, err := pd8.Decode([]byte(msg), pd11.Decode)
		if err != nil {
			logger_server_Greeting.Errorf("received request not decodeable (%v)", err)
			writer.WriteHeader(400)
			return
		}
		env := &AnonInductive4{}
		if err = env.Parse(n); err != nil {
			logger_server_Greeting.Errorf("parsing call envelope (%v)", err)
			writer.WriteHeader(400)
			return
		}

		writer.Header()["Content-Type"] = []string{
			"application/vnd.ipfs.rpc+dag-json; version=1",
		}

		// demultiplex request
		switch {

		case env.Hello != nil:
			ch := make(chan *Greeting_Hello_AsyncResult)
			if err = s.Hello(pd7.Background(), env.Hello, ch); err != nil {
				logger_server_Greeting.Errorf("get p2p provider rejected request (%v)", err)
				writer.WriteHeader(500)
				return
			}
			for resp := range ch {
				var env *AnonInductive5
				if resp.Err != nil {
					env = &AnonInductive5{ Error: &Greeting_Error{Code: pd2.String(resp.Err.Error())} }
				} else {
					env = &AnonInductive5{ Hello: resp.Resp }
				}
				var buf pd10.Buffer
				if err = pd8.EncodeStreaming(&buf, env, pd11.Encode); err != nil {
					logger_server_Greeting.Errorf("cannot encode response (%v)", err)
					continue
				}
				buf.WriteByte("\n"[0])
				writer.Write(buf.Bytes())
		}


		case env.Identify != nil:
			var env *AnonInductive5
			env = &AnonInductive5{
				Identify: &Greeting_IdentifyResult{
					Methods: []pd2.String{
						"Hello",

					},
				},
			}
			var buf pd10.Buffer
			if err = pd8.EncodeStreaming(&buf, env, pd11.Encode); err != nil {
				logger_server_Greeting.Errorf("cannot encode identify response (%v)", err)
				writer.WriteHeader(500)
				return
			}
			buf.WriteByte("\n"[0])
			writer.Write(buf.Bytes())

		default:
			logger_server_Greeting.Errorf("missing or unknown request")
			writer.WriteHeader(404)
		}
	}
}

// -- protocol type AnonList7 --

type AnonList7 []pd2.String

func (v AnonList7) Node() pd1.Node {
	return v
}

func (v *AnonList7) Parse(n pd1.Node) error {
	if n.Kind() == pd1.Kind_Null {
		*v = nil
		return nil
	}
	if n.Kind() != pd1.Kind_List {
		return pd2.ErrNA
	} else {
		*v = make(AnonList7, n.Length())
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

func (AnonList7) Kind() pd1.Kind {
	return pd1.Kind_List
}

func (AnonList7) LookupByString(string) (pd1.Node, error) {
	return nil, pd2.ErrNA
}

func (AnonList7) LookupByNode(key pd1.Node) (pd1.Node, error) {
	return nil, pd2.ErrNA
}

func (v AnonList7) LookupByIndex(i int64) (pd1.Node, error) {
	if i < 0 || i >= v.Length() {
		return nil, pd2.ErrBounds
	} else {
		return v[i].Node(), nil
	}
}

func (v AnonList7) LookupBySegment(seg pd1.PathSegment) (pd1.Node, error) {
	if i, err := seg.Index(); err != nil {
		return nil, pd2.ErrNA
	} else {
		return v.LookupByIndex(i)
	}
}

func (AnonList7) MapIterator() pd1.MapIterator {
	return nil
}

func (v AnonList7) ListIterator() pd1.ListIterator {
	return &AnonList7_ListIterator{v, 0}
}

func (v AnonList7) Length() int64 {
	return int64(len(v))
}

func (AnonList7) IsAbsent() bool {
	return false
}

func (AnonList7) IsNull() bool {
	return false
}

func (v AnonList7) AsBool() (bool, error) {
	return false, pd2.ErrNA
}

func (AnonList7) AsInt() (int64, error) {
	return 0, pd2.ErrNA
}

func (AnonList7) AsFloat() (float64, error) {
	return 0, pd2.ErrNA
}

func (AnonList7) AsString() (string, error) {
	return "", pd2.ErrNA
}

func (AnonList7) AsBytes() ([]byte, error) {
	return nil, pd2.ErrNA
}

func (AnonList7) AsLink() (pd1.Link, error) {
	return nil, pd2.ErrNA
}

func (AnonList7) Prototype() pd1.NodePrototype {
	return nil // not needed
}

type AnonList7_ListIterator struct {
	list AnonList7
	at   int64
}

func (iter *AnonList7_ListIterator) Next() (int64, pd1.Node, error) {
	if iter.Done() {
		return -1, nil, pd2.ErrBounds
	}
	v := iter.list[iter.at]
	i := int64(iter.at)
	iter.at++
	return i, v.Node(), nil
}

func (iter *AnonList7_ListIterator) Done() bool {
	return iter.at >= iter.list.Length()
}
// -- protocol type AnonInductive8 --

type AnonInductive8 struct {
		US *USAddress
		SK *SKAddress


		OtherCountry string
		OtherAddress *AnonList7

}

func (x *AnonInductive8) Parse(n pd1.Node) error {
	*x = AnonInductive8{}
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
	case "US":
		var y USAddress
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.US = &y
		return nil
	case "SouthKorea":
		var y SKAddress
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.SK = &y
		return nil


	default:
		var y AnonList7
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.OtherCountry = k
		x.OtherAddress = &y
		return nil

	}

}

type AnonInductive8_MapIterator struct {
	done bool
	s    *AnonInductive8
}

func (x *AnonInductive8_MapIterator) Next() (key pd1.Node, value pd1.Node, err error) {
	if x.done {
		return nil, nil, pd2.ErrNA
	} else {
		x.done = true
		switch {
			case x.s.US != nil:
			return pd2.String("US"), x.s.US.Node(), nil
			case x.s.SK != nil:
			return pd2.String("SouthKorea"), x.s.SK.Node(), nil


	case x.s.OtherAddress != nil:
		return pd2.String(x.s.OtherCountry), x.s.OtherAddress.Node(), nil

		default:
			return nil, nil, pd3.Errorf("no inductive cases are set")
		}
	}
}

func (x *AnonInductive8_MapIterator) Done() bool {
	return x.done
}

func (x AnonInductive8) Node() pd1.Node {
	return x
}

func (x AnonInductive8) Kind() pd1.Kind {
	return pd1.Kind_Map
}

func (x AnonInductive8) LookupByString(key string) (pd1.Node, error) {
	switch {
		case x.US != nil && key == "US":
		return x.US.Node(), nil
		case x.SK != nil && key == "SouthKorea":
		return x.SK.Node(), nil


	case x.OtherAddress != nil && key == x.OtherCountry:
		return x.OtherAddress.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x AnonInductive8) LookupByNode(key pd1.Node) (pd1.Node, error) {
	if key.Kind() != pd1.Kind_String {
		return nil, pd2.ErrNA
	}
	if s, err := key.AsString(); err != nil {
		return nil, err
	} else {
		return x.LookupByString(s)
	}
}

func (x AnonInductive8) LookupByIndex(idx int64) (pd1.Node, error) {
	return nil, pd2.ErrNA
}

func (x AnonInductive8) LookupBySegment(seg pd1.PathSegment) (pd1.Node, error) {
	switch seg.String() {
		case "US":
		return x.US.Node(), nil
		case "SouthKorea":
		return x.SK.Node(), nil


	case x.OtherCountry:
		return x.OtherAddress.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x AnonInductive8) MapIterator() pd1.MapIterator {
	return &AnonInductive8_MapIterator{false, &x}
}

func (x AnonInductive8) ListIterator() pd1.ListIterator {
	return nil
}

func (x AnonInductive8) Length() int64 {
	return 1
}

func (x AnonInductive8) IsAbsent() bool {
	return false
}

func (x AnonInductive8) IsNull() bool {
	return false
}

func (x AnonInductive8) AsBool() (bool, error) {
	return false, pd2.ErrNA
}

func (x AnonInductive8) AsInt() (int64, error) {
	return 0, pd2.ErrNA
}

func (x AnonInductive8) AsFloat() (float64, error) {
	return 0, pd2.ErrNA
}

func (x AnonInductive8) AsString() (string, error) {
	return "", pd2.ErrNA
}

func (x AnonInductive8) AsBytes() ([]byte, error) {
	return nil, pd2.ErrNA
}

func (x AnonInductive8) AsLink() (pd1.Link, error) {
	return nil, pd2.ErrNA
}

func (x AnonInductive8) Prototype() pd1.NodePrototype {
	return nil
}
// -- protocol type HelloRequest --

type HelloRequest struct {
		Name pd2.String
		Address AnonInductive8

}

func (x HelloRequest) Node() pd1.Node {
	return x
}

func (x *HelloRequest) Parse(n pd1.Node) error {
	if n.Kind() != pd1.Kind_Map {
		return pd2.ErrNA
	}
	iter := n.MapIterator()
	fieldMap := map[string]pd2.ParseFunc{
				"Name": x.Name.Parse,
		"Address": x.Address.Parse,

	}
	for !iter.Done() {
		if kn, vn, err := iter.Next(); err != nil {
			return err
		} else {
			if k, err := kn.AsString(); err != nil {
				return pd3.Errorf("structure map key is not a string")
			} else {
				_ = vn
				switch k {
			case "Name":
			if _, notParsed := fieldMap["Name"]; !notParsed {
				return pd3.Errorf("field %s already parsed", "Name")
			}
			if err := x.Name.Parse(vn); err != nil {
				return err
			}
			delete(fieldMap, "Name")
			case "Address":
			if _, notParsed := fieldMap["Address"]; !notParsed {
				return pd3.Errorf("field %s already parsed", "Address")
			}
			if err := x.Address.Parse(vn); err != nil {
				return err
			}
			delete(fieldMap, "Address")

				}
			}
		}
	}
	for _, fieldParse := range fieldMap {
		if err := fieldParse(pd1.Null); err != nil {
			return err
		}
	}
	return nil
}

type HelloRequest_MapIterator struct {
	i int64
	s *HelloRequest
}

func (x *HelloRequest_MapIterator) Next() (key pd1.Node, value pd1.Node, err error) {
	x.i++
	switch x.i {
			case 0:
			return pd2.String("Name"), x.s.Name.Node(), nil
			case 1:
			return pd2.String("Address"), x.s.Address.Node(), nil

	}
	return nil, nil, pd2.ErrNA
}

func (x *HelloRequest_MapIterator) Done() bool {
	return x.i + 1 >= 2
}

func (x HelloRequest) Kind() pd1.Kind {
	return pd1.Kind_Map
}

func (x HelloRequest) LookupByString(key string) (pd1.Node, error) {
	switch key {
		case "Name":
		return x.Name.Node(), nil
		case "Address":
		return x.Address.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x HelloRequest) LookupByNode(key pd1.Node) (pd1.Node, error) {
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

func (x HelloRequest) LookupByIndex(idx int64) (pd1.Node, error) {
	switch idx {
		case 0:
		return x.Name.Node(), nil
		case 1:
		return x.Address.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x HelloRequest) LookupBySegment(seg pd1.PathSegment) (pd1.Node, error) {
	switch seg.String() {
		case "0", "Name":
		return x.Name.Node(), nil
		case "1", "Address":
		return x.Address.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x HelloRequest) MapIterator() pd1.MapIterator {
	return &HelloRequest_MapIterator{-1, &x}
}

func (x HelloRequest) ListIterator() pd1.ListIterator {
	return nil
}

func (x HelloRequest) Length() int64 {
	return 2
}

func (x HelloRequest) IsAbsent() bool {
	return false
}

func (x HelloRequest) IsNull() bool {
	return false
}

func (x HelloRequest) AsBool() (bool, error) {
	return false, pd2.ErrNA
}

func (x HelloRequest) AsInt() (int64, error) {
	return 0, pd2.ErrNA
}

func (x HelloRequest) AsFloat() (float64, error) {
	return 0, pd2.ErrNA
}

func (x HelloRequest) AsString() (string, error) {
	return "", pd2.ErrNA
}

func (x HelloRequest) AsBytes() ([]byte, error) {
	return nil, pd2.ErrNA
}

func (x HelloRequest) AsLink() (pd1.Link, error) {
	return nil, pd2.ErrNA
}

func (x HelloRequest) Prototype() pd1.NodePrototype {
	return nil
}

// -- protocol type USAddress --

type USAddress struct {
		Street pd2.String
		City pd2.String
		State State
		ZIP pd2.Int

}

func (x USAddress) Node() pd1.Node {
	return x
}

func (x *USAddress) Parse(n pd1.Node) error {
	if n.Kind() != pd1.Kind_Map {
		return pd2.ErrNA
	}
	iter := n.MapIterator()
	fieldMap := map[string]pd2.ParseFunc{
				"street": x.Street.Parse,
		"city": x.City.Parse,
		"state": x.State.Parse,
		"zip": x.ZIP.Parse,

	}
	for !iter.Done() {
		if kn, vn, err := iter.Next(); err != nil {
			return err
		} else {
			if k, err := kn.AsString(); err != nil {
				return pd3.Errorf("structure map key is not a string")
			} else {
				_ = vn
				switch k {
			case "street":
			if _, notParsed := fieldMap["street"]; !notParsed {
				return pd3.Errorf("field %s already parsed", "street")
			}
			if err := x.Street.Parse(vn); err != nil {
				return err
			}
			delete(fieldMap, "street")
			case "city":
			if _, notParsed := fieldMap["city"]; !notParsed {
				return pd3.Errorf("field %s already parsed", "city")
			}
			if err := x.City.Parse(vn); err != nil {
				return err
			}
			delete(fieldMap, "city")
			case "state":
			if _, notParsed := fieldMap["state"]; !notParsed {
				return pd3.Errorf("field %s already parsed", "state")
			}
			if err := x.State.Parse(vn); err != nil {
				return err
			}
			delete(fieldMap, "state")
			case "zip":
			if _, notParsed := fieldMap["zip"]; !notParsed {
				return pd3.Errorf("field %s already parsed", "zip")
			}
			if err := x.ZIP.Parse(vn); err != nil {
				return err
			}
			delete(fieldMap, "zip")

				}
			}
		}
	}
	for _, fieldParse := range fieldMap {
		if err := fieldParse(pd1.Null); err != nil {
			return err
		}
	}
	return nil
}

type USAddress_MapIterator struct {
	i int64
	s *USAddress
}

func (x *USAddress_MapIterator) Next() (key pd1.Node, value pd1.Node, err error) {
	x.i++
	switch x.i {
			case 0:
			return pd2.String("street"), x.s.Street.Node(), nil
			case 1:
			return pd2.String("city"), x.s.City.Node(), nil
			case 2:
			return pd2.String("state"), x.s.State.Node(), nil
			case 3:
			return pd2.String("zip"), x.s.ZIP.Node(), nil

	}
	return nil, nil, pd2.ErrNA
}

func (x *USAddress_MapIterator) Done() bool {
	return x.i + 1 >= 4
}

func (x USAddress) Kind() pd1.Kind {
	return pd1.Kind_Map
}

func (x USAddress) LookupByString(key string) (pd1.Node, error) {
	switch key {
		case "street":
		return x.Street.Node(), nil
		case "city":
		return x.City.Node(), nil
		case "state":
		return x.State.Node(), nil
		case "zip":
		return x.ZIP.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x USAddress) LookupByNode(key pd1.Node) (pd1.Node, error) {
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

func (x USAddress) LookupByIndex(idx int64) (pd1.Node, error) {
	switch idx {
		case 0:
		return x.Street.Node(), nil
		case 1:
		return x.City.Node(), nil
		case 2:
		return x.State.Node(), nil
		case 3:
		return x.ZIP.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x USAddress) LookupBySegment(seg pd1.PathSegment) (pd1.Node, error) {
	switch seg.String() {
		case "0", "street":
		return x.Street.Node(), nil
		case "1", "city":
		return x.City.Node(), nil
		case "2", "state":
		return x.State.Node(), nil
		case "3", "zip":
		return x.ZIP.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x USAddress) MapIterator() pd1.MapIterator {
	return &USAddress_MapIterator{-1, &x}
}

func (x USAddress) ListIterator() pd1.ListIterator {
	return nil
}

func (x USAddress) Length() int64 {
	return 4
}

func (x USAddress) IsAbsent() bool {
	return false
}

func (x USAddress) IsNull() bool {
	return false
}

func (x USAddress) AsBool() (bool, error) {
	return false, pd2.ErrNA
}

func (x USAddress) AsInt() (int64, error) {
	return 0, pd2.ErrNA
}

func (x USAddress) AsFloat() (float64, error) {
	return 0, pd2.ErrNA
}

func (x USAddress) AsString() (string, error) {
	return "", pd2.ErrNA
}

func (x USAddress) AsBytes() ([]byte, error) {
	return nil, pd2.ErrNA
}

func (x USAddress) AsLink() (pd1.Link, error) {
	return nil, pd2.ErrNA
}

func (x USAddress) Prototype() pd1.NodePrototype {
	return nil
}

// -- protocol type AnonSingletonString11 --

type AnonSingletonString11 struct{}

func (AnonSingletonString11) Parse(n pd1.Node) error {
	if n.Kind() != pd1.Kind_String {
		return pd2.ErrNA
	}
	v, err := n.AsString()
	if err != nil {
		return err
	}
	if v != "CA" {
		return pd2.ErrNA
	}
	return nil
}

func (v AnonSingletonString11) Node() pd1.Node {
	return v
}

func (AnonSingletonString11) Kind() pd1.Kind {
	return pd1.Kind_String
}

func (AnonSingletonString11) LookupByString(string) (pd1.Node, error) {
	return nil, pd2.ErrNA
}

func (AnonSingletonString11) LookupByNode(key pd1.Node) (pd1.Node, error) {
	return nil, pd2.ErrNA
}

func (AnonSingletonString11) LookupByIndex(idx int64) (pd1.Node, error) {
	return nil, pd2.ErrNA
}

func (AnonSingletonString11) LookupBySegment(_ pd1.PathSegment) (pd1.Node, error) {
	return nil, pd2.ErrNA
}

func (AnonSingletonString11) MapIterator() pd1.MapIterator {
	return nil
}

func (AnonSingletonString11) ListIterator() pd1.ListIterator {
	return nil
}

func (AnonSingletonString11) Length() int64 {
	return -1
}

func (AnonSingletonString11) IsAbsent() bool {
	return false
}

func (AnonSingletonString11) IsNull() bool {
	return false
}

func (v AnonSingletonString11) AsBool() (bool, error) {
	return false, pd2.ErrNA
}

func (AnonSingletonString11) AsInt() (int64, error) {
	return 0, pd2.ErrNA
}

func (AnonSingletonString11) AsFloat() (float64, error) {
	return 0, pd2.ErrNA
}

func (AnonSingletonString11) AsString() (string, error) {
	return "CA", nil
}

func (AnonSingletonString11) AsBytes() ([]byte, error) {
	return nil, pd2.ErrNA
}

func (AnonSingletonString11) AsLink() (pd1.Link, error) {
	return nil, pd2.ErrNA
}

func (AnonSingletonString11) Prototype() pd1.NodePrototype {
	return nil
}

// -- protocol type AnonSingletonString12 --

type AnonSingletonString12 struct{}

func (AnonSingletonString12) Parse(n pd1.Node) error {
	if n.Kind() != pd1.Kind_String {
		return pd2.ErrNA
	}
	v, err := n.AsString()
	if err != nil {
		return err
	}
	if v != "NY" {
		return pd2.ErrNA
	}
	return nil
}

func (v AnonSingletonString12) Node() pd1.Node {
	return v
}

func (AnonSingletonString12) Kind() pd1.Kind {
	return pd1.Kind_String
}

func (AnonSingletonString12) LookupByString(string) (pd1.Node, error) {
	return nil, pd2.ErrNA
}

func (AnonSingletonString12) LookupByNode(key pd1.Node) (pd1.Node, error) {
	return nil, pd2.ErrNA
}

func (AnonSingletonString12) LookupByIndex(idx int64) (pd1.Node, error) {
	return nil, pd2.ErrNA
}

func (AnonSingletonString12) LookupBySegment(_ pd1.PathSegment) (pd1.Node, error) {
	return nil, pd2.ErrNA
}

func (AnonSingletonString12) MapIterator() pd1.MapIterator {
	return nil
}

func (AnonSingletonString12) ListIterator() pd1.ListIterator {
	return nil
}

func (AnonSingletonString12) Length() int64 {
	return -1
}

func (AnonSingletonString12) IsAbsent() bool {
	return false
}

func (AnonSingletonString12) IsNull() bool {
	return false
}

func (v AnonSingletonString12) AsBool() (bool, error) {
	return false, pd2.ErrNA
}

func (AnonSingletonString12) AsInt() (int64, error) {
	return 0, pd2.ErrNA
}

func (AnonSingletonString12) AsFloat() (float64, error) {
	return 0, pd2.ErrNA
}

func (AnonSingletonString12) AsString() (string, error) {
	return "NY", nil
}

func (AnonSingletonString12) AsBytes() ([]byte, error) {
	return nil, pd2.ErrNA
}

func (AnonSingletonString12) AsLink() (pd1.Link, error) {
	return nil, pd2.ErrNA
}

func (AnonSingletonString12) Prototype() pd1.NodePrototype {
	return nil
}

// -- protocol type State --

type State struct {
		CA *AnonSingletonString11
		NY *AnonSingletonString12
		Other *pd2.String

}

func (x *State) Parse(n pd1.Node) error {
	*x = State{}
	
			var  CA AnonSingletonString11
			if err := CA.Parse(n); err == nil {
				x.CA = &CA
				return nil
			}
	
			var  NY AnonSingletonString12
			if err := NY.Parse(n); err == nil {
				x.NY = &NY
				return nil
			}
	
			var  Other pd2.String
			if err := Other.Parse(n); err == nil {
				x.Other = &Other
				return nil
			}

	return pd3.Errorf("no union cases parses")
}

func (x State) Node() pd1.Node {
	if x.CA != nil {
		return x.CA
	}
	if x.NY != nil {
		return x.NY
	}
	if x.Other != nil {
		return x.Other
	}

	return nil
}

// proxy Node methods to active case

func (x State) Kind() pd1.Kind {
	if x.CA != nil { return x.CA.Kind() }
	if x.NY != nil { return x.NY.Kind() }
	if x.Other != nil { return x.Other.Kind() }

	return pd1.Kind_Invalid
}

func (x State) LookupByString(key string) (pd1.Node, error) {
	if x.CA != nil { return x.CA.LookupByString(key) }
	if x.NY != nil { return x.NY.LookupByString(key) }
	if x.Other != nil { return x.Other.LookupByString(key) }

	return nil, pd3.Errorf("no active union case found")
}

func (x State) LookupByNode(key pd1.Node) (pd1.Node, error) {
	if x.CA != nil { return x.CA.LookupByNode(key) }
	if x.NY != nil { return x.NY.LookupByNode(key) }
	if x.Other != nil { return x.Other.LookupByNode(key) }

	return nil, pd3.Errorf("no active union case found")
}

func (x State) LookupByIndex(idx int64) (pd1.Node, error) {
	if x.CA != nil { return x.CA.LookupByIndex(idx) }
	if x.NY != nil { return x.NY.LookupByIndex(idx) }
	if x.Other != nil { return x.Other.LookupByIndex(idx) }

	return nil, pd3.Errorf("no active union case found")
}

func (x State) LookupBySegment(seg pd1.PathSegment) (pd1.Node, error) {
	if x.CA != nil { return x.CA.LookupBySegment(seg) }
	if x.NY != nil { return x.NY.LookupBySegment(seg) }
	if x.Other != nil { return x.Other.LookupBySegment(seg) }

	return nil, pd3.Errorf("no active union case found")
}

func (x State) MapIterator() pd1.MapIterator {
	if x.CA != nil { return x.CA.MapIterator() }
	if x.NY != nil { return x.NY.MapIterator() }
	if x.Other != nil { return x.Other.MapIterator() }

	return nil
}

func (x State) ListIterator() pd1.ListIterator {
	if x.CA != nil { return x.CA.ListIterator() }
	if x.NY != nil { return x.NY.ListIterator() }
	if x.Other != nil { return x.Other.ListIterator() }

	return nil
}

func (x State) Length() int64 {
	if x.CA != nil { return x.CA.Length() }
	if x.NY != nil { return x.NY.Length() }
	if x.Other != nil { return x.Other.Length() }

	return -1
}

func (x State) IsAbsent() bool {
	if x.CA != nil { return x.CA.IsAbsent() }
	if x.NY != nil { return x.NY.IsAbsent() }
	if x.Other != nil { return x.Other.IsAbsent() }

	return false
}

func (x State) IsNull() bool {
	if x.CA != nil { return x.CA.IsNull() }
	if x.NY != nil { return x.NY.IsNull() }
	if x.Other != nil { return x.Other.IsNull() }

	return false
}

func (x State) AsBool() (bool, error) {
	if x.CA != nil { return x.CA.AsBool() }
	if x.NY != nil { return x.NY.AsBool() }
	if x.Other != nil { return x.Other.AsBool() }

	return false, pd3.Errorf("no active union case found")
}

func (x State) AsInt() (int64, error) {
	if x.CA != nil { return x.CA.AsInt() }
	if x.NY != nil { return x.NY.AsInt() }
	if x.Other != nil { return x.Other.AsInt() }

	return 0, pd3.Errorf("no active union case found")
}

func (x State) AsFloat() (float64, error) {
	if x.CA != nil { return x.CA.AsFloat() }
	if x.NY != nil { return x.NY.AsFloat() }
	if x.Other != nil { return x.Other.AsFloat() }

	return 0.0, pd3.Errorf("no active union case found")
}

func (x State) AsString() (string, error) {
	if x.CA != nil { return x.CA.AsString() }
	if x.NY != nil { return x.NY.AsString() }
	if x.Other != nil { return x.Other.AsString() }

	return "", pd3.Errorf("no active union case found")
}

func (x State) AsBytes() ([]byte, error) {
	if x.CA != nil { return x.CA.AsBytes() }
	if x.NY != nil { return x.NY.AsBytes() }
	if x.Other != nil { return x.Other.AsBytes() }

	return nil, pd3.Errorf("no active union case found")
}

func (x State) AsLink() (pd1.Link, error) {
	if x.CA != nil { return x.CA.AsLink() }
	if x.NY != nil { return x.NY.AsLink() }
	if x.Other != nil { return x.Other.AsLink() }

	return nil, pd3.Errorf("no active union case found")
}

func (x State) Prototype() pd1.NodePrototype {
	return nil
}
// -- protocol type SKAddress --

type SKAddress struct {
		Street pd2.String
		City pd2.String
		Province pd2.String
		PostalCode pd2.Int

}

func (x SKAddress) Node() pd1.Node {
	return x
}

func (x *SKAddress) Parse(n pd1.Node) error {
	if n.Kind() != pd1.Kind_Map {
		return pd2.ErrNA
	}
	iter := n.MapIterator()
	fieldMap := map[string]pd2.ParseFunc{
				"street": x.Street.Parse,
		"city": x.City.Parse,
		"province": x.Province.Parse,
		"postal_code": x.PostalCode.Parse,

	}
	for !iter.Done() {
		if kn, vn, err := iter.Next(); err != nil {
			return err
		} else {
			if k, err := kn.AsString(); err != nil {
				return pd3.Errorf("structure map key is not a string")
			} else {
				_ = vn
				switch k {
			case "street":
			if _, notParsed := fieldMap["street"]; !notParsed {
				return pd3.Errorf("field %s already parsed", "street")
			}
			if err := x.Street.Parse(vn); err != nil {
				return err
			}
			delete(fieldMap, "street")
			case "city":
			if _, notParsed := fieldMap["city"]; !notParsed {
				return pd3.Errorf("field %s already parsed", "city")
			}
			if err := x.City.Parse(vn); err != nil {
				return err
			}
			delete(fieldMap, "city")
			case "province":
			if _, notParsed := fieldMap["province"]; !notParsed {
				return pd3.Errorf("field %s already parsed", "province")
			}
			if err := x.Province.Parse(vn); err != nil {
				return err
			}
			delete(fieldMap, "province")
			case "postal_code":
			if _, notParsed := fieldMap["postal_code"]; !notParsed {
				return pd3.Errorf("field %s already parsed", "postal_code")
			}
			if err := x.PostalCode.Parse(vn); err != nil {
				return err
			}
			delete(fieldMap, "postal_code")

				}
			}
		}
	}
	for _, fieldParse := range fieldMap {
		if err := fieldParse(pd1.Null); err != nil {
			return err
		}
	}
	return nil
}

type SKAddress_MapIterator struct {
	i int64
	s *SKAddress
}

func (x *SKAddress_MapIterator) Next() (key pd1.Node, value pd1.Node, err error) {
	x.i++
	switch x.i {
			case 0:
			return pd2.String("street"), x.s.Street.Node(), nil
			case 1:
			return pd2.String("city"), x.s.City.Node(), nil
			case 2:
			return pd2.String("province"), x.s.Province.Node(), nil
			case 3:
			return pd2.String("postal_code"), x.s.PostalCode.Node(), nil

	}
	return nil, nil, pd2.ErrNA
}

func (x *SKAddress_MapIterator) Done() bool {
	return x.i + 1 >= 4
}

func (x SKAddress) Kind() pd1.Kind {
	return pd1.Kind_Map
}

func (x SKAddress) LookupByString(key string) (pd1.Node, error) {
	switch key {
		case "street":
		return x.Street.Node(), nil
		case "city":
		return x.City.Node(), nil
		case "province":
		return x.Province.Node(), nil
		case "postal_code":
		return x.PostalCode.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x SKAddress) LookupByNode(key pd1.Node) (pd1.Node, error) {
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

func (x SKAddress) LookupByIndex(idx int64) (pd1.Node, error) {
	switch idx {
		case 0:
		return x.Street.Node(), nil
		case 1:
		return x.City.Node(), nil
		case 2:
		return x.Province.Node(), nil
		case 3:
		return x.PostalCode.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x SKAddress) LookupBySegment(seg pd1.PathSegment) (pd1.Node, error) {
	switch seg.String() {
		case "0", "street":
		return x.Street.Node(), nil
		case "1", "city":
		return x.City.Node(), nil
		case "2", "province":
		return x.Province.Node(), nil
		case "3", "postal_code":
		return x.PostalCode.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x SKAddress) MapIterator() pd1.MapIterator {
	return &SKAddress_MapIterator{-1, &x}
}

func (x SKAddress) ListIterator() pd1.ListIterator {
	return nil
}

func (x SKAddress) Length() int64 {
	return 4
}

func (x SKAddress) IsAbsent() bool {
	return false
}

func (x SKAddress) IsNull() bool {
	return false
}

func (x SKAddress) AsBool() (bool, error) {
	return false, pd2.ErrNA
}

func (x SKAddress) AsInt() (int64, error) {
	return 0, pd2.ErrNA
}

func (x SKAddress) AsFloat() (float64, error) {
	return 0, pd2.ErrNA
}

func (x SKAddress) AsString() (string, error) {
	return "", pd2.ErrNA
}

func (x SKAddress) AsBytes() ([]byte, error) {
	return nil, pd2.ErrNA
}

func (x SKAddress) AsLink() (pd1.Link, error) {
	return nil, pd2.ErrNA
}

func (x SKAddress) Prototype() pd1.NodePrototype {
	return nil
}

// -- protocol type HelloResponse --

type HelloResponse struct {
		English *pd2.String
		Korean *pd2.String

}

func (x *HelloResponse) Parse(n pd1.Node) error {
	*x = HelloResponse{}
	
			var  English pd2.String
			if err := English.Parse(n); err == nil {
				x.English = &English
				return nil
			}
	
			var  Korean pd2.String
			if err := Korean.Parse(n); err == nil {
				x.Korean = &Korean
				return nil
			}

	return pd3.Errorf("no union cases parses")
}

func (x HelloResponse) Node() pd1.Node {
	if x.English != nil {
		return x.English
	}
	if x.Korean != nil {
		return x.Korean
	}

	return nil
}

// proxy Node methods to active case

func (x HelloResponse) Kind() pd1.Kind {
	if x.English != nil { return x.English.Kind() }
	if x.Korean != nil { return x.Korean.Kind() }

	return pd1.Kind_Invalid
}

func (x HelloResponse) LookupByString(key string) (pd1.Node, error) {
	if x.English != nil { return x.English.LookupByString(key) }
	if x.Korean != nil { return x.Korean.LookupByString(key) }

	return nil, pd3.Errorf("no active union case found")
}

func (x HelloResponse) LookupByNode(key pd1.Node) (pd1.Node, error) {
	if x.English != nil { return x.English.LookupByNode(key) }
	if x.Korean != nil { return x.Korean.LookupByNode(key) }

	return nil, pd3.Errorf("no active union case found")
}

func (x HelloResponse) LookupByIndex(idx int64) (pd1.Node, error) {
	if x.English != nil { return x.English.LookupByIndex(idx) }
	if x.Korean != nil { return x.Korean.LookupByIndex(idx) }

	return nil, pd3.Errorf("no active union case found")
}

func (x HelloResponse) LookupBySegment(seg pd1.PathSegment) (pd1.Node, error) {
	if x.English != nil { return x.English.LookupBySegment(seg) }
	if x.Korean != nil { return x.Korean.LookupBySegment(seg) }

	return nil, pd3.Errorf("no active union case found")
}

func (x HelloResponse) MapIterator() pd1.MapIterator {
	if x.English != nil { return x.English.MapIterator() }
	if x.Korean != nil { return x.Korean.MapIterator() }

	return nil
}

func (x HelloResponse) ListIterator() pd1.ListIterator {
	if x.English != nil { return x.English.ListIterator() }
	if x.Korean != nil { return x.Korean.ListIterator() }

	return nil
}

func (x HelloResponse) Length() int64 {
	if x.English != nil { return x.English.Length() }
	if x.Korean != nil { return x.Korean.Length() }

	return -1
}

func (x HelloResponse) IsAbsent() bool {
	if x.English != nil { return x.English.IsAbsent() }
	if x.Korean != nil { return x.Korean.IsAbsent() }

	return false
}

func (x HelloResponse) IsNull() bool {
	if x.English != nil { return x.English.IsNull() }
	if x.Korean != nil { return x.Korean.IsNull() }

	return false
}

func (x HelloResponse) AsBool() (bool, error) {
	if x.English != nil { return x.English.AsBool() }
	if x.Korean != nil { return x.Korean.AsBool() }

	return false, pd3.Errorf("no active union case found")
}

func (x HelloResponse) AsInt() (int64, error) {
	if x.English != nil { return x.English.AsInt() }
	if x.Korean != nil { return x.Korean.AsInt() }

	return 0, pd3.Errorf("no active union case found")
}

func (x HelloResponse) AsFloat() (float64, error) {
	if x.English != nil { return x.English.AsFloat() }
	if x.Korean != nil { return x.Korean.AsFloat() }

	return 0.0, pd3.Errorf("no active union case found")
}

func (x HelloResponse) AsString() (string, error) {
	if x.English != nil { return x.English.AsString() }
	if x.Korean != nil { return x.Korean.AsString() }

	return "", pd3.Errorf("no active union case found")
}

func (x HelloResponse) AsBytes() ([]byte, error) {
	if x.English != nil { return x.English.AsBytes() }
	if x.Korean != nil { return x.Korean.AsBytes() }

	return nil, pd3.Errorf("no active union case found")
}

func (x HelloResponse) AsLink() (pd1.Link, error) {
	if x.English != nil { return x.English.AsLink() }
	if x.Korean != nil { return x.Korean.AsLink() }

	return nil, pd3.Errorf("no active union case found")
}

func (x HelloResponse) Prototype() pd1.NodePrototype {
	return nil
}