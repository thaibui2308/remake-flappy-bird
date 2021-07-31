package features

import (
	"math/rand"
)

type Block struct {
	X      int
	Y      int
	WIDTH  int
	HEIGHT int
	GAP    int
}

func (b *Block) GenerateRandomPositions(xNum, yNum int) (X, Y int) {
	X = rand.Intn(xNum)
	Y = rand.Intn(yNum)
	return
}

func (b *Block) Move(distance, screenWidth int) {
	b.X -= distance
	if b.X < 0 {
		b.X = screenWidth - (-b.X % screenWidth)
	}
}

func (b *Block) ChangeGap(boundGridSize, yNumInScreen int) {
	gridSize := rand.Intn(boundGridSize) + 1
	b.HEIGHT = gridSize * yNumInScreen
}

func (b *Block) DetectCollision(x, y, yNumInScreen int) bool {
	if (x >= b.X && x <= b.X+1) && (y <= b.HEIGHT || y >= (b.HEIGHT+b.GAP*yNumInScreen)) {
		return true
	}
	return false
}
