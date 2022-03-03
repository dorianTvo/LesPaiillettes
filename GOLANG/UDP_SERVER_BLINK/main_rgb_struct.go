package main

import (
	"encoding/json"
	"log"
	"net"
	"time"
)

type Cellule struct {
	RgbColor [3]int
}

type Dalle struct {
	Cellules [4]Cellule
}

type Dalles struct {
	Matrice [8][8]Dalle
}

func main() {

	var dalle1 Dalle
	var dalle2 Dalle
	var dalle3 Dalle
	var dalle4 Dalle
	var dalles [4]Dalle

	pc, err := net.ListenPacket("udp", ":1053")
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()

	addr, err := net.ResolveUDPAddr("udp", "192.168.0.255:1053")
	if err != nil {
		log.Fatal(err)
	}

	dalle1.Cellules[0].RgbColor = [3]int{0, 0, 200}
	dalle1.Cellules[1].RgbColor = [3]int{0, 200, 200}
	dalle1.Cellules[2].RgbColor = [3]int{0, 200, 200}
	dalle1.Cellules[3].RgbColor = [3]int{0, 0, 200}

	dalle2.Cellules[0].RgbColor = [3]int{200, 0, 200}
	dalle2.Cellules[1].RgbColor = [3]int{200, 200, 200}
	dalle2.Cellules[2].RgbColor = [3]int{200, 200, 200}
	dalle2.Cellules[3].RgbColor = [3]int{200, 0, 200}

	dalle3.Cellules[0].RgbColor = [3]int{0, 200, 0}
	dalle3.Cellules[1].RgbColor = [3]int{200, 200, 0}
	dalle3.Cellules[2].RgbColor = [3]int{200, 200, 0}
	dalle3.Cellules[3].RgbColor = [3]int{0, 200, 0}

	dalle4.Cellules[0].RgbColor = [3]int{200, 0, 0}
	dalle4.Cellules[1].RgbColor = [3]int{200, 0, 200}
	dalle4.Cellules[2].RgbColor = [3]int{200, 0, 200}
	dalle4.Cellules[3].RgbColor = [3]int{200, 0, 0}

	for {

		dalles[0] = dalle1
		dalles[1] = dalle2
		dalles[2] = dalle3
		dalles[3] = dalle4
		data, _ := json.Marshal(dalles)
		go serve(pc, addr, data)

		time.Sleep(350 * time.Millisecond)

		dalles[0] = dalle2
		dalles[1] = dalle3
		dalles[2] = dalle4
		dalles[3] = dalle1
		data, _ = json.Marshal(dalles)
		go serve(pc, addr, data)

		time.Sleep(350 * time.Millisecond)

		dalles[0] = dalle3
		dalles[1] = dalle4
		dalles[2] = dalle1
		dalles[3] = dalle2
		data, _ = json.Marshal(dalles)
		go serve(pc, addr, data)

		time.Sleep(350 * time.Millisecond)

		dalles[0] = dalle4
		dalles[1] = dalle1
		dalles[2] = dalle2
		dalles[3] = dalle3
		data, _ = json.Marshal(dalles)
		go serve(pc, addr, data)

		time.Sleep(350 * time.Millisecond)

	}
}

func serve(pc net.PacketConn, addr net.Addr, message []byte) {
	pc.WriteTo(message, addr)
}

