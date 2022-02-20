package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type Sides int

const square Sides = 4
const triangle Sides = 3
const circle Sides = 0

func CalcSquare(sideLen float64, sidesNum Sides) float64 {
	switch sidesNum {
	case circle:
		return math.Pow(sideLen, 2) * math.Pi
	case triangle:
		return math.Pow(sideLen, 2) * math.Sqrt(3) / 4
	case square:
		return math.Pow(sideLen, 2)
	default:
		return 0
	}
}
