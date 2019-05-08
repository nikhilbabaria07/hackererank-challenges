package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type fileContents struct {
	source         string
	nilVal1        string
	nilVal2        string
	timestampVal   string
	timeZone       string
	requestType    string
	imgaeLoc       string
	imgaeName      string
	responseType   string
	responseStatus int
	fileSize       int
}

var readFile []fileContents

func splitFile(fileContent string) {

	var strSplit []string
	var tempFileDet fileContents
	scanner := bufio.NewScanner(strings.NewReader(fileContent))

	for scanner.Scan() {
		strSplit = append(strSplit, scanner.Text())
	}

	for _, value := range strSplit {
		tempSplit := strings.Split(value, " ")
		tempFileDet.source = tempSplit[0]
		tempFileDet.nilVal1 = tempSplit[1]
		tempFileDet.nilVal2 = tempSplit[2]
		tempFileDet.timestampVal = tempSplit[3]
		tempFileDet.timeZone = tempSplit[4]
		tempFileDet.requestType = strings.Replace(tempSplit[5], "\"", "", 1)
		tempFileDet.imgaeLoc = tempSplit[6]
		tempFileDet.imgaeName = filepath.Base(tempSplit[6])
		tempFileDet.responseType = tempSplit[7]
		tempFileDet.responseStatus, _ = strconv.Atoi(tempSplit[8])
		tempFileDet.fileSize, _ = strconv.Atoi(tempSplit[9])
		readFile = append(readFile, tempFileDet)
	}
}

func removeDups(elements []string) []string {
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {
		} else {
			encountered[elements[v]] = true
			result = append(result, elements[v])
		}
	}
	return result
}

func checkValidGifs() []string {

	var validGifs []string

	for _, value := range readFile {
		if (strings.Contains(value.imgaeName, ".gif")) || (strings.Contains(value.imgaeName, ".GIF")) {

			if value.requestType == "GET" && value.responseStatus == 200 {
				validGifs = append(validGifs, value.imgaeName)
			}

		}
	}

	validGifs = removeDups(validGifs)

	return validGifs

}

func main() {
	// read the string filename
	var filename string
	fmt.Scanf("%s", &filename)

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}

	splitFile(string(b))

	writeFile := checkValidGifs()

	outFile := "gifs_" + filename

	fo, err := os.Create(outFile)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	for _, line := range writeFile {
		if _, err = fo.WriteString(line + "\n"); err != nil {
			panic(err)
		}
		fmt.Println(line)
	}

}
