package voronoi

import (
	"image"
	"image/color"
	"math/rand"

	"gitlab.com/rileythomp14/voronoi/src/utils"
)

const (
	MaxSites = 500
)

type (
	Site struct {
		x     int
		y     int
		color color.RGBA
		slope Slope
	}

	Slope struct {
		dx int
		dy int
	}
)

func NewSite(tl, br image.Point, c1, c2 string) Site {
	r1, g1, b1 := uint8(rand.Intn(256)), uint8(rand.Intn(256)), uint8(rand.Intn(256))
	r2, g2, b2 := r1, g1, b1
	if c1 == "red" {
		r1, g1, b1 = 255, 34, 34
		r2, g2, b2 = 255, 187, 187
	} else if c1 == "green" {
		r1, g1, b1 = 0, 102, 34
		r2, g2, b2 = 153, 255, 153
	} else if c1 == "blue" {
		r1, g1, b1 = 34, 34, 255
		r2, g2, b2 = 187, 187, 255
	} else if utils.IsHexColor(c1) && utils.IsHexColor(c2) {
		r1, g1, b1 = getHexVal(c1[0:2]), getHexVal(c1[2:4]), getHexVal(c1[4:6])
		r2, g2, b2 = getHexVal(c2[0:2]), getHexVal(c2[2:4]), getHexVal(c2[4:6])
	}
	colour := getGradientColor(
		color.RGBA{R: r1, G: g1, B: b1},
		color.RGBA{R: r2, G: g2, B: b2},
	)
	x := rand.Intn(utils.Max(br.X-tl.X, 1)) + tl.X
	y := rand.Intn(utils.Max(br.Y-tl.Y, 1)) + tl.Y
	return Site{x, y, colour, NewSlope(5)}
}

func NewSlope(n int) Slope {
	dx, dy := rand.Intn(2*n+1)-n, rand.Intn(2*n+1)-n
	return Slope{dx: dx, dy: dy}
}

func getHexVal(hex string) uint8 {
	val := 0
	for _, h := range hex {
		if '0' <= h && h <= '9' {
			val = 16*val + int(h-'0')
		} else if 'a' <= h && h <= 'f' {
			val = 16*val + int(h-'a') + 10
		} else if 'A' <= h && h <= 'F' {
			val = 16*val + int(h-'A') + 10
		}
	}
	return uint8(val)
}

func getGradientColor(c1, c2 color.RGBA) color.RGBA {
	dr := float64(c2.R) - float64(c1.R)
	dg := float64(c2.G) - float64(c1.G)
	db := float64(c2.B) - float64(c1.B)
	percent := rand.Float64()
	return color.RGBA{
		R: uint8(float64(c1.R) + percent*dr),
		G: uint8(float64(c1.G) + percent*dg),
		B: uint8(float64(c1.B) + percent*db),
		A: 0xff,
	}
}
