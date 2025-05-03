package snake

import rl "github.com/gen2brain/raylib-go/raylib"

var direction rl.Vector2 = rl.NewVector2(1, 0)

var directionMap = map[int32]rl.Vector2{
	rl.KeyUp:    rl.NewVector2(0, -1),
	rl.KeyDown:  rl.NewVector2(0, 1),
	rl.KeyLeft:  rl.NewVector2(-1, 0),
	rl.KeyRight: rl.NewVector2(1, 0),
}

func (s *Snake) SetDirection(key int32) {
	if newDirection, exists := directionMap[key]; exists {
		direction = newDirection
	}
}

func (s *Snake) GetDirection() rl.Vector2 {
	return direction
}
