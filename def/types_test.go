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
		MakeService(
			Method{"Put1",
				Fn{
					Arg: Ref{"PutArgs"},
					Return: MakeUnion(
						Case{"ok", Any{}},
						Case{"error", String{}},
					),
				},
			},
			Method{"Put2",
				Fn{
					Arg: MakeTuple(Ref{"Key"}, Any{}),
					Return: MakeUnion(
						Case{"ok", Any{}},
						Case{"error1", Nothing{}},
						Case{"error2", Nothing{}},
					),
				},
			},
			Method{"Get",
				Fn{
					Arg: Ref{"Key"},
					Return: MakeUnion(
						Case{"found", Any{}},
						Case{"not_found", Nothing{}},
					),
				},
			},
		),
	}
}
