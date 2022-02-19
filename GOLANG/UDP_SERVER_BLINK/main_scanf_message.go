package main2

import (
	"fmt"
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

		var message string
		fmt.Scanf("%s", &message)

		addr, err := net.ResolveUDPAddr("udp", "10.42.0.255:1053")
		if err != nil {
			log.Fatal(err)
		}

		go serve(pc, addr, message)
	}
}

func serve(pc net.PacketConn, addr net.Addr, message string) {

	pc.WriteTo([]byte(message), addr)
	fmt.Println("Envoi de: " + message)
}
