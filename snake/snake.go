package snake

import (
	"app/snake-game/constants"
	"app/snake-game/deque"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Snake struct {
	body      *deque.Deque[rl.Vector2]
	isRunning bool
}

func NewSnake() *Snake {
	body := deque.NewDeque[rl.Vector2]()
	body.PushBack(rl.NewVector2(6, 9))
	body.PushBack(rl.NewVector2(5, 9))
	body.PushBack(rl.NewVector2(4, 9))

	return &Snake{
		body:      body,
		isRunning: true,
	}
}

func (s *Snake) Draw() {
	isHead := true
	for v := range s.body.Iter() {
		x := v.X
		y := v.Y
		segment := rl.NewRectangle(float32(x)*constants.CellSize, float32(y)*constants.CellSize, constants.CellSize, constants.CellSize)
		if isHead {
			rl.DrawRectangleRounded(segment, 0.5, 6, rl.DarkGreen)
			isHead = false
		} else {
			rl.DrawRectangleRounded(segment, 0.5, 6, rl.Green)
		}
	}
}

func (s *Snake) Update() {
	if !s.isRunning {
		return
	}
	newHead := rl.Vector2Add(s.body.Front(), direction)

	s.body.PopBack()
	s.body.PushFront(newHead)
}

func (s *Snake) Stop() {
	s.isRunning = false
}

func (s *Snake) IsRunning() bool {
	return s.isRunning
}

func (s *Snake) GetBody() *deque.Deque[rl.Vector2] {
	return s.body
}

func (s *Snake) Reset() {
	s.body = deque.NewDeque[rl.Vector2]()
	s.body.PushBack(rl.NewVector2(6, 9))
	s.body.PushBack(rl.NewVector2(5, 9))
	s.body.PushBack(rl.NewVector2(4, 9))
	direction = rl.NewVector2(1, 0)
	s.isRunning = true
}
