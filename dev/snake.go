package GO_TOYPRJ_1

import (
	tl "github.com/JoelOtter/termloop"
)

func NewSnake() *Snake {
	snake := new(Snake)
	// set a size
	snake.Entity = tl.NewEntity(5, 5, 1, 1)
	// set a direction, default = right
	snake.Direction = right
	// creates a snake containing of 3 entities
	snake.Bodylength = []Coordinates{
		{1, 6}, //tail
		{2, 6}, //body
		{3, 6}, //head
	}
	return snake
}

//head is a hitbox
func (snake *Snake) Head() *Coordinates {
	return &snake.Bodylength[len(snake.Bodylength)-1]
}

//Border collision checker  if the arena border contains then return true
func (snake *Snake) BorderCollision() bool {
	return gs.ArenaEntity.Contains(*snake.Head())
}

//food collision checker if the food contains the snakes head then return true
func (snake *Snake) FoodCollision() bool {
	return gs.FoodEntity.Contains(*snake.Head())
}

// snake collision checkers if the snakes body contains its head then return true
func (snake *Snake) SnakeCollision() bool {
	return snake.Contains()
}

// Draw will check every tick and draw the snake on the Screen
// check collision
func (snake *Snake) Draw(screen *tl.Screen) {
	// create new Head
	nHead := *snake.Head()
	// check the current direction
	switch snake.Direction {
	// snake go upside
	case up:
		nHead.Y-- // Y coordinates will be lowered.
	case down:
		nHead.Y++ // Y coordinates will be increased.
	case left:
		nHead.X-- // X coordinates will be lowered
	case right:
		nHead.X++ // X coordinates will be increased.
	}
	if snake.FoodCollision() {
		// change the food emoji is a special kind of food.
		switch gs.FoodEntity.Emoji {
		case 'R':
			switch ts.GameDifficulty {
			case easy:
				if gs.FPS-3 <= 8 {
					UpdateScore(5)
				} else {
					gs.FPS -= 3
					UpdateScore(5)
					UpdateFPS()
				}
			case normal:
				if gs.FPS-2 <= 12 {
					UpdateScore(5)
				} else {
					gs.FPS -= 2
					UpdateScore(5)
					UpdateFPS()
				}
			case hard:
				if gs.FPS-1 <= 20 {
					UpdateScore(5)
				} else {
					gs.FPS--
					UpdateScore(5)
					UpdateFPS()
				}
			}
			snake.Bodylength = append(snake.Bodylength, nHead)
		case 'S':
			switch ts.GameDifficulty {
			case easy:
				gs.FPS++
			case normal:
				gs.FPS += 3
			case hard:
				gs.FPS += 5
			}
			UpdateFPS()
		default:
			UpdateScore(1)
			snake.Bodylength = append(snake.Bodylength, nHead)
		}
		// if there is a food collision, it will call movefood fuc
		gs.FoodEntity.Movefood()
	} else {
		// if there is no collision with food add new head
		// but exclude the tail
		// keeping the snake the same size as before
		snake.Bodylength = append(snake.Bodylength[1:], nHead)
	}
	// position of snake will be moved
	snake.SetPosition(nHead.X, nHead.Y)

	if snake.BorderCollision() || snake.SnakeCollision() {
		Gameover()
	}

	for _, c := range snake.Bodylength {
		screen.RenderCell(c.X, c.Y, &tl.Cell{
			Fg: CheckSelectedColor(counterSnake),
			Ch: '$',
		})
	}
}

func (snake *Snake) Contains() bool {
	// i 0 부터 하나씩 더해서 길이 전 까지
	for i := 0; i < len(snake.Bodylength)-1; i++ {
		// 머리가 몸안에 있다면 트루 리턴
		if *snake.Head() == snake.Bodylength[i] {
			return true //game over
		}
	}
	// if the snake is not colliding with itself
	return false
}
