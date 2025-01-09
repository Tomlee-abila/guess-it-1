package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"guess-it/functions"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("To run this program a command similar to this should be run\ngo run ." )
		return
	}
	var data []float64
	var sum float64 = 0.0

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
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
		diff := standardDeviation * 2	
		uppr := average + diff
		lowr := average - diff
		if lowr < 0 {
			lowr = 0
		}
		fmt.Printf("%.0f %.0f\n", lowr, uppr)
	}
}
