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

	plateVideo = "video/video.mp4"
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

	img := gocv.NewMat()
	defer img.Close()

	frames, err := gocv.VideoCaptureFile(plateVideo)
	if logger.OnError(err) {
		return
	}

	trackers := make(map[int]*face.Tracker)

	for {
		if ok := frames.Read(&img); !ok {
			return
		}

		if len(trackers) == 0 {
			// try detect plates
			plates, err := rec.DetectFromMatCNN(img)
			if logger.OnError(err) {
				return
			}

			for _, p := range plates {
				tracker, err := face.NewTracker()
				if logger.OnError(err) {
					return
				}

				tracker.StartMat(img, p.Rectangle)

				trackers[len(trackers)] = tracker

				// draw rect
				gocv.Rectangle(&img, p.Rectangle, greenColor, 2)

				p.Close()
			}
		}

		for key, tracker := range trackers {
			conf, err := tracker.UpdateMat(img)
			if conf < 3 || err != nil {
				delete(trackers, key)
				continue
			}

			position, err := tracker.Position()
			if err != nil {
				delete(trackers, key)
				continue
			}

			// draw rect
			gocv.Rectangle(&img, position, greenColor, 2)
		}

		w.IMShow(img)

		if key := w.WaitKey(1); key != -1 {
			return
		}

	}

	fmt.Println("press any key to exit...")

	for {
		if key := w.WaitKey(1000); key != -1 {
			return
		}
	}
}
