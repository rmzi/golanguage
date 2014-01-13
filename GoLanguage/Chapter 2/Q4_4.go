package main

import (
	"fmt"
)

func main() {
	
	msg := "foobar"
	original := []rune(msg)

	for i, j := 0, len(original) - 1; i < j; i,j = i + 1, j - 1 {
    	original[i], original[j] = original[j], original[i]
    }

	res := string(original)

	fmt.Println("The new message is: ", res)
}