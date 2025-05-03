package colors

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	figuresColors = []rl.Color{
		rl.DarkGray,
		rl.Red,
		rl.Green,
	}
)

func GetColors() []rl.Color {
	return figuresColors
}
