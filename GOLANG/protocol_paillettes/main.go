package main

import (
	pa "paillettes/protocol"
	"time"
)

func main() {

	dt := new(pa.DataController)
	ledMatrix := new(pa.LedMatrixModel)

	ledMatrix.Matrix[2][3] = 200

	frameOut1, frameOut2 := ledMatrix.ConvertMatrixToFrame()

	dt.IP = "127.0.0.1"
	dt.Port = "1053"
	dt.Broadcast = "192.168.0.255"

	dt.InitConnection()

	dt.WriteData(frameOut1.ConvertFrameOutputToBytes())
	dt.WriteData(frameOut2.ConvertFrameOutputToBytes())

	data, _ := dt.ReadIncommingData() // Lectures des octects de la trame

	frameIn := new(pa.FrameModelInput)     // Definition de la structure de données
	frameIn.ConvertBytesToFrameInput(data) //Initialisation de la structure à partir des octets reçu

	time.Sleep(1 * time.Second)

	dt.StopConnection()
}
