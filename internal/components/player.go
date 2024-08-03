package components

import "github.com/hajimehoshi/ebiten/v2"

type Player struct {
	entityId int
	x        int
	y        int
	image    *ebiten.Image
}

func NewPlayer() *Player {
	return &Player{}
}

func (p *Player) SetPosition(x, y int) {
	p.x = x
	p.y = y
}

func (p *Player) UpdatePosition(x, y int) {
	p.x += x
	p.y += y
}

func (p *Player) GetPosition() (int, int) {
	return p.x, p.y
}

func (p *Player) SetImage(i *ebiten.Image) {
	p.image = i
}

func (p *Player) GetImage() *ebiten.Image {
	return p.image
}

func (p *Player) GetEntityId() int {
	return p.entityId
}

func (p *Player) SetEntityId(id int) {
	p.entityId = id
}
