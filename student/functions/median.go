package functions

import "sort"

func Median(intData []float64) float64 {
	sort.Float64s(intData) // Sort the slice

	var median float64
	n := len(intData)

	if n%2 == 0 {
		median = (intData[n/2-1] + intData[n/2]) / 2
	} else {
		median = (intData[n/2])
	}
	return median
}