package protocol

const size int = 16

type ColorModel struct {
	Color [3]uint8
}

type LedMatrixModel struct {
	Matrix [size][size]ColorModel
}

func (l *LedMatrixModel) ConvertMatrixToFrame() (FrameModelOutput, FrameModelOutput) {

	var frame FrameModelOutput
	var frame2 FrameModelOutput

	frame.FrameNumber = 0x01
	frame2.FrameNumber = 0x02

	for column := 0; column < 4; column++ { // On fait la hauteur de dalle (4)
		for lign := 0; lign < 8; lign++ { // On fait la longueur de dalle (8)

			frame.Data = append(frame.Data[:], l.Matrix[column*2][lign*2].Color[0])
			frame.Data = append(frame.Data[:], l.Matrix[column*2][lign*2].Color[1])
			frame.Data = append(frame.Data[:], l.Matrix[column*2][lign*2].Color[2])

			frame.Data = append(frame.Data[:], l.Matrix[column*2][lign*2+1].Color[0])
			frame.Data = append(frame.Data[:], l.Matrix[column*2][lign*2+1].Color[1])
			frame.Data = append(frame.Data[:], l.Matrix[column*2][lign*2+1].Color[2])

			frame.Data = append(frame.Data[:], l.Matrix[column*2+1][lign*2].Color[0])
			frame.Data = append(frame.Data[:], l.Matrix[column*2+1][lign*2].Color[1])
			frame.Data = append(frame.Data[:], l.Matrix[column*2+1][lign*2].Color[2])

			frame.Data = append(frame.Data[:], l.Matrix[column*2+1][lign*2+1].Color[0])
			frame.Data = append(frame.Data[:], l.Matrix[column*2+1][lign*2+1].Color[1])
			frame.Data = append(frame.Data[:], l.Matrix[column*2+1][lign*2+1].Color[2])

		}
	}

	for column := 4; column < 8; column++ { // On fait la hauteur de dalle (4)
		for lign := 0; lign < 8; lign++ { // On fait la longueur de dalle (8)

			frame2.Data = append(frame2.Data[:], l.Matrix[column*2][lign*2].Color[0])
			frame2.Data = append(frame2.Data[:], l.Matrix[column*2][lign*2].Color[1])
			frame2.Data = append(frame2.Data[:], l.Matrix[column*2][lign*2].Color[2])

			frame2.Data = append(frame2.Data[:], l.Matrix[column*2][lign*2+1].Color[0])
			frame2.Data = append(frame2.Data[:], l.Matrix[column*2][lign*2+1].Color[1])
			frame2.Data = append(frame.Data[:], l.Matrix[column*2][lign*2+1].Color[2])

			frame2.Data = append(frame2.Data[:], l.Matrix[column*2+1][lign*2].Color[0])
			frame2.Data = append(frame2.Data[:], l.Matrix[column*2+1][lign*2].Color[1])
			frame2.Data = append(frame2.Data[:], l.Matrix[column*2+1][lign*2].Color[2])

			frame2.Data = append(frame2.Data[:], l.Matrix[column*2+1][lign*2+1].Color[0])
			frame2.Data = append(frame2.Data[:], l.Matrix[column*2+1][lign*2+1].Color[1])
			frame2.Data = append(frame2.Data[:], l.Matrix[column*2+1][lign*2+1].Color[2])

		}
	}

	return frame, frame2
}
