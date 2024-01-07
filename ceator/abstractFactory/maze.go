package abstractFactory

type Direction = string

const (
	North Direction = "North"
	South Direction = "South"
	East  Direction = "East"
	West  Direction = "West"
)

// 所有迷宫组件的公共抽象类
type MapSite interface {
	Enter()
}

type Mazer interface {
	AddRoom(room Roomer)
	RoomNo(i int) Roomer
	Clone() Mazer
}

// 迷宫
type Maze struct {
	rooms []Roomer
}

func (r *Maze) AddRoom(room Roomer) {
	r.rooms = append(r.rooms, room)
}

func (r *Maze) RoomNo(i int) Roomer {
	return r.rooms[i]
}

func (r *Maze) Clone() Mazer {
	return &Maze{}
}

// 迷宫
type EnchantedMaze struct {
	rooms []Roomer
}

func (r *EnchantedMaze) AddRoom(room Roomer) {
	r.rooms = append(r.rooms, room)
}

func (r *EnchantedMaze) RoomNo(i int) Roomer {
	return r.rooms[i]
}

func (r *EnchantedMaze) Clone() Mazer {
	return &Maze{}
}
