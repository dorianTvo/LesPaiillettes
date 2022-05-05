package protocol

const size int = 16

type LedMatrixModel struct {
	Matrix [size][size]uint32
}

func (l *LedMatrixModel) ConvertMatrixToFrame() (FrameModelOutput, FrameModelOutput) {

	var frame FrameModelOutput
	var frame2 FrameModelOutput

	frame.FrameNumber = 1
	frame.FrameNumber = 2

	for j := 0; j < size/2; j++ {
		for i := 0; i < size; i++ {
			frame.Data = append(frame.Data[:], l.Matrix[j][i])
		}
	}
	for j := size / 2; j < size; j++ {
		for i := 0; i < size; i++ {
			frame2.Data = append(frame2.Data[:], l.Matrix[j][i])
		}
	}

	return frame, frame2
}
