package distance

import "math"

const earthRadius = 6371008.8
const piOver180 = math.Pi / 180

// GreatCircle returns the great-circle distance between 2 coordinates
// using the [Haversine formula](https://en.wikipedia.org/wiki/Haversine_formula)
// and approximates using a spherical (non-ellipsoid) Earth with a
// mean radius of 6,371,008.8 meters derived from the WGS84 datum.
func GreatCircle(pt1, pt2 []float64) float64 {

	a := math.Sin(((pt2[1] - pt1[1]) * piOver180) / 2)

	b := math.Sin(((pt2[0] - pt1[0]) * piOver180) / 2)

	c := a*a + b*b*math.Cos(pt1[1]*piOver180)*math.Cos(pt2[1]*piOver180)

	return 2 * earthRadius * math.Atan2(math.Sqrt(c), math.Sqrt(1-c))
}

// PointSegment returns the distance between a point the closest point on a line segment
func PointSegment(pt, pt1, pt2 []float64) float64 {
	x := pt1[0]
	y := pt1[1]
	dx := pt2[0] - x
	dy := pt2[1] - y

	if dx != 0 || dy != 0 {
		t := ((pt[0]-x)*dx + (pt[1]-y)*dy) / (dx*dx + dy*dy)
		if t > 1 {
			x = pt2[0]
			y = pt2[1]
		} else if t > 0 {
			x += dx * t
			y += dy * t
		}
	}

	dx = pt[0] - x
	dy = pt[1] - y

	return math.Sqrt(dx*dx + dy*dy)
}
