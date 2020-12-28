package main

import (
	"fmt"
	"gocv.io/x/gocv"
)

func main() {
	img := gocv.NewMat()
	cap, err := gocv.VideoCaptureFile("/Users/subway/Downloads/small.mp4")
	if err != nil {
		return
	}
	defer cap.Close()
	if cap.IsOpened() {
		fps := cap.Get(gocv.VideoCaptureFPS)
		wid := int(cap.Get(gocv.VideoCaptureFrameWidth))
		heig := int(cap.Get(gocv.VideoCaptureFrameHeight))

		fmt.Println(fps, wid, heig)
	}

	var i int
	for i = 0; i < 10; i++ {
		if ok := cap.Read(&img); !ok {
			fmt.Println(err)
			return
		}
		filename := "image" +string(i)+ ".jpg"
		if img.Empty() {
			fmt.Println(err)
			return
		}
		gocv.IMWrite(filename, img)
	}
}
