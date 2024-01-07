package factoryMethod

import "testing"

func TestBombedMazeGame(t *testing.T) {
	game := &BombedMazeGame{}
	game.CreateMaze()
}

func TestEnchantedMazeGame(t *testing.T) {
	game := &EnchantedMazeGame{}
	game.CreateMaze()
}
