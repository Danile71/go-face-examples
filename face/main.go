package main

import (
	"fmt"
	"image"
	"image/color"

	"github.com/Danile71/go-face"
	"github.com/Danile71/go-logger"
	"gocv.io/x/gocv"
)

const (
	cnnModel = "models/mmod_human_face_detector.dat"

	shapeModel  = "models/shape_predictor_68_face_landmarks.dat" // "models/shape_predictor_5_face_landmarks.dat"
	descrModel  = "models/dlib_face_recognition_resnet_model_v1.dat"
	ageModel    = "models/dnn_age_predictor_v1.dat"
	genderModel = "models/dnn_gender_classifier_v1.dat"

	faceImage  = "images/face.jpg"
	facesImage = "images/faces.jpg"
)

var (
	greenColor = color.RGBA{0, 255, 0, 255}
	redColor   = color.RGBA{255, 0, 0, 255}
)

func init() {
	logger.SetLevel(logger.DEBUG)
}

func main() {
	// try to find him
	var descriptor face.Descriptor

	// craete window
	w := gocv.NewWindow("example")
	defer w.Close()

	// Init recognizer
	rec, err := face.NewRecognizer()
	if logger.OnError(err) {
		return
	}
	// close it
	defer rec.Close()

	// Load shape model
	if err = rec.SetShapeModel(shapeModel); logger.OnError(err) {
		return
	}

	// Load description model
	if err = rec.SetDescriptorModel(descrModel); logger.OnError(err) {
		return
	}

	// Load age model
	if err = rec.SetAgeModel(ageModel); logger.OnError(err) {
		return
	}

	// Load gener model
	if err = rec.SetGenderModel(genderModel); logger.OnError(err) {
		return
	}

	// Load CNN model
	if err = rec.SetCNNModel(cnnModel); logger.OnError(err) {
		return
	}

	// load first image
	img1 := gocv.IMRead(faceImage, gocv.IMReadUnchanged)
	defer img1.Close()

	// load second image
	img2 := gocv.IMRead(facesImage, gocv.IMReadUnchanged)
	defer img2.Close()

	// copy bg to draw
	background := img2.Clone()
	defer background.Close()

	// try detect faces
	faces, err := rec.DetectFromMatCNN(img1)
	if logger.OnError(err) {
		return
	}

	for _, f := range faces {
		defer f.Close()

		// get face description
		if err = rec.Recognize(&f); err != nil {
			return
		}

		// predict face age
		rec.GetAge(&f)

		// predict face gender
		rec.GetGender(&f)

		// set descriptor
		descriptor = f.Descriptor

		// draw rect
		gocv.Rectangle(&img1, f.Rectangle, greenColor, 2)

		gocv.PutText(&img1, fmt.Sprintf("%s:%d y.o.", f.Gender, f.Age), image.Point{f.Rectangle.Min.X - 20, f.Rectangle.Min.Y - 5}, gocv.FontHersheyPlain, 1, redColor, 1)

	}
	gocv.PutText(&img1, "press any key...", image.Point{0, img1.Cols() - 30}, gocv.FontHersheyPlain, 1, redColor, 1)

	w.IMShow(img1)

	fmt.Println("press any key to continue...")

	for {
		if key := w.WaitKey(1000); key != -1 {
			break
		}
	}

	faces, err = rec.DetectFromMatCNN(img2)
	if logger.OnError(err) {
		return
	}

	for _, f := range faces {
		defer f.Close()

		if err = rec.Recognize(&f); err != nil {
			return
		}

		rec.GetAge(&f)
		rec.GetGender(&f)
		gocv.Rectangle(&background, f.Rectangle, greenColor, 2)

		gocv.PutText(&background, fmt.Sprintf("%s:%d y.o.", f.Gender, f.Age), image.Point{f.Rectangle.Min.X - 20, f.Rectangle.Min.Y - 5}, gocv.FontHersheyPlain, 1, redColor, 1)

		dist := face.SquaredEuclideanDistance(f.Descriptor, descriptor)

		c := redColor
		if dist < 0.1 {
			c = greenColor
		}

		gocv.PutText(&background, fmt.Sprintf("%f", dist), image.Point{f.Rectangle.Min.X, f.Rectangle.Max.Y}, gocv.FontHersheyPlain, 1, c, 1)
	}

	w.IMShow(background)

	fmt.Println("press any key to exit...")

	for {
		if key := w.WaitKey(1000); key != -1 {
			return
		}
	}
}
