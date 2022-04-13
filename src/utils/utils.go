package utils

import (
	"bufio"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
	"strconv"
	"strings"
)

func Contains(strs []string, str string) bool {
	for i := range strs {
		if strs[i] == str {
			return true
		}
	}
	return false
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Abs(n int) int {
	if n < 0 {
		n *= -1
	}
	return n
}

func IsHexColor(c string) bool {
	if len(c) != 6 {
		return false
	}
	for _, hex := range c {
		if !strings.Contains("0123456789ABCDEFabcdef", string(hex)) {
			return false
		}
	}
	return true
}

func GetUserInput(msg string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(msg)
	input, _ := reader.ReadString('\n')
	return strings.Trim(input, " \n")
}

func GetImageFromPath(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	return img, err
}

// imageToRGBA taken from https://stackoverflow.com/questions/61720744/how-to-convert-picture-to-image-rgba
func ImageToRGBA(src image.Image) *image.RGBA {
	// No conversion needed if image is an *image.RGBA.
	if dst, ok := src.(*image.RGBA); ok {
		return dst
	}
	// Use the image/draw package to convert to *image.RGBA.
	b := src.Bounds()
	dst := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(dst, dst.Bounds(), src, b.Min, draw.Src)
	return dst
}

// Outputs a png image
func OutputPngFrame(dir string, n int, img *image.RGBA) error {
	name := strconv.Itoa(n)
	if n < 100 {
		name = "0" + name
		if n < 10 {
			name = "0" + name
		}
	}
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		if err := os.Mkdir(dir, os.ModePerm); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	imgName := dir + "/" + name + ".png"
	f, err := os.Create(imgName)
	if err != nil {
		return fmt.Errorf("error creating %s: %s", imgName, err.Error())
	}
	defer f.Close()
	if err = png.Encode(f, img); err != nil {
		return fmt.Errorf("error encoding png %s: %s", imgName, err.Error())
	}
	return nil
}
