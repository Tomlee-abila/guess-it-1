package main

import (
	"fmt"
	"os"

	"guess-it/functions"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("To run this program a command similar to this should be run\ngo run . <your data text file>")
		return
	}
	var data []float64
	var number float64
	var sum float64 = 0.0

	for {

		_, err := fmt.Scan(&number)

		if err != nil {
			fmt.Println("Ivalid input", err)
		} else {
			data = append(data, number)
			sum += number
		}

		average := functions.Avarage(sum, data)

		variance := functions.Variance(average, data)

		standardDeviation := functions.StandardDeviation(variance)
		diff := standardDeviation * 2
		fmt.Printf("%.0f %.0f\n", (average - diff), (average + diff))

		// fmt.Printf("Average: %0.0f\n", average)
		// fmt.Printf("Variance: %0.0f\n", variance)
		// fmt.Printf("Standard Deviation: %0.0f\n", standardDeviation)
	}
}
