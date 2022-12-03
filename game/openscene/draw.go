package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"gomokupot/internal"
	"image/color"
)

func (g *OScene) Draw(screen *ebiten.Image) {
	sw, _ := internal.InitSceneScreen(screen)
	lines := []string{"五子棋 Gomoku"}
	for _, line := range lines {
		f := internal.SpaceAgeBig
		r := text.BoundString(f, line)
		x := (sw-r.Dx())/2 - r.Min.X - 30
		y := 150
		text.Draw(screen, line, f, x, y, internal.CTitle)
	}
	lines = []string{"开始游戏 Start"}
	for n, line := range lines {
		f := internal.SpaceAgeSmall
		r := text.BoundString(f, line)
		x := (sw-r.Dx())/2 - r.Min.X - 30
		y := 350 + (n * 150)
		text.Draw(screen, line, f, x, y, color.CMYK{61, 49, 0, 0})
	}
}
