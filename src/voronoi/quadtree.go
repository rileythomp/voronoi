package voronoi

import (
	"image"
	"image/color"
	"math"

	"gitlab.com/rileythomp14/voronoi/src/utils"
)

type Quadtree struct {
	name             string
	upLeft, lowRight image.Point
	nw, ne, se, sw   *Quadtree
	site             Site
	color            color.RGBA
	err              float64
	depth            int
}

func NewQuadtree(name string, upLeft image.Point, lowRight image.Point, depth int, img *image.RGBA) *Quadtree {
	qt := &Quadtree{
		name:     name,
		upLeft:   upLeft,
		lowRight: lowRight,
		site:     NewSite(upLeft, lowRight, "", ""),
		depth:    depth,
	}
	qt.color = qt.averageColor(img)
	qt.err = qt.colorError(img)
	return qt
}

func (qt *Quadtree) splitTree(img *image.RGBA) []Site {
	var (
		sites           []Site
		maxColorErr     float64
		maxColorErrNode *Quadtree
	)
	nodeStack := Stack{qt}
	for !nodeStack.IsEmpty() {
		curNode, _ := nodeStack.Pop()
		// is leaf node
		if curNode.nw == nil {
			if curNode.err > maxColorErr {
				maxColorErr = curNode.err
				maxColorErrNode = curNode
			}
			curNode.site.color = curNode.color
			sites = append(sites, curNode.site)
		} else {
			// not leaf, add children to stack
			nodeStack.Push(curNode.nw)
			nodeStack.Push(curNode.ne)
			nodeStack.Push(curNode.se)
			nodeStack.Push(curNode.sw)
		}

	}
	splitNode := maxColorErrNode
	midX := int((splitNode.lowRight.X-splitNode.upLeft.X)/2) + splitNode.upLeft.X
	midY := int((splitNode.lowRight.Y-splitNode.upLeft.Y)/2) + splitNode.upLeft.Y
	depth := splitNode.depth + 1
	splitNode.nw = NewQuadtree("nw", splitNode.upLeft, image.Point{midX, midY}, depth, img)
	splitNode.ne = NewQuadtree("ne", image.Point{midX, splitNode.upLeft.Y}, image.Point{splitNode.lowRight.X, midY}, depth, img)
	splitNode.se = NewQuadtree("se", image.Point{midX, midY}, splitNode.lowRight, depth, img)
	splitNode.sw = NewQuadtree("sw", image.Point{splitNode.upLeft.X, midY}, image.Point{midX, splitNode.lowRight.Y}, depth, img)
	return sites
}

func (qt *Quadtree) colorError(img *image.RGBA) float64 {
	avgRed, avgGreen, avgBlue := int(qt.color.R), int(qt.color.G), int(qt.color.B)
	rSum, gSum, bSum := 0, 0, 0
	for y := qt.upLeft.Y; y < qt.lowRight.Y; y++ {
		for x := qt.upLeft.X; x < qt.lowRight.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			red, green, blue := int(r/257), int(g/257), int(b/257)
			rSum += (avgRed - red) * (avgRed - red)
			gSum += (avgGreen - green) * (avgGreen - green)
			bSum += (avgBlue - blue) * (avgGreen - blue)
		}
	}
	numPix := float64((qt.lowRight.X - qt.upLeft.X) * (qt.lowRight.Y - qt.upLeft.Y))
	rStdDev, gStdDev, bStdDev := math.Sqrt(float64(rSum)/numPix), math.Sqrt(float64(gSum)/numPix), math.Sqrt(float64(bSum)/numPix)
	return float64(rStdDev+gStdDev+bStdDev) / float64(qt.depth*qt.depth)
}

func (qt *Quadtree) averageColor(img *image.RGBA) color.RGBA {
	red, green, blue, alpha := 0, 0, 0, 0
	for y := qt.upLeft.Y; y < qt.lowRight.Y; y++ {
		for x := qt.upLeft.X; x < qt.lowRight.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			red += int(r / 257)
			green += int(g / 257)
			blue += int(b / 257)
			alpha += int(a / 257)
		}
	}
	numPixels := (qt.lowRight.X - qt.upLeft.X) * (qt.lowRight.Y - qt.upLeft.Y)
	if numPixels == 0 {
		return color.RGBA{}
	}
	avgRed := int(red / numPixels)
	avgGreen := int(green / numPixels)
	avgBlue := int(blue / numPixels)
	avgAlpha := int(alpha / numPixels)
	return color.RGBA{
		R: uint8(avgRed),
		G: uint8(avgGreen),
		B: uint8(avgBlue),
		A: uint8(avgAlpha),
	}
}

func (qt *Quadtree) createQuadImg2(lines bool) *image.RGBA {
	img := image.NewRGBA(image.Rectangle{qt.upLeft, qt.lowRight})
	nodeStack := Stack{qt}
	for !nodeStack.IsEmpty() {
		curNode, _ := nodeStack.Pop()
		// is leaf node
		if curNode.nw == nil {
			for y := curNode.upLeft.Y; y <= curNode.lowRight.Y; y++ {
				for x := curNode.upLeft.X; x <= curNode.lowRight.X; x++ {
					if lines && (y == curNode.upLeft.Y || y == curNode.lowRight.Y || x == curNode.upLeft.X || x == curNode.lowRight.X) {
						img.Set(x, y, Black)
					} else {
						img.Set(x, y, curNode.color)
					}
				}
			}
		} else {
			// not leaf, add children to stack
			nodeStack.Push(curNode.nw)
			nodeStack.Push(curNode.ne)
			nodeStack.Push(curNode.se)
			nodeStack.Push(curNode.sw)
		}

	}
	return img
}

func (qt *Quadtree) createCircleImg(lines bool) *image.RGBA {
	img := image.NewRGBA(image.Rectangle{qt.upLeft, qt.lowRight})
	nodeStack := Stack{qt}
	for !nodeStack.IsEmpty() {
		curNode, _ := nodeStack.Pop()
		if curNode.nw == nil {
			width, height := curNode.lowRight.X-curNode.upLeft.X, curNode.lowRight.Y-curNode.upLeft.Y
			midX, midY := curNode.upLeft.X+width/2, curNode.upLeft.Y+height/2
			maxRadius := math.Min(float64(width)/2, float64(height)/2) + 1
			for y := curNode.upLeft.Y; y <= curNode.lowRight.Y; y++ {
				for x := curNode.upLeft.X; x <= curNode.lowRight.X; x++ {
					dy, dx := utils.Abs(y-midY), utils.Abs(x-midX)
					if lines && (y == curNode.upLeft.Y || y == curNode.lowRight.Y || x == curNode.upLeft.X || x == curNode.lowRight.X) {
						img.Set(x, y, Black)
					} else if EuclideanDistance(dy, dx) < int(maxRadius*maxRadius) {
						img.SetRGBA(x, y, curNode.color)
					} else {
						img.SetRGBA(x, y, White)
					}
				}
			}
		} else {
			nodeStack.Push(curNode.nw)
			nodeStack.Push(curNode.ne)
			nodeStack.Push(curNode.se)
			nodeStack.Push(curNode.sw)
		}
	}
	return img
}

// Stack code taken from https://www.educative.io/edpresso/how-to-implement-a-stack-in-golang

type Stack []*Quadtree

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(qt *Quadtree) {
	*s = append(*s, qt) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (*Quadtree, bool) {
	if s.IsEmpty() {
		return nil, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}
