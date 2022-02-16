# TYPE SYSTEM DISCUSSION

## OBJECTIVES

At the highest-level, we perceive that the main design objectives for a protocol compiler are:

(O1) _Facilitate the decentralized development of protocols._ In practice, this means that different developer teams should be able to work independently and asynchronously on the client and the server side of a network protocol, while maintaining runtime interoperability.

(O2) _Facilitate the development of middleware._ In practice, middleware services aim to understand a subset of a more complex network protocol, and therefore they necessitate programming tools to define that subset rigorously.

We think that both of these objectives can be captured entirely within a type system abstraction, which shields the developer from the complex details of on-the-wire interoperability. Specifically, to address (O1) we propose a type system which captures all on-the-wire interoperability concerns within a formal type relationship which is understood and enforced by the compiler. To address (O2), we make sure that the type system incorporates types that describe optionality and opaqueness in a flexible manner.

## INTRO

The type system in a protocol language provides a framework for describing data and functions.

Types serve three related purposes:
(1) Representation: They dictate the serialized representation of data and functions
(2) Semantics: They dictate the developer-facing representation and properties of data and functions
(3) Interoperability: They dictate which data and functions can be used in place of others

To fully define a type system, we need to address all three aspects of each individual type. We set out to accomplish this in this document.

## INTEROPERABILITY AS A TYPE RELATION

Interoperability ultimately tries to address the question:

     Can a client using protocol definition C make a network call to a server using protocol definition S?

There are a various reasons why C and S may be different:
- client and server are compiled at different times and use different versions of the same protocol definition.
- a server may need to serve clients with different capabilities. E.g. the PL indexer service must process lookup requests from delegated routing clients and finder clients. One lookup response must capture common and different information sought after by each client.

We propose to capture interoperability at the granularity of individual data and function types in a protocol definition.

To this end, we introduce a relation "readable as" between types, which reflects on-the-wire interoperability between individual protocol objects like data or functions. Having this relation ultimately allows a compiler to:
- verify that two different protocol definitions are interoperable
- generate RPC and serialization code which is guaranteed to fulfill the interoperability properties of types (data and functions) at runtime

Specifically, we would like to have:

	"type1 is readable as type2" if and only if "values serialized as type1 can be deserialized as type2" for all supported on-the-wire serializations (dagjson, dagcbor, etc).

In practice, it is too difficult to ensure that a large and growing set of type serialization strategies (dagjson, dagcbor, etc) each correctly implement the desired interoperability relations. To resolve this problem, we choose to represent (i.e. encode and decode) typed protocol values as IPLD Data Model values, and benefit from the IPLD infrastructure which can serialize the IPLD Data Model to different serialization formats.

This enables us to work with an equivalent and simpler definition of the type relation:

	(X) "type1 is readable as type2" if and only if "values encoded in the IPLD data model of type1 can be decoded as values in the IPLD data model of type2".

We use the shorthand "type1 > type2" to denote "type1 is readable as type2". This notation is justified by the intuition that type1 must have more information than type2, if a value of type1 is to be used in places where a value of type2 is expected.

We use the shorthand "type1 = type2" to denote that both "type1 > type2" and "type1 < type2" hold.

Definition (X) sets the tone for the remainder of this document. Specifically, it makes type relations the primary focus of design discussions. Whereas, any choice of type representations is acceptable, as long as is does not violate (X).

The rest of this document focuses on justifying our choices of type relations (where the justification criteria is singularly aimed at facilitating evolution of protocol definitions while maintaining interoperability). We mostly avoid prescribing reresentations to types in what follows, except in cases where it helps intuition.

Our intention is that once type relations are set in stone, choosing corresponding representations that meet type constraints and optimize for side concerns (like readability or overhead) should be an easier and self-contained task.

## THE LEGACY OF IPLD SCHEMA

In IPLD Schema, each IPLD type is defined by a pair of a type kind (aka schema kind) and a type representation:
     (T1) the type kind describes data schema, independent of representation
     (T2) the type representation describes how the schema is mapped to an IPLD data model

Footnote: Unfortunately, (T2) is not entirely true. There are corner cases where IPLD representations dictate serialization beyond the IPLD data model. One such corner case, for example, is the stringpairs representation of maps.

