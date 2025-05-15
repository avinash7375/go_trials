//this is some trials of arrays in golang

package main

import "fmt"

func main() {
	nums := []float64{1,2,3,4,5,6,7,8}

	for i, j := range nums {
		fmt.Printf("%d : %f\n",i, j)
	}
}
