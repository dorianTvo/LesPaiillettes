package main

import (
	pa "paillettesprotocol"
	"time"
)

func main2() {

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

	time.Sleep(1 * time.Second)

	dt.StopConnection()
}
