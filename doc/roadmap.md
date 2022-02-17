# Roadmap for the decentralized protocol compiler, 2022

This roadmap details the first three major milestones that we see for the decentralized protocol compiler. Milestones included in this roadmap are squarely focused on adding immediate value to users, either in the form of new, requested features or in the form of significant performance improvements over existing tools.

For the time being, we are not yet planning (or including here) milestones that focus on transforming the protocol compiler itself into a polished tool for wide-spread self-serve public use. Such polishings — which include custom syntax, packaging and versioning, github integration, etc — will be scheduled after we gain some experience with our first adopters.

Milestones 2 and 3 below may be re-ordered. We will determine which is more critical to the PL ecosystem after polling/interviewing a larger audience of stakeholders.

## Milestone 1: An RPC protocol compiler for IPLD

The goal of this milestone is to accomplish a production-grade code generator for RPC APIs based on IPLD. We plan to support a core type system, which is a superset of the familiar IPLD schema, wherein each type uses a canonical representation to IPLD. Canonical representations will match chosen pre-existing IPLD Schema representations.

Key goals of this milestone are to establish flexible and extensible frameworks for compilation and code-generation that enable rapid and safe addition of new types and code-generation strategies.

Furthermore, this milestone introduces code generation of RPC code, which will initially target a networking stack of DAGJSON-over-HTTP, modelled after the stacks used by the Delegated Routing and Indexer projects within PL. However, the RPC framework will be extensible enabling users to implement their own RPC networking backends.

### Summary of included features:

- Supported types:
  - Primitive: Bool, Float, Int, Byte, Char, String, Bytes
  - User-defined: Link, List, Map, Structure, Inductive (previously known as Union), and Singleton
  - Functional types: Function, Call, Return, and Service
- An extensible code-generation framework that enables safe rapid prototyping
- Fully-static (no reflection) code-generation in Go:
  - Serialization based on code-generated implementations of the IPLD `Node` interface.
  - Deserialization based on parsing from a deserialized IPLD data model.
- RPC service code-generation using a DAGJSON-over-HTTP networking stack

### Timeline

We expect this milestone to be ready for production use by the end of Q1 of 2022.

### Stakeholders and applications

We have commitments from the Delegated Routing and Indexer projects in PL to be the first users of this milestone. Delegated Routing will enable all IPFS nodes to delegate routing to third parties, using a code-generated RPC API. Indexers will use the same framework to act as a third-party routing provider for the IPFS Hydra nodes.

## Milestone 2: Feature parity with IPLD schema and the fastest IPLD de/serializer

In this milestone, we plan to introduce the notion of _transforms_ which are a generalization of IPLD schema representations. Transforms are user-defined middleware, decoupled from types, which changes IPLD schema on-the-fly with zero-allocations. Combinging types with transforms enables users to de/serialize any specific IPLD representation strategy to a desired user-facing type schema. This feature will bring the protocol compiler to feature parity with IPLD schema.

Additionally, we plan to upgrade the protocol compiler's type deserialization strategy from parsing the IPLD data model (from Mileston 1) to implementing the native IPLD `NodeAssembler` interface. This will result in the fastest IPLD de/serialization code to date. Both encoding and decoding paths will be fully-static native IPLD code.

### Timeline

We expect this milestone to be ready for production use by the end of Q2 of 2022.

### Stakeholders and applications

At this milestone, we speculate that speed-sensitive users of IPLD schema would be interested in using the protocol compiler, because we expect that the removal of reflection would bring an order of magnitude improvement in speed and CPU utilization.

## Milestone 3: Lambda passing across network boundaries

Preliminary conversations with the Filecoin/FVM team have indicated that Filecoin Actors necessitate an RPC compiler, which additionally must support "passing callbacks" across network boundaries (between a client and a server).

This milestone targets this requirement by incorporating type-safe lambdas (a generalization of callbacks) in the protocol type system. Lambdas — which are references to function implementation instances — must support user-defined methods of refering to an implementation.

There are a few design alternatives in this space and we plan to converge on the right choice by close collaboration with Filecoin, FVM and actors developers.

### Stakeholders and applications

Our hope is that Filecoin actors will adopt the protocol compiler as the standard method of defining Filecoin actor interfaces.