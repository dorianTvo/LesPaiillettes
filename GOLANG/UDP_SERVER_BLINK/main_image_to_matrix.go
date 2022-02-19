package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
)

func main() {
	/*filename := "VIDEO.mp4"
	width := 256
	height := 256
	cmd := exec.Command("ffmpeg", "-ss", "TIME", "-i", filename, "-vframes", "1", "-s", fmt.Sprintf("%dx%d", width, height), "-f", "singlejpeg", "-")
	var buffer bytes.Buffer
	cmd.Stdout = &buffer
	if cmd.Run() != nil {
		panic("could not generate frame")
	}*/

	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	file, err := os.Open("./image.jpg")

	if err != nil {
		fmt.Println("Error Open File")
		os.Exit(1)
	}

	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Error lecture image")
	}

	r, g, b, _ := img.At(0, 0).RGBA()
	println("red: " + fmt.Sprint(r/256) + "green: " + fmt.Sprint(g/256) + "blue: " + fmt.Sprint(b/256))

}
