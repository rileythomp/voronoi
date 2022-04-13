package voronoi

import (
	"gitlab.com/rileythomp14/voronoi/src/utils"
)

type (
	DistanceFunc func(int, int) int
)

var (
	distances = map[string]DistanceFunc{
		"euclidean": EuclideanDistance,
		"manhattan": ManhattanDistance,
		"chebyshev": ChebyshevDistance,
	}
)

func EuclideanDistance(dx, dy int) int {
	return dx*dx + dy*dy
}

func ManhattanDistance(dx, dy int) int {
	return utils.Abs(dx) + utils.Abs(dy)
}

func ChebyshevDistance(dx, dy int) int {
	return utils.Max(utils.Abs(dx), utils.Abs(dy))
}
