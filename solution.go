package square

import (
	"math"
)

type numOfSides uint8

const (
	SidesCircle numOfSides = iota
	_
	_
	SidesTriangle
	SidesSquare
)

func CalcSquare(sideLen float64, sidesNum numOfSides) float64 {
	var area float64 = 0
	if sidesNum == SidesTriangle {
		area = ((math.Sqrt(3) / 4) * sideLen * sideLen)
	} else if sidesNum == SidesSquare {
		area = (sideLen * sideLen)
	} else if sidesNum == SidesCircle {
		area = math.Pi * sideLen * sideLen
	} else {
		area = 0
	}
	return area
}