Coupling definition of schema with definition of representation is a design choice unique to IPLD Schema. This choice has been motivated by a specific use case: Writing IPLD schemas to retro-fit to pre-existing serialization formats. In this case, having options of representation is helpful to the user.

A protocol compiler, on the other hand, has objectives that are odds with options in representation. A protocol framework aims to facilitate the user by removing the choice of representation, and picking the "most interoperable" representations for each type.

The problem that a protocol language is trying to solve is to enable users to author forward-looking protocols, which are safely and maximally interoperable with future versions or variants. Interoperability tomorrow is dictated by choice of representation today. Therefore a protocol compiler aims to pick a single "best" representation for each type kind, so that the resulting choices for all type kinds work well _as a whole_, from the standpoint of interoperability.

We dedicate the next section to a detailed discussion of requirements for interoperable representations, motivated by applications.

The key point here is that, unlike IPLD Schema, a protocol language aspires to remove the representation component from the user-facing type definition and instead, represent the interoperability of representations in the form of a type compatibility relation. I.e. the "readable as" relation defined earlier.

## INTEROPERABLE REPRESENTATIONS

In designing type representations, there are a few requirements that arose from our experience with PL applications (like delegated routing and cache middleware, for instance):

### Composite types should have _distinguishable representations_ (R1)

For instance, a struct should not be decodable from the encoding of a map; a list should not be decodable from the encoding of a struct. This rule is generally violated by existing IPLD schema representations, as this is not a design requirement in the IPLD schema use case. For instance, the dagjson encoding:

     [["f", "a"], ["g", "b"]]

Could be decoded either as an IPLD schema struct type:

     type X struct {
          f String
          g String
     } representation listpairs

Or as an IPLD schema list type:

     [[String]]

In other words, decoding of types is ambiguous when a schema definition is absent.

On the other hand, in middleware applications (like caches), a decoder must be able to decode values whose schema is not known ahead of time. For instance, a cache proxy for content routing requests may be aware of the top-level schema of content lookup requests and responses, without being aware of the schema of specific query types or provider record types. Nevertheless, the cache should still be able to decode, compare and store these values.

### Representations should leave room for new type kinds (R2)

The representation of kinds should be crafted in a way that leaves room for new kinds to be added in the future, and represented in ways that are not in conflict (of ambiguity) with existing kinds. 

For instance, we think (and argue in a later section) that a function call — which is required for RPC semantics — should be modelled as a new type kind. More generally, languages need room to evolve. Case in point, protobuf has added a few new type kinds since v1.

## TYPE SEMANTICS

This section tries to provide a complete list of required types, discuss their applications to interoperability and define their formal interoperability relations.

At high level, we describe the following type taxonomy:

     Boolean, Integer, Float, String, Byte
     Struct, Map, List
     Function
     Singleton
     Link
     Union
     Any

Below we use pseudo-notation to discuss types, because in some cases a corresponding IPLD Schema notation is either missing or not compatible with our intentions. This pseudo-notation is used for the sake of facilitating this discussion and is not meant as a syntax proposal.

### PRIMITIVES

The IPLD data model primitives (Boolean, Integer, Float, String), excluding Bytes, are a good foundation.

Along with most protocol and programming languages, we view these types as mutually incompatible.
Formally, if P is primitive and Q is any type:

     "P is readable as Q" if and only if "P and Q are the same primitive type".
     "Q is readable as P" if and only if "P and Q are the same primitive type".

We also need to accommodate type arrays of bytes, which are widely used in IPLD data model, as the type Bytes.

Our recommendation is that from a language point of view, bytes should be treated as a list of individual byte values. Note that this does not preclude byte lists from being represented in a custom manner (e.g. as IPLD data model type Bytes) that differs from how other types of lists are represented (see Lists in a section below).

In order to accommodate bytes as a list of individual byte values, we would have to introduce a primitive type corresponding to an individual byte into the language, as well.

### STRUCTURES

We set forth a design tenet, which we think is important in the context of decentralized development:

     (X1) Type names and type representations are not dependent on each other.

