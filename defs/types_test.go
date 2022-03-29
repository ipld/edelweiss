package defs

import (
	"testing"
)

func TestServiceDef(t *testing.T) {
	_ = Defs{

		Named{"Key", List{Byte{}}},

		Named{"PutArgs", Structure{
			Fields: Fields{
				Field{Name: "key", Type: Ref{"Key"}},
				Field{Name: "value", Type: Any{}},
			}}},

		Named{"ResultOk", Structure{
			Fields: Fields{
				Field{Name: "status", Type: SingletonString{"ok"}},
				Field{Name: "value", Type: Any{}},
			}}},

		Named{"ResultError", Structure{
			Fields{
				Field{Name: "status", Type: SingletonString{"error"}},
				Field{Name: "value", Type: String{}},
			}}},

		Named{"RoutingService",
			Service{
				Methods: Methods{
					Method{"Put1",
						Fn{
							Arg: Ref{"PutArgs"},
							Return: Union{
								Cases: Cases{
									Case{Name: "ok", Type: Ref{"ResultOk"}},
									Case{Name: "error", Type: Ref{"ResultError"}},
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
									Case{Name: "ok", Type: Ref{"ResultOk"}},
									Case{Name: "error", Type: Ref{"ResultError"}},
								},
							},
						},
					},
					Method{"Get",
						Fn{
							Arg: Ref{"Key"},
							Return: Union{
								Cases: Cases{
									Case{Name: "found", Type: Ref{"ResultOk"}},
									Case{Name: "not_found", Type: Ref{"ResultError"}},
								},
							},
						},
					},
				},
			},
		},
	}
}

func TestServiceDef2(t *testing.T) {
	_ = Defs{
		// Delegated Routing service definition
		Named{"DelegatedRoutingService",
			Service{
				Methods: Methods{
					Method{"PutP2PProvider", Fn{Arg: Ref{"PutP2PProviderRequest"}, Return: Ref{"PutP2PProviderResponse"}}},
					Method{"GetP2PProviders", Fn{Arg: Ref{"GetP2PProvidersRequest"}, Return: Ref{"GetP2PProvidersResponse"}}},
				},
			},
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
