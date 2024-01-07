package maze

// 房间
type Room struct {
	roomNumber int
	sides      map[Direction]MapSite
}

func NewRoom(n int) *Room {
	return &Room{
		roomNumber: n,
		sides:      nil,
	}
}

func (r *Room) GetSide(direction Direction) MapSite {
	return r.sides[direction]
}

func (r *Room) SetSide(direction Direction, site MapSite) {
	if r.sides == nil {
		r.sides = make(map[Direction]MapSite, 4)
	}
	r.sides[direction] = site
}

func (r *Room) Enter() {

}
