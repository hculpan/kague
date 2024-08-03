package game

import (
	"kague/internal/components"
	"kague/internal/dungeon"
	"kague/internal/gameconfig"
	"kague/internal/tiles"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	dungeon *dungeon.Dungeon
}

func NewGame() *Game {
	g := &Game{
		dungeon: dungeon.NewDungeon(),
	}

	x, y := g.dungeon.ActiveLevel().Rooms[0].Center()

	p := components.NewPlayer()
	p.SetPosition(x, y)
	p.SetImage(tiles.GetTileImage(byte(tiles.PLAYER)))

	components.AddEntity(p)

	return g
}

func (g *Game) Update() error {
	g.MovePlayer()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.dungeon.Draw(screen)
	g.DrawEntities(screen)
}

func (g *Game) MovePlayer() {
	x := 0
	y := 0

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		y = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		y = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		x = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		x = 1
	}

	for i := 0; i < components.GetEntityCount(); i++ {
		entity := components.GetNextEntity(i)
		if components.CheckInterface(entity, (*components.Movable)(nil)) {
			cx, cy := entity.(components.Renderable).GetPosition()
			tile := g.dungeon.ActiveLevel().Tiles[g.dungeon.ActiveLevel().GetIndexFromXY(cx+x, cy+y)]
			if !tile.Blocked {
				entity.(components.Movable).UpdatePosition(x, y)
			}
		}
	}
}

func (g *Game) DrawEntities(screen *ebiten.Image) {
	entity, index := components.FindNextEntity(0, (*components.Renderable)(nil))
	for entity != nil {
		r := entity.(components.Renderable)
		op := &ebiten.DrawImageOptions{}
		x, y := r.GetPosition()
		op.GeoM.Translate(float64(x*gameconfig.GameConfig.TileWidth), float64(y*gameconfig.GameConfig.TileHeight))
		screen.DrawImage(r.GetImage(), op)
		entity, index = components.FindNextEntity(index+1, (*components.Renderable)(nil))
	}
}

// Layout will return the screen dimensions.
func (g *Game) Layout(w, h int) (int, int) {
	return gameconfig.GameConfig.ScreenWidth * gameconfig.GameConfig.TileWidth,
		gameconfig.GameConfig.ScreenHeight * gameconfig.GameConfig.TileHeight
}