/*

for {

		dalle1.Cellules[0].RgbColor = [3]int{150, 150, 150}
		dalle1.Cellules[1].RgbColor = [3]int{150, 0, 0}
		dalle1.Cellules[2].RgbColor = [3]int{0, 150, 0}
		dalle1.Cellules[3].RgbColor = [3]int{0, 0, 150}

		dalle2.Cellules[0].RgbColor = [3]int{0, 0, 150}
		dalle2.Cellules[1].RgbColor = [3]int{150, 150, 150}
		dalle2.Cellules[2].RgbColor = [3]int{150, 0, 0}
		dalle2.Cellules[3].RgbColor = [3]int{0, 150, 0}

		dalle3.Cellules[0].RgbColor = [3]int{0, 150, 0}
		dalle3.Cellules[1].RgbColor = [3]int{0, 0, 150}
		dalle3.Cellules[2].RgbColor = [3]int{150, 150, 150}
		dalle3.Cellules[3].RgbColor = [3]int{150, 0, 0}

		dalle4.Cellules[0].RgbColor = [3]int{150, 0, 0}
		dalle4.Cellules[1].RgbColor = [3]int{0, 150, 0}
		dalle4.Cellules[2].RgbColor = [3]int{0, 0, 150}
		dalle4.Cellules[3].RgbColor = [3]int{150, 150, 150}

		dalles[0] = dalle1
		dalles[1] = dalle2
		dalles[2] = dalle3
		dalles[3] = dalle4

		data, _ := json.Marshal(dalles)
		go serve(pc, addr, data)

		time.Sleep(350 * time.Millisecond)

		dalle1.Cellules[0].RgbColor = [3]int{150, 150, 150}
		dalle1.Cellules[1].RgbColor = [3]int{150, 0, 0}
		dalle1.Cellules[2].RgbColor = [3]int{0, 150, 0}
		dalle1.Cellules[3].RgbColor = [3]int{0, 0, 150}

		dalle2.Cellules[0].RgbColor = [3]int{0, 0, 150}
		dalle2.Cellules[1].RgbColor = [3]int{150, 150, 150}
		dalle2.Cellules[2].RgbColor = [3]int{150, 0, 0}
		dalle2.Cellules[3].RgbColor = [3]int{0, 150, 0}

		dalle3.Cellules[0].RgbColor = [3]int{0, 150, 0}
		dalle3.Cellules[1].RgbColor = [3]int{0, 0, 150}
		dalle3.Cellules[2].RgbColor = [3]int{150, 150, 150}
		dalle3.Cellules[3].RgbColor = [3]int{150, 0, 0}

		dalle4.Cellules[0].RgbColor = [3]int{150, 0, 0}
		dalle4.Cellules[1].RgbColor = [3]int{0, 150, 0}
		dalle4.Cellules[2].RgbColor = [3]int{0, 0, 150}
		dalle4.Cellules[3].RgbColor = [3]int{150, 150, 150}

		dalles[0] = dalle4
		dalles[1] = dalle1
		dalles[2] = dalle2
		dalles[3] = dalle3

		data, _ = json.Marshal(dalles)
		go serve(pc, addr, data)

		time.Sleep(350 * time.Millisecond)

		dalle1.Cellules[0].RgbColor = [3]int{150, 150, 150}
		dalle1.Cellules[1].RgbColor = [3]int{150, 0, 0}
		dalle1.Cellules[2].RgbColor = [3]int{0, 150, 0}
		dalle1.Cellules[3].RgbColor = [3]int{0, 0, 150}

		dalle2.Cellules[0].RgbColor = [3]int{0, 0, 150}
		dalle2.Cellules[1].RgbColor = [3]int{150, 150, 150}
		dalle2.Cellules[2].RgbColor = [3]int{150, 0, 0}
		dalle2.Cellules[3].RgbColor = [3]int{0, 150, 0}

		dalle3.Cellules[0].RgbColor = [3]int{0, 150, 0}
		dalle3.Cellules[1].RgbColor = [3]int{0, 0, 150}
		dalle3.Cellules[2].RgbColor = [3]int{150, 150, 150}
		dalle3.Cellules[3].RgbColor = [3]int{150, 0, 0}

		dalle4.Cellules[0].RgbColor = [3]int{150, 0, 0}
		dalle4.Cellules[1].RgbColor = [3]int{0, 150, 0}
		dalle4.Cellules[2].RgbColor = [3]int{0, 0, 150}
		dalle4.Cellules[3].RgbColor = [3]int{150, 150, 150}

		dalles[0] = dalle3
		dalles[1] = dalle4
		dalles[2] = dalle1
		dalles[3] = dalle2

		data, _ = json.Marshal(dalles)
		go serve(pc, addr, data)

		time.Sleep(350 * time.Millisecond)

		dalle1.Cellules[0].RgbColor = [3]int{150, 150, 150}
		dalle1.Cellules[1].RgbColor = [3]int{150, 0, 0}
		dalle1.Cellules[2].RgbColor = [3]int{0, 150, 0}
		dalle1.Cellules[3].RgbColor = [3]int{0, 0, 150}

		dalle2.Cellules[0].RgbColor = [3]int{0, 0, 150}
		dalle2.Cellules[1].RgbColor = [3]int{150, 150, 150}
		dalle2.Cellules[2].RgbColor = [3]int{150, 0, 0}
		dalle2.Cellules[3].RgbColor = [3]int{0, 150, 0}

		dalle3.Cellules[0].RgbColor = [3]int{0, 150, 0}
		dalle3.Cellules[1].RgbColor = [3]int{0, 0, 150}
		dalle3.Cellules[2].RgbColor = [3]int{150, 150, 150}
		dalle3.Cellules[3].RgbColor = [3]int{150, 0, 0}

		dalle4.Cellules[0].RgbColor = [3]int{150, 0, 0}
		dalle4.Cellules[1].RgbColor = [3]int{0, 150, 0}
		dalle4.Cellules[2].RgbColor = [3]int{0, 0, 150}
		dalle4.Cellules[3].RgbColor = [3]int{150, 150, 150}

		dalles[0] = dalle2
		dalles[1] = dalle3
		dalles[2] = dalle4
		dalles[3] = dalle1

		data, _ = json.Marshal(dalles)
		go serve(pc, addr, data)

		time.Sleep(350 * time.Millisecond)
	}

*/
