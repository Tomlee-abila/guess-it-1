package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"guess-it/functions"
	// "strconv"
	// "strings"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("To run this program a command similar to this should be run\ngo run . <your data text file>")
		return
	}
	var data []float64
	// var number float64
	var sum float64 = 0.0

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		// _, err := fmt.Scan(&number)
		value := scanner.Text()
		number, err := strconv.ParseFloat(value, 64)

		if err != nil {
			fmt.Println("Ivalid input", err)
		} else {
			data = append(data, number)
			sum += number
		}

		average := functions.Avarage(sum, data)

		variance := functions.Variance(average, data)

		standardDeviation := functions.StandardDeviation(variance)
		diff := standardDeviation * 3
		fmt.Printf("%.0f %.0f\n", (average - diff), (average + diff))

		// fmt.Printf("Average: %0.0f\n", average)
		// fmt.Printf("Variance: %0.0f\n", variance)
		// fmt.Printf("Standard Deviation: %0.0f\n", standardDeviation)
	}
}
