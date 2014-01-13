package main

import "fmt"

func average(slice []float64){
	sum := 0.0
	res := 0.0

	for i := 0; i < cap(slice); i++{
		sum += slice[i]
	}

	res = sum / float64(cap(slice))
	fmt.Println(res)
}

func main() {
	test := []float64{1.0, 2.5, 3.0}
	average(test)
}