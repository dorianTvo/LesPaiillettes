package main

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
		buf := make([]byte, 255)
		var err error
		var addrRecv net.Addr

		_, addrRecv, err = pc.ReadFrom(buf)
		if err != nil {
			continue
		}

		if addrRecv.String() != "10.42.0.1:1053" {

			message_client := string(buf)
			fmt.Println("Reception de: " + message_client)

			var addr net.Addr
			fmt.Println(addrRecv.String())
			addr, err = net.ResolveUDPAddr("udp", "10.42.0.255:1053")
			if err != nil {
				log.Fatal(err)
			}

			go serve(pc, addr)
		}
	}
}

func serve(pc net.PacketConn, addr net.Addr) {

	message_server := "hello from server\n"
	pc.WriteTo([]byte(message_server), addr)
	fmt.Println("Envoi de: " + message_server)
}
