package main

import (
	"math"
)

const R_EARTH = 6371e3 // in meters

func haversin(x float64) float64 {
	s := math.Sin(x / 2)
	return s * s
}

func archaversin(x float64) float64 {
	return 2 * math.Asin(math.Sqrt(x))
}

func GCDist(lat1, lon1, lat2, lon2 float64) float64 {
	h := haversin(lat2 - lat1)
	h += math.Cos(lat1) * math.Cos(lat2) * haversin(lon2-lon1)
	return R_EARTH * archaversin(h)
}
