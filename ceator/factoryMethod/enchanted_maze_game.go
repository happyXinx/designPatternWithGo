package factoryMethod

import "design-pattern-go/ceator/abstractFactory"

type EnchantedMazeGame struct {
}

func (r *EnchantedMazeGame) MakeMaze() abstractFactory.Mazer {
	return &abstractFactory.EnchantedMaze{}
}

func (r *EnchantedMazeGame) MakeRoom(int2 int) abstractFactory.Roomer {
	return &abstractFactory.Room{}
}

func (r *EnchantedMazeGame) MakeDoor(roomer abstractFactory.Roomer, doorer2 abstractFactory.Roomer) abstractFactory.Doorer {
	return &abstractFactory.Door{}
}

func (r *EnchantedMazeGame) MakeWall() abstractFactory.Waller {
	return &abstractFactory.Wall{}
}

func (r *EnchantedMazeGame) CreateMaze() abstractFactory.Mazer {
	maze := r.MakeMaze()
	r1 := r.MakeRoom(1)
	r2 := r.MakeRoom(2)
	maze.AddRoom(r1)
	maze.AddRoom(r2)
	return maze
}
