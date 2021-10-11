package main

import (
	"fmt"
	"image/color"

	"github.com/Danile71/go-face"
	"github.com/Danile71/go-logger"
	"gocv.io/x/gocv"
)

const (
	cnnModel = "models/mmod_plate_detector.dat"

	plateImage = "images/numbers.jpg"
)

var (
	greenColor = color.RGBA{0, 255, 0, 255}
	redColor   = color.RGBA{255, 0, 0, 255}
)

func init() {
	logger.SetLevel(logger.DEBUG)
}

func main() {
	// create window
	w := gocv.NewWindow("example")
	defer w.Close()

	// Init recognizer
	rec, err := face.NewRecognizer()
	if logger.OnError(err) {
		return
	}
	// close it
	defer rec.Close()

	// Load CNN model
	if err = rec.SetCNNModel(cnnModel); logger.OnError(err) {
		return
	}

	// load first image
	img := gocv.IMRead(plateImage, gocv.IMReadUnchanged)
	defer img.Close()

	// try detect plates
	plates, err := rec.DetectFromMatCNN(img)
	if logger.OnError(err) {
		return
	}

	for _, p := range plates {
		defer p.Close()

		// draw rect
		gocv.Rectangle(&img, p.Rectangle, greenColor, 2)
	}

	w.IMShow(img)

	fmt.Println("press any key to exit...")

	for {
		if key := w.WaitKey(1000); key != -1 {
			return
		}
	}
}
