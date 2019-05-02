package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type charCnt struct {
	char byte
	cnt  int32
}

var charDB []charCnt

func charCntr(c byte) {

	flag := false
	flagInd := -1

	for index, value := range charDB {
		if value.char == c {
			flag = true
			flagInd = index
		}
	}

	if flag == true {
		charDB[flagInd].cnt++
	} else {
		var addChar charCnt
		addChar.char = c
		addChar.cnt = 1

		charDB = append(charDB, addChar)
	}

}

func checkValidity() string {

	flagValid := 1
	charArrLen := int32(len(charDB))

	for i := int32(1); i < charArrLen; i++ {
		if charDB[i].cnt != charDB[i-1].cnt {
			flagValid--
		}

	}

	if flagValid == 0 || flagValid == 1 {
		return "YES"
	} else {
		return "NO"
	}

}

// Complete the isValid function below.
func isValid(s string) string {

	charArr := []byte(s)

	for _, value := range charArr {
		charCntr(value)
	}

	return checkValidity()

}

func main() {

	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	result := isValid(s)

	fmt.Fprintf(writer, "%s\n", result)

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
