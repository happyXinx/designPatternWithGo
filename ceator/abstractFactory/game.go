package abstractFactory

func CreateGame() Mazer {
	factory := &EnchantedMazeFactory{}

	maze := factory.MakeMaze()
	r1 := factory.MakeRoom(1)
	r2 := factory.MakeRoom(2)
	door := factory.MakeDoor(r1, r2)

	maze.AddRoom(r1)
	maze.AddRoom(r2)

	r1.SetSide(North, factory.MakeWall())
	r1.SetSide(East, door)
	r1.SetSide(South, factory.MakeWall())
	r1.SetSide(West, factory.MakeWall())

	r2.SetSide(North, factory.MakeWall())
	r2.SetSide(East, factory.MakeWall())
	r2.SetSide(South, factory.MakeWall())
	r2.SetSide(West, door)

	return maze
}
