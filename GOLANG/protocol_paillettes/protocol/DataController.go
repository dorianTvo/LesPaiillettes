package protocol

import (
	"fmt"
	"log"
	"net"
)

type DataController struct {
	IP        string
	Port      string
	Broadcast string
	Pc        net.PacketConn
	Err       error

	Broadcast_addr net.Addr
}

func (d *DataController) InitConnection() {
	var err error
	d.Pc, err = net.ListenPacket("udp", ":"+d.Port)
	fmt.Println(":" + d.Port)
	if err != nil {
		log.Fatal(err)
	}

	d.Broadcast_addr, d.Err = net.ResolveUDPAddr("udp", d.Broadcast+":"+d.Port)
	fmt.Println(d.Broadcast + ":" + d.Port)
	if d.Err != nil {
		log.Fatal(d.Err)
	}
}

func (d *DataController) ReadIncommingData() ([]byte, error) {

	buf := make([]byte, 255)
	var addrRecv net.Addr

	_, addrRecv, d.Err = d.Pc.ReadFrom(buf)
	if d.Err != nil {
		return nil, d.Err
	}

	if addrRecv.String() != d.IP+":"+d.Port {
		return buf, nil
	} else {
		return nil, d.Err
	}
}

func (d *DataController) WriteData(data []byte) {
	go d.Pc.WriteTo(data, d.Broadcast_addr)
}

func (d *DataController) StopConnection() {
	d.Pc.Close()
}
