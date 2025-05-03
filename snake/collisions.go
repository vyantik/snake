package snake

import (
	"app/snake-game/constants"
	"app/snake-game/food"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (s *Snake) CheckCollisionWithFood(food *food.Food) bool {
	head := s.body.Front()
	if rl.Vector2Equals(head, food.GetPosition()) {
		s.body.PushBack(food.GetPosition())
		food.SetRandomPosition(*s.GetBody())
		return true
	}
	return false
}

func (s *Snake) CheckCollisionWithEdges() bool {
	head := s.body.Front()
	if head.X == constants.CellsCount || head.X == -1 {
		return true
	}
	if head.Y == constants.CellsCount || head.Y == -1 {
		return true
	}
	return false
}

func (s *Snake) CheckCollisionWithSelf() bool {
	head := s.body.Front()
	isHead := true
	for v := range s.body.Iter() {
		if isHead {
			isHead = false
			continue
		}
		if rl.Vector2Equals(v, head) {
			return true
		}
	}
	return false
}
