package game

import (
	game "gomokupot/game/openscene"
	"gomokupot/internal"
)

type Game struct {
	Scene  internal.Scene
	InitOK bool
}

func (g *Game) InitScenes() {
	g.Scene = &game.OScene{}
}
