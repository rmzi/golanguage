package main

import "fmt"

func main() {

  	i := 0
  	
	Lūp:    
  		fmt.Println(i)
  		i++
		if i < 10{
  			goto Lūp
		}
}