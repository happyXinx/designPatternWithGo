package prototype

import "design-pattern-go/ceator/abstractFactory"

type MazeProtoTypeFactory struct {
	protoTypeMaze abstractFactory.Mazer
	protoTypeRoom abstractFactory.Roomer
	protoTypeWall abstractFactory.Waller
	protoTypeDoor abstractFactory.Doorer
}

func (r MazeProtoTypeFactory) MakeMaze() abstractFactory.Mazer {
	return r.protoTypeMaze.Clone()
}

func (r MazeProtoTypeFactory) MakeWall() abstractFactory.Waller {
	return r.protoTypeWall.Clone()
}

func (r MazeProtoTypeFactory) MakeRoom(n int) abstractFactory.Roomer {
	return r.protoTypeRoom.Clone()
}

func (r MazeProtoTypeFactory) MakeDoor(r1 abstractFactory.Roomer, r2 abstractFactory.Roomer) abstractFactory.Doorer {
	return r.protoTypeDoor.Clone()
}

func NewMazeProtoTypeFactory(mazer abstractFactory.Mazer, doorer abstractFactory.Doorer, waller abstractFactory.Waller, roomer abstractFactory.Roomer) *MazeProtoTypeFactory {
	return &MazeProtoTypeFactory{
		protoTypeMaze: mazer,
		protoTypeRoom: roomer,
		protoTypeWall: waller,
		protoTypeDoor: doorer,
	}
}
