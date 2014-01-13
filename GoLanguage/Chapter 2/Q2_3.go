package main

import "fmt"

func main() {

  	i := 0
	var a [10]int
  	
	Lūp:    
  		a[i] = i
		i++
		if i < 10{
  			goto Lūp
		} else {
			goto PrintOut
		}
		
	PrintOut:
		for _,v := range a {
			fmt.Println(v)
		}
}