---
marp: true
title: Decentralized Protocol Compiler
paginate: true
theme:  default
---

# **Edelweiss**: Decentralized Protocol Compiler
## Milestone 1 (MVP): RPC compiler for Go

_Petar Maymounkov_
petar@protocol.ai

---
# Roadmap

- **Milestone 1 (Q1 2022, MVP)**: RPC compiler for Go __(THIS TALK)__
- **Milestone 2 (Q2 2022)**: Feature parity with IPLD schema, performance squared, policies and transforms
- **Milestone 3 (Q3 2022)**: Lambdas across networks, Filecoin/FVM actors API

- **Milestone X**: Multiple target languages, packaging, github integration, doc generation, cli generation, type interoperability checks at compile time, etc.

---
# Plan

- Definitions
- Types (semantics, representations, generated runtime code)
- Type interoperability

---
# Significance of types

1. Semantics of data (agnostic to programming language)
2. Representation of data in IPLD Data Model (encoding/decoding)
   Note: _"Transforms" (introduced later) can alter representation._
3. Representation of data in user's programming language

---
# Type

- **Builtin:** Bool, Float, Int, Byte, _Char_, _String_, _Bytes_
- **Parametric:** Link, List, Map, Structure, _Inductive_, _Singleton_, _Union_
- **Special:** Any, Nothing
- **Functional:** _Function_, _Service_, _Method_

---
# __Definitions__

---
# AST interface

User builds type definition AST in Go.
_Syntax to come later, when the language matures._

<br>

```go
import "github.com/ipld/edelweiss/defs"

Types{
   Named{ Name: "MyLink", Type: Link{To: Int{}} }, // link to int
   Named{ Name: "MyList", Type: List{Element: Any{}} }, // list of any
   ...
}
```

---
# Named definitions

Wrap a type definition in `Named{}` 
<br>
```go
Named{
   Name: "MyStructure"
   Type: Structure{
      Fields: Fields{
         Field{ Name: "Foo", Type: Int{} },
         Field{ Name: "Bar", Type: Any{} },
      }
   }
}
```

---
# Inline definitions

```go
Named{
   Name: "MyStructure"
   Type: Structure{
      Fields: Fields{
         Field{
            Name: "MyFieldFoo",
            Type: List{Element: Int{}}, // <-- inline type definition, list of int
         },
      }
   }
}
```

Inline types are named generically, e.g. `AnonListXXX`

---
# Named inline definitions

```go
Named{
   Name: "MyStructure"
   Type: Structure{
      Fields: Fields{
         Field{
            Name: "MyFieldFoo",
            Type: Named { // <-- named inline type definition, list of int
               Name: "MyInlineListOfInt",
               Type: List{Element: Int{}},
            },
         },
      }
   }
}
```

---
# Type references

Use `Ref{Name: "TypeName"}` to refer to any `Named` type
<br>
```go
Named{
     Name: "MyList",
     Type: List{ Element: Int{} }
}

Named{
     Name: "MyListOfList",
     Type: List{ Element: Ref{Name: "MyList"} }
}
```

---
# __Non-parametric types__

---
# Builtin types

Definitions:
```go
Bool{} // represented as IPLD bool
Float{} // represented as IPLD float
Int{} // represented as IPLD int
Byte{} // represented as IPLD int
Char{}
String{}
Bytes{}
```
Runtime implementations in package `github.com/ipld/edelweiss/values`:
```go
type Byte byte
// etc.
```

---
# Char

Semantically:
- a character is not an integer

Representationally:
- encoded as an IPLD integer which is a valid UTF8

Programmatically:
- Implemented by `type Char rune` in package `edelweiss/values`

---
# String

Semantically:
- `String{}` is equivalent to `List{Element: Char{}}`

Representationally:
- Encodes to IPLD string
- Decodes from a UTF8 IPLD string or the IPLD encoding of `List{Element: Char{}}`

Programmatically:
- Implemented by `type String string` in package `edelweiss/values`

---
# Bytes

Semantically:
- `Bytes{}` is equivalent to `List{Element: Byte{}}`

Representationally:
- Encodes to IPLD bytes
- Decodes from IPLD bytes or the IPLD encoding of `List{Element: Byte{}}`

Programmatically:
- Implemented by `type Bytes []byte` in package `edelweiss/values`

---
# __Special types__

---
# Nothing

Semantically:
- `Nothing{}` holds no value

Representationally:
- Encodes as IPLD nothing

Programmatically:
- Implemented by `type Nothing struct{}`

E.g. use in conjunction with `Inductive` types to describe enumerations.
E.g. use in conjunction with `Union` types to describe optional values.

---
# Any

Semantically:
- `Any{}` can hold any IPLD value
- IPLD kinds are in one-to-one mapping with types in this type sytem:
  - IPLD bool, int, float, string, bytes map to `Bool{}`, `Int{}`, `Float{}`, `String{}`, `Bytes{}`
  - IPLD link maps to `Link{To: Any{}}`
  - IPLD list maps to `List{To: Any{}}`
  - IPLD map maps to `Map{Key: Any{}, Value: Any{}}`
  - IPLD nothing maps to `Nothing{}`

Programmatically:
- Implemented by `type Any struct{ Value }`  where `Value` is an interface

---
# __Parametric types__

---
# Link

Semantically:
- `Link{To: TYPE_DEF_OR_REF}`

Representationally:
- Encodes as IPLD link

Programmatically:
- Code-generated Go `struct` which holds a `Cid`

Use `Link{To: Any{}}` when the link target is of unknown type.

---
# List

Semantically:
- `List{Element: TYPE_DEF_OR_REF}`

Representationally:
- Encodes as IPLD list

Programmatically:
- Code-generated Go alias for a slice type

---
# Map

Semantically:
- `Map{Key: TYPE_DEF_OR_REF, Value: TYPE_DEF_OR_REF}`

Representationally:
- Encodes as IPLD list of key/value pairs or an IPLD map, if the key is a string

Programmatically:
- Code-generated Go slice of key/value pairs or a Go map, if the key is a string

---
# Structure

Semantically:
- `Structure{Fields: []Field}` where `Field` is
```go
type Field struct {
     Name string
     Value TYPE_DEF_OR_REF
}
```

Representationally:
- Encodes as IPLD map

Programmatically:
- Code-generated Go `struct`
