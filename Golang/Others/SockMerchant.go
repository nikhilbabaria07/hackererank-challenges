package Others

import (
	"fmt"
)

type sockCntr struct {
	sockClr int32
	cntr    int32
}

var sockDet sockCntr

func sockCntrVal(sockDet *sockCntr) int32 {

}

// Complete the sockMerchant function below.
func sockMerchant(n int32, ar []int32) int32 {

	var tempSockDet sockCntr
	for i = 0; i < n; i++ {
		tempSockDet.sockClr = ar[i]
		tempSockDet.cntr = sockCntrVal(&sockDet)
	}
}

func main() {
	socks := []int32{1, 2, 1, 3, 1, 1, 2, 2, 3, 4}
	fmt.Println(sockMerchant(10, socks))
}

//	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
//
//	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
//	checkError(err)
//
//	defer stdout.Close()
//
//	writer := bufio.NewWriterSize(stdout, 1024 * 1024)
//
//	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
//	checkError(err)
//	n := int32(nTemp)
//
//	arTemp := strings.Split(readLine(reader), " ")
//
//	var ar []int32
//
//	for i := 0; i < int(n); i++ {
//		arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
//		checkError(err)
//		arItem := int32(arItemTemp)
//		ar = append(ar, arItem)
//	}
//
//	result := sockMerchant(n, ar)
//
//	fmt.Fprintf(writer, "%d\n", result)
//
//	writer.Flush()
//}
//
//func readLine(reader *bufio.Reader) string {
//	str, _, err := reader.ReadLine()
//	if err == io.EOF {
//		return ""
//	}
//
//	return strings.TrimRight(string(str), "\r\n")
//}
//
//func checkError(err error) {
//	if err != nil {
//		panic(err)
//	}
//}
