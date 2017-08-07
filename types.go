package topogo

import "math"

const piOver180 = math.Pi / 180

// Coordinate holds longitude and latitude pair in degrees
type Coordinate struct {
	Lon float64
	Lat float64
}

// ToRadians converts values from degrees to radians
func (c *Coordinate) ToRadians() {
	c.Lon = c.Lon * piOver180
	c.Lat = c.Lat * piOver180
}
