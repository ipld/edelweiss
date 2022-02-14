
# Decentralized Protocol Compiler (WIP)

For a detailed and longer-term roadmap of this project, refer to the [Protocol Compiler Roadmap 2022](design/roadmap.md).

## Target user experience for Mileston 1: MVP RPC protocol compiler

The MVP for an RPC compiler targets to enable the following workflow.

(1) Define services and types in Go. This example defines a simplified Delegated Routing protocol:

```go
defs = Types{
     // Delegated Routing service definition
     Named{"DelegatedRoutingService",
          MakeService(
               Method{"PutP2PProvider", Fn{Arg: Ref{"PutP2PProviderRequest"}, Return: Ref{"PutP2PProviderResponse"}}},
               Method{"GetP2PProviders", Fn{Arg: Ref{"GetP2PProvidersRequest"}, Return: Ref{"GetP2PProvidersResponse"}}},
          ),
     },

     // PutP2PProvider argument and result types
     Named{"PutP2PProviderRequest",
          MakeStructure(
               Field{Name: "Key", Type: List{Byte{}}},
               Field{Name: "Providers", Type: List{String{}}},
          ),
     },
     Named{"PutP2PProviderResponse",
          MakeUnion(
               Case{Name: "Success", Type: Nothing{}},
               Case{Name: "Error", Type: String{}},
          ),
     },

     // GetP2PProviders argument and result types
     Named{"GetP2PProvidersRequest",
          MakeStructure(
               Field{Name: "Key", Type: List{Byte{}}},
          ),
     },
     Named{"GetP2PProvidersResponse",
          MakeUnion(
               Case{Name: "Success", Type: List{Ref{"PeerAddr"}}},
               Case{Name: "Error", Type: String{}},
          ),
     },

     // Libp2p types
     Named{"PeerAddr",
          MakeStructure(
               Field{Name: "ID", Type: List{Byte{}}},
               Field{Name: "Multiaddresses", Type: List{String{}}},
          ),
     },
}
```

(2) Generate the service implementation (client and server). For instance:

```go
build := GenerateGo{
     Defs: defs,
     PkgPath: "github.com/ipfs/go-delegated-routing", // go package path
     PkgDir: "/home/petar/src/github.com/ipfs/go-delegated-routing", // local directory
}
if err := build.Build(); err != nil {
     // ...
}
// ...
```

### Generated code

The generated code will be entirely static — no use of Go reflection.

The generated code will supports IPLD serialization by implementing the native `Node` interface. Due to the lack of reflection, we expect this to be the fastest IPLD encoder to date.

For the MVP stage, the generated code will support IPLD deserialization by parsing from a deserialized IPLD data model format, rather than implementing the native `NodeAssembler` interface. Following the MVP, we expect to transition to implementing `NodeAssembler` relatively easily — due to the novel code generation framework we are utilizing. At this point, we expect that the generated code will be the fastest IPLD decoder yet.
