# distance

[![CircleCI](https://circleci.com/gh/ehis/distance/tree/master.svg?style=svg)](https://circleci.com/gh/ehis/distance/tree/master)

Using the [Haversine formula](https://en.wikipedia.org/wiki/Haversine_formula) to calculate the
great-circle distance between 2 points.

## Installation

```sh
go get -u github.com/ehis/distance
```

## Usage

 - `distance.GreatCircle()` - returns the great-circle distance between 2 points.

 ```go
// []float64{longitude, latitude}
pt1 := []float64{-180, 40.71427}
pt2 := []float64{180, 40.71427}

d := distance.GreatCircle(pt1, pt2) 

 ```

 - `distance.PointSegment()` - return the distance between a point and the closest point on a line segment.


 ```go
// []float64{x, y}
pt := []float64{3, 2}
pt1 := []float64{-2, 1}
pt2 := []float64{5, 3}
d := distance.PointSegment(pt, pt1, pt2)

 ```

**Inspiration from [_distance_ - an elixir library](https://github.com/pkinney/distance)**
