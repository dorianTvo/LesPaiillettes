package main

import (
	"github.com/hajimehoshi/ebiten/v2"

	"image/color"
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

					couleur := color.RGBA{led.Matrix[j][i].Color[0], led.Matrix[j][i].Color[1], led.Matrix[j][i].Color[2], 0}
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
				//led.Matrix[j][i] = uint32(rand.Intn(16777215))
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

	for i := 0; i < 16; i++ {

			for j := 0; j < 16; j++ {

				led.Matrix[i][j].Hold = 0
			

			}
			
		}

	go pross3()

	for {

		data, _ := dt.ReadIncommingData() // Lectures des octects de la trame

		frameIn := new(pa.FrameModelInput)     // Definition de la structure de données
		frameIn.ConvertBytesToFrameInput(data) //Initialisation de la structure à partir des octets reçu

		print(frameIn.ID)
		go cross_color(frameIn.ID, dt)


		time.Sleep(10 * time.Millisecond)
	}
}

func pross3() {

	//Couleur de fond

	for {

		for i := 0; i < 16; i++ {

			for j := 0; j < 16; j++ {

				if(led.Matrix[i][j].Hold == 0) {
					led.Matrix[i][j].Color = [3]uint8{0, 0, 255}
				}

			}
			
		}

		time.Sleep(200 * time.Millisecond)

		for i := 0; i < 16; i++ {

			for j := 0; j < 16; j++ {

				if(led.Matrix[i][j].Hold == 0) {
					led.Matrix[i][j].Color = [3]uint8{0, 255, 0}
				}

			}
			
		}

		time.Sleep(200 * time.Millisecond)

	}
}



func cross_color(ID uint8, dt *pa.DataController) {

	x := int8(ID)
	y := int8(ID)
	

	for i := 0; i < 17; i++ {

			if x+1*int8(i) < 16 {

				led.Matrix[x+1*int8(i)][y].Color = [3]uint8{255, 0, 0}
				led.Matrix[x+1*int8(i)][y].Hold = 1

			}
			if (x+1*int8(i)) <= 16 && (x+1*int8(i)) >= 1 {

				led.Matrix[x+1*(int8(i)-1)][y].Color = [3]uint8{0, 0, 0}
				led.Matrix[x+1*(int8(i)-1)][y].Hold = 0

			}

			if x-1*int8(i) >= 0 {

				led.Matrix[x-1*int8(i)][y].Color = [3]uint8{255, 0, 0}
				led.Matrix[x-1*int8(i)][y].Hold = 1
				

			}

			if (x-int8(i)) <= 14 && (x-int8(i)) >= -1  {

				//print("test")
				led.Matrix[x-1*(int8(i)-1)][y].Color = [3]uint8{0, 0, 0}
				led.Matrix[x-1*(int8(i)-1)][y].Hold = 0

			}

			if y+1*int8(i) < 16 {

				led.Matrix[x][y+1*int8(i)].Color = [3]uint8{255, 0, 0}
				led.Matrix[x][y+1*int8(i)].Hold = 1
				

			}

			if (y+1*int8(i)) <= 16 && (y+1*int8(i)) >= 1 {

				led.Matrix[x][y+1*(int8(i)-1)].Color = [3]uint8{0, 0, 0}
				led.Matrix[x][y+1*(int8(i)-1)].Hold = 0

			}

			if y-1*int8(i) >= 0 {

				led.Matrix[x][y-1*int8(i)].Color = [3]uint8{255, 0, 0}
				led.Matrix[x][y-1*int8(i)].Hold = 1
				

			}		

			if (y-1*int8(i)) <= 14 && (y-1*int8(i)) >= -1  {

				led.Matrix[x][y-1*(int8(i)-1)].Color = [3]uint8{0, 0, 0}
				led.Matrix[x][y-1*(int8(i)-1)].Hold = 0

			}	

			frameOut1, frameOut2 := led.ConvertMatrixToFrame()

			dt.WriteData(frameOut1.ConvertFrameOutputToBytes())
			dt.WriteData(frameOut2.ConvertFrameOutputToBytes())

			time.Sleep(200 * time.Millisecond)

		}

}