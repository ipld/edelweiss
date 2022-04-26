package proto

import (
	pd11 "bytes"
	pd8 "context"
	pd9 "errors"
	pd3 "fmt"
	pd5 "io"
	pd10 "net/http"
	pd7 "net/url"

	pd12 "github.com/ipfs/go-log"
	pd13 "github.com/ipld/edelweiss/services"
	pd2 "github.com/ipld/edelweiss/values"
	pd6 "github.com/ipld/go-ipld-prime"
	pd4 "github.com/ipld/go-ipld-prime/codec/dagjson"
	pd1 "github.com/ipld/go-ipld-prime/datamodel"
)

// -- protocol type DelegatedRouting_IdentifyArg --

type DelegatedRouting_IdentifyArg struct {
}

func (x DelegatedRouting_IdentifyArg) Node() pd1.Node {
	return x
}

func (x *DelegatedRouting_IdentifyArg) Parse(n pd1.Node) error {
	if n.Kind() != pd1.Kind_Map {
		return pd2.ErrNA
	}
	iter := n.MapIterator()
	fieldMap := map[string]pd2.ParseFunc{}
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

type DelegatedRouting_IdentifyArg_MapIterator struct {
	i int64
	s *DelegatedRouting_IdentifyArg
}

func (x *DelegatedRouting_IdentifyArg_MapIterator) Next() (key pd1.Node, value pd1.Node, err error) {
	x.i++
	switch x.i {

	}
	return nil, nil, pd2.ErrNA
}

func (x *DelegatedRouting_IdentifyArg_MapIterator) Done() bool {
	return x.i+1 >= 0
}

func (x DelegatedRouting_IdentifyArg) Kind() pd1.Kind {
	return pd1.Kind_Map
}

func (x DelegatedRouting_IdentifyArg) LookupByString(key string) (pd1.Node, error) {
	switch key {

	}
	return nil, pd2.ErrNA
}

func (x DelegatedRouting_IdentifyArg) LookupByNode(key pd1.Node) (pd1.Node, error) {
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

func (x DelegatedRouting_IdentifyArg) LookupByIndex(idx int64) (pd1.Node, error) {
	switch idx {

	}
	return nil, pd2.ErrNA
}

func (x DelegatedRouting_IdentifyArg) LookupBySegment(seg pd1.PathSegment) (pd1.Node, error) {
	switch seg.String() {

	}
	return nil, pd2.ErrNA
}

func (x DelegatedRouting_IdentifyArg) MapIterator() pd1.MapIterator {
	return &DelegatedRouting_IdentifyArg_MapIterator{-1, &x}
}

func (x DelegatedRouting_IdentifyArg) ListIterator() pd1.ListIterator {
	return nil
}

func (x DelegatedRouting_IdentifyArg) Length() int64 {
	return 0
}

func (x DelegatedRouting_IdentifyArg) IsAbsent() bool {
	return false
}

func (x DelegatedRouting_IdentifyArg) IsNull() bool {
	return false
}

func (x DelegatedRouting_IdentifyArg) AsBool() (bool, error) {
	return false, pd2.ErrNA
}

func (x DelegatedRouting_IdentifyArg) AsInt() (int64, error) {
	return 0, pd2.ErrNA
}

func (x DelegatedRouting_IdentifyArg) AsFloat() (float64, error) {
	return 0, pd2.ErrNA
}

func (x DelegatedRouting_IdentifyArg) AsString() (string, error) {
	return "", pd2.ErrNA
}

func (x DelegatedRouting_IdentifyArg) AsBytes() ([]byte, error) {
	return nil, pd2.ErrNA
}

func (x DelegatedRouting_IdentifyArg) AsLink() (pd1.Link, error) {
	return nil, pd2.ErrNA
}

func (x DelegatedRouting_IdentifyArg) Prototype() pd1.NodePrototype {
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

// -- protocol type DelegatedRouting_IdentifyResult --

type DelegatedRouting_IdentifyResult struct {
	Methods AnonList1
}

func (x DelegatedRouting_IdentifyResult) Node() pd1.Node {
	return x
}

func (x *DelegatedRouting_IdentifyResult) Parse(n pd1.Node) error {
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

type DelegatedRouting_IdentifyResult_MapIterator struct {
	i int64
	s *DelegatedRouting_IdentifyResult
}

func (x *DelegatedRouting_IdentifyResult_MapIterator) Next() (key pd1.Node, value pd1.Node, err error) {
	x.i++
	switch x.i {
	case 0:
		return pd2.String("Methods"), x.s.Methods.Node(), nil

	}
	return nil, nil, pd2.ErrNA
}

func (x *DelegatedRouting_IdentifyResult_MapIterator) Done() bool {
	return x.i+1 >= 1
}

func (x DelegatedRouting_IdentifyResult) Kind() pd1.Kind {
	return pd1.Kind_Map
}

func (x DelegatedRouting_IdentifyResult) LookupByString(key string) (pd1.Node, error) {
	switch key {
	case "Methods":
		return x.Methods.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x DelegatedRouting_IdentifyResult) LookupByNode(key pd1.Node) (pd1.Node, error) {
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

func (x DelegatedRouting_IdentifyResult) LookupByIndex(idx int64) (pd1.Node, error) {
	switch idx {
	case 0:
		return x.Methods.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x DelegatedRouting_IdentifyResult) LookupBySegment(seg pd1.PathSegment) (pd1.Node, error) {
	switch seg.String() {
	case "0", "Methods":
		return x.Methods.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x DelegatedRouting_IdentifyResult) MapIterator() pd1.MapIterator {
	return &DelegatedRouting_IdentifyResult_MapIterator{-1, &x}
}

func (x DelegatedRouting_IdentifyResult) ListIterator() pd1.ListIterator {
	return nil
}

func (x DelegatedRouting_IdentifyResult) Length() int64 {
	return 1
}

func (x DelegatedRouting_IdentifyResult) IsAbsent() bool {
	return false
}

func (x DelegatedRouting_IdentifyResult) IsNull() bool {
	return false
}

func (x DelegatedRouting_IdentifyResult) AsBool() (bool, error) {
	return false, pd2.ErrNA
}

func (x DelegatedRouting_IdentifyResult) AsInt() (int64, error) {
	return 0, pd2.ErrNA
}

func (x DelegatedRouting_IdentifyResult) AsFloat() (float64, error) {
	return 0, pd2.ErrNA
}

func (x DelegatedRouting_IdentifyResult) AsString() (string, error) {
	return "", pd2.ErrNA
}

func (x DelegatedRouting_IdentifyResult) AsBytes() ([]byte, error) {
	return nil, pd2.ErrNA
}

func (x DelegatedRouting_IdentifyResult) AsLink() (pd1.Link, error) {
	return nil, pd2.ErrNA
}

func (x DelegatedRouting_IdentifyResult) Prototype() pd1.NodePrototype {
	return nil
}

// -- protocol type DelegatedRouting_Error --

type DelegatedRouting_Error struct {
	Code pd2.String
}

func (x DelegatedRouting_Error) Node() pd1.Node {
	return x
}

func (x *DelegatedRouting_Error) Parse(n pd1.Node) error {
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

type DelegatedRouting_Error_MapIterator struct {
	i int64
	s *DelegatedRouting_Error
}

func (x *DelegatedRouting_Error_MapIterator) Next() (key pd1.Node, value pd1.Node, err error) {
	x.i++
	switch x.i {
	case 0:
		return pd2.String("Code"), x.s.Code.Node(), nil

	}
	return nil, nil, pd2.ErrNA
}

func (x *DelegatedRouting_Error_MapIterator) Done() bool {
	return x.i+1 >= 1
}

func (x DelegatedRouting_Error) Kind() pd1.Kind {
	return pd1.Kind_Map
}

func (x DelegatedRouting_Error) LookupByString(key string) (pd1.Node, error) {
	switch key {
	case "Code":
		return x.Code.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x DelegatedRouting_Error) LookupByNode(key pd1.Node) (pd1.Node, error) {
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

func (x DelegatedRouting_Error) LookupByIndex(idx int64) (pd1.Node, error) {
	switch idx {
	case 0:
		return x.Code.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x DelegatedRouting_Error) LookupBySegment(seg pd1.PathSegment) (pd1.Node, error) {
	switch seg.String() {
	case "0", "Code":
		return x.Code.Node(), nil

	}
	return nil, pd2.ErrNA
}

func (x DelegatedRouting_Error) MapIterator() pd1.MapIterator {
	return &DelegatedRouting_Error_MapIterator{-1, &x}
}

func (x DelegatedRouting_Error) ListIterator() pd1.ListIterator {
	return nil
}

func (x DelegatedRouting_Error) Length() int64 {
	return 1
}

func (x DelegatedRouting_Error) IsAbsent() bool {
	return false
}

func (x DelegatedRouting_Error) IsNull() bool {
	return false
}

func (x DelegatedRouting_Error) AsBool() (bool, error) {
	return false, pd2.ErrNA
}

func (x DelegatedRouting_Error) AsInt() (int64, error) {
	return 0, pd2.ErrNA
}

func (x DelegatedRouting_Error) AsFloat() (float64, error) {
	return 0, pd2.ErrNA
}

func (x DelegatedRouting_Error) AsString() (string, error) {
	return "", pd2.ErrNA
}

func (x DelegatedRouting_Error) AsBytes() ([]byte, error) {
	return nil, pd2.ErrNA
}

func (x DelegatedRouting_Error) AsLink() (pd1.Link, error) {
	return nil, pd2.ErrNA
}

func (x DelegatedRouting_Error) Prototype() pd1.NodePrototype {
	return nil
}

// -- protocol type AnonInductive4 --

type AnonInductive4 struct {
	Identify      *DelegatedRouting_IdentifyArg
	GetP2PProvide *GetP2PProvideRequest
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
		var y DelegatedRouting_IdentifyArg
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.Identify = &y
		return nil
	case "GetP2PProvideRequest":
		var y GetP2PProvideRequest
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.GetP2PProvide = &y
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
		case x.s.GetP2PProvide != nil:
			return pd2.String("GetP2PProvideRequest"), x.s.GetP2PProvide.Node(), nil

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
	case x.GetP2PProvide != nil && key == "GetP2PProvideRequest":
		return x.GetP2PProvide.Node(), nil

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
	case "GetP2PProvideRequest":
		return x.GetP2PProvide.Node(), nil

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
	Identify      *DelegatedRouting_IdentifyResult
	GetP2PProvide *GetP2PProvideResponse
	Error         *DelegatedRouting_Error
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
		var y DelegatedRouting_IdentifyResult
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.Identify = &y
		return nil
	case "GetP2PProvideResponse":
		var y GetP2PProvideResponse
		if err := y.Parse(vn); err != nil {
			return err
		}
		x.GetP2PProvide = &y
		return nil
	case "Error":
		var y DelegatedRouting_Error
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
		case x.s.GetP2PProvide != nil:
			return pd2.String("GetP2PProvideResponse"), x.s.GetP2PProvide.Node(), nil
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
	case x.GetP2PProvide != nil && key == "GetP2PProvideResponse":
		return x.GetP2PProvide.Node(), nil
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
	case "GetP2PProvideResponse":
		return x.GetP2PProvide.Node(), nil
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

var logger_client_DelegatedRouting = pd12.Logger("service/client/delegatedrouting")

type DelegatedRouting_Client interface {
	Identify(ctx pd8.Context, req *DelegatedRouting_IdentifyArg) ([]*DelegatedRouting_IdentifyResult, error)

	GetP2PProvide(ctx pd8.Context, req *GetP2PProvideRequest) ([]*GetP2PProvideResponse, error)

	Identify_Async(ctx pd8.Context, req *DelegatedRouting_IdentifyArg) (<-chan DelegatedRouting_Identify_AsyncResult, error)

	GetP2PProvide_Async(ctx pd8.Context, req *GetP2PProvideRequest) (<-chan DelegatedRouting_GetP2PProvide_AsyncResult, error)
}

type DelegatedRouting_Identify_AsyncResult struct {
	Resp *DelegatedRouting_IdentifyResult
	Err  error
}

type DelegatedRouting_GetP2PProvide_AsyncResult struct {
	Resp *GetP2PProvideResponse
	Err  error
}

type DelegatedRouting_ClientOption func(*client_DelegatedRouting) error

type client_DelegatedRouting struct {
	httpClient *pd10.Client
	endpoint   *pd7.URL
}

func DelegatedRouting_Client_WithHTTPClient(hc *pd10.Client) DelegatedRouting_ClientOption {
	return func(c *client_DelegatedRouting) error {
		c.httpClient = hc
		return nil
	}
}

func New_DelegatedRouting_Client(endpoint string, opts ...DelegatedRouting_ClientOption) (*client_DelegatedRouting, error) {
	u, err := pd7.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	c := &client_DelegatedRouting{endpoint: u, httpClient: pd10.DefaultClient}
	for _, o := range opts {
		if err := o(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (c *client_DelegatedRouting) Identify(ctx pd8.Context, req *DelegatedRouting_IdentifyArg) ([]*DelegatedRouting_IdentifyResult, error) {
	ctx, cancel := pd8.WithCancel(ctx)
	defer cancel()
	ch, err := c.Identify_Async(ctx, req)
	if err != nil {
		return nil, err
	}
	var resps []*DelegatedRouting_IdentifyResult
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
					logger_client_DelegatedRouting.Errorf("client received error response (%v)", r.Err)
					cancel()
					return resps, r.Err
				}
			}
		case <-ctx.Done():
			return resps, ctx.Err()
		}
	}
}

func (c *client_DelegatedRouting) Identify_Async(ctx pd8.Context, req *DelegatedRouting_IdentifyArg) (<-chan DelegatedRouting_Identify_AsyncResult, error) {
	envelope := &AnonInductive4{
		Identify: req,
	}

	buf, err := pd6.Encode(envelope, pd4.Encode)
	if err != nil {
		return nil, pd3.Errorf("unexpected serialization error (%v)", err)
	}

	// encode request in URL
	u := *c.endpoint
	q := pd7.Values{}
	q.Set("q", string(buf))
	u.RawQuery = q.Encode()
	httpReq, err := pd10.NewRequestWithContext(ctx, "POST", u.String(), pd11.NewReader(buf))
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

	ch := make(chan DelegatedRouting_Identify_AsyncResult, 1)
	go process_DelegatedRouting_Identify_AsyncResult(ctx, ch, resp.Body)
	return ch, nil
}

func process_DelegatedRouting_Identify_AsyncResult(ctx pd8.Context, ch chan<- DelegatedRouting_Identify_AsyncResult, r pd5.Reader) {
	defer close(ch)
	for {
		if ctx.Err() != nil {
			ch <- DelegatedRouting_Identify_AsyncResult{Err: pd13.ErrContext{Cause: ctx.Err()}} // context cancelled
			return
		}

		n, err := pd6.DecodeStreaming(r, pd4.Decode)
		if pd9.Is(err, pd5.EOF) || pd9.Is(err, pd5.ErrUnexpectedEOF) {
			return
		}
		if err != nil {
			ch <- DelegatedRouting_Identify_AsyncResult{Err: pd13.ErrProto{Cause: err}} // IPLD decode error
			return
		}
		env := &AnonInductive5{}
		if err = env.Parse(n); err != nil {
			ch <- DelegatedRouting_Identify_AsyncResult{Err: pd13.ErrProto{Cause: err}} // schema decode error
			return
		}

		if env.Error != nil {
			ch <- DelegatedRouting_Identify_AsyncResult{Err: pd13.ErrService{Cause: pd9.New(string(env.Error.Code))}} // service-level error
			return
		}
		if env.Identify == nil {
			continue
		}
		ch <- DelegatedRouting_Identify_AsyncResult{Resp: env.Identify}
	}
}

func (c *client_DelegatedRouting) GetP2PProvide(ctx pd8.Context, req *GetP2PProvideRequest) ([]*GetP2PProvideResponse, error) {
	ctx, cancel := pd8.WithCancel(ctx)
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
					logger_client_DelegatedRouting.Errorf("client received error response (%v)", r.Err)
					cancel()
					return resps, r.Err
				}
			}
		case <-ctx.Done():
			return resps, ctx.Err()
		}
	}
}

func (c *client_DelegatedRouting) GetP2PProvide_Async(ctx pd8.Context, req *GetP2PProvideRequest) (<-chan DelegatedRouting_GetP2PProvide_AsyncResult, error) {
	envelope := &AnonInductive4{
		GetP2PProvide: req,
	}

	buf, err := pd6.Encode(envelope, pd4.Encode)
	if err != nil {
		return nil, pd3.Errorf("unexpected serialization error (%v)", err)
	}

	// encode request in URL
	u := *c.endpoint
	q := pd7.Values{}
	q.Set("q", string(buf))
	u.RawQuery = q.Encode()
	httpReq, err := pd10.NewRequestWithContext(ctx, "POST", u.String(), pd11.NewReader(buf))
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

	ch := make(chan DelegatedRouting_GetP2PProvide_AsyncResult, 1)
	go process_DelegatedRouting_GetP2PProvide_AsyncResult(ctx, ch, resp.Body)
	return ch, nil
}

func process_DelegatedRouting_GetP2PProvide_AsyncResult(ctx pd8.Context, ch chan<- DelegatedRouting_GetP2PProvide_AsyncResult, r pd5.Reader) {
	defer close(ch)
	for {
		if ctx.Err() != nil {
			ch <- DelegatedRouting_GetP2PProvide_AsyncResult{Err: pd13.ErrContext{Cause: ctx.Err()}} // context cancelled
			return
		}

		n, err := pd6.DecodeStreaming(r, pd4.Decode)
		if pd9.Is(err, pd5.EOF) || pd9.Is(err, pd5.ErrUnexpectedEOF) {
			return
		}
		if err != nil {
			ch <- DelegatedRouting_GetP2PProvide_AsyncResult{Err: pd13.ErrProto{Cause: err}} // IPLD decode error
			return
		}
		env := &AnonInductive5{}
		if err = env.Parse(n); err != nil {
			ch <- DelegatedRouting_GetP2PProvide_AsyncResult{Err: pd13.ErrProto{Cause: err}} // schema decode error
			return
		}

		if env.Error != nil {
			ch <- DelegatedRouting_GetP2PProvide_AsyncResult{Err: pd13.ErrService{Cause: pd9.New(string(env.Error.Code))}} // service-level error
			return
		}
		if env.GetP2PProvide == nil {
			continue
		}
		ch <- DelegatedRouting_GetP2PProvide_AsyncResult{Resp: env.GetP2PProvide}
	}
}

var logger_server_DelegatedRouting = pd12.Logger("service/server/delegatedrouting")

type DelegatedRouting_Server interface {
	GetP2PProvide(ctx pd8.Context, req *GetP2PProvideRequest) (<-chan *DelegatedRouting_GetP2PProvide_AsyncResult, error)
}

func DelegatedRouting_AsyncHandler(s DelegatedRouting_Server) pd10.HandlerFunc {
	return func(writer pd10.ResponseWriter, request *pd10.Request) {
		// parse request
		msg := request.URL.Query().Get("q")
		n, err := pd6.Decode([]byte(msg), pd4.Decode)
		if err != nil {
			logger_server_DelegatedRouting.Errorf("received request not decodeable (%v)", err)
			writer.WriteHeader(400)
			return
		}
		env := &AnonInductive4{}
		if err = env.Parse(n); err != nil {
			logger_server_DelegatedRouting.Errorf("parsing call envelope (%v)", err)
			writer.WriteHeader(400)
			return
		}

		writer.Header()["Content-Type"] = []string{
			"application/vnd.ipfs.rpc+dag-json; version=1",
		}

		// demultiplex request
		switch {

		case env.GetP2PProvide != nil:
			ch, err := s.GetP2PProvide(pd8.Background(), env.GetP2PProvide)
			if err != nil {
				logger_server_DelegatedRouting.Errorf("get p2p provider rejected request (%v)", err)
				writer.WriteHeader(500)
				return
			}
			for resp := range ch {
				var env *AnonInductive5
				if resp.Err != nil {
					env = &AnonInductive5{Error: &DelegatedRouting_Error{Code: pd2.String(resp.Err.Error())}}
				} else {
					env = &AnonInductive5{GetP2PProvide: resp.Resp}
				}
				var buf pd11.Buffer
				if err = pd6.EncodeStreaming(&buf, env, pd4.Encode); err != nil {
					logger_server_DelegatedRouting.Errorf("cannot encode response (%v)", err)
					continue
				}
				buf.WriteByte("\n"[0])
				writer.Write(buf.Bytes())
			}

		case env.Identify != nil:
			var env *AnonInductive5
			env = &AnonInductive5{
				Identify: &DelegatedRouting_IdentifyResult{
					Methods: []pd2.String{
						"GetP2PProvide",
					},
				},
			}
			var buf pd11.Buffer
			if err = pd6.EncodeStreaming(&buf, env, pd4.Encode); err != nil {
				logger_server_DelegatedRouting.Errorf("cannot encode identify response (%v)", err)
				writer.WriteHeader(500)
				return
			}
			buf.WriteByte("\n"[0])
			writer.Write(buf.Bytes())

		default:
			logger_server_DelegatedRouting.Errorf("missing or unknown request")
			writer.WriteHeader(404)
		}
	}
}

// -- protocol type AnonList7 --

type AnonList7 []Multihash

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

// -- protocol type GetP2PProvideRequest --

type GetP2PProvideRequest struct {
	Keys AnonList7
}

func (x GetP2PProvideRequest) Node() pd1.Node {
	return x
}

func (x *GetP2PProvideRequest) Parse(n pd1.Node) error {
	if n.Kind() != pd1.Kind_Map {
		return pd2.ErrNA
	}
	iter := n.MapIterator()
	fieldMap := map[string]pd2.ParseFunc{
		"Keys": x.Keys.Parse,
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
				case "Keys":
					if _, notParsed := fieldMap["Keys"]; !notParsed {
						return pd3.Errorf("field %s already parsed", "Keys")
					}
					if err := x.Keys.Parse(vn); err != nil {
						return err
					}
					delete(fieldMap, "Keys")

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
	return x.i+1 >= 1
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

// -- protocol type AnonList9 --

type AnonList9 []ProvidersByKey

func (v AnonList9) Node() pd1.Node {
	return v
}

func (v *AnonList9) Parse(n pd1.Node) error {
	if n.Kind() == pd1.Kind_Null {
		*v = nil
		return nil
	}
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

// -- protocol type GetP2PProvideResponse --

type GetP2PProvideResponse struct {
	ProvidersByKey AnonList9
}

func (x GetP2PProvideResponse) Node() pd1.Node {
	return x
}

func (x *GetP2PProvideResponse) Parse(n pd1.Node) error {
	if n.Kind() != pd1.Kind_Map {
		return pd2.ErrNA
	}
	iter := n.MapIterator()
	fieldMap := map[string]pd2.ParseFunc{
		"ProvidersByKey": x.ProvidersByKey.Parse,
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
				case "ProvidersByKey":
					if _, notParsed := fieldMap["ProvidersByKey"]; !notParsed {
						return pd3.Errorf("field %s already parsed", "ProvidersByKey")
					}
					if err := x.ProvidersByKey.Parse(vn); err != nil {
						return err
					}
					delete(fieldMap, "ProvidersByKey")

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
	return x.i+1 >= 1
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
	Key      Multihash
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
	fieldMap := map[string]pd2.ParseFunc{
		"Key":      x.Key.Parse,
		"Provider": x.Provider.Parse,
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
				case "Key":
					if _, notParsed := fieldMap["Key"]; !notParsed {
						return pd3.Errorf("field %s already parsed", "Key")
					}
					if err := x.Key.Parse(vn); err != nil {
						return err
					}
					delete(fieldMap, "Key")
				case "Provider":
					if _, notParsed := fieldMap["Provider"]; !notParsed {
						return pd3.Errorf("field %s already parsed", "Provider")
					}
					if err := x.Provider.Parse(vn); err != nil {
						return err
					}
					delete(fieldMap, "Provider")

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
	return x.i+1 >= 2
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
	fieldMap := map[string]pd2.ParseFunc{
		"Bytes": x.Bytes.Parse,
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
				case "Bytes":
					if _, notParsed := fieldMap["Bytes"]; !notParsed {
						return pd3.Errorf("field %s already parsed", "Bytes")
					}
					if err := x.Bytes.Parse(vn); err != nil {
						return err
					}
					delete(fieldMap, "Bytes")

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
	return x.i+1 >= 1
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

// -- protocol type AnonList13 --

type AnonList13 []Node

func (v AnonList13) Node() pd1.Node {
	return v
}

func (v *AnonList13) Parse(n pd1.Node) error {
	if n.Kind() == pd1.Kind_Null {
		*v = nil
		return nil
	}
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

// -- protocol type AnonList14 --

type AnonList14 []TransferProto

func (v AnonList14) Node() pd1.Node {
	return v
}

func (v *AnonList14) Parse(n pd1.Node) error {
	if n.Kind() == pd1.Kind_Null {
		*v = nil
		return nil
	}
	if n.Kind() != pd1.Kind_List {
		return pd2.ErrNA
	} else {
		*v = make(AnonList14, n.Length())
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

func (AnonList14) Kind() pd1.Kind {
	return pd1.Kind_List
}

func (AnonList14) LookupByString(string) (pd1.Node, error) {
	return nil, pd2.ErrNA
}

func (AnonList14) LookupByNode(key pd1.Node) (pd1.Node, error) {
	return nil, pd2.ErrNA
}

func (v AnonList14) LookupByIndex(i int64) (pd1.Node, error) {
	if i < 0 || i >= v.Length() {
		return nil, pd2.ErrBounds
	} else {
		return v[i].Node(), nil
	}
}

func (v AnonList14) LookupBySegment(seg pd1.PathSegment) (pd1.Node, error) {
	if i, err := seg.Index(); err != nil {
		return nil, pd2.ErrNA
	} else {
		return v.LookupByIndex(i)
	}
}

func (AnonList14) MapIterator() pd1.MapIterator {
	return nil
}

func (v AnonList14) ListIterator() pd1.ListIterator {
	return &AnonList14_ListIterator{v, 0}
}

func (v AnonList14) Length() int64 {
	return int64(len(v))
}

func (AnonList14) IsAbsent() bool {
	return false
}

func (AnonList14) IsNull() bool {
	return false
}

func (v AnonList14) AsBool() (bool, error) {
	return false, pd2.ErrNA
}

func (AnonList14) AsInt() (int64, error) {
	return 0, pd2.ErrNA
}

func (AnonList14) AsFloat() (float64, error) {
	return 0, pd2.ErrNA
}

func (AnonList14) AsString() (string, error) {
	return "", pd2.ErrNA
}

func (AnonList14) AsBytes() ([]byte, error) {
	return nil, pd2.ErrNA
}

func (AnonList14) AsLink() (pd1.Link, error) {
	return nil, pd2.ErrNA
}

func (AnonList14) Prototype() pd1.NodePrototype {
	return nil // not needed
}

type AnonList14_ListIterator struct {
	list AnonList14
	at   int64
}

func (iter *AnonList14_ListIterator) Next() (int64, pd1.Node, error) {
	if iter.Done() {
		return -1, nil, pd2.ErrBounds
	}
	v := iter.list[iter.at]
	i := int64(iter.at)
	iter.at++
	return i, v.Node(), nil
}

func (iter *AnonList14_ListIterator) Done() bool {
	return iter.at >= iter.list.Length()
}

// -- protocol type Provider --

type Provider struct {
	Nodes AnonList13
	Proto AnonList14
}

func (x Provider) Node() pd1.Node {
	return x
}

func (x *Provider) Parse(n pd1.Node) error {
	if n.Kind() != pd1.Kind_Map {
		return pd2.ErrNA
	}
	iter := n.MapIterator()
	fieldMap := map[string]pd2.ParseFunc{
		"Nodes": x.Nodes.Parse,
		"Proto": x.Proto.Parse,
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
				case "Nodes":
					if _, notParsed := fieldMap["Nodes"]; !notParsed {
						return pd3.Errorf("field %s already parsed", "Nodes")
					}
					if err := x.Nodes.Parse(vn); err != nil {
						return err
					}
					delete(fieldMap, "Nodes")
				case "Proto":
					if _, notParsed := fieldMap["Proto"]; !notParsed {
						return pd3.Errorf("field %s already parsed", "Proto")
					}
					if err := x.Proto.Parse(vn); err != nil {
						return err
					}
					delete(fieldMap, "Proto")

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
	return x.i+1 >= 2
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

// -- protocol type AnonList17 --

type AnonList17 []pd2.Bytes

func (v AnonList17) Node() pd1.Node {
	return v
}

func (v *AnonList17) Parse(n pd1.Node) error {
	if n.Kind() == pd1.Kind_Null {
		*v = nil
		return nil
	}
	if n.Kind() != pd1.Kind_List {
		return pd2.ErrNA
	} else {
		*v = make(AnonList17, n.Length())
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

func (AnonList17) Kind() pd1.Kind {
	return pd1.Kind_List
}

func (AnonList17) LookupByString(string) (pd1.Node, error) {
	return nil, pd2.ErrNA
}

func (AnonList17) LookupByNode(key pd1.Node) (pd1.Node, error) {
	return nil, pd2.ErrNA
}

func (v AnonList17) LookupByIndex(i int64) (pd1.Node, error) {
	if i < 0 || i >= v.Length() {
		return nil, pd2.ErrBounds
	} else {
		return v[i].Node(), nil
	}
}

func (v AnonList17) LookupBySegment(seg pd1.PathSegment) (pd1.Node, error) {
	if i, err := seg.Index(); err != nil {
		return nil, pd2.ErrNA
	} else {
		return v.LookupByIndex(i)
	}
}

func (AnonList17) MapIterator() pd1.MapIterator {
	return nil
}

func (v AnonList17) ListIterator() pd1.ListIterator {
	return &AnonList17_ListIterator{v, 0}
}

func (v AnonList17) Length() int64 {
	return int64(len(v))
}

func (AnonList17) IsAbsent() bool {
	return false
}

func (AnonList17) IsNull() bool {
	return false
}

func (v AnonList17) AsBool() (bool, error) {
	return false, pd2.ErrNA
}

func (AnonList17) AsInt() (int64, error) {
	return 0, pd2.ErrNA
}

func (AnonList17) AsFloat() (float64, error) {
	return 0, pd2.ErrNA
}

func (AnonList17) AsString() (string, error) {
	return "", pd2.ErrNA
}

func (AnonList17) AsBytes() ([]byte, error) {
	return nil, pd2.ErrNA
}

func (AnonList17) AsLink() (pd1.Link, error) {
	return nil, pd2.ErrNA
}

func (AnonList17) Prototype() pd1.NodePrototype {
	return nil // not needed
}

type AnonList17_ListIterator struct {
	list AnonList17
	at   int64
}

func (iter *AnonList17_ListIterator) Next() (int64, pd1.Node, error) {
	if iter.Done() {
		return -1, nil, pd2.ErrBounds
	}
	v := iter.list[iter.at]
	i := int64(iter.at)
	iter.at++
	return i, v.Node(), nil
}

func (iter *AnonList17_ListIterator) Done() bool {
	return iter.at >= iter.list.Length()
}

// -- protocol type Peer --

type Peer struct {
	ID             pd2.Bytes
	Multiaddresses AnonList17
}

func (x Peer) Node() pd1.Node {
	return x
}

func (x *Peer) Parse(n pd1.Node) error {
	if n.Kind() != pd1.Kind_Map {
		return pd2.ErrNA
	}
	iter := n.MapIterator()
	fieldMap := map[string]pd2.ParseFunc{
		"ID":             x.ID.Parse,
		"Multiaddresses": x.Multiaddresses.Parse,
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
				case "ID":
					if _, notParsed := fieldMap["ID"]; !notParsed {
						return pd3.Errorf("field %s already parsed", "ID")
					}
					if err := x.ID.Parse(vn); err != nil {
						return err
					}
					delete(fieldMap, "ID")
				case "Multiaddresses":
					if _, notParsed := fieldMap["Multiaddresses"]; !notParsed {
						return pd3.Errorf("field %s already parsed", "Multiaddresses")
					}
					if err := x.Multiaddresses.Parse(vn); err != nil {
						return err
					}
					delete(fieldMap, "Multiaddresses")

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
	return x.i+1 >= 2
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
	fieldMap := map[string]pd2.ParseFunc{}
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
	return x.i+1 >= 0
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
