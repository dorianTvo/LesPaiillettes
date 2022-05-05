package protocol

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

type FrameModelOutput struct {
	FrameNumber uint8    // Decomposition of the Frame to avoid the maximum limit of data length
	Data        []uint32 // data to send
}

type FrameModelInput struct {
	ID   uint8   // ID of cell
	Data float32 // data received (mass)
}

func int32ToByte(f uint32) []byte {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, f)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	return buf.Bytes()
}

func (f *FrameModelOutput) ConvertFrameOutputToBytes() []byte {

	var result []byte
	var data []byte

	for i := 0; i < len(f.Data); i++ {

		res := int32ToByte(f.Data[i])
		data = append(data, res...)
	}

	result = append(result, f.FrameNumber)
	result = append(result, data...)

	return result
}

func (f *FrameModelInput) ConvertBytesToFrameInput(data []byte) {
	f.ID = data[0]
	bits := binary.BigEndian.Uint32(data[1:])
	f.Data = math.Float32frombits(bits)
}
