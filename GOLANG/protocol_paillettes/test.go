package main

import (
	"github.com/hajimehoshi/ebiten/v2"

	"image/color"
	"math/rand"
	"time"

	pa "paillettes/protocol"
)

var led pa.LedMatrixModel

const (
	screenWidth  = 280
	screenHeight = 280
)

type Game struct{}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {

	return screenWidth, screenHeight
}

func (g *Game) Update() error {

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//ebitenutil.DebugPrint(screen, "The dancefloor")

	x_start := 20
	x_stop := 50

	for j := 0; j < 8; j++ {

		for i := 0; i < 8; i++ {

			for x := x_start; x < x_stop; x++ {

				for y := i*30 + 20; y < i*30+50; y++ {

					couleur := color.RGBA{uint8((led.Matrix[j][i] >> 16) & 0xFF), uint8((led.Matrix[j][i] >> 8) & 0xFF), uint8((led.Matrix[j][i]) & 0xFF), 0}
					screen.Set(x, y, couleur)
				}
			}
		}

		x_start += 30
		x_stop += 30
	}
}

func main() {

	go pross()

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("The dancefloor")

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}

}

func pross() {

	for {

		for i := 0; i < 8; i++ {
			for j := 0; j < 8; j++ {
				led.Matrix[j][i] = uint32(rand.Intn(16777215))
			}
		}

		time.Sleep(300 * time.Millisecond)

	}

}
