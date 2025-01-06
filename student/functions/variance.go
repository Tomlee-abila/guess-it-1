package functions

import "math"

func Variance(avarage float64, data []float64)float64{
	var sum float64 = 0.0

	for _, number := range data{
		sum += math.Pow((number - avarage), 2)		
	}
	variance := sum/float64(len(data))
	return variance
}