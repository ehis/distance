package distance_test

import (
	. "github.com/ehis/distance"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Distance", func() {
	It("should return the great-circle distance given 2 coordinates", func() {
		Expect(GreatCircle(
			[]float64{-96.796667, 32.775833},
			[]float64{126.967583, 37.566776},
		)).To(BeNumerically("~", 10974883, 1))

		Expect(GreatCircle(
			[]float64{-180, 40.71427},
			[]float64{180, 40.71427},
		)).To(BeNumerically("~", 0, 0.1))

		Expect(GreatCircle(
			[]float64{-180, 0},
			[]float64{180, 0},
		)).To(BeNumerically("~", 0, 0.1))

		Expect(GreatCircle(
			[]float64{-142, 90},
			[]float64{26, -90},
		)).To(BeNumerically("~", 6371008.8*3.14159265359, 1))

		Expect(GreatCircle(
			[]float64{-90, 0},
			[]float64{90, 0},
		)).To(BeNumerically("~", 6371008.8*3.14159265359, 1))

		Expect(GreatCircle(
			[]float64{-180, 0},
			[]float64{0, 0},
		)).To(BeNumerically("~", 6371008.8*3.14159265359, 1))

		Expect(GreatCircle(
			[]float64{-96.796667, 32.775833},
			[]float64{126.967583, 37.566776},
		)).To(BeNumerically("~", 10974883, 1))

		Expect(GreatCircle(
			[]float64{126.967583, 37.566776},
			[]float64{151.215158, -33.857406},
		)).To(BeNumerically("~", 8328610, 1))

		Expect(GreatCircle(
			[]float64{151.215158, -33.857406},
			[]float64{55.274180, 25.197229},
		)).To(BeNumerically("~", 12048941, 1))

		Expect(GreatCircle(
			[]float64{55.274180, 25.197229},
			[]float64{6.942661, 50.334057},
		)).To(BeNumerically("~", 4962217, 1))

		Expect(GreatCircle(
			[]float64{6.942661, 50.334057},
			[]float64{-97.635926, 30.134442},
		)).To(BeNumerically("~", 8414177, 1))

		Expect(GreatCircle(
			[]float64{-90, 30},
			[]float64{-87, 32},
		)).To(BeNumerically("~", 362210, 1))

		Expect(GreatCircle(
			[]float64{-75.343, 39.984},
			[]float64{-75.534, 39.123},
		)).To(BeNumerically("~", 97129, 1))

		Expect(GreatCircle(
			[]float64{-130, 42},
			[]float64{130, -42},
		)).To(BeNumerically("~", 13669374, 1))

		Expect(GreatCircle(
			[]float64{-74.00597, 40.71427},
			[]float64{-70.56656, -33.42628},
		)).To(BeNumerically("~", 8251609, 1))
	})

	It("should return the point segment distance given a point and a line segement", func() {
		Expect(PointSegment(
			[]float64{0.5, 0.5},
			[]float64{0, 0},
			[]float64{1, 0},
		)).To(BeNumerically("~", 0.5, 0.001))

		Expect(PointSegment(
			[]float64{0.5, 0.5},
			[]float64{0, 0},
			[]float64{-1, 0},
		)).To(BeNumerically("~", 0.707, 0.001))

		Expect(PointSegment(
			[]float64{0.5, 0.5},
			[]float64{-1, 0},
			[]float64{0, 0},
		)).To(BeNumerically("~", 0.707, 0.001))

		Expect(PointSegment(
			[]float64{0.5, 0.5},
			[]float64{0, 0},
			[]float64{0, 0},
		)).To(BeNumerically("~", 0.707, 0.001))

		Expect(PointSegment(
			[]float64{3, 2},
			[]float64{-2, 1},
			[]float64{5, 3},
		)).To(BeNumerically("~", 0.412, 0.001))

		Expect(PointSegment(
			[]float64{7, 1},
			[]float64{-2, 1},
			[]float64{5, 1},
		)).To(BeNumerically("~", 2, 0.001))

		Expect(PointSegment(
			[]float64{17, 1},
			[]float64{-2, 1},
			[]float64{5, 1},
		)).To(BeNumerically("~", 12, 0.001))

		Expect(PointSegment(
			[]float64{3, 1},
			[]float64{-2, 1},
			[]float64{5, 1},
		)).To(BeNumerically("~", 0, 0.001))

		Expect(PointSegment(
			[]float64{-2, 1},
			[]float64{-2, 1},
			[]float64{5, 1},
		)).To(BeNumerically("~", 0, 0.001))

		Expect(PointSegment(
			[]float64{5, 1},
			[]float64{-2, 1},
			[]float64{5, 1},
		)).To(BeNumerically("~", 0, 0.001))

		Expect(PointSegment(
			[]float64{3, 3},
			[]float64{-2, 1},
			[]float64{5, 1},
		)).To(BeNumerically("~", 2, 0.001))

		Expect(PointSegment(
			[]float64{3, -1},
			[]float64{-2, 1},
			[]float64{5, 1},
		)).To(BeNumerically("~", 2, 0.001))

		Expect(PointSegment(
			[]float64{-2, -1},
			[]float64{-2, 1},
			[]float64{5, 1},
		)).To(BeNumerically("~", 2, 0.001))

		Expect(PointSegment(
			[]float64{7, 1},
			[]float64{-2, 1},
			[]float64{-2, 5},
		)).To(BeNumerically("~", 9, 0.001))

		Expect(PointSegment(
			[]float64{17, 1},
			[]float64{-2, 1},
			[]float64{-2, 5},
		)).To(BeNumerically("~", 19, 0.001))

		Expect(PointSegment(
			[]float64{-2, 3},
			[]float64{-2, 1},
			[]float64{-2, 5},
		)).To(BeNumerically("~", 0, 0.001))

		Expect(PointSegment(
			[]float64{-2, 1},
			[]float64{-2, 1},
			[]float64{-2, 5},
		)).To(BeNumerically("~", 0, 0.001))

		Expect(PointSegment(
			[]float64{-2, 5},
			[]float64{-2, 1},
			[]float64{-2, 5},
		)).To(BeNumerically("~", 0, 0.001))

		Expect(PointSegment(
			[]float64{3, 3},
			[]float64{-2, 1},
			[]float64{-2, 5},
		)).To(BeNumerically("~", 5, 0.001))

		Expect(PointSegment(
			[]float64{2, 1},
			[]float64{-2, 1},
			[]float64{-2, 5},
		)).To(BeNumerically("~", 4, 0.001))

		Expect(PointSegment(
			[]float64{-2, -1},
			[]float64{-2, 1},
			[]float64{-2, 5},
		)).To(BeNumerically("~", 2, 0.001))

		Expect(PointSegment(
			[]float64{7, 1},
			[]float64{-2, 1},
			[]float64{-2, 1},
		)).To(BeNumerically("~", 9, 0.001))

		Expect(PointSegment(
			[]float64{17, 1},
			[]float64{-2, 1},
			[]float64{-2, 1},
		)).To(BeNumerically("~", 19, 0.001))

		Expect(PointSegment(
			[]float64{-2, 3},
			[]float64{-2, 1},
			[]float64{-2, 1},
		)).To(BeNumerically("~", 2, 0.001))
	})
})
