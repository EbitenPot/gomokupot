package game

func (g *Game) Update() error {
	if !g.InitOK {
		g.InitScenes()
	}
	return nil
}
