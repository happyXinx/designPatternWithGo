package factoryMethod

import "design-pattern-go/ceator/abstractFactory"

type MazeGame interface {
	MakeMaze() abstractFactory.Mazer
	MakeRoom(int2 int) abstractFactory.Roomer
	MakeWall() abstractFactory.Waller
	MakeDoor(roomer abstractFactory.Roomer, roomer2 abstractFactory.Roomer) abstractFactory.Doorer
}
