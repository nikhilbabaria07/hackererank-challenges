package Others

import (
	"fmt"
	"math"
)

func main() {

	var x float64
	fmt.Scanln(&x)

	for i := 1; i <= int(x); i++ {
		if math.Mod(float64(i), 3) == 0 {
			if math.Mod(float64(i), 5) == 0 {
				fmt.Println("FizzBuzz")
				break
			}
			fmt.Println("Fizz")

		} else if math.Mod(float64(i), 5) == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

}
