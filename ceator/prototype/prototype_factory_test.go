package prototype

import (
	"design-pattern-go/ceator/abstractFactory"
	"testing"
)

func TestPrototype(t *testing.T) {
	game := &MazeGame{}
	simpleFactory := NewMazeProtoTypeFactory(&abstractFactory.Maze{}, &abstractFactory.Door{}, &abstractFactory.Wall{}, &abstractFactory.Room{})
	game.CreateMaze(simpleFactory)

	bombedMazeFactory := NewMazeProtoTypeFactory(&abstractFactory.Maze{}, &abstractFactory.EnchantedDoor{}, &abstractFactory.Wall{}, &abstractFactory.Room{})
	game.CreateMaze(bombedMazeFactory)
}
