package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gomokupot/game"
	"log"
)

func main() {
	ebiten.SetWindowTitle("Gomokupot")
	if err := ebiten.RunGame(&game.Game{}); err != nil {
		log.Fatal(err)
	}

}
