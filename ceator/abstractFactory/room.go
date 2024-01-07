package abstractFactory

type Roomer interface {
	MapSite
	GetSide(direction Direction) MapSite
	SetSide(direction Direction, site MapSite)
	Clone() Roomer
}

// 房间
type Room struct {
	roomNumber int
	sides      map[Direction]MapSite
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

func (r *Room) Clone() Roomer {
	return &Room{}
}

// 房间
type EnchantedRoom struct {
	roomNumber int
	sides      map[Direction]MapSite
}

func (r *EnchantedRoom) GetSide(direction Direction) MapSite {
	return r.sides[direction]
}

func (r *EnchantedRoom) SetSide(direction Direction, site MapSite) {
	if r.sides == nil {
		r.sides = make(map[Direction]MapSite, 4)
	}
	r.sides[direction] = site
}

func (r *EnchantedRoom) Enter() {

}

func (r *EnchantedRoom) Clone() Roomer {
	return &EnchantedRoom{}
}
