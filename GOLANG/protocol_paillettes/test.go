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
	x_stop := 35

	for j := 0; j < 16; j++ {

		for i := 0; i < 16; i++ {

			for x := x_start; x < x_stop; x++ {

				for y := i*15 + 20; y < i*15+35; y++ {

					couleur := color.RGBA{uint8((led.Matrix[j][i] >> 16) & 0xFF), uint8((led.Matrix[j][i] >> 8) & 0xFF), uint8((led.Matrix[j][i]) & 0xFF), 0}
					screen.Set(x, y, couleur)
				}
			}
		}

		x_start += 15
		x_stop += 15
	}
}

func main() {

	go pross2()

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("The dancefloor")

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}

}

func pross() {

	for {

		for i := 0; i < 16; i++ {
			for j := 0; j < 16; j++ {
				led.Matrix[j][i] = uint32(rand.Intn(16777215))
			}
		}

		time.Sleep(300 * time.Millisecond)

	}

}

func pross2() {

	dt := new(pa.DataController)
	dt.IP = "127.0.0.1"
	dt.Port = "1053"
	dt.Broadcast = "192.168.0.255"

	dt.InitConnection()

	i := 0
	j := 0

	for {
		j = rand.Intn(15)
		i = rand.Intn(15)

		led.Matrix[j][i] = uint32(rand.Intn(16777215))

		frameOut1, frameOut2 := led.ConvertMatrixToFrame()

		dt.WriteData(frameOut1.ConvertFrameOutputToBytes())
		dt.WriteData(frameOut2.ConvertFrameOutputToBytes())

		time.Sleep(300 * time.Millisecond)
	}
}
