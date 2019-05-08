package main

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

type callRedord struct {
	callDur  string
	phoneNum int
	durSec   int
}

type totalDurPhone struct {
	phoneNum  int
	totDurSec int
}

var callDetails []callRedord
var highestCaller []totalDurPhone

func stringMan(S string) {

	var strSplit []string
	var tempDet callRedord
	scanner := bufio.NewScanner(strings.NewReader(S))

	for scanner.Scan() {
		strSplit = append(strSplit, scanner.Text())
	}

	for _, value := range strSplit {
		tempSplit := strings.Split(value, ",")
		tempDet.callDur = tempSplit[0]
		tempDet.phoneNum, _ = strconv.Atoi(strings.Replace(tempSplit[1], "-", "", -1))
		tempDet.durSec = callToSec(tempSplit[0])
		callDetails = append(callDetails, tempDet)
	}
}

func callToSec(cD string) int {

	var cdSecs, HH, MM, SS int
	HMS := strings.Split(cD, ":")

	HH, _ = strconv.Atoi(HMS[0])

	MM, _ = strconv.Atoi(HMS[1])

	SS, _ = strconv.Atoi(HMS[2])

	cdSecs = (HH * 60 * 60) + (MM * 60) + (SS)

	return cdSecs

}

func findTotDur() {
	var tempCallRed totalDurPhone

	flag := false

	for _, value := range callDetails {
		tempCallRed.phoneNum = value.phoneNum
		tempCallRed.totDurSec = value.durSec

		for _, value := range highestCaller {
			if value.phoneNum == tempCallRed.phoneNum {
				flag = true
			}
		}

		if flag == true {
			for index, value := range highestCaller {
				if value.phoneNum == tempCallRed.phoneNum {
					highestCaller[index].totDurSec += tempCallRed.totDurSec
				}
			}
		} else {
			highestCaller = append(highestCaller, tempCallRed)
		}
	}
}

func findMaxDurCaller() {

	var maxDur, maxCaller int

	for _, value := range highestCaller {
		if value.totDurSec > maxDur {
			maxDur = value.totDurSec
			maxCaller = value.phoneNum
		} else if value.totDurSec == maxDur {
			if maxCaller > value.phoneNum {
				maxCaller = value.phoneNum
			}
		}
	}

	for index, value := range callDetails {
		if value.phoneNum == maxCaller {
			callDetails[index].durSec = 0
		}
	}
}

func calBillAmount() int {

	var billAmount int

	for _, value := range callDetails {
		if value.durSec <= 300 {
			billAmount += value.durSec * 3
		} else {
			billAmount += int(math.Ceil(float64(value.durSec)/float64(60))) * 150
		}
	}

	return billAmount

}

func moneyCalc() int {

	findTotDur()

	findMaxDurCaller()

	return calBillAmount()

}

func Solution(S string) int {
	// write your code in Go 1.4
	stringMan(S)

	return moneyCalc()
}

func main() {
	x := `00:01:07,400-234-090
00:05:01,701-080-080
00:05:00,400-234-090`

	fmt.Println(Solution(x))
}
