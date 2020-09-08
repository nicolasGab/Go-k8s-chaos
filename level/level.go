package level

import tl "github.com/JoelOtter/termloop"

// Level -
type Level struct {
	*tl.BaseLevel
}

// Init -
func Init() *tl.BaseLevel {
	return tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorGreen,
		Fg: tl.ColorBlack,
		Ch: 'v',
	})
}

// Construct - Add elements to a level
func (level *Level) Construct() {
	level.AddEntity(tl.NewRectangle(10, 10, 50, 20, tl.ColorBlue))
}
