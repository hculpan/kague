package components

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Entity interface {
	GetEntityId() int
	SetEntityId(id int)
}

type Renderable interface {
	GetPosition() (int, int)
	GetImage() *ebiten.Image
}

type Movable interface {
	SetPosition(x, y int)
	UpdatePosition(x, y int)
}
