package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"math-skills/functions"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("To run this program a command similar to this should be run\ngo run . <your data text file>")
		return
	}

	dataFile := os.Args[1]

	dataFileContent, err := os.ReadFile(dataFile)
	if err != nil {
		fmt.Println("error reading", dataFile, err)
	}

	dataString := strings.Split(string(dataFileContent), "\n")
	var data []float64
	var sum float64 = 0.0
	for _, str := range dataString {
		str = strings.TrimSpace(str)
		if !(str == "") {
			num, err := strconv.ParseFloat(str, 64)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return
			}
			data = append(data, num)
			sum += num
		}
	}
	if len(data) == 0{
		fmt.Println("The file", dataFile,"is empty")
		return
	}

	avarage := functions.Avarage(sum, data)

	median := functions.Median(data)

	variance := functions.Variance(avarage, data)

	fmt.Println("Average:", avarage)
	fmt.Println("Variance:", variance)
	fmt.Println("Median:", median)
}
