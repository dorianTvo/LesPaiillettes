package protocol

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type FrameModelOutput struct {
	FrameNumber uint8   // Decomposition of the Frame to avoid the maximum limit of data length
	Data        []uint8 // data to send
}

type FrameModelInput struct {
	ID   uint8 // ID of cell
	Data uint8 // data received (mass)
}

func int8ToByte(f uint8) []byte {
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

		res := int8ToByte(f.Data[i])
		data = append(data, res...)
	}

	result = append(result, f.FrameNumber)
	result = append(result, data...)

	return result
}

func (f *FrameModelInput) ConvertBytesToFrameInput(data []byte) {
	f.ID = data[0]
	f.Data = data[1]
}
