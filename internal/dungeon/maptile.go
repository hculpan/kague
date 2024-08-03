package dungeon

import "kague/internal/tiles"

type TileType int

type MapTile struct {
	PixelX  int
	PixelY  int
	Blocked bool
	Tile    tiles.TileIdentifier
}
