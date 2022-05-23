package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	//"fmt"

	"math/rand"

	"image/color"
	"time"

	pa "paillettes/protocol"
)

var led pa.LedMatrixModel

var mode = 0
var mode_press = 0

const (
	screenWidth  = 440
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

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
        //ebitenutil.DebugPrint(screen, "You're pressing the 'LEFT' mouse button.")
	    // Get the x, y position of the cursor from the CursorPosition() function
	    x, y := ebiten.CursorPosition()
	    
	    // Display the information with "X: xx, Y: xx" format
	    //ebitenutil.DebugPrint(screen, fmt.Sprintf("X: %d, Y: %d", x, y))

	    if(x > 300 && x<353 && y>40 && y<56){
	    	mode = 0
	    }

	    if(x > 300 && x<353 && y>59 && y<75){
	    	mode = 1
	    }

	    if(x > 300 && x<353 && y>78 && y<93){
	    	mode = 2
	    }

	    if(x > 300 && x<353 && y>100 && y<115){
	    	mode = 3
	    }

	    if(x > 300 && x<353 && y>121 && y<135){
	    	mode = 4
	    }

	    if(x > 300 && x<400 && y>180 && y<193){
	    	mode_press = 0
	    }

	    if(x > 300 && x<400 && y>202 && y<214){
	    	mode_press = 1
	    }

	    if(x > 300 && x<400 && y>220 && y<237){
	    	mode_press = 2
	    }

	    if(x > 300 && x<400 && y>240 && y<260){
	    	mode_press = 3
	    }

	    



    }


	ebitenutil.DebugPrintAt(screen,"Select Mode :", 300, 20)
	ebitenutil.DebugPrintAt(screen,"- OFF", 310, 40)
	ebitenutil.DebugPrintAt(screen,"- Mode 1", 310, 60)
	ebitenutil.DebugPrintAt(screen,"- Mode 2", 310, 80)
	ebitenutil.DebugPrintAt(screen,"- Mode 3", 310, 100)
	ebitenutil.DebugPrintAt(screen,"- Mode 4", 310, 120)

	ebitenutil.DebugPrintAt(screen,"Select Mode Pression :", 300, 160)
	ebitenutil.DebugPrintAt(screen,"- OFF", 310, 180)
	ebitenutil.DebugPrintAt(screen,"- Cross", 310, 200)
	ebitenutil.DebugPrintAt(screen,"- Turn off", 310, 220)
	ebitenutil.DebugPrintAt(screen,"- Random Color", 310, 240)


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

	mode = 0

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
	//go envoie_trame(dt)

	for {

		data, _ := dt.ReadIncommingData() // Lectures des octects de la trame

		frameIn := new(pa.FrameModelInput)     // Definition de la structure de données
		frameIn.ConvertBytesToFrameInput(data) //Initialisation de la structure à partir des octets reçu

		print(frameIn.ID)
		go cross_color(frameIn.ID, dt)


		time.Sleep(10 * time.Millisecond)
	}
}

func envoie_trame(dt *pa.DataController) {

	for{
		frameOut1, frameOut2 := led.ConvertMatrixToFrame()

		dt.WriteData(frameOut1.ConvertFrameOutputToBytes())
		dt.WriteData(frameOut2.ConvertFrameOutputToBytes())

		time.Sleep(20 * time.Millisecond)
	}
}

