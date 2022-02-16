# PACKAGING SYSTEM FOR A PROTOCOL COMPILER


## SUMMARY

- We propose requirements for the packaging system of a decentralized protocol compiler.
- We argue that the design of the Go packaging system meets our requirements. 
- We propose an abstraction that allows us to build the packaging system as an independent reusable component.

This document's intention is to provide a framework for a concrete and rigorous discussion, which leads to improvements, elaborations and final agreement on a packaging system design.

## GOALS

We put forth the following high-level requirements for a packaging system, intended for decentralized development of protocols:

(G1) Developers must be able to utilize the protocol compiler from different programming languages (e.g. Go, JavaScript, Julia). By implication, the protocol packaging system must be agnostic to the packaging system of the programming language. In particular, it must be agnostic to whether it is used in a mono-repo or a multi-repo environment.

(G2) Packages must be able to depend on packages in other repos on the Internet. Furthermore, all build considerations such as resolving and fetching dependencies should be streamlined and automated by the compiler.

## RELATED WORK

We make references to these articles:

[1] Go versioning and the drawbacks of Rust versioning
[2] Minimal version selection

## DESIGN

The design spec here covers the user-facing aspect of the packaging system. The design is mostly identical to that of Go, but we choose to spell it out explicitly to facilitate discussion and modification.

What is the rationale for this design choice?

First and foremost, this design meets our goals (G1), (G2) and (G3). Arguments for (G1) and (G2) are partially addressed by use case examples in a later section. Regarding (G3). While the Go language does not support cyclical package dependencies, the packaging system itself does. This is the case because package resolution happens before compilation and does not depend on the contents of source files.

We identify the following user-facing properties of the packaging system:

(D1) Packages are named after the URL of a repo directory. For instance,

	"github.com/protocol/routing/cid"

(D2) A package comprises all source files with extension ".proto" residing in the package directory.

(D3) In source, dependence on a package is codified with an import statement which specifies the dependent package's name as well as its local alias. Here are examples of import statements:

	import "github.com/protocol/routing/cid" as cid

	import (
		cid "github.com/protocol/routing/cid"
		mh "github.com/protocol/routing/multihash"
	)

(D4) Local aliases are used to refer to objects defined in imported packages, e.g.

	cid.CID

refers to "CID" defined in package "github.com/protocol/routing/cid"

(D5) A directory tree of packages can be organized into a module by placing a module definition file, "proto.mod", in the root directory.

(D6) A module definition file specifies the minimum acceptable version of each package imported by the source files in the module. This is accomplished using the familiar Go syntax:

	require (
		github.com/protocol/routing v0.3.0 # specific github release
		github.com/filecoin/indexer 745127c # specific github commit
		github.com/libp2p/formats /Users/petar/src/github.com/libp2p/formats # local repo
	)

## VERSIONING AND RESOLUTION

Package resolution (aka version resolution) is the task of determining versions for each module (and package), such that all module versioning requirements in a compilation (which may include multiple modules) are satisfied simultaneously. There are many designs for package resolution in prior art.

We think that Go's approach, called Minimal Version Selection, is most flexible and understandable.
This approach is driven by two user-facing design tenets:

(V1) Import uniqueness rule:
In a compilation (which may comprise multiple modules), every import statement for a package name should resolve to the same concrete implementation (i.e. version of the package).

(V2) Import compatibility rule:
Packages from a newer version of a module should work in place of packages from an older version.

(V3) Gradual repair rule:
It should be possible for a module to depend simultaneously on two different versions (or variants) of a package.

The algorithm for implementing Minimal Version Selection is given in [2]. This algorithm resolves every package in a compilation to the minimal version that fulfills every module’s requirements.

## PACKAGES AND WIRE REPRESENTATIONS

In our context, the packaging system will serve a language for defining protocol types and functions and their relation to on-the-wire representation. We expect that the protocol language and its type system are yet to evolve.

It is therefore worth pointing a rule that language design must follow, in order to ensure that developers can reorganize the packaging of their projects without affecting the interpretation of protocols defined in source files:

On-the-wire representation of protocols (data types and function calls) is not dependent on package names.

## USE CASE EXAMPLES

Here are a few examples, aiming to justify that the packaging system meets the goals (G1) and (G2).

### EXAMPLE 1: DEDICATED PROTOCOL REPO USED BY A GO APP REPO

- Multiple protocol packages are located in the same repo, using a shared "proto.mod" file.
For instance, the repo "github.com/protocol/routing" may contain the following directory structure:

	proto.mod
	routing/find.proto
	routing/provide.proto
	multiformats/cid.proto
	multiformats/multihash.proto

- A Go application utilizes generated code for the definitions in the protocol repo.
For instance, the repo "github.com/ipfs/go-delegated-routing" may need to generate Go code implementations of the definitions in "github.com/protocol/routing". To accomplish this,
	- Create a package inside "github.com/ipfs/go-delegated-routing" dedicated to code bindings, e.g. "github.com/ipfs/go-delegated-routing/proto"
	- Inside that package, include a "gen.go" file with a go generate directive, akin to:

		// go:generate proto gen --pkg=github.com/protocol/routing --version=v0.3.1

	In other words, the go generate directive directly specifies the package name and implementation version of what is being generated. The protocol compiler takes responsibility to resolve and provision all dependencies.

### EXAMPLE 2: A GO APP REPO INCORPORATING PROTOCOL DEFINITIONS

Alternatively, ".proto" definitions can be colocated inside a Go repository. For instance, one might use a dedicated directory (inside a Go repo) for proto definitions and code bindings generation, as in:

	github.com/ipfs/go-delegated-routing/proto/find.proto
	github.com/ipfs/go-delegated-routing/proto/provide.proto
	github.com/ipfs/go-delegated-routing/proto/proto.mod
	github.com/ipfs/go-delegated-routing/proto/gen.go

Where the contents of "gen.go" is:

	// go:generate proto gen --pkg=github.com/ipfs/go-delegated-routing/proto

## MODULAR INTERFACE TO PACKAGING

The reader may notice that the packaging system described here (effectively Go's system) is independent of the contents of the source files. In that sense, it is agnostic to language semantics. In particular, it could also be applied retro-actively to existing languages like IPLD schema or protobuf, for instance.

Our goal here is to abstract the packaging system as a reusable software component.

The packaging system performs one key function — package resolution:

	Given a target package name, package resolution computes (and provisions) a list of all dependent packages together with specific package implementations, which meet the versioning constraints.

Therefore, package resolution can be viewed as a function, with:

	- Input: A package name and a version number.

	- Output: A single directory containing all dependent packages, whereby package names correspond to sub-directory paths. Furthermore, each package directory contains the package's source files for the appropriate package version, according to the package resolution rules.

In the context of an end-to-end compilation workflow, package resolution fits as the first stage, to be followed by parsing,compilation, and so on.

## TOOLING

There are two basic operations that developers need in order to have a productive interaction with packaging:

- Tidying a module. This ensures that all packages imported by source files in a module have a corresponding required version in the module definition.

- Performing resolution.

## IMPLEMENTATION STRATEGY

We defer implementation strategy discussions until after agreement on the goals.
One open question was: Can the code of the Go packaging system be reused to implement a packaging system with the modular interface proposed in this document? Our initial investigation of the Go codebase suggests that this is not easy to do cleanly.
