package builder

import "design-pattern-go/ceator/maze"

type StandardMazeBuilder struct {
	currentMaze *maze.Maze
}

func (r *StandardMazeBuilder) BuildMaze() {
	r.currentMaze = &maze.Maze{}
}

func (r *StandardMazeBuilder) BuildRoom(int2 int) {

}

func (r *StandardMazeBuilder) BuildDoor(int2 int, int3 int) {

}

func (r *StandardMazeBuilder) GetMaze() *maze.Maze {
	return r.currentMaze
}
