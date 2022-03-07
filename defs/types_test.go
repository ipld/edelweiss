package defs

import (
	"testing"
)

func TestServiceDef(t *testing.T) {
	_ = Defs{

		Named{"Key", List{Byte{}}},

		Named{"PutArgs", Structure{
			Fields: Fields{
				Field{"key", Ref{"Key"}},
				Field{"value", Any{}},
			}}},

		Named{"ResultOk", Structure{
			Fields: Fields{
				Field{"status", SingletonString{"ok"}},
				Field{"value", Any{}},
			}}},

		Named{"ResultError", Structure{
			Fields{
				Field{"status", SingletonString{"error"}},
				Field{"value", String{}},
			}}},

		Named{"RoutingService",
			MakeService(
				Method{"Put1",
					Fn{
						Arg: Ref{"PutArgs"},
						Return: Union{
							Cases: Cases{
								Case{"ok", Ref{"ResultOk"}},
								Case{"error", Ref{"ResultError"}},
							},
						},
					},
				},
				Method{"Put2",
					Fn{
						Arg: Tuple{
							Slots: Slots{Ref{"Key"}, Any{}},
						},
						Return: Union{
							Cases: Cases{
								Case{"ok", Ref{"ResultOk"}},
								Case{"error", Ref{"ResultError"}},
							},
						},
					},
				},
				Method{"Get",
					Fn{
						Arg: Ref{"Key"},
						Return: Union{
							Cases: Cases{
								Case{"found", Ref{"ResultOk"}},
								Case{"not_found", Ref{"ResultError"}},
							},
						},
					},
				},
			),
		},
	}
}

func TestServiceDef2(t *testing.T) {
	_ = Defs{
		// Delegated Routing service definition
		Named{"DelegatedRoutingService",
			MakeService(
				Method{"PutP2PProvider", Fn{Arg: Ref{"PutP2PProviderRequest"}, Return: Ref{"PutP2PProviderResponse"}}},
				Method{"GetP2PProviders", Fn{Arg: Ref{"GetP2PProvidersRequest"}, Return: Ref{"GetP2PProvidersResponse"}}},
			),
		},

		// PutP2PProvider argument and result types
		Named{"PutP2PProviderRequest",
			Structure{
				Fields: Fields{
					Field{Name: "Key", Type: List{Byte{}}},
					Field{Name: "Providers", Type: List{String{}}},
				},
			},
		},
		Named{"PutP2PProviderResponse",
			Union{
				Cases: Cases{
					Case{Name: "Success", Type: Nothing{}},
					Case{Name: "Error", Type: String{}},
				},
			},
		},

		// GetP2PProviders argument and result types
		Named{"GetP2PProvidersRequest",
			Structure{
				Fields: Fields{
					Field{Name: "Key", Type: List{Byte{}}},
				},
			},
		},
		Named{"GetP2PProvidersResponse",
			Union{
				Cases: Cases{
					Case{Name: "Success", Type: List{Ref{"PeerAddr"}}},
					Case{Name: "Error", Type: String{}},
				},
			},
		},

		// Libp2p types
		Named{"PeerAddr",
			Structure{
				Fields: Fields{
					Field{Name: "ID", Type: List{Byte{}}},
					Field{Name: "Multiaddresses", Type: List{String{}}},
				},
			},
		},
	}
}
