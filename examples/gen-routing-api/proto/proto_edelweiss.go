package proto

import (
	pd7 "bytes"
	pd8 "context"
	pd11 "errors"
	pd3 "fmt"
	pd6 "io"
	pd12 "net/http"
	pd4 "net/url"
	pd5 "sync"

	pd13 "github.com/ipfs/go-log"
	pd14 "github.com/ipld/edelweiss/services"
	pd1 "github.com/ipld/edelweiss/values"
	pd10 "github.com/ipld/go-ipld-prime"
	pd9 "github.com/ipld/go-ipld-prime/codec/dagjson"
	pd2 "github.com/ipld/go-ipld-prime/datamodel"
)

// -- protocol type DelegatedRouting_IdentifyArg --

type DelegatedRouting_IdentifyArg struct {
}

func (x DelegatedRouting_IdentifyArg) Node() pd2.Node {
	return x
}

func (x *DelegatedRouting_IdentifyArg) Parse(n pd2.Node) error {
	if n.Kind() != pd2.Kind_Map {
		return pd1.ErrNA
	}
	iter := n.MapIterator()
	fieldMap := map[string]pd1.ParseFunc{}
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
		if err := fieldParse(pd2.Null); err != nil {
			return err
		}
	}
	return nil
}

type DelegatedRouting_IdentifyArg_MapIterator struct {
	i int64
	s *DelegatedRouting_IdentifyArg
}

func (x *DelegatedRouting_IdentifyArg_MapIterator) Next() (key pd2.Node, value pd2.Node, err error) {
	x.i++
	switch x.i {

	}
	return nil, nil, pd1.ErrNA
}

func (x *DelegatedRouting_IdentifyArg_MapIterator) Done() bool {
	return x.i+1 >= 0
}

func (x DelegatedRouting_IdentifyArg) Kind() pd2.Kind {
	return pd2.Kind_Map
}

func (x DelegatedRouting_IdentifyArg) LookupByString(key string) (pd2.Node, error) {
	switch key {

	}
	return nil, pd1.ErrNA
}

func (x DelegatedRouting_IdentifyArg) LookupByNode(key pd2.Node) (pd2.Node, error) {
	switch key.Kind() {
	case pd2.Kind_String:
		if s, err := key.AsString(); err != nil {
			return nil, err
		} else {
			return x.LookupByString(s)
		}
	case pd2.Kind_Int:
		if i, err := key.AsInt(); err != nil {
			return nil, err
		} else {
			return x.LookupByIndex(i)
		}
	}
	return nil, pd1.ErrNA
}

func (x DelegatedRouting_IdentifyArg) LookupByIndex(idx int64) (pd2.Node, error) {
	switch idx {

	}
	return nil, pd1.ErrNA
}

func (x DelegatedRouting_IdentifyArg) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	switch seg.String() {

	}
	return nil, pd1.ErrNA
}

func (x DelegatedRouting_IdentifyArg) MapIterator() pd2.MapIterator {
	return &DelegatedRouting_IdentifyArg_MapIterator{-1, &x}
}

func (x DelegatedRouting_IdentifyArg) ListIterator() pd2.ListIterator {
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
	return false, pd1.ErrNA
}

func (x DelegatedRouting_IdentifyArg) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (x DelegatedRouting_IdentifyArg) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (x DelegatedRouting_IdentifyArg) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (x DelegatedRouting_IdentifyArg) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (x DelegatedRouting_IdentifyArg) AsLink() (pd2.Link, error) {
	return nil, pd1.ErrNA
}

func (x DelegatedRouting_IdentifyArg) Prototype() pd2.NodePrototype {
	return nil
}

// -- protocol type AnonList1 --

type AnonList1 []pd1.String

func (v AnonList1) Node() pd2.Node {
	return v
}

func (v *AnonList1) Parse(n pd2.Node) error {
	if n.Kind() == pd2.Kind_Null {
		*v = nil
		return nil
	}
	if n.Kind() != pd2.Kind_List {
		return pd1.ErrNA
	} else {
		*v = make(AnonList1, n.Length())
		iter := n.ListIterator()
		for !iter.Done() {
			if i, n, err := iter.Next(); err != nil {
				return pd1.ErrNA
			} else if err = (*v)[i].Parse(n); err != nil {
				return err
			}
		}
		return nil
	}
}

func (AnonList1) Kind() pd2.Kind {
	return pd2.Kind_List
}

func (AnonList1) LookupByString(string) (pd2.Node, error) {
	return nil, pd1.ErrNA
}

func (AnonList1) LookupByNode(key pd2.Node) (pd2.Node, error) {
	return nil, pd1.ErrNA
}

func (v AnonList1) LookupByIndex(i int64) (pd2.Node, error) {
	if i < 0 || i >= v.Length() {
		return nil, pd1.ErrBounds
	} else {
		return v[i].Node(), nil
	}
}

func (v AnonList1) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	if i, err := seg.Index(); err != nil {
		return nil, pd1.ErrNA
	} else {
		return v.LookupByIndex(i)
	}
}

func (AnonList1) MapIterator() pd2.MapIterator {
	return nil
}

func (v AnonList1) ListIterator() pd2.ListIterator {
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
	return false, pd1.ErrNA
}

func (AnonList1) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (AnonList1) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (AnonList1) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (AnonList1) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (AnonList1) AsLink() (pd2.Link, error) {
	return nil, pd1.ErrNA
}

func (AnonList1) Prototype() pd2.NodePrototype {
	return nil // not needed
}

type AnonList1_ListIterator struct {
	list AnonList1
	at   int64
}

