package abstractFactory

type MazeFactory interface {
	MakeMaze() Mazer
	MakeWall() Waller
	MakeRoom(n int) Roomer
	MakeDoor(room Roomer, room2 Roomer) Doorer
}

type DefaultMazeFactory struct{}

func (r DefaultMazeFactory) MakeMaze() Mazer {
	return &Maze{}
}

func (r DefaultMazeFactory) MakeWall() Waller {
	return &Wall{}
}

func (r DefaultMazeFactory) MakeRoom(n int) Roomer {
	return &Room{}
}

func (r DefaultMazeFactory) MakeDoor(r1 Roomer, r2 Roomer) Doorer {
	return &Door{}
}