This rule has multiple benefits:
- It removes any ambiguities about the interaction between type aliases (a programming aid) and type interoperability (a runtime property)
- It affords developers the flexibility to maintain multiple variants of a type definition in the same source environment, under different type names
- It ensures that package refactoring, which alters the fully-qualified names of types, does not have unexpected effects on the representations of user types defined inside packages.

So, for example, the following two types define the same protocol (i.e. the same representation):

     struct X
          s : String
     end

     struct Y
          s : String
     end

In terms of type interoperability relations, both X < Y and Y < X hold.

More generally, the representation of structures should induce the following natural type relation:

     Let P and Q be structure types. Then P > Q, if for every field in Q named F of type S, there is a field in P of the same name and type T and T > S.

For instance, the default IPLD Schema representation of structures induces exactly this relationship. 

There are applications where protocol designers will desire to add additional rigidity to type interoperability by placing a name or a magic value in data representations. We can accommodate this use case, without violating (X1), by means of singleton types (which are described in a later section).

### LISTS

The list type models a list of values of a given element type. The type interoperability relation for lists is natural:

     - A list with elements of type P is readable as a list with elements of type Q, if and only if P is readable as Q.
     - Lists are not readable as other types, and other types are not readable as lists.

### LINKS

Like IPLD Schema, we view links as references to content, which is also within the type system.
Accordingly, we view links as parametric types. In IPLD Schema notation, they are written as `&TypeOfContent`. In our pseudo-notation here, they could be written as `Link{TypeOfContent}`.

Link interoperability could come in two forms. First, links that point to compatible content should be compatible themselves. This is captured as the relation:

     Link{P} > Link{Q}, if P > Q.

Second, we propose that it is useful to allow links and content to be interoperable, in the following sense:

     - A link can be passed where content is expected. Formally, Link{P} > Q, if P > Q.
     - Content can be passed where a link is expected. Formally, P > Link{Q}, if P > Q.

Both of these rules extend interoperability without sacrificing correctness.

### FUNCTIONS

An RPC framework needs to specify how function calls are expressed on the wire.
We propose that function calls are modelled as the values of a dedicated functional type kind, because their interoperability semantics fit naturally in the "readable as" relation.

Suppose a function declaration is a type kind with three parameters: a name f, an argument type A and a return type R. Let's refer to this type as "f(A)->R".

From an interoperability standbpoint, note that a client can invoke f(A1)->R1 on a server that provides f(A2)->R2, as long as:
     
     A1 is readable as A2, and R2 is readable as R1.

Therefore, it is natural to model a function as a type, where "f(A1)->R1 is readable as f(A2)->R2" if and only if "A1 is readable as A2" and "R2 is readable as R1".

Function calls are the values of the functional type. A function call specifies a concrete argument value.
Analogously to other type values, function calls would have a distinguishable representation.

Thus, an RPC request can be described by a single value in the type system. For instance, using pseudo-notation:

     GetP2PProviders({Key: "123"})

There are various benefits from this integration of function calls with the type system:

- Middleware. Implementing middleware solutions amounts to decoding values in the type system. A lot of applications can be implemented as generic middleware, which is either unaware of schema or aware of portions of domain-specific schemas. A few examples:
  - RPC call caching
  - load-balancing, sharding
  - authentication
  - access control

