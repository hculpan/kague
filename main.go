package main

import (
	_ "image/png"
	"kague/internal/game"
	"kague/internal/gameconfig"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	SCREENWIDTH  = 80
	SCREENHEIGHT = 50
)

func main() {
	g := game.NewGame()
	gameconfig.GameConfig.ScreenWidth = SCREENWIDTH
	gameconfig.GameConfig.ScreenHeight = SCREENHEIGHT

	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeDisabled)
	ebiten.SetWindowTitle("Kague")
	ebiten.SetWindowSize(gameconfig.GameConfig.ScreenWidth*gameconfig.GameConfig.TileWidth,
		gameconfig.GameConfig.ScreenHeight*gameconfig.GameConfig.TileHeight)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
