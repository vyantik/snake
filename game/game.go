package game

import (
	"app/snake-game/constants"
	"app/snake-game/food"
	"app/snake-game/snake"
	"fmt"
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type game struct {
	cellSize   int32
	cellsCount int32

	score int

	food  *food.Food
	snake *snake.Snake
}

func newGame() *game {
	return &game{
		cellSize:   constants.CellSize,
		cellsCount: constants.CellsCount,
		food:       food.NewFood(),
		snake:      snake.NewSnake(),
		score:      0,
	}
}

func (g *game) unload() {
	g.food.Unload()
}

func (g *game) draw() {
	g.food.Draw()
	g.snake.Draw()
}

func (g *game) update() {
	if g.snake.IsRunning() {
		head := g.snake.GetBody().Front()
		if rl.Vector2Equals(head, g.food.GetPosition()) {
			if g.snake.CheckCollisionWithFood(g.food) {
				g.score++
			}
		} else {
			g.snake.Update()
			if g.snake.CheckCollisionWithEdges() || g.snake.CheckCollisionWithSelf() {
				g.snake.Stop()
				return
			}
		}
	}
}

func (g *game) reset() {
	g.snake.Reset()
	g.food.SetRandomPosition(*g.snake.GetBody())
	g.score = 0
}

var lastUpdate float64 = rl.GetTime()

func Start() {
	screenWidth := int32(rl.GetScreenWidth())
	screenHeight := int32(rl.GetScreenHeight())
	log.Println(screenWidth, screenHeight)
	rl.InitWindow(constants.CellSize*constants.CellsCount, constants.CellSize*constants.CellsCount, "Snake game")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	backgroundColor := rl.NewColor(30, 30, 30, 255)

	game := newGame()
	defer game.unload()

	for !rl.WindowShouldClose() {
		rl.SetWindowTitle(fmt.Sprintf("Snake game - FPS: %d; Score: %d", rl.GetFPS(), game.score))
		if eventTriggered(0.2) {
			game.update()
		}
		game.snake.SetDirection(game.getInput())

		rl.BeginDrawing()
		rl.ClearBackground(backgroundColor)
		game.draw()
		rl.EndDrawing()
	}
}

func (g *game) getInput() int32 {
	direction := g.snake.GetDirection()
	switch {
	case rl.IsKeyDown(rl.KeyUp):
		if direction.Y != 1 {
			if !g.snake.IsRunning() {
				g.reset()
			}
			return rl.KeyUp
		}
	case rl.IsKeyDown(rl.KeyDown):
		if direction.Y != -1 {
			if !g.snake.IsRunning() {
				g.reset()
			}
			return rl.KeyDown
		}
	case rl.IsKeyDown(rl.KeyLeft):
		if direction.X != 1 {
			if !g.snake.IsRunning() {
				g.reset()
			}
			return rl.KeyLeft
		}
	case rl.IsKeyDown(rl.KeyRight):
		if direction.X != -1 {
			if !g.snake.IsRunning() {
				g.reset()
			}
			return rl.KeyRight
		}
	}
	return 0
}

func eventTriggered(interval float64) bool {
	now := rl.GetTime()
	if now-lastUpdate > interval {
		lastUpdate = now
		return true
	}
	return false
}
