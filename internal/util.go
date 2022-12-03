// Created by EldersJavas(EldersJavas&gmail.com)

package internal

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func IsInPos(x, y, x1, y1, x2, y2 float64) bool {
	if x >= x1 && x <= x2 {
		if y >= y1 && y <= y2 {
			return true
		}
	}
	return false
}

func InitSceneScreen(image *ebiten.Image) (int, int) {
	image.Fill(CBG)
	return image.Size()
}
func LogP(a ...interface{}) {
	log.Println(a)
}
