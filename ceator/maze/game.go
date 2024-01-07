package maze

// 创建迷宫的类
type MazeGame struct {
}

// 使用一系列操作将构件增加到迷宫中
func (r *MazeGame) CreateMaze() *Maze {
	aMaze := NewMaze()
	r1 := NewRoom(1)
	r2 := NewRoom(2)

	theDoor := NewDoor(r1, r2)
	aMaze.AddRoom(r1)
	aMaze.AddRoom(r2)

	r1.SetSide(North, NewWall())
	r1.SetSide(East, theDoor)
	r1.SetSide(South, NewWall())
	r1.SetSide(West, NewWall())

	r2.SetSide(North, NewWall())
	r2.SetSide(East, NewWall())
	r2.SetSide(South, NewWall())
	r2.SetSide(West, theDoor)

	return aMaze
}
