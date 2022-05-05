package protocol

const size int = 16

type LedMatrixModel struct {
	Matrix [size][size]uint32
}

func (l *LedMatrixModel) ConvertMatrixToFrame() (FrameModelOutput, FrameModelOutput) {

	var frame FrameModelOutput
	var frame2 FrameModelOutput

	frame.FrameNumber = 0x01
	frame2.FrameNumber = 0x02

	for column := 0; column < 4; column++ { // On fait la hauteur de dalle (4)
		for lign := 0; lign < 8; lign++ { // On fait la longueur de dalle (8)

			frame.Data = append(frame.Data[:], l.Matrix[column*2][lign*2])
			frame.Data = append(frame.Data[:], l.Matrix[column*2][lign*2+1])
			frame.Data = append(frame.Data[:], l.Matrix[column*2+1][lign*2])
			frame.Data = append(frame.Data[:], l.Matrix[column*2+1][lign*2+1])

		}
	}

	for column := 4; column < 8; column++ { // On fait la hauteur de dalle (4)
		for lign := 0; lign < 8; lign++ { // On fait la longueur de dalle (8)

			frame2.Data = append(frame2.Data[:], l.Matrix[column*2][lign*2])
			frame2.Data = append(frame2.Data[:], l.Matrix[column*2][lign*2+1])
			frame2.Data = append(frame2.Data[:], l.Matrix[column*2+1][lign*2])
			frame2.Data = append(frame2.Data[:], l.Matrix[column*2+1][lign*2+1])

		}
	}

	return frame, frame2
}
