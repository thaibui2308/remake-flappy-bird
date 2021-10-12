package main

import (
	"fmt"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"

	"example.com/golang-game/floppy-bob/features"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var score int
var birdImage *ebiten.Image
var backgroundImage *ebiten.Image
var startGame = false

const (
	screenWidth  = 640
	screenHeight = 480
	gridSize     = 20
	xNumInScreen = screenWidth / gridSize
	yNumInScreen = screenHeight / gridSize
)

type Position struct {
	IMG *ebiten.Image
	X   int
	Y   int
}

var blocksArray = []*features.Block{
	&features.Block{xNumInScreen * gridSize / 3, yNumInScreen * gridSize / 5, 3 * gridSize, 3 * gridSize, gridSize * 3 / 10},
	&features.Block{xNumInScreen * 2 * gridSize / 3, yNumInScreen * 2 * gridSize / 5, 3 * gridSize, 3 * gridSize, gridSize * 0.3},
	&features.Block{xNumInScreen * gridSize, yNumInScreen * 3 * gridSize / 5, 3 * gridSize, 3 * gridSize, gridSize * 0.3},
}

type Game struct {
	point  Position
	blocks []*features.Block
	timer  int
	level  int
	score  int
}

func init() {
	var err error
	birdImage, err = ebitenutil.NewImageFromURL("https://www.toyjoy.com/wp-content/uploads/2019/04/catface3rotate-320x320.png")
	if err != nil {
		log.Fatal(err)
	}
	backgroundImage, err = ebitenutil.NewImageFromURL("https://www.zachwhalen.net/wp/wp-content/uploads/2015/06/flap.png")
	if err != nil {
		log.Fatal(err)
	}

}

func needsToMove(g *Game) bool {
	return g.timer/g.level > 0
}

func (g *Game) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		startGame = true
		g.point.Y -= (yNumInScreen / 2) * gridSize
	}

	if needsToMove(g) && startGame {
		for _, v := range g.blocks {
			// Detect collision
			if (g.point.X/3) >= v.X && (g.point.X/3) <= v.X+(3*gridSize) {
				if g.point.Y/4 <= v.HEIGHT || g.point.Y/4 >= v.HEIGHT+v.GAP*yNumInScreen {
					g.reset()
				}
				if (g.point.X / 3) == v.X+(3*gridSize) {
					g.score++
				}
			}
			if v.X == 0 {
				v.ChangeGap(gridSize/3, yNumInScreen)
			}

			v.Move(gridSize/5, screenWidth)
		}

		g.point.Y += screenHeight / (yNumInScreen * 2)
	}

	g.timer++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	// Draw backgroundImage
	var bgOp = &ebiten.DrawImageOptions{}
	bgOp.GeoM.Scale(.963, .722)
	screen.DrawImage(backgroundImage, bgOp)

	// Draw bird
	var op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.point.X), float64(g.point.Y))
	op.GeoM.Scale(0.2, 0.2)
	screen.DrawImage(g.point.IMG, op)

	// Draw blocks
	for i, _ := range g.blocks {
		ebitenutil.DrawRect(screen, float64(g.blocks[i].X), 0, float64(g.blocks[i].WIDTH), float64(g.blocks[i].HEIGHT), color.RGBA{45, 90, 39, 0xff})
		ebitenutil.DrawRect(screen, float64(g.blocks[i].X), float64(g.blocks[i].HEIGHT+g.blocks[i].GAP*yNumInScreen), float64(g.blocks[i].WIDTH), float64(screenWidth), color.RGBA{45, 90, 39, 0xff})
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %d", g.score))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *Game) reset() {
	g.point = Position{
		IMG: birdImage,
		X:   screenWidth * 0.4,
		Y:   2 * screenHeight / 3,
	}
	g.blocks = []*features.Block{
		&features.Block{xNumInScreen * gridSize / 3, yNumInScreen * gridSize / 5, 3 * gridSize, 3 * gridSize, gridSize * 3 / 10},
		&features.Block{xNumInScreen * 2 * gridSize / 3, yNumInScreen * 2 * gridSize / 5, 3 * gridSize, 3 * gridSize, gridSize * 0.3},
		&features.Block{xNumInScreen * gridSize, yNumInScreen * 3 * gridSize / 5, 3 * gridSize, 3 * gridSize, gridSize * 0.3},
	}
	g.timer = 1
	g.level = 1
	g.score = 0
	startGame = false
}

func newGame() *Game {
	g := &Game{
		point: Position{
			IMG: birdImage,
			X:   screenWidth * 0.4,
			Y:   2 * screenHeight / 3,
		},
		blocks: blocksArray,
		timer:  1,
		level:  1,
		score:  0,
	}

	return g
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Floppy Bob")

	if err := ebiten.RunGame(newGame()); err != nil {
		log.Fatal(err)
	}
}
