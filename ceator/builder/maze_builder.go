package builder

import (
	"design-pattern-go/ceator/maze"
)

type MazeBuilder interface {
	BuildMaze()
	BuildRoom(room int)
	BuildDoor(roomFrom int, roomTo int)
	GetMaze() *maze.Maze
}
