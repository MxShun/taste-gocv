package main

import (
	"gocv.io/x/gocv"
)

func main() {
	sample1()
}

// sample1 https://analyticsindiamag.com/converting-an-image-to-a-cartoon/
func sample1() {
	img := gocv.IMRead("./input.jpg", gocv.IMReadColor)

	imgRGB := img.Clone()
	gocv.CvtColor(img, &imgRGB, gocv.ColorBGRToRGB)
	gocv.IMWrite("./sample1_rgb.jpg", imgRGB)

	imgEdges := imgRGB.Clone()
	gocv.CvtColor(img, &imgEdges, gocv.ColorBGRToGray)
	gocv.MedianBlur(imgEdges, &imgEdges, 5)
	gocv.AdaptiveThreshold(imgEdges, &imgEdges, 255, gocv.AdaptiveThresholdMean, gocv.ThresholdBinary, 9, 9)
	gocv.IMWrite("./sample1_edges.jpg", imgEdges)

	imgColor := imgRGB.Clone()
	gocv.BilateralFilter(img, &imgColor, 5, 250, 250)
	gocv.IMWrite("./sample1_color.jpg", imgColor)

	gocv.BitwiseAndWithMask(imgColor, imgColor, &img, imgEdges)
	gocv.IMWrite("./sample1_output.jpg", img)
}
