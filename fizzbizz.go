//  Write a function that will accept a positive number n as input. It will
//       print numbers m from 1 to n. If m is divisible by 3, it will
//       print fizz. If it's divisible by 5, it will print bizz. If it's
//       divisible by 15, it will print fizzbizz. Otherwise, it will just
//       print the number
package main

import (
	"fmt"
)

func fizzbizz(n int) {
	for i := 1; i < n; i++ {

		if i%15 == 0 {
			fmt.Println("fizzbizz")
		} else if i%5 == 0 {
			fmt.Println("bizz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else {

			fmt.Println(i)
		}

	}
}

func main() {

	fizzbizz(18)

}
