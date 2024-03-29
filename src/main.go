package main

import (
	"image/color"
	"log"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 360
	boidsCount   = 500
	viewRadius   = 13
	adjRate      = 0.015
)

var (
	green   = color.RGBA{10, 255, 50, 255}
	boids   [boidsCount]*Boid
	boidMap [screenWidth + 1][screenHeight + 1]int
	rwMutex = sync.RWMutex{}
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, boid := range boids {
		screen.Set(int(boid.Position.x+1), int(boid.Position.y), green)
		screen.Set(int(boid.Position.x-1), int(boid.Position.y), green)
		screen.Set(int(boid.Position.x), int(boid.Position.y+1), green)
		screen.Set(int(boid.Position.x), int(boid.Position.y-1), green)
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

	for i := 0; i < boidsCount; i++ {
		createBoid(i)
	}

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Boids In Box")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
