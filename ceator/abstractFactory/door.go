package abstractFactory

type Doorer interface {
	MapSite
	OtherSideFrom(room Roomer) Roomer
	Clone() Doorer
}

// 门
type Door struct {
	room1  Roomer
	room2  Roomer
	isOpen bool
}

func (r Door) Enter() {

}

func (r Door) OtherSideFrom(room Roomer) Roomer {
	return r.room1
}

func (r Door) Clone() Doorer {
	return &Door{}
}

// 门
type EnchantedDoor struct {
	room1  Roomer
	room2  Roomer
	isOpen bool
}

func (r EnchantedDoor) Enter() {

}

func (r EnchantedDoor) OtherSideFrom(room Roomer) Roomer {
	return r.room1
}

func (r EnchantedDoor) Clone() Doorer {
	return &EnchantedDoor{}
}
