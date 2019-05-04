package Implmentation

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the countingValleys function below.
func countingValleys(n int32, s string) int32 {

	var seaLevel int32
	var flagU, flagD [2]int32

	for _, value := range s {
		if string(value) == "U" {
			if seaLevel == 0 {
				flagU[0]++
			}
			seaLevel++
			if (seaLevel == 0) && (flagD[1]+1 == flagD[0]) {
				flagD[1]++
			}
		} else {
			if seaLevel == 0 {
				flagD[0]++
			}
			seaLevel--
			if (seaLevel == 0) && (flagU[1]+1 == flagU[0]) {
				flagU[1]++
			}
		}
	}
	return flagD[1]

}

func main() {

	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	s := readLine(reader)

	result := countingValleys(n, s)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
