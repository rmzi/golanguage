package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	
	msg := "asSASA ddd dsjkdsjs dk"

	fmt.Println("The number of characters in the message is: ", utf8.RuneCountInString(msg))

}