
# Types, transforms and representations

An explicit objective of the Edelweiss language and compiler is to give the user full control over the representation of data on the wire. This is necessary both in order to provide full backwards support for pre-existing protocols based on IPLD Schema, as well as to accommodate use cases of communication with non-Edelweiss counterparties, like manually coded JavaScript-based JSON programs.

To accomplish this we adopt the following paradigm.

## Canonical type representation

Every Edelweiss type has a canonical representation strategy to (and from) the IPLD data model.
In other words, each Edelweiss type is associated with a reversible (i.e. bijective) function that maps typed values to IPLD data model values (encoding) and back (decoding).

## Transforms

In addition to types, we introduce the notion of transforms. A transform is also a reversible transformation that maps IPLD data model values to IPLD data model values. Therefore, transforms can be viewed as "middleware" which can be chained after the canonical representation of types. This is illustrated below.

![](transforms/model.svg)

## Compile-time type awareness of transforms

One important design aspect of transforms is that: _transforms_ are aware of the Edelweiss type that they are applied to. In other words, while functionally a transform maps one IPLD data value into another, it is aware of the Edelweiss type that underlies. This property is crucial to enable the implementation of generic reversible transforms, even in cases when the manipulated IPLD data model values do not incorporate type information in themselves.

This is illustrated in the following example:

![](transforms/example.svg)

- On the left, the user (a developer using a chosen programming language) manipulates typed Edelweiss values, in this case — a structure.
- In the middle, the user's structure is canonically represented as an IPLD map.
- On the right, when the "tuple" transform is applied to the canonical representation, it is converted into an IPLD list.

The encoding path (left-to-right) in this diagram is possible using a generic — i.e. unaware of the Edelweiss type structure — "tuple" transform implementation. However, the decoding path (right-to-left) is clearly not.

By providing the Edelweiss type information to the "tuple" transform, we enable it to implement the decoding path as well. More generally, however, we enable the implementation of different _kinds of generic_ transforms:

> A generic transform is an object which, given an Edelweiss type, can produce an encoder and a decoder function corresponding to that type.

In the context of the compiler, the Edelweiss type information is available at _compile time_. Knowing this, the compiler code-generates the transform's encoder and decoder functions for the user to use in their own code.

## Decoupling of types and transforms

Note that types and transforms are _decoupled_ in the following sense. At compile-time, the compiler generates separate implementations of:
- the user's type (and its canonical representation encoder and decoder), as well as
- the encoder and decoder for any number of desired transforms

At run-time, this enables the user to use different combinations of transforms with the same type in different places or times in their program.

### Comparison with IPLD Schema

If you are a user of IPLD Schema, it is instructive to notice the difference.

IPLD Schema _couples_ the user's type to a chosen representation, in the sense that it code-generates a single object that implements the data type and one specific representation strategy.
This precludes the user from using different representation strategies with the same type, without passing through a costly runtime conversion.
