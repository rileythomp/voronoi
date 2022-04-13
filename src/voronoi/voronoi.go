package voronoi

import (
	"bytes"
	"image"
	"image/color"
	"math"
	"os/exec"

	"gitlab.com/rileythomp14/voronoi/src/utils"
)

const (
	MaxInt        = int(^uint32(0))
	DefaultWidth  = 512
	DefaultHeight = 512
	DefaultFrames = 100
)

var (
	Black = color.RGBA{0, 0, 0, 0}
	White = color.RGBA{255, 255, 255, 255}
)

type (
	AnimationFunc func(string, int, bool, DistanceFunc) error

	StainedGlass struct {
		image *image.RGBA
		sites []Site
		qt    *Quadtree
	}
)

func VoronoiDiagram(numSites int, sitesVisible bool, c1, c2, dist string) *image.RGBA {
	distance, ok := distances[dist]
	if !ok {
		return nil
	}
	w, h := DefaultWidth, DefaultHeight
	tl, br := image.Point{0, 0}, image.Point{w, h}
	numSites = utils.Min(MaxSites, numSites)
	sites := make([]Site, numSites)
	for i := range sites {
		sites[i] = NewSite(tl, br, c1, c2)
	}
	return generateVoronoiImg(w, h, sites, distance, sitesVisible)
}

func HTreeImage(splits int) *image.RGBA {
	w, h := DefaultWidth, DefaultHeight
	tl, br := image.Point{0, 0}, image.Point{w, h}
	img := image.NewRGBA(image.Rectangle{tl, br})
	htree := &Fractal{w / 2, int(float64(h/2) / math.Sqrt(2)), w / 2, h / 2, nil, nil, nil, nil}
	for i := 0; i < splits; i++ {
		htree.Split(HTree)
	}
	stack := FStack{htree}
	for !stack.IsEmpty() {
		cur, _ := stack.Pop()
		cur.Draw(img, HTree)
		if cur.ne != nil {
			stack.Push(cur.ne)
		}
		if cur.nw != nil {
			stack.Push(cur.nw)
		}
		if cur.se != nil {
			stack.Push(cur.se)
		}
		if cur.sw != nil {
			stack.Push(cur.sw)
		}
	}
	return img
}

func TSquareImage(splits int) *image.RGBA {
	w, h := DefaultWidth, DefaultHeight
	tl, br := image.Point{0, 0}, image.Point{w, h}
	img := image.NewRGBA(image.Rectangle{tl, br})
	tsquare := &Fractal{w / 3, h / 3, w / 2, h / 2, nil, nil, nil, nil}
	for i := 0; i < splits; i++ {
		tsquare.Split(TSquare3)
	}
	stack := FStack{tsquare}
	for !stack.IsEmpty() {
		cur, _ := stack.Pop()
		cur.Draw(img, TSquare3)
		if cur.ne != nil {
			stack.Push(cur.ne)
		}
		if cur.nw != nil {
			stack.Push(cur.nw)
		}
		if cur.se != nil {
			stack.Push(cur.se)
		}
		if cur.sw != nil {
			stack.Push(cur.sw)
		}
	}
	return img
}

func FractalImage(splits int) *image.RGBA {
	w, h := DefaultWidth, DefaultHeight
	tl, br := image.Point{0, 0}, image.Point{w, h}
	img := image.NewRGBA(image.Rectangle{tl, br})
	tsquare := &Fractal{w / 2, h / 2, w / 2, h / 2, nil, nil, nil, nil}
	for i := 0; i < splits; i++ {
		tsquare.Split(TSquare2)
	}
	stack := FStack{tsquare}
	for !stack.IsEmpty() {
		cur, _ := stack.Pop()
		cur.Draw(img, TSquare2)
		if cur.ne != nil {
			stack.Push(cur.ne)
		}
		if cur.nw != nil {
			stack.Push(cur.nw)
		}
		if cur.se != nil {
			stack.Push(cur.se)
		}
		if cur.sw != nil {
			stack.Push(cur.sw)
		}
	}
	return img
}

func MovingVoronoi(numSites int, sitesVisible bool, c1, c2, dist string) (string, string) {
	distance, ok := distances[dist]
	if !ok {
		return "", ""
	}
	w, h := DefaultWidth, DefaultHeight
	tl, br := image.Point{0, 0}, image.Point{w, h}
	numSites = utils.Min(MaxSites, numSites)
	sites := make([]Site, numSites)
	for i := range sites {
		sites[i] = NewSite(tl, br, c1, c2)
	}
	for j := 0; j < DefaultFrames; j++ {
		for i := range sites {
			sites[i].x += sites[i].slope.dx
			if sites[i].x < 0 {
				sites[i].x = w - 1
			} else if sites[i].x >= h {
				sites[i].x = 0
			}
			sites[i].y += sites[i].slope.dy
			if sites[i].y < 0 {
				sites[i].y = h - 1
			} else if sites[i].y >= h {
				sites[i].y = 0
			}
		}
		err := createVoronoiImg(j, w, h, "frames", sitesVisible, sites, distance)
		if err != nil {
			return "", ""
		}
	}
	format := "mp4"
	vidName := "voronoi." + format
	cmd := exec.Command("/bin/sh", "./mp4.sh", vidName)
	cmd.Stdin = bytes.NewReader([]byte("y")) // to overwrite existing file
	if _, err := cmd.Output(); err != nil {
		return "", ""
	}
	return vidName, format
}

