package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
	"sync"
)

const (
	screenWidth, screenHeight = 640, 360
	birdCount                 = 500
	viewRadius                = 13
	adjRate                   = 0.015
)

var (
	green   = color.RGBA{10, 255, 50, 255}
	birds   [birdCount]*Bird
	boidMap [screenWidth + 1][screenHeight + 1]int
	rWlock  = sync.RWMutex{}
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, bird := range birds {
		screen.Set(int(bird.position.x+1), int(bird.position.y), green)
		screen.Set(int(bird.position.x-1), int(bird.position.y), green)
		screen.Set(int(bird.position.x), int(bird.position.y-1), green)
		screen.Set(int(bird.position.x), int(bird.position.y+1), green)
	}
}

func (g *Game) Layout(_, _ int) (w, h int) {
	return screenWidth, screenHeight
}

func main() {
	for i, row := range boidMap {
		for j := range row {
			boidMap[i][j] = -1
		}
	}

	for i := 0; i < birdCount; i++ {
		createBird(i)
	}
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("BigBrotherBirdBag in a box")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
