package def

type Tuple struct {
	Slots SlotListOrNone
}

func (Tuple) Kind() string {
	return "Tuple"
}

type SlotListOrNone interface{}

type SlotList struct {
	Slot Type
	Rest SlotListOrNone
}

func MakeTuple(slots ...Type) Tuple {
	return Tuple{
		Slots: makeSlots(slots),
	}
}

func makeSlots(slots []Type) SlotListOrNone {
	if len(slots) == 0 {
		return nil
	} else {
		return SlotList{
			Slot: slots[0],
			Rest: makeSlots(slots[1:]),
		}
	}
}
