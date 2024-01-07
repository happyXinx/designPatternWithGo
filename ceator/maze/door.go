package maze

// 门
type Door struct {
	room1  *Room
	room2  *Room
	isOpen bool
}

func NewDoor(room *Room, room2 *Room) *Door {
	return &Door{
		room1:  room,
		room2:  room2,
		isOpen: false,
	}
}

func (r Door) Enter() {

}

func (r Door) OtherSideFrom(room Room) *Room {
	return r.room1
}
