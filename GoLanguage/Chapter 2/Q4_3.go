package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	
	msg := "asSASA ddd dsjkdsjs dk"

	c := []rune(msg)
	c[4] = 'a'
	c[5] = 'b'
	c[6] = 'c'

	res := string(c)

	fmt.Println("The new message is: ", res)

	fmt.Println("The number of characters in the message is: ", utf8.RuneCountInString(msg))

}