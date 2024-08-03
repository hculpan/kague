package dungeon

import (
	"kague/internal/gameconfig"
	"kague/internal/tiles"
	"kague/internal/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type Level struct {
	Tiles []MapTile
	Rooms []Rect
}

func NewLevel() *Level {
	l := Level{}
	rooms := make([]Rect, 0)
	l.Rooms = rooms
	l.GenerateLevelTiles()
	return &l
}

func (level *Level) createRoom(room Rect) {
	for y := room.Y1 + 1; y < room.Y2; y++ {
		for x := room.X1 + 1; x < room.X2; x++ {
			index := level.GetIndexFromXY(x, y)
			level.Tiles[index].Blocked = false
			level.Tiles[index].Tile = tiles.FLOOR
		}
	}
}

func (level *Level) GenerateLevelTiles() {
	MIN_SIZE := 6
	MAX_SIZE := 10
	MAX_ROOMS := 30

	gd := gameconfig.GameConfig
	tiles := level.createTiles()
	level.Tiles = tiles

	for idx := 0; idx < MAX_ROOMS; idx++ {
		w := utils.GetRandomBetween(MIN_SIZE, MAX_SIZE)
		h := utils.GetRandomBetween(MIN_SIZE, MAX_SIZE)
		x := utils.GetDiceRoll(gd.ScreenWidth-w-1) - 1
		y := utils.GetDiceRoll(gd.ScreenHeight-h-1) - 1

		new_room := NewRect(x, y, w, h)
		okToAdd := true
		for _, otherRoom := range level.Rooms {
			if new_room.Intersect(otherRoom) {
				okToAdd = false
				break
			}
		}
		if okToAdd {
			level.createRoom(new_room)
			level.Rooms = append(level.Rooms, new_room)
		}
	}
}

func (l *Level) GetIndexFromXY(x int, y int) int {
	return (y * gameconfig.GameConfig.ScreenWidth) + x
}

func (l *Level) createTiles() []MapTile {
	screenWidth := gameconfig.GameConfig.ScreenWidth
	screenHeight := gameconfig.GameConfig.ScreenHeight
	tileMap := make([]MapTile, screenWidth*screenHeight)

	for x := 0; x < screenWidth; x++ {
		for y := 0; y < screenHeight; y++ {
			index := l.GetIndexFromXY(x, y)
			tile := MapTile{
				PixelX:  x * gameconfig.GameConfig.TileWidth,
				PixelY:  y * gameconfig.GameConfig.TileHeight,
				Blocked: true,
				Tile:    tiles.WALL,
			}
			tileMap[index] = tile
		}
	}

	return tileMap
}

// Draw is called each draw cycle and is where we will blit.
func (l *Level) Draw(screen *ebiten.Image) {
	for x := 0; x < gameconfig.GameConfig.ScreenWidth; x++ {
		for y := 0; y < gameconfig.GameConfig.ScreenHeight; y++ {
			tile := l.Tiles[l.GetIndexFromXY(x, y)]
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
			subImage := tiles.GetTileImage(byte(tile.Tile))
			screen.DrawImage(subImage, op)
		}
	}
}
