
# Edelweiss: Decentralized Protocol Compiler

Edelweiss is a code-generating compiler. Currently, it supports:
- a comprehensive type system for modeling data
- service and method definitions
- Go language code-generation of data types and associated fast, static encoders and decoders to and from the IPLD data model
- Go language code-generation of RPC services based on a DAGJSON-over-HTTP networking stack
- the ability to rapidly add user-defined code-generating backends for custom RPC networking stacks

For a detailed longer-term roadmap and planned features refer to the [Protocol Compiler Roadmap 2022](doc/roadmap.md).

# Documentation

The current state of the language and how to use the compiler is covered in the [Edelweiss for users](doc/slides/user-milestone1-slides.pdf) slides. Examples of the canonical representation of Edelweiss types are provided in [Representations of types](doc/representations.md).

A [complete working example](examples/greeting-service/) — which defines a greeting service API, and implements a sample client and server cli tools — is provided in this repo. The [resulting generated code](examples/greeting-service/api/proto/proto_edelweiss.go) is also included in the repo for your viewing convenience.

# Projects using Edelweiss

Currently, Edelweiss is in the process of being onboarded as the [RPC framework for the Delegated Routing API](https://github.com/ipfs/go-delegated-routing/pull/11) of IPFS and Hydra nodes.
