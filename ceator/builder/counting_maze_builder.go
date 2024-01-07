package builder

import "design-pattern-go/ceator/maze"

type CountingMazeBuilder struct {
	doors int
	rooms int
}

func (r *CountingMazeBuilder) BuildMaze() {

}

func (r *CountingMazeBuilder) BuildRoom(int2 int) {
	r.rooms++
}

func (r *CountingMazeBuilder) BuildDoor(int2 int, int3 int) {
	r.doors++
}

func (r *CountingMazeBuilder) GetMaze() *maze.Maze {
	return maze.NewMaze()
}
