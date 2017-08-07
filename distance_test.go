package topogo_test

import (
	. "github.com/topogo"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Distance", func() {
	It("should return the great-circle distance given 2 coordinates", func() {
		Expect(Distance(
			Coordinate{Lon: -96.796667, Lat: 32.775833},
			Coordinate{Lon: 126.967583, Lat: 37.566776},
		)).To(BeNumerically("~", 10974883, 1))

		Expect(Distance(
			Coordinate{Lon: -180, Lat: 40.71427},
			Coordinate{Lon: 180, Lat: 40.71427},
		)).To(BeNumerically("~", 0, 0.1))

		Expect(Distance(
			Coordinate{Lon: -180, Lat: 0},
			Coordinate{Lon: 180, Lat: 0},
		)).To(BeNumerically("~", 0, 0.1))

		Expect(Distance(
			Coordinate{Lon: -142, Lat: 90},
			Coordinate{Lon: 26, Lat: -90},
		)).To(BeNumerically("~", 6371008.8*3.14159265359, 1))

		Expect(Distance(
			Coordinate{Lon: -90, Lat: 0},
			Coordinate{Lon: 90, Lat: 0},
		)).To(BeNumerically("~", 6371008.8*3.14159265359, 1))

		Expect(Distance(
			Coordinate{Lon: -180, Lat: 0},
			Coordinate{Lon: 0, Lat: 0},
		)).To(BeNumerically("~", 6371008.8*3.14159265359, 1))

		Expect(Distance(
			Coordinate{Lon: -96.796667, Lat: 32.775833},
			Coordinate{Lon: 126.967583, Lat: 37.566776},
		)).To(BeNumerically("~", 10974883, 1))

		Expect(Distance(
			Coordinate{Lon: 126.967583, Lat: 37.566776},
			Coordinate{Lon: 151.215158, Lat: -33.857406},
		)).To(BeNumerically("~", 8328610, 1))

		Expect(Distance(
			Coordinate{Lon: 151.215158, Lat: -33.857406},
			Coordinate{Lon: 55.274180, Lat: 25.197229},
		)).To(BeNumerically("~", 12048941, 1))

		Expect(Distance(
			Coordinate{Lon: 55.274180, Lat: 25.197229},
			Coordinate{Lon: 6.942661, Lat: 50.334057},
		)).To(BeNumerically("~", 4962217, 1))

		Expect(Distance(
			Coordinate{Lon: 6.942661, Lat: 50.334057},
			Coordinate{Lon: -97.635926, Lat: 30.134442},
		)).To(BeNumerically("~", 8414177, 1))

		Expect(Distance(
			Coordinate{Lon: -90, Lat: 30},
			Coordinate{Lon: -87, Lat: 32},
		)).To(BeNumerically("~", 362210, 1))

		Expect(Distance(
			Coordinate{Lon: -75.343, Lat: 39.984},
			Coordinate{Lon: -75.534, Lat: 39.123},
		)).To(BeNumerically("~", 97129, 1))

		Expect(Distance(
			Coordinate{Lon: -130, Lat: 42},
			Coordinate{Lon: 130, Lat: -42},
		)).To(BeNumerically("~", 13669374, 1))

		Expect(Distance(
			Coordinate{Lon: -74.00597, Lat: 40.71427},
			Coordinate{Lon: -70.56656, Lat: -33.42628},
		)).To(BeNumerically("~", 8251609, 1))
	})
})
