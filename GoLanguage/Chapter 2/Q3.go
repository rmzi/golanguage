package main

import "fmt"

func main() {
	
	for i := 0; i < 100; i++{
		if i == 0 {
			fmt.Println(i)
		} else if i % 5 == 0 && i % 3 == 0 {
			fmt.Println("Fizzbuzz")
		} else if i % 5 == 0 {
			fmt.Println("Buzz")
		} else if i % 3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}	
}