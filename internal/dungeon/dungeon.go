package dungeon

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Dungeon is a container for all the levels that make up a particular dungeon in the world.
type Dungeon struct {
	Name   string
	Levels []*Level

	activeLevel int
}

func NewDungeon() *Dungeon {
	result := &Dungeon{
		Levels: make([]*Level, 0),
	}
	result.Levels = append(result.Levels, NewLevel())
	result.activeLevel = 0

	return result
}

func (d *Dungeon) ActiveLevel() *Level {
	return d.Levels[d.activeLevel]
}

func (d *Dungeon) Draw(screen *ebiten.Image) {
	d.ActiveLevel().Draw(screen)
}
