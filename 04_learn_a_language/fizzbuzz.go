package main

import "fmt"

// – 3で割り切れるときは「Fizz」を発言
// – 5で割り切れるときは「Buzz」を発言
// – 両方で割り切れるときは「FizzBuzz」を発言

func sub() {
	for i := 1; i <= 30; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}
