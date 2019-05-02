package Others

import "fmt"

func main() {
	var inStr string
	fmt.Scanln(&inStr)

	var b1, b2, b3 int

	for _, i := range inStr {
		if i == '(' {
			b1++
		}
		if i == '[' {
			b2++
		}
		if i == '{' {
			b3++
		}
		if i == ')' {
			b1--
		}
		if i == ']' {
			b2--
		}
		if i == '}' {
			b3--
		}
	}

	if b1+b2+b3 == 0 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