func (iter *AnonList1_ListIterator) Next() (int64, pd2.Node, error) {
	if iter.Done() {
		return -1, nil, pd1.ErrBounds
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

func (x DelegatedRouting_IdentifyResult) Node() pd2.Node {
	return x
}

func (x *DelegatedRouting_IdentifyResult) Parse(n pd2.Node) error {
	if n.Kind() != pd2.Kind_Map {
		return pd1.ErrNA
	}
	iter := n.MapIterator()
	fieldMap := map[string]pd1.ParseFunc{
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
		if err := fieldParse(pd2.Null); err != nil {
			return err
		}
	}
	return nil
}

type DelegatedRouting_IdentifyResult_MapIterator struct {
	i int64
	s *DelegatedRouting_IdentifyResult
}

func (x *DelegatedRouting_IdentifyResult_MapIterator) Next() (key pd2.Node, value pd2.Node, err error) {
	x.i++
	switch x.i {
	case 0:
		return pd1.String("Methods"), x.s.Methods.Node(), nil

	}
	return nil, nil, pd1.ErrNA
}

func (x *DelegatedRouting_IdentifyResult_MapIterator) Done() bool {
	return x.i+1 >= 1
}

func (x DelegatedRouting_IdentifyResult) Kind() pd2.Kind {
	return pd2.Kind_Map
}

func (x DelegatedRouting_IdentifyResult) LookupByString(key string) (pd2.Node, error) {
	switch key {
	case "Methods":
		return x.Methods.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x DelegatedRouting_IdentifyResult) LookupByNode(key pd2.Node) (pd2.Node, error) {
	switch key.Kind() {
	case pd2.Kind_String:
		if s, err := key.AsString(); err != nil {
			return nil, err
		} else {
			return x.LookupByString(s)
		}
	case pd2.Kind_Int:
		if i, err := key.AsInt(); err != nil {
			return nil, err
		} else {
			return x.LookupByIndex(i)
		}
	}
	return nil, pd1.ErrNA
}

func (x DelegatedRouting_IdentifyResult) LookupByIndex(idx int64) (pd2.Node, error) {
	switch idx {
	case 0:
		return x.Methods.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x DelegatedRouting_IdentifyResult) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	switch seg.String() {
	case "0", "Methods":
		return x.Methods.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x DelegatedRouting_IdentifyResult) MapIterator() pd2.MapIterator {
	return &DelegatedRouting_IdentifyResult_MapIterator{-1, &x}
}

func (x DelegatedRouting_IdentifyResult) ListIterator() pd2.ListIterator {
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
	return false, pd1.ErrNA
}

func (x DelegatedRouting_IdentifyResult) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (x DelegatedRouting_IdentifyResult) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (x DelegatedRouting_IdentifyResult) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (x DelegatedRouting_IdentifyResult) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (x DelegatedRouting_IdentifyResult) AsLink() (pd2.Link, error) {
	return nil, pd1.ErrNA
}

func (x DelegatedRouting_IdentifyResult) Prototype() pd2.NodePrototype {
	return nil
}

// -- protocol type DelegatedRouting_Error --

type DelegatedRouting_Error struct {
	Code pd1.String
}

func (x DelegatedRouting_Error) Node() pd2.Node {
	return x
}

func (x *DelegatedRouting_Error) Parse(n pd2.Node) error {
	if n.Kind() != pd2.Kind_Map {
		return pd1.ErrNA
	}
	iter := n.MapIterator()
	fieldMap := map[string]pd1.ParseFunc{
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
		if err := fieldParse(pd2.Null); err != nil {
			return err
		}
	}
	return nil
}

type DelegatedRouting_Error_MapIterator struct {
	i int64
	s *DelegatedRouting_Error
}

func (x *DelegatedRouting_Error_MapIterator) Next() (key pd2.Node, value pd2.Node, err error) {
	x.i++
	switch x.i {
	case 0:
		return pd1.String("Code"), x.s.Code.Node(), nil

	}
	return nil, nil, pd1.ErrNA
}

func (x *DelegatedRouting_Error_MapIterator) Done() bool {
	return x.i+1 >= 1
}

func (x DelegatedRouting_Error) Kind() pd2.Kind {
	return pd2.Kind_Map
}

func (x DelegatedRouting_Error) LookupByString(key string) (pd2.Node, error) {
	switch key {
	case "Code":
		return x.Code.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x DelegatedRouting_Error) LookupByNode(key pd2.Node) (pd2.Node, error) {
	switch key.Kind() {
	case pd2.Kind_String:
		if s, err := key.AsString(); err != nil {
			return nil, err
		} else {
			return x.LookupByString(s)
		}
	case pd2.Kind_Int:
		if i, err := key.AsInt(); err != nil {
			return nil, err
		} else {
			return x.LookupByIndex(i)
		}
	}
	return nil, pd1.ErrNA
}

func (x DelegatedRouting_Error) LookupByIndex(idx int64) (pd2.Node, error) {
	switch idx {
	case 0:
		return x.Code.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x DelegatedRouting_Error) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	switch seg.String() {
	case "0", "Code":
		return x.Code.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x DelegatedRouting_Error) MapIterator() pd2.MapIterator {
	return &DelegatedRouting_Error_MapIterator{-1, &x}
}

func (x DelegatedRouting_Error) ListIterator() pd2.ListIterator {
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
	return false, pd1.ErrNA
}

func (x DelegatedRouting_Error) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (x DelegatedRouting_Error) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (x DelegatedRouting_Error) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (x DelegatedRouting_Error) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (x DelegatedRouting_Error) AsLink() (pd2.Link, error) {
	return nil, pd1.ErrNA
}

func (x DelegatedRouting_Error) Prototype() pd2.NodePrototype {
	return nil
}

// -- protocol type AnonInductive4 --

type AnonInductive4 struct {
	Identify      *DelegatedRouting_IdentifyArg
	GetP2PProvide *GetP2PProvideRequest
}

func (x *AnonInductive4) Parse(n pd2.Node) error {
	*x = AnonInductive4{}
	if n.Kind() != pd2.Kind_Map {
		return pd1.ErrNA
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

func (x *AnonInductive4_MapIterator) Next() (key pd2.Node, value pd2.Node, err error) {
	if x.done {
		return nil, nil, pd1.ErrNA
	} else {
		x.done = true
		switch {
		case x.s.Identify != nil:
			return pd1.String("IdentifyRequest"), x.s.Identify.Node(), nil
		case x.s.GetP2PProvide != nil:
			return pd1.String("GetP2PProvideRequest"), x.s.GetP2PProvide.Node(), nil

		default:
			return nil, nil, pd3.Errorf("no inductive cases are set")
		}
	}
}

func (x *AnonInductive4_MapIterator) Done() bool {
	return x.done
}

func (x AnonInductive4) Node() pd2.Node {
	return x
}

func (x AnonInductive4) Kind() pd2.Kind {
	return pd2.Kind_Map
}

func (x AnonInductive4) LookupByString(key string) (pd2.Node, error) {
	switch {
	case x.Identify != nil && key == "IdentifyRequest":
		return x.Identify.Node(), nil
	case x.GetP2PProvide != nil && key == "GetP2PProvideRequest":
		return x.GetP2PProvide.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x AnonInductive4) LookupByNode(key pd2.Node) (pd2.Node, error) {
	if key.Kind() != pd2.Kind_String {
		return nil, pd1.ErrNA
	}
	if s, err := key.AsString(); err != nil {
		return nil, err
	} else {
		return x.LookupByString(s)
	}
}

func (x AnonInductive4) LookupByIndex(idx int64) (pd2.Node, error) {
	return nil, pd1.ErrNA
}

func (x AnonInductive4) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	switch seg.String() {
	case "IdentifyRequest":
		return x.Identify.Node(), nil
	case "GetP2PProvideRequest":
		return x.GetP2PProvide.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x AnonInductive4) MapIterator() pd2.MapIterator {
	return &AnonInductive4_MapIterator{false, &x}
}

func (x AnonInductive4) ListIterator() pd2.ListIterator {
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
	return false, pd1.ErrNA
}

func (x AnonInductive4) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (x AnonInductive4) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (x AnonInductive4) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (x AnonInductive4) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (x AnonInductive4) AsLink() (pd2.Link, error) {
	return nil, pd1.ErrNA
}

func (x AnonInductive4) Prototype() pd2.NodePrototype {
	return nil
}

// -- protocol type AnonInductive5 --

type AnonInductive5 struct {
	Identify      *DelegatedRouting_IdentifyResult
	GetP2PProvide *GetP2PProvideResponse
	Error         *DelegatedRouting_Error
}

func (x *AnonInductive5) Parse(n pd2.Node) error {
	*x = AnonInductive5{}
	if n.Kind() != pd2.Kind_Map {
		return pd1.ErrNA
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

func (x *AnonInductive5_MapIterator) Next() (key pd2.Node, value pd2.Node, err error) {
	if x.done {
		return nil, nil, pd1.ErrNA
	} else {
		x.done = true
		switch {
		case x.s.Identify != nil:
			return pd1.String("IdentifyResponse"), x.s.Identify.Node(), nil
		case x.s.GetP2PProvide != nil:
			return pd1.String("GetP2PProvideResponse"), x.s.GetP2PProvide.Node(), nil
		case x.s.Error != nil:
			return pd1.String("Error"), x.s.Error.Node(), nil

		default:
			return nil, nil, pd3.Errorf("no inductive cases are set")
		}
	}
}

func (x *AnonInductive5_MapIterator) Done() bool {
	return x.done
}

func (x AnonInductive5) Node() pd2.Node {
	return x
}

func (x AnonInductive5) Kind() pd2.Kind {
	return pd2.Kind_Map
}

func (x AnonInductive5) LookupByString(key string) (pd2.Node, error) {
	switch {
	case x.Identify != nil && key == "IdentifyResponse":
		return x.Identify.Node(), nil
	case x.GetP2PProvide != nil && key == "GetP2PProvideResponse":
		return x.GetP2PProvide.Node(), nil
	case x.Error != nil && key == "Error":
		return x.Error.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x AnonInductive5) LookupByNode(key pd2.Node) (pd2.Node, error) {
	if key.Kind() != pd2.Kind_String {
		return nil, pd1.ErrNA
	}
	if s, err := key.AsString(); err != nil {
		return nil, err
	} else {
		return x.LookupByString(s)
	}
}

func (x AnonInductive5) LookupByIndex(idx int64) (pd2.Node, error) {
	return nil, pd1.ErrNA
}

func (x AnonInductive5) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	switch seg.String() {
	case "IdentifyResponse":
		return x.Identify.Node(), nil
	case "GetP2PProvideResponse":
		return x.GetP2PProvide.Node(), nil
	case "Error":
		return x.Error.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x AnonInductive5) MapIterator() pd2.MapIterator {
	return &AnonInductive5_MapIterator{false, &x}
}

func (x AnonInductive5) ListIterator() pd2.ListIterator {
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
	return false, pd1.ErrNA
}

func (x AnonInductive5) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (x AnonInductive5) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (x AnonInductive5) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (x AnonInductive5) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (x AnonInductive5) AsLink() (pd2.Link, error) {
	return nil, pd1.ErrNA
}

func (x AnonInductive5) Prototype() pd2.NodePrototype {
	return nil
}

var logger_client_DelegatedRouting = pd13.Logger("service/client/delegatedrouting")

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
	httpClient  *pd12.Client
	endpoint    *pd4.URL
	ulk         pd5.Mutex
	unsupported map[string]bool // cache of methods not supported by server
}

func DelegatedRouting_Client_WithHTTPClient(hc *pd12.Client) DelegatedRouting_ClientOption {
	return func(c *client_DelegatedRouting) error {
		c.httpClient = hc
		return nil
	}
}

func New_DelegatedRouting_Client(endpoint string, opts ...DelegatedRouting_ClientOption) (*client_DelegatedRouting, error) {
	u, err := pd4.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	c := &client_DelegatedRouting{endpoint: u, httpClient: pd12.DefaultClient, unsupported: make(map[string]bool)}
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
	// check if we have memoized that this method is not supported by the server
	c.ulk.Lock()
	notSupported := c.unsupported["Identify"]
	c.ulk.Unlock()
	if notSupported {
		return nil, pd14.ErrSchema
	}

	envelope := &AnonInductive4{
		Identify: req,
	}

	buf, err := pd10.Encode(envelope, pd9.Encode)
	if err != nil {
		return nil, pd3.Errorf("unexpected serialization error (%v)", err)
	}

	// encode request in URL
	u := *c.endpoint
	q := pd4.Values{}
	q.Set("q", string(buf))
	u.RawQuery = q.Encode()
	httpReq, err := pd12.NewRequestWithContext(ctx, "POST", u.String(), pd7.NewReader(buf))
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

	// HTTP codes 400 and 404 correspond to unrecognized method or request schema
	if resp.StatusCode == 400 || resp.StatusCode == 404 {
		resp.Body.Close()
		// memoize that this method is not supported by the server
		c.ulk.Lock()
		c.unsupported["Identify"] = true
		c.ulk.Unlock()
		return nil, pd14.ErrSchema
	}

	ch := make(chan DelegatedRouting_Identify_AsyncResult, 1)
	go process_DelegatedRouting_Identify_AsyncResult(ctx, ch, resp.Body)
	return ch, nil
}

func process_DelegatedRouting_Identify_AsyncResult(ctx pd8.Context, ch chan<- DelegatedRouting_Identify_AsyncResult, r pd6.Reader) {
	defer close(ch)
	for {
		if ctx.Err() != nil {
			ch <- DelegatedRouting_Identify_AsyncResult{Err: pd14.ErrContext{Cause: ctx.Err()}} // context cancelled
			return
		}

		n, err := pd10.DecodeStreaming(r, pd9.Decode)
		if pd11.Is(err, pd6.EOF) || pd11.Is(err, pd6.ErrUnexpectedEOF) {
			return
		}
		if err != nil {
			ch <- DelegatedRouting_Identify_AsyncResult{Err: pd14.ErrProto{Cause: err}} // IPLD decode error
			return
		}
		env := &AnonInductive5{}
		if err = env.Parse(n); err != nil {
			ch <- DelegatedRouting_Identify_AsyncResult{Err: pd14.ErrProto{Cause: err}} // schema decode error
			return
		}

		if env.Error != nil {
			ch <- DelegatedRouting_Identify_AsyncResult{Err: pd14.ErrService{Cause: pd11.New(string(env.Error.Code))}} // service-level error
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
	// check if we have memoized that this method is not supported by the server
	c.ulk.Lock()
	notSupported := c.unsupported["GetP2PProvide"]
	c.ulk.Unlock()
	if notSupported {
		return nil, pd14.ErrSchema
	}

	envelope := &AnonInductive4{
		GetP2PProvide: req,
	}

	buf, err := pd10.Encode(envelope, pd9.Encode)
	if err != nil {
		return nil, pd3.Errorf("unexpected serialization error (%v)", err)
	}

	// encode request in URL
	u := *c.endpoint
	q := pd4.Values{}
	q.Set("q", string(buf))
	u.RawQuery = q.Encode()
	httpReq, err := pd12.NewRequestWithContext(ctx, "POST", u.String(), pd7.NewReader(buf))
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

	// HTTP codes 400 and 404 correspond to unrecognized method or request schema
	if resp.StatusCode == 400 || resp.StatusCode == 404 {
		resp.Body.Close()
		// memoize that this method is not supported by the server
		c.ulk.Lock()
		c.unsupported["GetP2PProvide"] = true
		c.ulk.Unlock()
		return nil, pd14.ErrSchema
	}

	ch := make(chan DelegatedRouting_GetP2PProvide_AsyncResult, 1)
	go process_DelegatedRouting_GetP2PProvide_AsyncResult(ctx, ch, resp.Body)
	return ch, nil
}

func process_DelegatedRouting_GetP2PProvide_AsyncResult(ctx pd8.Context, ch chan<- DelegatedRouting_GetP2PProvide_AsyncResult, r pd6.Reader) {
	defer close(ch)
	for {
		if ctx.Err() != nil {
			ch <- DelegatedRouting_GetP2PProvide_AsyncResult{Err: pd14.ErrContext{Cause: ctx.Err()}} // context cancelled
			return
		}

		n, err := pd10.DecodeStreaming(r, pd9.Decode)
		if pd11.Is(err, pd6.EOF) || pd11.Is(err, pd6.ErrUnexpectedEOF) {
			return
		}
		if err != nil {
			ch <- DelegatedRouting_GetP2PProvide_AsyncResult{Err: pd14.ErrProto{Cause: err}} // IPLD decode error
			return
		}
		env := &AnonInductive5{}
		if err = env.Parse(n); err != nil {
			ch <- DelegatedRouting_GetP2PProvide_AsyncResult{Err: pd14.ErrProto{Cause: err}} // schema decode error
			return
		}

		if env.Error != nil {
			ch <- DelegatedRouting_GetP2PProvide_AsyncResult{Err: pd14.ErrService{Cause: pd11.New(string(env.Error.Code))}} // service-level error
			return
		}
		if env.GetP2PProvide == nil {
			continue
		}
		ch <- DelegatedRouting_GetP2PProvide_AsyncResult{Resp: env.GetP2PProvide}
	}
}

var logger_server_DelegatedRouting = pd13.Logger("service/server/delegatedrouting")

type DelegatedRouting_Server interface {
	GetP2PProvide(ctx pd8.Context, req *GetP2PProvideRequest) (<-chan *DelegatedRouting_GetP2PProvide_AsyncResult, error)
}

func DelegatedRouting_AsyncHandler(s DelegatedRouting_Server) pd12.HandlerFunc {
	return func(writer pd12.ResponseWriter, request *pd12.Request) {
		// parse request
		msg := request.URL.Query().Get("q")
		n, err := pd10.Decode([]byte(msg), pd9.Decode)
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
					env = &AnonInductive5{Error: &DelegatedRouting_Error{Code: pd1.String(resp.Err.Error())}}
				} else {
					env = &AnonInductive5{GetP2PProvide: resp.Resp}
				}
				var buf pd7.Buffer
				if err = pd10.EncodeStreaming(&buf, env, pd9.Encode); err != nil {
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
					Methods: []pd1.String{
						"GetP2PProvide",
					},
				},
			}
			var buf pd7.Buffer
			if err = pd10.EncodeStreaming(&buf, env, pd9.Encode); err != nil {
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

func (v AnonList7) Node() pd2.Node {
	return v
}

func (v *AnonList7) Parse(n pd2.Node) error {
	if n.Kind() == pd2.Kind_Null {
		*v = nil
		return nil
	}
	if n.Kind() != pd2.Kind_List {
		return pd1.ErrNA
	} else {
		*v = make(AnonList7, n.Length())
		iter := n.ListIterator()
		for !iter.Done() {
			if i, n, err := iter.Next(); err != nil {
				return pd1.ErrNA
			} else if err = (*v)[i].Parse(n); err != nil {
				return err
			}
		}
		return nil
	}
}

func (AnonList7) Kind() pd2.Kind {
	return pd2.Kind_List
}

func (AnonList7) LookupByString(string) (pd2.Node, error) {
	return nil, pd1.ErrNA
}

func (AnonList7) LookupByNode(key pd2.Node) (pd2.Node, error) {
	return nil, pd1.ErrNA
}

func (v AnonList7) LookupByIndex(i int64) (pd2.Node, error) {
	if i < 0 || i >= v.Length() {
		return nil, pd1.ErrBounds
	} else {
		return v[i].Node(), nil
	}
}

func (v AnonList7) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	if i, err := seg.Index(); err != nil {
		return nil, pd1.ErrNA
	} else {
		return v.LookupByIndex(i)
	}
}

func (AnonList7) MapIterator() pd2.MapIterator {
	return nil
}

func (v AnonList7) ListIterator() pd2.ListIterator {
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
	return false, pd1.ErrNA
}

func (AnonList7) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (AnonList7) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (AnonList7) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (AnonList7) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (AnonList7) AsLink() (pd2.Link, error) {
	return nil, pd1.ErrNA
}

func (AnonList7) Prototype() pd2.NodePrototype {
	return nil // not needed
}

type AnonList7_ListIterator struct {
	list AnonList7
	at   int64
}

func (iter *AnonList7_ListIterator) Next() (int64, pd2.Node, error) {
	if iter.Done() {
		return -1, nil, pd1.ErrBounds
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

func (x GetP2PProvideRequest) Node() pd2.Node {
	return x
}

func (x *GetP2PProvideRequest) Parse(n pd2.Node) error {
	if n.Kind() != pd2.Kind_Map {
		return pd1.ErrNA
	}
	iter := n.MapIterator()
	fieldMap := map[string]pd1.ParseFunc{
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
		if err := fieldParse(pd2.Null); err != nil {
			return err
		}
	}
	return nil
}

type GetP2PProvideRequest_MapIterator struct {
	i int64
	s *GetP2PProvideRequest
}

func (x *GetP2PProvideRequest_MapIterator) Next() (key pd2.Node, value pd2.Node, err error) {
	x.i++
	switch x.i {
	case 0:
		return pd1.String("Keys"), x.s.Keys.Node(), nil

	}
	return nil, nil, pd1.ErrNA
}

func (x *GetP2PProvideRequest_MapIterator) Done() bool {
	return x.i+1 >= 1
}

func (x GetP2PProvideRequest) Kind() pd2.Kind {
	return pd2.Kind_Map
}

func (x GetP2PProvideRequest) LookupByString(key string) (pd2.Node, error) {
	switch key {
	case "Keys":
		return x.Keys.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x GetP2PProvideRequest) LookupByNode(key pd2.Node) (pd2.Node, error) {
	switch key.Kind() {
	case pd2.Kind_String:
		if s, err := key.AsString(); err != nil {
			return nil, err
		} else {
			return x.LookupByString(s)
		}
	case pd2.Kind_Int:
		if i, err := key.AsInt(); err != nil {
			return nil, err
		} else {
			return x.LookupByIndex(i)
		}
	}
	return nil, pd1.ErrNA
}

func (x GetP2PProvideRequest) LookupByIndex(idx int64) (pd2.Node, error) {
	switch idx {
	case 0:
		return x.Keys.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x GetP2PProvideRequest) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	switch seg.String() {
	case "0", "Keys":
		return x.Keys.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x GetP2PProvideRequest) MapIterator() pd2.MapIterator {
	return &GetP2PProvideRequest_MapIterator{-1, &x}
}

func (x GetP2PProvideRequest) ListIterator() pd2.ListIterator {
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
	return false, pd1.ErrNA
}

func (x GetP2PProvideRequest) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (x GetP2PProvideRequest) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (x GetP2PProvideRequest) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (x GetP2PProvideRequest) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (x GetP2PProvideRequest) AsLink() (pd2.Link, error) {
	return nil, pd1.ErrNA
}

func (x GetP2PProvideRequest) Prototype() pd2.NodePrototype {
	return nil
}

// -- protocol type AnonList9 --

type AnonList9 []ProvidersByKey

func (v AnonList9) Node() pd2.Node {
	return v
}

func (v *AnonList9) Parse(n pd2.Node) error {
	if n.Kind() == pd2.Kind_Null {
		*v = nil
		return nil
	}
	if n.Kind() != pd2.Kind_List {
		return pd1.ErrNA
	} else {
		*v = make(AnonList9, n.Length())
		iter := n.ListIterator()
		for !iter.Done() {
			if i, n, err := iter.Next(); err != nil {
				return pd1.ErrNA
			} else if err = (*v)[i].Parse(n); err != nil {
				return err
			}
		}
		return nil
	}
}

func (AnonList9) Kind() pd2.Kind {
	return pd2.Kind_List
}

func (AnonList9) LookupByString(string) (pd2.Node, error) {
	return nil, pd1.ErrNA
}

func (AnonList9) LookupByNode(key pd2.Node) (pd2.Node, error) {
	return nil, pd1.ErrNA
}

func (v AnonList9) LookupByIndex(i int64) (pd2.Node, error) {
	if i < 0 || i >= v.Length() {
		return nil, pd1.ErrBounds
	} else {
		return v[i].Node(), nil
	}
}

func (v AnonList9) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	if i, err := seg.Index(); err != nil {
		return nil, pd1.ErrNA
	} else {
		return v.LookupByIndex(i)
	}
}

func (AnonList9) MapIterator() pd2.MapIterator {
	return nil
}

func (v AnonList9) ListIterator() pd2.ListIterator {
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
	return false, pd1.ErrNA
}

func (AnonList9) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (AnonList9) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (AnonList9) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (AnonList9) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (AnonList9) AsLink() (pd2.Link, error) {
	return nil, pd1.ErrNA
}

func (AnonList9) Prototype() pd2.NodePrototype {
	return nil // not needed
}

type AnonList9_ListIterator struct {
	list AnonList9
	at   int64
}

func (iter *AnonList9_ListIterator) Next() (int64, pd2.Node, error) {
	if iter.Done() {
		return -1, nil, pd1.ErrBounds
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

func (x GetP2PProvideResponse) Node() pd2.Node {
	return x
}

func (x *GetP2PProvideResponse) Parse(n pd2.Node) error {
	if n.Kind() != pd2.Kind_Map {
		return pd1.ErrNA
	}
	iter := n.MapIterator()
	fieldMap := map[string]pd1.ParseFunc{
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
		if err := fieldParse(pd2.Null); err != nil {
			return err
		}
	}
	return nil
}

type GetP2PProvideResponse_MapIterator struct {
	i int64
	s *GetP2PProvideResponse
}

func (x *GetP2PProvideResponse_MapIterator) Next() (key pd2.Node, value pd2.Node, err error) {
	x.i++
	switch x.i {
	case 0:
		return pd1.String("ProvidersByKey"), x.s.ProvidersByKey.Node(), nil

	}
	return nil, nil, pd1.ErrNA
}

func (x *GetP2PProvideResponse_MapIterator) Done() bool {
	return x.i+1 >= 1
}

func (x GetP2PProvideResponse) Kind() pd2.Kind {
	return pd2.Kind_Map
}

func (x GetP2PProvideResponse) LookupByString(key string) (pd2.Node, error) {
	switch key {
	case "ProvidersByKey":
		return x.ProvidersByKey.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x GetP2PProvideResponse) LookupByNode(key pd2.Node) (pd2.Node, error) {
	switch key.Kind() {
	case pd2.Kind_String:
		if s, err := key.AsString(); err != nil {
			return nil, err
		} else {
			return x.LookupByString(s)
		}
	case pd2.Kind_Int:
		if i, err := key.AsInt(); err != nil {
			return nil, err
		} else {
			return x.LookupByIndex(i)
		}
	}
	return nil, pd1.ErrNA
}

func (x GetP2PProvideResponse) LookupByIndex(idx int64) (pd2.Node, error) {
	switch idx {
	case 0:
		return x.ProvidersByKey.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x GetP2PProvideResponse) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	switch seg.String() {
	case "0", "ProvidersByKey":
		return x.ProvidersByKey.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x GetP2PProvideResponse) MapIterator() pd2.MapIterator {
	return &GetP2PProvideResponse_MapIterator{-1, &x}
}

func (x GetP2PProvideResponse) ListIterator() pd2.ListIterator {
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
	return false, pd1.ErrNA
}

func (x GetP2PProvideResponse) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (x GetP2PProvideResponse) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (x GetP2PProvideResponse) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (x GetP2PProvideResponse) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (x GetP2PProvideResponse) AsLink() (pd2.Link, error) {
	return nil, pd1.ErrNA
}

func (x GetP2PProvideResponse) Prototype() pd2.NodePrototype {
	return nil
}

// -- protocol type ProvidersByKey --

type ProvidersByKey struct {
	Key      Multihash
	Provider Provider
}

func (x ProvidersByKey) Node() pd2.Node {
	return x
}

func (x *ProvidersByKey) Parse(n pd2.Node) error {
	if n.Kind() != pd2.Kind_Map {
		return pd1.ErrNA
	}
	iter := n.MapIterator()
	fieldMap := map[string]pd1.ParseFunc{
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
		if err := fieldParse(pd2.Null); err != nil {
			return err
		}
	}
	return nil
}

type ProvidersByKey_MapIterator struct {
	i int64
	s *ProvidersByKey
}

func (x *ProvidersByKey_MapIterator) Next() (key pd2.Node, value pd2.Node, err error) {
	x.i++
	switch x.i {
	case 0:
		return pd1.String("Key"), x.s.Key.Node(), nil
	case 1:
		return pd1.String("Provider"), x.s.Provider.Node(), nil

	}
	return nil, nil, pd1.ErrNA
}

func (x *ProvidersByKey_MapIterator) Done() bool {
	return x.i+1 >= 2
}

func (x ProvidersByKey) Kind() pd2.Kind {
	return pd2.Kind_Map
}

func (x ProvidersByKey) LookupByString(key string) (pd2.Node, error) {
	switch key {
	case "Key":
		return x.Key.Node(), nil
	case "Provider":
		return x.Provider.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x ProvidersByKey) LookupByNode(key pd2.Node) (pd2.Node, error) {
	switch key.Kind() {
	case pd2.Kind_String:
		if s, err := key.AsString(); err != nil {
			return nil, err
		} else {
			return x.LookupByString(s)
		}
	case pd2.Kind_Int:
		if i, err := key.AsInt(); err != nil {
			return nil, err
		} else {
			return x.LookupByIndex(i)
		}
	}
	return nil, pd1.ErrNA
}

func (x ProvidersByKey) LookupByIndex(idx int64) (pd2.Node, error) {
	switch idx {
	case 0:
		return x.Key.Node(), nil
	case 1:
		return x.Provider.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x ProvidersByKey) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	switch seg.String() {
	case "0", "Key":
		return x.Key.Node(), nil
	case "1", "Provider":
		return x.Provider.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x ProvidersByKey) MapIterator() pd2.MapIterator {
	return &ProvidersByKey_MapIterator{-1, &x}
}

func (x ProvidersByKey) ListIterator() pd2.ListIterator {
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
	return false, pd1.ErrNA
}

func (x ProvidersByKey) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (x ProvidersByKey) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (x ProvidersByKey) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (x ProvidersByKey) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (x ProvidersByKey) AsLink() (pd2.Link, error) {
	return nil, pd1.ErrNA
}

func (x ProvidersByKey) Prototype() pd2.NodePrototype {
	return nil
}

// -- protocol type Multihash --

type Multihash struct {
	Bytes pd1.Bytes
}

func (x Multihash) Node() pd2.Node {
	return x
}

func (x *Multihash) Parse(n pd2.Node) error {
	if n.Kind() != pd2.Kind_Map {
		return pd1.ErrNA
	}
	iter := n.MapIterator()
	fieldMap := map[string]pd1.ParseFunc{
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
		if err := fieldParse(pd2.Null); err != nil {
			return err
		}
	}
	return nil
}

type Multihash_MapIterator struct {
	i int64
	s *Multihash
}

func (x *Multihash_MapIterator) Next() (key pd2.Node, value pd2.Node, err error) {
	x.i++
	switch x.i {
	case 0:
		return pd1.String("Bytes"), x.s.Bytes.Node(), nil

	}
	return nil, nil, pd1.ErrNA
}

func (x *Multihash_MapIterator) Done() bool {
	return x.i+1 >= 1
}

func (x Multihash) Kind() pd2.Kind {
	return pd2.Kind_Map
}

func (x Multihash) LookupByString(key string) (pd2.Node, error) {
	switch key {
	case "Bytes":
		return x.Bytes.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x Multihash) LookupByNode(key pd2.Node) (pd2.Node, error) {
	switch key.Kind() {
	case pd2.Kind_String:
		if s, err := key.AsString(); err != nil {
			return nil, err
		} else {
			return x.LookupByString(s)
		}
	case pd2.Kind_Int:
		if i, err := key.AsInt(); err != nil {
			return nil, err
		} else {
			return x.LookupByIndex(i)
		}
	}
	return nil, pd1.ErrNA
}

func (x Multihash) LookupByIndex(idx int64) (pd2.Node, error) {
	switch idx {
	case 0:
		return x.Bytes.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x Multihash) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	switch seg.String() {
	case "0", "Bytes":
		return x.Bytes.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x Multihash) MapIterator() pd2.MapIterator {
	return &Multihash_MapIterator{-1, &x}
}

func (x Multihash) ListIterator() pd2.ListIterator {
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
	return false, pd1.ErrNA
}

func (x Multihash) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (x Multihash) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (x Multihash) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (x Multihash) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (x Multihash) AsLink() (pd2.Link, error) {
	return nil, pd1.ErrNA
}

func (x Multihash) Prototype() pd2.NodePrototype {
	return nil
}

// -- protocol type AnonList13 --

type AnonList13 []Node

func (v AnonList13) Node() pd2.Node {
	return v
}

func (v *AnonList13) Parse(n pd2.Node) error {
	if n.Kind() == pd2.Kind_Null {
		*v = nil
		return nil
	}
	if n.Kind() != pd2.Kind_List {
		return pd1.ErrNA
	} else {
		*v = make(AnonList13, n.Length())
		iter := n.ListIterator()
		for !iter.Done() {
			if i, n, err := iter.Next(); err != nil {
				return pd1.ErrNA
			} else if err = (*v)[i].Parse(n); err != nil {
				return err
			}
		}
		return nil
	}
}

func (AnonList13) Kind() pd2.Kind {
	return pd2.Kind_List
}

func (AnonList13) LookupByString(string) (pd2.Node, error) {
	return nil, pd1.ErrNA
}

func (AnonList13) LookupByNode(key pd2.Node) (pd2.Node, error) {
	return nil, pd1.ErrNA
}

func (v AnonList13) LookupByIndex(i int64) (pd2.Node, error) {
	if i < 0 || i >= v.Length() {
		return nil, pd1.ErrBounds
	} else {
		return v[i].Node(), nil
	}
}

func (v AnonList13) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	if i, err := seg.Index(); err != nil {
		return nil, pd1.ErrNA
	} else {
		return v.LookupByIndex(i)
	}
}

func (AnonList13) MapIterator() pd2.MapIterator {
	return nil
}

func (v AnonList13) ListIterator() pd2.ListIterator {
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
	return false, pd1.ErrNA
}

func (AnonList13) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (AnonList13) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (AnonList13) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (AnonList13) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (AnonList13) AsLink() (pd2.Link, error) {
	return nil, pd1.ErrNA
}

func (AnonList13) Prototype() pd2.NodePrototype {
	return nil // not needed
}

type AnonList13_ListIterator struct {
	list AnonList13
	at   int64
}

func (iter *AnonList13_ListIterator) Next() (int64, pd2.Node, error) {
	if iter.Done() {
		return -1, nil, pd1.ErrBounds
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

func (v AnonList14) Node() pd2.Node {
	return v
}

func (v *AnonList14) Parse(n pd2.Node) error {
	if n.Kind() == pd2.Kind_Null {
		*v = nil
		return nil
	}
	if n.Kind() != pd2.Kind_List {
		return pd1.ErrNA
	} else {
		*v = make(AnonList14, n.Length())
		iter := n.ListIterator()
		for !iter.Done() {
			if i, n, err := iter.Next(); err != nil {
				return pd1.ErrNA
			} else if err = (*v)[i].Parse(n); err != nil {
				return err
			}
		}
		return nil
	}
}

func (AnonList14) Kind() pd2.Kind {
	return pd2.Kind_List
}

func (AnonList14) LookupByString(string) (pd2.Node, error) {
	return nil, pd1.ErrNA
}

func (AnonList14) LookupByNode(key pd2.Node) (pd2.Node, error) {
	return nil, pd1.ErrNA
}

func (v AnonList14) LookupByIndex(i int64) (pd2.Node, error) {
	if i < 0 || i >= v.Length() {
		return nil, pd1.ErrBounds
	} else {
		return v[i].Node(), nil
	}
}

func (v AnonList14) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	if i, err := seg.Index(); err != nil {
		return nil, pd1.ErrNA
	} else {
		return v.LookupByIndex(i)
	}
}

func (AnonList14) MapIterator() pd2.MapIterator {
	return nil
}

func (v AnonList14) ListIterator() pd2.ListIterator {
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
	return false, pd1.ErrNA
}

func (AnonList14) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (AnonList14) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (AnonList14) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (AnonList14) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (AnonList14) AsLink() (pd2.Link, error) {
	return nil, pd1.ErrNA
}

func (AnonList14) Prototype() pd2.NodePrototype {
	return nil // not needed
}

type AnonList14_ListIterator struct {
	list AnonList14
	at   int64
}

func (iter *AnonList14_ListIterator) Next() (int64, pd2.Node, error) {
	if iter.Done() {
		return -1, nil, pd1.ErrBounds
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

func (x Provider) Node() pd2.Node {
	return x
}

func (x *Provider) Parse(n pd2.Node) error {
	if n.Kind() != pd2.Kind_Map {
		return pd1.ErrNA
	}
	iter := n.MapIterator()
	fieldMap := map[string]pd1.ParseFunc{
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
		if err := fieldParse(pd2.Null); err != nil {
			return err
		}
	}
	return nil
}

type Provider_MapIterator struct {
	i int64
	s *Provider
}

func (x *Provider_MapIterator) Next() (key pd2.Node, value pd2.Node, err error) {
	x.i++
	switch x.i {
	case 0:
		return pd1.String("Nodes"), x.s.Nodes.Node(), nil
	case 1:
		return pd1.String("Proto"), x.s.Proto.Node(), nil

	}
	return nil, nil, pd1.ErrNA
}

func (x *Provider_MapIterator) Done() bool {
	return x.i+1 >= 2
}

func (x Provider) Kind() pd2.Kind {
	return pd2.Kind_Map
}

func (x Provider) LookupByString(key string) (pd2.Node, error) {
	switch key {
	case "Nodes":
		return x.Nodes.Node(), nil
	case "Proto":
		return x.Proto.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x Provider) LookupByNode(key pd2.Node) (pd2.Node, error) {
	switch key.Kind() {
	case pd2.Kind_String:
		if s, err := key.AsString(); err != nil {
			return nil, err
		} else {
			return x.LookupByString(s)
		}
	case pd2.Kind_Int:
		if i, err := key.AsInt(); err != nil {
			return nil, err
		} else {
			return x.LookupByIndex(i)
		}
	}
	return nil, pd1.ErrNA
}

func (x Provider) LookupByIndex(idx int64) (pd2.Node, error) {
	switch idx {
	case 0:
		return x.Nodes.Node(), nil
	case 1:
		return x.Proto.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x Provider) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	switch seg.String() {
	case "0", "Nodes":
		return x.Nodes.Node(), nil
	case "1", "Proto":
		return x.Proto.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x Provider) MapIterator() pd2.MapIterator {
	return &Provider_MapIterator{-1, &x}
}

func (x Provider) ListIterator() pd2.ListIterator {
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
	return false, pd1.ErrNA
}

func (x Provider) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (x Provider) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (x Provider) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (x Provider) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (x Provider) AsLink() (pd2.Link, error) {
	return nil, pd1.ErrNA
}

func (x Provider) Prototype() pd2.NodePrototype {
	return nil
}

// -- protocol type Node --

type Node struct {
	Peer *Peer
}

func (x *Node) Parse(n pd2.Node) error {
	*x = Node{}
	if n.Kind() != pd2.Kind_Map {
		return pd1.ErrNA
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

func (x *Node_MapIterator) Next() (key pd2.Node, value pd2.Node, err error) {
	if x.done {
		return nil, nil, pd1.ErrNA
	} else {
		x.done = true
		switch {
		case x.s.Peer != nil:
			return pd1.String("Peer"), x.s.Peer.Node(), nil

		default:
			return nil, nil, pd3.Errorf("no inductive cases are set")
		}
	}
}

func (x *Node_MapIterator) Done() bool {
	return x.done
}

func (x Node) Node() pd2.Node {
	return x
}

func (x Node) Kind() pd2.Kind {
	return pd2.Kind_Map
}

func (x Node) LookupByString(key string) (pd2.Node, error) {
	switch {
	case x.Peer != nil && key == "Peer":
		return x.Peer.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x Node) LookupByNode(key pd2.Node) (pd2.Node, error) {
	if key.Kind() != pd2.Kind_String {
		return nil, pd1.ErrNA
	}
	if s, err := key.AsString(); err != nil {
		return nil, err
	} else {
		return x.LookupByString(s)
	}
}

func (x Node) LookupByIndex(idx int64) (pd2.Node, error) {
	return nil, pd1.ErrNA
}

func (x Node) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	switch seg.String() {
	case "Peer":
		return x.Peer.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x Node) MapIterator() pd2.MapIterator {
	return &Node_MapIterator{false, &x}
}

func (x Node) ListIterator() pd2.ListIterator {
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
	return false, pd1.ErrNA
}

func (x Node) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (x Node) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (x Node) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (x Node) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (x Node) AsLink() (pd2.Link, error) {
	return nil, pd1.ErrNA
}

func (x Node) Prototype() pd2.NodePrototype {
	return nil
}

// -- protocol type AnonList17 --

type AnonList17 []pd1.Bytes

func (v AnonList17) Node() pd2.Node {
	return v
}

func (v *AnonList17) Parse(n pd2.Node) error {
	if n.Kind() == pd2.Kind_Null {
		*v = nil
		return nil
	}
	if n.Kind() != pd2.Kind_List {
		return pd1.ErrNA
	} else {
		*v = make(AnonList17, n.Length())
		iter := n.ListIterator()
		for !iter.Done() {
			if i, n, err := iter.Next(); err != nil {
				return pd1.ErrNA
			} else if err = (*v)[i].Parse(n); err != nil {
				return err
			}
		}
		return nil
	}
}

func (AnonList17) Kind() pd2.Kind {
	return pd2.Kind_List
}

func (AnonList17) LookupByString(string) (pd2.Node, error) {
	return nil, pd1.ErrNA
}

func (AnonList17) LookupByNode(key pd2.Node) (pd2.Node, error) {
	return nil, pd1.ErrNA
}

func (v AnonList17) LookupByIndex(i int64) (pd2.Node, error) {
	if i < 0 || i >= v.Length() {
		return nil, pd1.ErrBounds
	} else {
		return v[i].Node(), nil
	}
}

func (v AnonList17) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	if i, err := seg.Index(); err != nil {
		return nil, pd1.ErrNA
	} else {
		return v.LookupByIndex(i)
	}
}

func (AnonList17) MapIterator() pd2.MapIterator {
	return nil
}

func (v AnonList17) ListIterator() pd2.ListIterator {
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
	return false, pd1.ErrNA
}

func (AnonList17) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (AnonList17) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (AnonList17) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (AnonList17) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (AnonList17) AsLink() (pd2.Link, error) {
	return nil, pd1.ErrNA
}

func (AnonList17) Prototype() pd2.NodePrototype {
	return nil // not needed
}

type AnonList17_ListIterator struct {
	list AnonList17
	at   int64
}

func (iter *AnonList17_ListIterator) Next() (int64, pd2.Node, error) {
	if iter.Done() {
		return -1, nil, pd1.ErrBounds
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
	ID             pd1.Bytes
	Multiaddresses AnonList17
}

func (x Peer) Node() pd2.Node {
	return x
}

func (x *Peer) Parse(n pd2.Node) error {
	if n.Kind() != pd2.Kind_Map {
		return pd1.ErrNA
	}
	iter := n.MapIterator()
	fieldMap := map[string]pd1.ParseFunc{
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
		if err := fieldParse(pd2.Null); err != nil {
			return err
		}
	}
	return nil
}

type Peer_MapIterator struct {
	i int64
	s *Peer
}

func (x *Peer_MapIterator) Next() (key pd2.Node, value pd2.Node, err error) {
	x.i++
	switch x.i {
	case 0:
		return pd1.String("ID"), x.s.ID.Node(), nil
	case 1:
		return pd1.String("Multiaddresses"), x.s.Multiaddresses.Node(), nil

	}
	return nil, nil, pd1.ErrNA
}

func (x *Peer_MapIterator) Done() bool {
	return x.i+1 >= 2
}

func (x Peer) Kind() pd2.Kind {
	return pd2.Kind_Map
}

func (x Peer) LookupByString(key string) (pd2.Node, error) {
	switch key {
	case "ID":
		return x.ID.Node(), nil
	case "Multiaddresses":
		return x.Multiaddresses.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x Peer) LookupByNode(key pd2.Node) (pd2.Node, error) {
	switch key.Kind() {
	case pd2.Kind_String:
		if s, err := key.AsString(); err != nil {
			return nil, err
		} else {
			return x.LookupByString(s)
		}
	case pd2.Kind_Int:
		if i, err := key.AsInt(); err != nil {
			return nil, err
		} else {
			return x.LookupByIndex(i)
		}
	}
	return nil, pd1.ErrNA
}

func (x Peer) LookupByIndex(idx int64) (pd2.Node, error) {
	switch idx {
	case 0:
		return x.ID.Node(), nil
	case 1:
		return x.Multiaddresses.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x Peer) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	switch seg.String() {
	case "0", "ID":
		return x.ID.Node(), nil
	case "1", "Multiaddresses":
		return x.Multiaddresses.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x Peer) MapIterator() pd2.MapIterator {
	return &Peer_MapIterator{-1, &x}
}

func (x Peer) ListIterator() pd2.ListIterator {
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
	return false, pd1.ErrNA
}

func (x Peer) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (x Peer) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (x Peer) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (x Peer) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (x Peer) AsLink() (pd2.Link, error) {
	return nil, pd1.ErrNA
}

func (x Peer) Prototype() pd2.NodePrototype {
	return nil
}

// -- protocol type TransferProto --

type TransferProto struct {
	Bitswap *BitswapTransfer
}

func (x *TransferProto) Parse(n pd2.Node) error {
	*x = TransferProto{}
	if n.Kind() != pd2.Kind_Map {
		return pd1.ErrNA
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

func (x *TransferProto_MapIterator) Next() (key pd2.Node, value pd2.Node, err error) {
	if x.done {
		return nil, nil, pd1.ErrNA
	} else {
		x.done = true
		switch {
		case x.s.Bitswap != nil:
			return pd1.String("Bitswap"), x.s.Bitswap.Node(), nil

		default:
			return nil, nil, pd3.Errorf("no inductive cases are set")
		}
	}
}

func (x *TransferProto_MapIterator) Done() bool {
	return x.done
}

func (x TransferProto) Node() pd2.Node {
	return x
}

func (x TransferProto) Kind() pd2.Kind {
	return pd2.Kind_Map
}

func (x TransferProto) LookupByString(key string) (pd2.Node, error) {
	switch {
	case x.Bitswap != nil && key == "Bitswap":
		return x.Bitswap.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x TransferProto) LookupByNode(key pd2.Node) (pd2.Node, error) {
	if key.Kind() != pd2.Kind_String {
		return nil, pd1.ErrNA
	}
	if s, err := key.AsString(); err != nil {
		return nil, err
	} else {
		return x.LookupByString(s)
	}
}

func (x TransferProto) LookupByIndex(idx int64) (pd2.Node, error) {
	return nil, pd1.ErrNA
}

func (x TransferProto) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	switch seg.String() {
	case "Bitswap":
		return x.Bitswap.Node(), nil

	}
	return nil, pd1.ErrNA
}

func (x TransferProto) MapIterator() pd2.MapIterator {
	return &TransferProto_MapIterator{false, &x}
}

func (x TransferProto) ListIterator() pd2.ListIterator {
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
	return false, pd1.ErrNA
}

func (x TransferProto) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (x TransferProto) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (x TransferProto) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (x TransferProto) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (x TransferProto) AsLink() (pd2.Link, error) {
	return nil, pd1.ErrNA
}

func (x TransferProto) Prototype() pd2.NodePrototype {
	return nil
}

// -- protocol type BitswapTransfer --

type BitswapTransfer struct {
}

func (x BitswapTransfer) Node() pd2.Node {
	return x
}

func (x *BitswapTransfer) Parse(n pd2.Node) error {
	if n.Kind() != pd2.Kind_Map {
		return pd1.ErrNA
	}
	iter := n.MapIterator()
	fieldMap := map[string]pd1.ParseFunc{}
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
		if err := fieldParse(pd2.Null); err != nil {
			return err
		}
	}
	return nil
}

type BitswapTransfer_MapIterator struct {
	i int64
	s *BitswapTransfer
}

func (x *BitswapTransfer_MapIterator) Next() (key pd2.Node, value pd2.Node, err error) {
	x.i++
	switch x.i {

	}
	return nil, nil, pd1.ErrNA
}

func (x *BitswapTransfer_MapIterator) Done() bool {
	return x.i+1 >= 0
}

func (x BitswapTransfer) Kind() pd2.Kind {
	return pd2.Kind_Map
}

func (x BitswapTransfer) LookupByString(key string) (pd2.Node, error) {
	switch key {

	}
	return nil, pd1.ErrNA
}

func (x BitswapTransfer) LookupByNode(key pd2.Node) (pd2.Node, error) {
	switch key.Kind() {
	case pd2.Kind_String:
		if s, err := key.AsString(); err != nil {
			return nil, err
		} else {
			return x.LookupByString(s)
		}
	case pd2.Kind_Int:
		if i, err := key.AsInt(); err != nil {
			return nil, err
		} else {
			return x.LookupByIndex(i)
		}
	}
	return nil, pd1.ErrNA
}

func (x BitswapTransfer) LookupByIndex(idx int64) (pd2.Node, error) {
	switch idx {

	}
	return nil, pd1.ErrNA
}

func (x BitswapTransfer) LookupBySegment(seg pd2.PathSegment) (pd2.Node, error) {
	switch seg.String() {

	}
	return nil, pd1.ErrNA
}

func (x BitswapTransfer) MapIterator() pd2.MapIterator {
	return &BitswapTransfer_MapIterator{-1, &x}
}

func (x BitswapTransfer) ListIterator() pd2.ListIterator {
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
	return false, pd1.ErrNA
}

func (x BitswapTransfer) AsInt() (int64, error) {
	return 0, pd1.ErrNA
}

func (x BitswapTransfer) AsFloat() (float64, error) {
	return 0, pd1.ErrNA
}

func (x BitswapTransfer) AsString() (string, error) {
	return "", pd1.ErrNA
}

func (x BitswapTransfer) AsBytes() ([]byte, error) {
	return nil, pd1.ErrNA
}

func (x BitswapTransfer) AsLink() (pd2.Link, error) {
	return nil, pd1.ErrNA
}

func (x BitswapTransfer) Prototype() pd2.NodePrototype {
	return nil
}