func StainGlass(inputImg image.Image, numSites int, sitesVisible bool, dist string) *image.RGBA {
	distance, ok := distances[dist]
	if !ok {
		return nil
	}
	img := utils.ImageToRGBA(inputImg)
	tl, br := img.Bounds().Min, img.Bounds().Max
	h, w := img.Bounds().Dy(), img.Bounds().Dx()
	qt := NewQuadtree("root", tl, br, 0, img)
	sg := StainedGlass{img, []Site{}, qt}
	splits := utils.Min(MaxSites, numSites)
	for i := 0; i < splits; i++ {
		sg.sites = sg.qt.splitTree(img)
	}
	return generateVoronoiImg(w, h, sg.sites, distance, sitesVisible)
}

func StainGlass2(inputImg image.Image, numSites int, sitesVisible bool, dist string) *image.RGBA {
	distance, ok := distances[dist]
	if !ok {
		return nil
	}
	img := utils.ImageToRGBA(inputImg)
	tl, br := img.Bounds().Min, img.Bounds().Max
	h, w := img.Bounds().Dy(), img.Bounds().Dx()
	qt := NewQuadtree("root", tl, br, 0, img)
	sg := StainedGlass{img, []Site{}, qt}
	splits := utils.Min(MaxSites, numSites)
	for i := 0; i < splits; i++ {
		sg.sites = sg.qt.splitTree(img)
	}
	return generateVoronoiImg2(w, h, sg.sites, distance, sitesVisible, img)
}

func QuadImage(inputImg image.Image, numSites int, sitesVisible bool) *image.RGBA {
	img := utils.ImageToRGBA(inputImg)
	tl, br := img.Bounds().Min, img.Bounds().Max
	qt := NewQuadtree("root", tl, br, 0, img)
	splits := utils.Min(MaxSites, numSites)
	for i := 0; i < splits; i++ {
		qt.splitTree(img)
	}
	return qt.createQuadImg2(sitesVisible)
}

func DrawCircles(inputImg image.Image, numSites int, sitesVisible bool) *image.RGBA {
	img := utils.ImageToRGBA(inputImg)
	tl, br := img.Bounds().Min, img.Bounds().Max
	qt := NewQuadtree("root", tl, br, 0, img)
	splits := utils.Min(MaxSites, numSites)
	for i := 0; i < splits; i++ {
		qt.splitTree(img)
	}
	return qt.createCircleImg(sitesVisible)
}

// Takes a blank image and set of points, returns a voronoi image
func generateVoronoiImg(w, h int, sites []Site, distance DistanceFunc, sitesVisible bool) *image.RGBA {
	tl, br := image.Point{0, 0}, image.Point{w, h}
	img := image.NewRGBA(image.Rectangle{tl, br})
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			minDist, minColor := MaxInt, Black
			for _, s := range sites {
				dist := distance(s.x-x, s.y-y)
				if dist < minDist {
					minDist, minColor = dist, s.color
				}
			}
			// if sitesVisible && minDist <= 1 {
			// 	minColor = Black
			// }
			img.Set(x, y, minColor)
			if sitesVisible && onBorder(x, y, img) {
				img.Set(x, y, Black)
			}
		}
	}
	return img
}

func generateVoronoiImg2(w, h int, sites []Site, distance DistanceFunc, sitesVisible bool, inputImg *image.RGBA) *image.RGBA {
	tl, br := image.Point{0, 0}, image.Point{w, h}
	img := image.NewRGBA(image.Rectangle{tl, br})
	sitesMap := make(map[Site][][2]int, len(sites))
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			minDist, minSite := MaxInt, Site{}
			for _, s := range sites {
				dist := distance(s.x-x, s.y-y)
				if dist < minDist {
					minDist, minSite = dist, s
				}
			}
			sitesMap[minSite] = append(sitesMap[minSite], [2]int{x, y})
		}
	}
	for _, coords := range sitesMap {
		colour := avgRegionColor(coords, inputImg)
		for _, coord := range coords {
			x, y := coord[0], coord[1]
			img.Set(x, y, colour)
			if sitesVisible && onBorder(x, y, img) {
				img.Set(x, y, Black)
			}
		}
	}
	return img
}

func avgRegionColor(coords [][2]int, img *image.RGBA) color.RGBA {
	red, green, blue, alpha := 0, 0, 0, 0
	for _, coord := range coords {
		x, y := coord[0], coord[1]
		r, g, b, a := img.At(x, y).RGBA()
		red += int(r / 257)
		green += int(g / 257)
		blue += int(b / 257)
		alpha += int(a / 257)
	}
	numPixels := len(coords)
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

func onBorder(x, y int, img *image.RGBA) bool {
	h, w := img.Rect.Max.Y-img.Rect.Min.Y, img.Rect.Max.X-img.Rect.Min.X
	sy, ey := utils.Max(y-1, 0), utils.Min(y+1, h-1)
	sx, ex := utils.Max(x-1, 0), utils.Min(x+1, w-1)
	for ix := sx; ix <= ex; ix++ {
		for iy := sy; iy <= ey; iy++ {
			if img.RGBAAt(x, y) != img.RGBAAt(ix, iy) && img.RGBAAt(ix, iy) != Black {
				return true
			}
		}
	}
	return false
}

func createVoronoiImg(i, w, h int, dir string, sitesVisible bool, sites []Site, distance DistanceFunc) error {
	img := generateVoronoiImg(w, h, sites, distance, sitesVisible)
	if err := utils.OutputPngFrame(dir, i, img); err != nil {
		return err
	}
	return nil
}
