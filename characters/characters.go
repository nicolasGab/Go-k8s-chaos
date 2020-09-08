package characters

import tl "github.com/JoelOtter/termloop"

// Player -
type Player struct {
	*tl.Entity
	PrevX int
	PrevY int
	Level *tl.BaseLevel
}

// Draw -
func (player *Player) Draw(screen *tl.Screen) {
	screenWidth, screenHeight := screen.Size()
	x, y := player.Position()
	player.Level.SetOffset(screenWidth/2-x, screenHeight/2-y)
	player.Entity.Draw(screen)
}

// Collide -
func (player *Player) Collide(collision tl.Physical) {
	// Check if it's a Rectangle we're colliding with
	if _, ok := collision.(*tl.Rectangle); ok {
		player.SetPosition(player.PrevX, player.PrevY)
	}
}

// Tick -
func (player *Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey { // Is it a keyboard event?
		player.PrevX, player.PrevY = player.Position()
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyArrowRight:
			player.SetPosition(player.PrevX+1, player.PrevY)
		case tl.KeyArrowLeft:
			player.SetPosition(player.PrevX-1, player.PrevY)
		case tl.KeyArrowUp:
			player.SetPosition(player.PrevX, player.PrevY-1)
		case tl.KeyArrowDown:
			player.SetPosition(player.PrevX, player.PrevY+1)
		}
	}
}