func pross3() {

	//Couleur de fond

	ic := -1
	jc := 0
	couleur := 1

	for {

		if(mode == 0){


			for i := 0; i < 16; i++ {

				for j := 0; j < 16; j++ {

					if(led.Matrix[i][j].Hold == 0) {
						led.Matrix[i][j].Color = [3]uint8{0, 0, 0}
					}

				}
				
			}

			time.Sleep(200 * time.Millisecond)

			
		}

		if(mode == 1){


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

		if(mode == 2){


			for i := 0; i < 16; i++ {

				for j := 0; j < 16; j++ {

					if(led.Matrix[i][j].Hold == 0) {
						led.Matrix[i][j].Color = [3]uint8{0, 255, 255}
					}

				}
				
			}

			time.Sleep(200 * time.Millisecond)

			for i := 0; i < 16; i++ {

				for j := 0; j < 16; j++ {

					if(led.Matrix[i][j].Hold == 0) {
						led.Matrix[i][j].Color = [3]uint8{255, 255, 0}
					}

				}
				
			}

			time.Sleep(200 * time.Millisecond)
		}

		if(mode == 3){

			ic++

			if(ic>=16){
				jc++
				ic = 0
			}
				
			if(jc>=16){
				jc = 0
			}

			

			if(couleur == 2){
				if(led.Matrix[ic][jc].Hold == 0) {
				led.Matrix[ic][jc].Color = [3]uint8{0, 0, 255}
				}
			}

			if(couleur == 1){
				if(led.Matrix[ic][jc].Hold == 0) {
				led.Matrix[ic][jc].Color = [3]uint8{0, 255, 0}
				}
			}

			if(ic == 15 && jc==15){
				if(couleur == 2){
					couleur = 1
				} else {
					couleur = 2
				}
			}
			

			time.Sleep(20 * time.Millisecond)

				
		}

		if(mode == 4){

			i := uint8(rand.Intn(16))
			j := uint8(rand.Intn(16))

	

			if(led.Matrix[i][j].Hold == 0) {
				led.Matrix[i][j].Color = [3]uint8{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255))}
			}	

			time.Sleep(20 * time.Millisecond)

				
		}


	}
}



func cross_color(ID uint8, dt *pa.DataController) {

	x := int8(ID)
	y := int8(ID)
	

	if(mode_press == 0){
		//OFF Mode
	}

	if(mode_press == 1){

		for i := 0; i < 17; i++ {

			if x+1*int8(i) < 16 {

				led.Matrix[x+1*int8(i)][y].Color = [3]uint8{255, 0, 0}
				led.Matrix[x+1*int8(i)][y].Hold = 1

			}
			if (x+1*int8(i)) <= 16 && (x+1*int8(i)) >= 1 {

				//led.Matrix[x+1*(int8(i)-1)][y].Color = [3]uint8{0, 0, 0}
				led.Matrix[x+1*(int8(i)-1)][y].Hold = 0

			}

			if x-1*int8(i) >= 0 {

				led.Matrix[x-1*int8(i)][y].Color = [3]uint8{255, 0, 0}
				led.Matrix[x-1*int8(i)][y].Hold = 1
				

			}

			if (x-int8(i)) <= 14 && (x-int8(i)) >= -1  {

				//print("test")
				//led.Matrix[x-1*(int8(i)-1)][y].Color = [3]uint8{0, 0, 0}
				led.Matrix[x-1*(int8(i)-1)][y].Hold = 0

			}

			if y+1*int8(i) < 16 {

				led.Matrix[x][y+1*int8(i)].Color = [3]uint8{255, 0, 0}
				led.Matrix[x][y+1*int8(i)].Hold = 1
				

			}

			if (y+1*int8(i)) <= 16 && (y+1*int8(i)) >= 1 {

				//led.Matrix[x][y+1*(int8(i)-1)].Color = [3]uint8{0, 0, 0}
				led.Matrix[x][y+1*(int8(i)-1)].Hold = 0

			}

			if y-1*int8(i) >= 0 {

				led.Matrix[x][y-1*int8(i)].Color = [3]uint8{255, 0, 0}
				led.Matrix[x][y-1*int8(i)].Hold = 1
				

			}		

			if (y-1*int8(i)) <= 14 && (y-1*int8(i)) >= -1  {

				//led.Matrix[x][y-1*(int8(i)-1)].Color = [3]uint8{0, 0, 0}
				led.Matrix[x][y-1*(int8(i)-1)].Hold = 0

			}	

			

			time.Sleep(200 * time.Millisecond)

		}

	}

	if(mode_press == 2){

		led.Matrix[x][y].Color = [3]uint8{0,0,0}
	}

	if(mode_press == 3){
		led.Matrix[x][y].Color = [3]uint8{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255))}
	}


	

}