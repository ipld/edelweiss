
# Decentralized Protocol Compiler (WIP)

## MVP RPC protocol compiler

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
     PkgPath: "",
     PkgDir: "",
}
if err := build.Build(); err != nil {
     // ...
}
// ...
```
