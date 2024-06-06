package functions

func Avarage(sum float64, data []float64)float64{
	if len(data) == 0{
		return 0.0
	}
	return sum/float64(len(data))
}