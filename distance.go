package topogo

import "math"

const earthRadius = 6371008.8

// Distance determines the great-circle distance between 2 coordinates
// using the [Haversine formula](https://en.wikipedia.org/wiki/Haversine_formula)
// and approximates using a spherical (non-ellipsoid) Earth with a
// mean radius of 6,371,008.8 meters derived from the WGS84 datum.
func Distance(pt1, pt2 Coordinate) float64 {
	pt1.ToRadians()
	pt2.ToRadians()

	a := math.Sin((pt2.Lat - pt1.Lat) / 2)

	b := math.Sin((pt2.Lon - pt1.Lon) / 2)

	c := a*a + b*b*math.Cos(pt1.Lat)*math.Cos(pt2.Lat)

	return 2 * earthRadius * math.Atan2(math.Sqrt(c), math.Sqrt(1-c))
}
