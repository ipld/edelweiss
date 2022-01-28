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
