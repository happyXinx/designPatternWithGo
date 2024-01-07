package abstractFactory

type EnchantedMazeFactory struct {
	DefaultMazeFactory
}

func (r EnchantedMazeFactory) MakeRoom(n int) Roomer {
	return &EnchantedRoom{}
}

func (r EnchantedMazeFactory) MakeDoor(r1 Roomer, r2 Roomer) Doorer {
	return &EnchantedDoor{room1: r1, room2: r2}
}
