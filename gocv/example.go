package main

import (
	"fmt"
	"gocv.io/x/gocv"
)

func main() {
	deviceID := 0

	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer  webcam.Close()

	window := gocv.NewWindow("Face Detect")
	defer window.Close()

	img := gocv.NewMat()
	defer  img.Close()

	// blue := color.RGBA{0,0,255,0}

	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	if !classifier.Load("") {
		fmt.Println("")
	}

	for {

	}
}
