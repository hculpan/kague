package tiles

import (
	"image"
	"kague/internal/gameconfig"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type TileIdentifier byte

const (
	FLOOR  TileIdentifier = '.'
	WALL   TileIdentifier = '#'
	PLAYER TileIdentifier = '@'
)

var (
	tilesImage *ebiten.Image
)

func init() {
	// Decode an image from the image file's byte slice.
	img, _, err := ebitenutil.NewImageFromFile("assets/ascii-tileset.png")
	if err != nil {
		log.Fatal(err)
	}
	tilesImage = ebiten.NewImageFromImage(img)
}

func GetTileImage(index byte) *ebiten.Image {
	w := tilesImage.Bounds().Dx()
	tileXCount := w / gameconfig.GameConfig.TileWidth

	sx := (int(index) % tileXCount) * gameconfig.GameConfig.TileWidth
	sy := (int(index) / tileXCount) * gameconfig.GameConfig.TileWidth

	return tilesImage.SubImage(image.Rect(sx, sy, sx+gameconfig.GameConfig.TileWidth, sy+gameconfig.GameConfig.TileWidth)).(*ebiten.Image)
}
