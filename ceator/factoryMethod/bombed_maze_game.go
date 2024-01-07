package factoryMethod

import "design-pattern-go/ceator/abstractFactory"

type BombedMazeGame struct {
}

func (r *BombedMazeGame) MakeMaze() abstractFactory.Mazer {
	return &abstractFactory.Maze{}
}

func (r *BombedMazeGame) MakeRoom(int2 int) abstractFactory.Roomer {
	return &abstractFactory.Room{}
}

func (r *BombedMazeGame) MakeDoor(roomer abstractFactory.Roomer, doorer2 abstractFactory.Roomer) abstractFactory.Doorer {
	return &abstractFactory.Door{}
}

func (r *BombedMazeGame) MakeWall() abstractFactory.Waller {
	return &abstractFactory.Wall{}
}

func (r *BombedMazeGame) CreateMaze() abstractFactory.Mazer {
	maze := r.MakeMaze()
	r1 := r.MakeRoom(1)
	r2 := r.MakeRoom(2)
	maze.AddRoom(r1)
	maze.AddRoom(r2)
	return maze
}
