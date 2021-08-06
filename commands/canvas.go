package commands

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers"
)
func RenderCanvas(imgPath string, ppi float64) {
	width := float64(4 * ppi)
	height := float64(6 * ppi)
	buffer := 2.0
	ppi = 100.0
	imageSize := float64(width / 2.0 - buffer * 2) // 2 px buffer on each side

	// load image
	var img image.Image
	var imgConfig image.Config

	if imgReader, err := os.Open(imgPath); err == nil {
		defer imgReader.Close()
		if stat ,_ := imgReader.Stat(); stat.IsDir() {
			fmt.Printf("file not correct format, %s\n", imgPath)
			return
		}
		img, err = jpeg.Decode(imgReader)
		if err != nil {
			fmt.Printf("image is not JPEG formatted.\n")
			return
		}
	} else if err != nil {
		fmt.Printf("error opening file : %s \nerror: %v \n", imgPath, err)
		return
	}
	
	if imgReader, err := os.Open(imgPath); err == nil {
		defer imgReader.Close()
		imgConfig, err = jpeg.DecodeConfig(imgReader)
		if err != nil {
			fmt.Println("Could not decode image.", err)
			return
		}

		if imgConfig.Width != imgConfig.Height {
			fmt.Println("Image is not in 1:1 aspect ratio.")
			return
		}
	}
	
	fmt.Printf("Input image actual size: %dx%d \n",imgConfig.Width, imgConfig.Height)
	fmt.Println("Image Resize", imageSize)
	scalingFactor := float64(imageSize / float64(imgConfig.Width))
	
	// create canvas
	c := canvas.New(width, height)
	ctx := canvas.NewContext(c)
	ctx.SetFillColor(canvas.White)
	ctx.DrawPath(0,0, canvas.Rectangle(c.W, c.H))
	
	ctx.DrawImage(buffer,height/2.0, img, canvas.DPMM(1/scalingFactor))
	ctx.DrawImage(width/2.0+buffer,height/2.0, img, canvas.DPMM(1/scalingFactor))

	ctx.DrawImage(buffer,height/2.0 - imageSize - buffer*2, img, canvas.DPMM(1/scalingFactor))
	ctx.DrawImage(width/2.0+buffer,height/2.0 - imageSize - buffer*2, img, canvas.DPMM(1/scalingFactor))
	
	newFileName := filepath.Dir(imgPath) +
		string(filepath.Separator) +
		"passport_sheet_" +
		filepath.Base(imgPath)

	renderers.Write(newFileName, c)
}