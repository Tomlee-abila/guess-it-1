package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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
	var data []int
	sum := 0
	for _, str := range dataString {
		str = strings.TrimSpace(str)
		if !(str == "") {
			num, err := strconv.Atoi(str)
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

	avarage := sum/len(data)

	median := Median(data)

	fmt.Println("Average:", avarage)
	fmt.Println("Median:", median)
}

func Median(intData []int) int {
	sort.Ints(intData) // Sort the slice

	var median int
	n := len(intData)

	if n%2 == 0 {
		median = (intData[n/2-1] + intData[n/2]) / 2
	} else {
		median = (intData[n/2])
	}
	return median
}
