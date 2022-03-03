package main

import (
	"log"
	"net"
)

func main() {

	pc, err := net.ListenPacket("udp", ":1053")
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()

	for {
		buf := make([]byte, 255)
		var err error
		var addrRecv net.Addr

		_, addrRecv, err = pc.ReadFrom(buf)
		if err != nil {
			continue
		}

		if addrRecv.String() != "10.42.0.1:1053" {

			message_client := string(buf)
			//fmt.Println("Reception de: " + message_client)

			var addr net.Addr
			addr, err = net.ResolveUDPAddr("udp", "10.42.0.255:1053")
			if err != nil {
				log.Fatal(err)
			}

			go serve(pc, addr, []byte(message_client))
		}
	}
}

func serve(pc net.PacketConn, addr net.Addr, message []byte) {
	pc.WriteTo(message, addr)
}
