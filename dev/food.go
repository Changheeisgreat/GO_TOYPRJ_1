package GO_TOYPRJ_1

import (
	"math/rand"
	"time"

	tl "github.com/JoelOtter/termloop"
)

// border = arena -1

var insideborderW = 70 - 1
var insideborderH = 25 - 1

func NewFood() *Food {
	food := new(Food)
	// create a new entity food with a standard position
	food.Entity = tl.NewEntity(1, 1, 1, 1)
	food.Movefood()

	return food
}

//move food into random place
func (food *Food) Movefood() {
	NewX := RandomInsideArena(insideborderW, 1)
	NewY := RandomInsideArena(insideborderH, 1)

	food.Foodposition.X = NewX
	food.Foodposition.Y = NewY
	food.Emoji = RandomFood()

	food.SetPosition(food.Foodposition.X, food.Foodposition.Y)
}

func RandomFood() rune {
	// This slice contains all of the possible food icons.
	emoji := []rune{
		'R', // Favourite dish, extra points!!!
		'■', // 1 point
		'■', // 1 point
		'■', // 1 point
		'■', // 1 point
		'■', // 1 point
		'■', // 1 point
		'■', // 1 point
		'■', // 1 point
		'■', // 1 point
		'■', // 1 point
		'S', // You do not want to eat the skull
	}

	rand.Seed(time.Now().UnixNano())

	// Return a random rune picked from the slice
	return emoji[rand.Intn(len(emoji))]
}

// print food
func (food *Food) Draw(screen *tl.Screen) {
	screen.RenderCell(food.Foodposition.X, food.Foodposition.Y, &tl.Cell{
		Ch: food.Emoji,
	})
}

// Contains checks if food contains the coordinates, then return true

func (food *Food) Contains(c Coordinates) bool {
	return c.X == food.Foodposition.X && c.Y == food.Foodposition.Y
}

// RandomInsideArena inside the Border
func RandomInsideArena(iMax int, iMin int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(iMax-iMin) + iMin
}