- Service composition. Treating function calls as values enables the expression of service composition in RPC calls. For instance, the value

     GetP2PProviders({Key: Resolve({URL: "http://madonna.nft")})

is calling for the providers of a given key, whose value is the result of resolving the URL "http://madonna.nft".

Service composition has many applications in the PL ecosystem, which are beyond the scope of this document. They are covered by our prior work on "Scaling Routing" and "Composable Routing". One notable benefit is that service composition allows for finer-grain caching at middleboxes. In a [prior work on scaling routing](???), we demonstrated that combining the DHT with service composition and caching can result in order of magnitude speed up in IPFS content delivery.

### MAPS

A map is a set of key/value pairs with unique keys. The keys and values must be typed.
Unlike lower level type systems, like JSON or IPLD Data Model, we think that the maps of a protocol language should allow both keys and values to be of any type within the protocol type system.

We have two arguments for it:

- The purpose of a protocol type — in this case, maps — is to describe a property of the data itself (independent of programming technology) — in this case, the uniqueness of keys. Since uniqueness is well-defined across all other type kinds, it seems unnatural to restrict the type of keys in any way. Perhaps more importantly, there are many applications where maps with composite key types are the correct abstraction (e.g. a set of typed AST trees).

- The programming languages which will interact with our protocols have varying capabilities, and thus we don't think they should be used as the deciding factor. Julia (along with all functional languages), for instance, has complete native support for dictionary keys with composite types. Go, on the other hand, has partial support: map key values can be composite, as long as pointers and slices are not used. Java Script has no support for composite key types at all: keys must be strings.

We recognize that implementing code-generators targetting Go (or JS) for maps with non-string keys requires more effort. It is nevertheless possible to accomplish this as an incremental extension to a codegen that supports string-only keys, like bindnode.

### ANY

Enabling middleware is a key goal in the protocol language. The main functionality required by a middlebox is to be able to ingest data with unknown schema. This motivates the inclusion of an "any" type into the system.

The "any" type has the following interoperability:
     - T > Any, for all T.
     - Any > T, if and only if T = Any.

We note here that the requirement (R1) — that type kinds have distinguishable representations — is what makes it possible to have an "any" type. Distinguishable kind representations enable the decoding of values whose schemas are not known ahead of time.

### UNIONS

Union types play a special role in a protocol language, because they provide a unique opportunity to shape interoperability.

From a programmer's standpoint, a union is a collection of cases, each distinguished by a unique name.
Each case is associated with a type. Case types are unrestricted, however there should be no "shadowed" cases. In pseudo-code, we might write a union as:

     union MyUnion
          Name1 : Type1
          Name2 : Type2
          Name3 : Type3
     end

There are different strategies to represent a union — aka to encode/decode a union to an IPLD Data Model value. Each strategy induces different interoperability properties of the union type, as reflected in the "readable as" relationship.

We propose a strategy, which is modelled after the "sequence composition" rule in Parsing Expression Grammars. We argue that this strategy offers a flexible building block for protocol development.

We propose the following representation strategy for unions:
(U1) when encoded, a union value is represented as the representation of its active case value.
(U2) when decoded, a union value is set to the first case, according to the order of case definitions, whose value can be decoded (recursively) successfully from the input. In other words, the decoder attempts to decode the type of each case in order until one succeeds.

Notably, the union itself has no footprint in the representation, only the active case type does.

This representation strategy induces the following simple type interoperability relationships:

- Let P and Q be two union types. Then:

     P > Q, if for every case in P of type S, there is a case in Q of type T with S > T

- Let T be any type and P be a union type. Then:

     T > P, if P has a case with type S such that T > S
     P > T, if for every case in P of type S, S > T

There are a few benefits to this particular design of unions.

- This union type facilitates the natural growth of protocols. We've witnessed that it is not uncommon for a protocol to evolve from supporting one type of application objects to supporting multiple. For instance, V0 of the Delegated Routing protocol may support only one type of provider records, which might be captured as in this protocol snippet (in pseudo-notation):

     struct GetP2PProvidersRequest
          Key : Bytes
     end

     struct GetP2PProvidersResponse
          Providers : List{Provider}
     end

     struct Provider // in V0 all provider records are of this type
          PeerID : Bytes
          Multiaddresses : List{Bytes}
     end

Later on, a new type of provider records is added, e.g. a miner, resulting in protocol V1. An updated version of the protocol might look like this:

     // unchanged

     struct GetP2PProvidersRequest
          Key : Bytes
     end

     struct GetP2PProvidersResponse
          Providers : List{Provider}
     end

     // changed

     union Provider
          Peer : PeerProvider
          Miner : MinerProvider
     end

     struct MinerProvider
          MinerID : Bytes
     end

     struct PeerProvider
          PeerID : Bytes
          Multiaddresses : List{Bytes}
     end

Note that in V0, Provider is a struct, while in V1 provider is a union. Nevertheless, a server using V0 of the protocol can serve requests from a client using V1. In the reverse situation, if the client is using V0 and the server is using V1, we still retain the maximum possible interoperability: The server will be able to decode and serve all client requests. The client, on the other hand, will be able to decode responses partially: It would recognize all MinerProvider records as undecodable (and pass them to the programmer out of band as generic values). After dropping the unrecognizable records, the client will be able to decode the remaining records and store them in the V0 structures as desired. (All of this would be handled by the compiler automatically.)

It is instructive to contrast this zero-footprint representation of unions with the `keyed` representation of unions in IPLD Schema. You will notice that, in general, when a union has a non-zero footprint in the representation, it is not possible to accommodate protocol growth "in the middle of the schema", as shown above.

- This union type can be used in conjunction with singleton type (described below) to express enums and optional values.

### SINGLETONS

Singleton types are a natural addition to the type system, which solves two problems:
- In conjunction with unions, it enables the expression of enums, alleviating the need for a dedicated type
- It provides a mechanism for controlling protocol interoperability, by representing "magic values" on the wire.

A singleton type can hold only one value, which must be of one of the primitive types.
The representation of a singleton type must induce the following type interoperability:

     - If S is a singleton type, then no type (other than S) is readable as S.
     - Singleton S of primitive type P is readable as P.

Singleton types can be used to embed required "magic values" in protocols. For instance, (in pseudo-notation):

     struct Named
          name : "abc"
          other : Bool
     end

Valid representations of this struct must have the field `name` always equal to `"abc"`.

Enums can be modelled as unions of singleton types. For instance, (in pseudo-notation):

     union MyEnum
          Case1 : "abc"
          Case2 : "def"
     end

Optional values can also be modeled via unions and singletons, for example:

     union OptionalAddress
          Some : Bytes
          None: "none"
     end

     struct PeerRecord
          PeerID : Bytes
          Address : OptionalAddress
     end

### STRINGS

A unicode character is a category of its own, having no semantic relationship to other primitive types. While many legacy languages equate characters to integral types, this is a coincidental relation arising from a specific way to represent a character. As a result, we think that a unicode character is appropriately modelled by a dedicated primitive type, let's call it `Char`, which is not interoperable with other primitive types. 

Formally, if P is any non-Char primitive type, then P is not readable as Char and Char is not readable as P.

Representation-wise, there are two possibilities:

(C1) For practicallity, a single Char could be encoded/decoded as an IPLD Integer. This would introduce the semantically undesirable relationships Char > Int and Char < Int. However, their impact of introducing protocol incompatibilities in practice is low.

(C2) Alternatively, a stricter approach would use a Char IPLD representation which cannot be confused with any other type. This would require wrapping a char in an IPLD struct, which adds space overhead. On the other hand, individual chars are rarely used in protocols and if so a large fraction of representation overhead is alleviated by compressing algorithms.

Lists of Unicode characters are, by far, the most commonly used data types. And so it is justified to give them a type name, `String`. However, in order to make sure that semantics are clearly reflected, `String` should be no more than an alias for `List{Char}`.

We have already postulated that a `Byte` is not interoperable with a `Char`:

     Neither `Byte > Char`, nor `Char > Byte`.

Appealing to the relations for lists, this implies:

     Neither `List{Byte} > List{Char}`, nor `List{Char} > List{Byte}`.

Substituting for aliases results in:

     Neither `String > Bytes`, nor `Bytes > String`.

This is what we want, because from a data semantics point of view. An unconstrained list of bytes has no canonical interpretation as a list of characters, and vice versa. Indeed, many modern high-level languages (which focus on correctness of data semantics, as opposed to representational convenience) deem strings and bytes as not interoperable (e.g. Julia, Lean, Rust, etc).

Now let's turn our attention to the IPLD representation of strings and bytes.

We have already postulated that `String = List{Char}`, which implies that `String` must be decodable from an encoding of `List{Char}`. However, type relations do not preclude a `String` (as well as, `List{Char}`) from also being encodable/decodable to/from an IPLD String that holds a valid unicode encoding, as well. (A similar logic applies to bytes.)

As a result, `String = List{Char}` (respectively `Bytes = List{Byte}`) can be encoded to an IPLD String (respectively, IPLD Bytes) and decoded from both an IPLD String (respectively, IPLD Bytes) and the IPLD representation of `List{Char}` (respectively, `List{Byte}`).
