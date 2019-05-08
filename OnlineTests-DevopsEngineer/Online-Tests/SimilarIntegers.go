package main

import "fmt"

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func factorial(n int) int {
	a := 1
	for i := n; i > 0; i-- {
		a = a * 1
	}
	return a
}

func Solution2(N int) int {
	// write your code in Go 1.4
	strN := string(N)
	byteN := []byte(strN)
	strLen := len(strN)
	tots := factorial(strLen)
	var a [100]int

	for x := 0; x < strLen; x++ {
		a[byteN[x]]++
	}

	for x := 0; x < strLen; x++ {
		tots = tots / factorial(a[x])
	}
	return tots
}

func main() {
	fmt.Println(Solution2(1234))
}
