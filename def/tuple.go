package def

type Tuple struct {
	Slots SlotListOrNone
}

func (t Tuple) Deps() Types {
	if t.Slots == nil {
		return nil
	} else {
		return t.Slots.Deps()
	}
}

func (Tuple) Kind() string {
	return "Tuple"
}

type SlotListOrNone interface {
	Deps() Types
}

type SlotList struct {
	Slot Type
	Rest SlotListOrNone
}

func (sl SlotList) Deps() Types {
	if sl.Rest == nil {
		return Types{sl.Slot}
	} else {
		return append(Types{sl.Slot}, sl.Rest.Deps()...)
	}
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

func FlattenSlotList(x SlotListOrNone) []Type {
	r, cur := []Type{}, x
	for cur != nil {
		l := cur.(SlotList)
		r = append(r, l.Slot)
		cur = l.Rest
	}
	return r
}
