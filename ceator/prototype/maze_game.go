package prototype

import "design-pattern-go/ceator/abstractFactory"

type MazeGame struct {
}

func (m MazeGame) CreateMaze(factory abstractFactory.MazeFactory) abstractFactory.Mazer {
	return factory.MakeMaze()
}
