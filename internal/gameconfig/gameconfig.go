package gameconfig

var GameConfig GameConfigType = GameConfigType{
	ScreenWidth:  80,
	ScreenHeight: 50,
	TileWidth:    16,
	TileHeight:   16,
}

type GameConfigType struct {
	ScreenWidth  int
	ScreenHeight int
	TileWidth    int
	TileHeight   int
}
