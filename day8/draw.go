package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	width  = 640
	height = 480
)

type Game struct {
	grid       [][]string
	highlights []Pos
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return width, height
}

func (g *Game) Draw(screen *ebiten.Image) {
	// totalRows := len(g.grid)
	// totalCols := len(g.grid[0])
	// cellWidth, cellHeight := width/totalCols, height/totalRows
	// x, y := 0, 0
	//
	// vector.DrawFilledRect(screen, float32(x), float32(y), float32(cellWidth), float32(cellHeight), color.RGBA{20, 45, 22, 1}, false)
	// vector.DrawFilledRect(screen, float32(x), float32(y), float32(cellWidth), float32(cellHeight), color.RGBA{20, 45, 22, 1}, false)
	//
	// // for y, row := range g.grid {
	// // 	for x := range row {
	// // 		vector.DrawFilledRect(screen, float32(x), float32(y), float32(cellWidth), float32(cellHeight), color.RGBA{20, 45, 22, 1}, false)
	// // 	}
	// // }
}

func (g *Game) Update() error {
	return nil
}

func DrawMap(grid [][]string, highlights []Pos) {
	ebiten.SetWindowTitle("Resonant Collinearity")
	ebiten.SetWindowSize(width, height)

	g := &Game{
		grid, highlights,
	}

	err := ebiten.RunGame(g)
	if err != nil {
		log.Fatal(err)
	}
}
