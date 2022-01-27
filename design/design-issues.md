# Compiler design issues

We are dedicating this document to listing the controversial aspects of the protocol compiler design and implementation proposals.

_How to use this document:_ If you find issues in the design, list them here and explain why they are problematic. We wouldn't be able to address issues productively, if there is no specific justification.

The document contains a starter list of potential issues. But if you agree with any of them, you'd have to add some text explaining why.

## Design issues

This is a summary of prominent issues deserving discussion. We assume the reader is familiar with the actual Type System Design Proposal details before reading this section.

### Users do not need to choose representations, as in IPLD schema

*Why is this an issue?*

### Maps support keys of all types

*Why is this an issue?*

### Strings are lists of unicode characters

*Why is this an issue?*

## Implementation issues

### Adopt a new type system implementation strategy

*Why this is needed:*

The IPLD Schema type system implementation has two drawbacks:

- _Type definitions cannot be used as Go map keys._ Implementing recursive type compatibility checks requires memoizing intermediate results. Memoization requires using recursive type structures as keys. This is not possible in Go, when type structures have pointers or slices, as is the case with IPLD Schema.
- _Anonymous types are not supported well._ IPLD Schema enforces (and relies on) all types having names. This makes it very difficult/awkward to support anonymous types (which may result from user code or from intermediate computations in the type system).

Both problems can be resolved by adopting a different implementation strategy, where:
- No Go pointers or slices are used in type definitions
- Naming a type is modeled as a type itself, e.g. 

     type Named struct {
          Name string
          WhatIsBeingNamed Type
     }

*Why is this an issue?*
