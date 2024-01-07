package builder

import "design-pattern-go/ceator/maze"

type MazeGame struct {
}

func (r *MazeGame) CreateMaze(builder MazeBuilder) *maze.Maze {
	builder.BuildMaze()
	builder.BuildRoom(1)
	builder.BuildDoor(1, 2)

	return builder.GetMaze()
}
