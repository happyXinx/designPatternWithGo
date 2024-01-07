package abstractFactory

type Waller interface {
	MapSite
	IsWall()
	Clone() Waller
}

// 墙壁
type Wall struct {
}

func (r *Wall) Enter() {

}

func (r *Wall) IsWall() {

}

func (r *Wall) Clone() Waller {
	return &Wall{}
}

// 墙壁
type EnchantedWall struct {
}

func (r *EnchantedWall) Enter() {

}

func (r *EnchantedWall) IsWall() {

}

func (r *EnchantedWall) Clone() Waller {
	return &EnchantedWall{}
}
