package def

import (
	"testing"
)

func TestServiceDef(t *testing.T) {
	_ = Types{

		Named{"Key", List{Byte{}}},

		Named{"PutArgs", MakeStructure(
			Field{"key", Ref{"Key"}},
			Field{"value", Any{}},
		)},

		Named{"ResultOk", MakeStructure(
			Field{"status", SingletonString{"ok"}},
			Field{"value", Any{}},
		)},

		Named{"ResultError", MakeStructure(
			Field{"status", SingletonString{"error"}},
			Field{"value", String{}},
		)},

		Named{"RoutingService",
			MakeService(
				Method{"Put1",
					Fn{
						Arg: Ref{"PutArgs"},
						Return: MakeUnion(
							Case{"ok", Ref{"ResultOk"}},
							Case{"error", Ref{"ResultError"}},
						),
					},
				},
				Method{"Put2",
					Fn{
						Arg: MakeTuple(Ref{"Key"}, Any{}),
						Return: MakeUnion(
							Case{"ok", Ref{"ResultOk"}},
							Case{"error", Ref{"ResultError"}},
						),
					},
				},
				Method{"Get",
					Fn{
						Arg: Ref{"Key"},
						Return: MakeUnion(
							Case{"found", Ref{"ResultOk"}},
							Case{"not_found", Ref{"ResultError"}},
						),
					},
				},
			),
		},
	}
}

func TestServiceDef2(t *testing.T) {
	_ = Types{
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
}
