package food

import (
	"app/snake-game/constants"
	"app/snake-game/deque"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Food struct {
	position rl.Vector2
	texture  rl.Texture2D
}

func NewFood() *Food {
	return &Food{
		position: randomPosition(),
		texture:  rl.LoadTexture("assets/food.png"),
	}
}

func (f *Food) Draw() {
	rl.DrawTexture(f.texture, int32(f.position.X)*constants.CellSize, int32(f.position.Y)*constants.CellSize, rl.White)
}

func (f *Food) GetPosition() rl.Vector2 {
	return f.position
}

func randomPosition() rl.Vector2 {
	randomX := float32(rand.Intn(constants.CellsCount))
	randomY := float32(rand.Intn(constants.CellsCount))

	return rl.NewVector2(randomX, randomY)
}

func (f *Food) SetRandomPosition(snakeBody deque.Deque[rl.Vector2]) {
	for {
		f.position = randomPosition()
		if !deque.Vector2InDeque(&snakeBody, f.position) {
			break
		}
	}
}

func (f *Food) Unload() {
	rl.UnloadTexture(f.texture)
}
