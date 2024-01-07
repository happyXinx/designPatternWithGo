package maze

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

// 迷宫
type Maze struct {
	rooms []*Room
}

func NewMaze() *Maze {
	return &Maze{}
}

func (r *Maze) AddRoom(room *Room) {
	r.rooms = append(r.rooms, room)
}

func (r *Maze) RoomNo(i int) *Room {
	return r.rooms[i]
}
