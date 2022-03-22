# Representations of types

Edelweiss data types have _canonical_ representations in the IPLD data model. This document describes the representation of each type and illustrates it via examples based on the DAG-JSON encoding of IPLD values.
To learn about the semantics of types, please refer to the [User manual](slides/user-milestone1-slides.pdf) for Milestone 1.

For context, note that users of Edelweiss _can_ modify the ultimate representation of types by applying [_transforms_](transforms.md) to the canonical representations. Transforms are a feature [slated for Milestone 2](roadmap.md) of the Edelweiss compiler.

# Canonical representations

## Bool

Example type definition:
```go
Bool{}
```
Example value in IPLD DAG-JSON format:
```json
true
```

## Byte

Example type definition:
```go
Byte{}
```
Example value in IPLD DAG-JSON format:
```json
211
```

## Char

Example type definition:
```go
Char{}
```
Example value in IPLD DAG-JSON format:
```json
22345
```

## Int
Example type definition:
```go
Int{}
```
Example value in IPLD DAG-JSON format:
```json
123456789
```

## Float
Example type definition:
```go
Float{}
```
Example value in IPLD DAG-JSON format:
```json
123.456
```

## String
Example type definition:
```go
String{}
```
Example value in IPLD DAG-JSON format:
```json
"abc" /* valid UTF8 string */
```

## Bytes
Example type definition:
```go
Bytes{}
```
Example value in IPLD DAG-JSON format:
```json
{"/": { "bytes": "XXX" /* Base64 encoded binary */ }}
```

## Nothing
Example type definition:
```go
Nothing{}
```
Example value in IPLD DAG-JSON format:
```json
{}
```

## Link
Example type definition:
```go
Link{ To: Any{} }
```
Example value in IPLD DAG-JSON format:
```json
{"/": "XXX" /* Base58 encoded CIDv0 or Multibase Base32 encoded CIDv1 */}
```

## List
Example type definition:
```go
List{ Element: Int{} }
```
Example value in IPLD DAG-JSON format:
```json
[ 1, 2, 3 ]
```

## Map
Example type definition:
```go
Map{ Key: String{}, Value: Int{} }
```
Example value in IPLD DAG-JSON format:
```json
{ "abc": 123, "def": 456 }
```

Unfortunately, IPLD encoders generally do not support maps with non-string keys even though the IPLD data model does. To overcome this gap — short of waiting for such support — we plan (for Milestone 2) to represent such maps as IPLD lists of pairs.

Example type definition:
```go
Map{ Key: Float{}, Value: String{} }
```
Example value in IPLD DAG-JSON format:
```json
[ [ 123.456, "abc" ], [ 456.789, "def" ] ]
```

## Structure
Example type definition:
```go
Structure{
     Fields: Fields{
          Field{Name: "Abc", Value: Int{}},
          Field{Name: "Def", Value: Float{}},
     }
}
```
Example value in IPLD DAG-JSON format:
```json
{
     "Abc": 123,
     "Def": 456.789,
}
```

## Singleton

### Singleton Bool
Example type definition:
```go
SingletonBool{true}
```
Example value in IPLD DAG-JSON format:
```json
true
```

### Singleton Int
Example type definition:
```go
SingletonInt{123}
```
Example value in IPLD DAG-JSON format:
```json
123
```

### Singleton Float
Example type definition:
```go
SingletonFloat{123.456}
```
Example value in IPLD DAG-JSON format:
```json
123.456
```
### Singleton String
Example type definition:
```go
SingletonString{"abc"}
```
Example value in IPLD DAG-JSON format:
```json
"abc"
```

## Inductive (aka "tagged union")
Example type definition:
```go
Inductive{
     Cases: Cases{
          Case{Name: "C1", Value: Int{}},
          Case{Name: "C2", Value: Float{}},
          Case{Name: "C3", Value: String{}},
     }
}
```
Example values in IPLD DAG-JSON format:
```json
{ "C1": 123 }
```

```json
{ "C2": 456.789 }
```

```json
{ "C3": "abc" }
```

## Union (aka "untagged union")
Example type definition:
```go
Union{
     Cases: Cases{
          Case{Name: "C1", Value: Int{}},
          Case{Name: "C2", Value: Float{}},
          Case{Name: "C3", Value: String{}},
     }
}
```
Example values in IPLD DAG-JSON format:
```json
123
```

```json
456.789
```

```json
"abc"
```

## Service method calls

Example service type definition:
```go
Service{
     Methods: Methods{
          Method{Name: "Method1", Type: Fn{Arg: Int{}, Return: Bool{}}},
          Method{Name: "Method2", Type: Fn{Arg: String{}, Return: Float{}}},
     },
},
```

### Request

Based on the service definition (above), Edelweiss generates a type definition for a method call request. In this case:

```go
Inductive{
     Cases: Cases{
          Case{ Name: "Method1", Value: Int{} }
          Case{ Name: "Method2", Value: String{} }
     },
}
```

### Response

Based on the service definition (above), Edelweiss generates a type definition for a method call response. In this case:

```go
Inductive{
     Cases: Cases{
          Case{ Name: "Method1", Value: Bool{} }
          Case{ Name: "Method2", Value: Float{} }
     },
}
```
