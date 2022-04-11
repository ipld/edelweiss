package plans

import (
	"github.com/ipld/edelweiss/defs"
)

type Plan interface {
	IAmPlan()
	Kind() string
}

type Builtin interface {
	IAmBuiltin()
}

type BuiltinOrRefPlan interface {
	Plan
}

type Ref struct {
	Name string
}

func (Ref) IAmPlan()     {}
func (Ref) IAmRef()      {}
func (Ref) Kind() string { return "Ref" }

type Bool struct{}

func (Bool) IAmPlan()     {}
func (Bool) IAmBuiltin()  {}
func (Bool) Kind() string { return "Bool" }

type Int struct{}

func (Int) IAmPlan()     {}
func (Int) IAmBuiltin()  {}
func (Int) Kind() string { return "Int" }

type Float struct{}

func (Float) IAmPlan()     {}
func (Float) IAmBuiltin()  {}
func (Float) Kind() string { return "Float" }

type Byte struct{}

func (Byte) IAmPlan()     {}
func (Byte) IAmBuiltin()  {}
func (Byte) Kind() string { return "Byte" }

type Char struct{}

func (Char) IAmPlan()     {}
func (Char) IAmBuiltin()  {}
func (Char) Kind() string { return "Char" }

type String struct{}

func (String) IAmPlan()     {}
func (String) IAmBuiltin()  {}
func (String) Kind() string { return "String" }

type Bytes struct{}

func (Bytes) IAmPlan()     {}
func (Bytes) IAmBuiltin()  {}
func (Bytes) Kind() string { return "Bytes" }

type Any struct{}

func (Any) IAmPlan()     {}
func (Any) IAmBuiltin()  {}
func (Any) Kind() string { return "Any" }

type Nothing struct{}

func (Nothing) IAmPlan()     {}
func (Nothing) IAmBuiltin()  {}
func (Nothing) Kind() string { return "Nothing" }

type Structure struct {
	Fields Fields
}

func (Structure) IAmPlan()     {}
func (Structure) Kind() string { return "Structure" }

type Fields []Field

type Field struct {
	Name   string
	GoName string
	Type   BuiltinOrRefPlan
}

type Inductive struct {
	Cases   Cases
	Default DefaultCase
}

type DefaultCase struct {
	GoKeyName   string
	GoValueName string
	Type        BuiltinOrRefPlan
}

func (Inductive) IAmPlan()     {}
func (Inductive) Kind() string { return "Inductive" }

type Cases []Case

type Case struct {
	Name   string
	GoName string
	Type   BuiltinOrRefPlan
}

type List struct {
	Element BuiltinOrRefPlan
}

func (List) IAmPlan()     {}
func (List) Kind() string { return "List" }

type Union struct {
	Cases Cases
}

func (Union) IAmPlan()     {}
func (Union) Kind() string { return "Union" }

type Tuple struct {
	Slots Slots
}

func (Tuple) IAmPlan()     {}
func (Tuple) Kind() string { return "Tuple" }

type Slots []BuiltinOrRefPlan

type Link struct {
	To BuiltinOrRefPlan
}

func (Link) IAmPlan()     {}
func (Link) Kind() string { return "Link" }

type Map struct {
	Key   BuiltinOrRefPlan
	Value BuiltinOrRefPlan
}

func (Map) IAmPlan()     {}
func (Map) Kind() string { return "Map" }

type SingletonBool defs.SingletonBool

func (SingletonBool) IAmPlan()     {}
func (SingletonBool) Kind() string { return "SingletonBool" }

type SingletonByte defs.SingletonByte

func (SingletonByte) IAmPlan()     {}
func (SingletonByte) Kind() string { return "SingletonByte" }

type SingletonChar defs.SingletonChar

func (SingletonChar) IAmPlan()     {}
func (SingletonChar) Kind() string { return "SingletonChar" }

type SingletonFloat defs.SingletonFloat

func (SingletonFloat) IAmPlan()     {}
func (SingletonFloat) Kind() string { return "SingletonFloat" }

type SingletonInt defs.SingletonInt

func (SingletonInt) IAmPlan()     {}
func (SingletonInt) Kind() string { return "SingletonInt" }

type SingletonString defs.SingletonString

func (SingletonString) IAmPlan()     {}
func (SingletonString) Kind() string { return "SingletonString" }

type Call struct {
	Fn Fn
	ID BuiltinOrRefPlan
}

func (Call) IAmPlan()     {}
func (Call) Kind() string { return "Call" }

type Return struct {
	Fn Fn
	ID BuiltinOrRefPlan
}

func (Return) IAmPlan()     {}
func (Return) Kind() string { return "Return" }

type Fn struct {
	Arg    BuiltinOrRefPlan
	Return BuiltinOrRefPlan
}
