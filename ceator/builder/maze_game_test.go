package builder

import (
	"testing"
)

func TestStandardMazeGame(t *testing.T) {
	game := &MazeGame{}
	builder := &StandardMazeBuilder{}
	game.CreateMaze(builder)
	builder.GetMaze()
}

func TestCountingMazeGame(t *testing.T) {
	game := &MazeGame{}
	builder := &CountingMazeBuilder{}
	game.CreateMaze(builder)
	builder.GetMaze()
}
