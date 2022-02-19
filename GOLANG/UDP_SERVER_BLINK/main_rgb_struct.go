package main

import (
	"encoding/json"
	"log"
	"net"
	"time"
)

type Message struct {
	RgbLedArray [64][3]int `json:"rgb"`
}

func main() {

	var mess Message
	rgb := [3]int{255, 255, 255}

	for i := 0; i < 64; i++ {
		mess.RgbLedArray[i] = rgb
	}

	pc, err := net.ListenPacket("udp", ":1053")
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()

	addr, err := net.ResolveUDPAddr("udp", "10.42.0.255:1053")
	if err != nil {
		log.Fatal(err)
	}

	for {

		/*go serve(pc, addr, "ON")
		time.Sleep(33 * time.Millisecond)
		go serve(pc, addr, "OFF")
		time.Sleep(33 * time.Millisecond)*/

		data, _ := json.Marshal(mess)
		go serve(pc, addr, data)
		time.Sleep(1 * time.Millisecond)
	}
}

func serve(pc net.PacketConn, addr net.Addr, message []byte) {
	pc.WriteTo(message, addr)
}
