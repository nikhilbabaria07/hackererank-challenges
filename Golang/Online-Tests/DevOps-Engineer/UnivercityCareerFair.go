package main

import (
	"fmt"
)

func checkCurrEvents(currHrs int32, arrival []int32) bool {

}

func maxEvents(arrival []int32, duration []int32) int32 {

	var hrsCntr, tempDep int32
	var departure []int32

	for index := range duration {
		tempDep = arrival[index] + duration[index]
		departure = append(departure, tempDep)
	}

	for index := range arrival {
		hrsCntr = arrival[index]

		if checkCurrEvents(hrsCntr, arrival, departure) {
			return 0
		}

	}

}

func main() {

	arrival := []int32{1, 3, 3, 5, 7}
	duration := []int32{1, 1, 2, 2, 1}

	fmt.Println(maxEvents(arrival, duration))
}
