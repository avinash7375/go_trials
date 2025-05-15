package main

import "fmt"

func main() {
	var i int = 42
	f := 3.14
	s := "Hello"
	b := true
	fmt.Printf("int: %d, float: %f, string: %s, boolean: %t\n", i,f,s,b)

	const pi = 12.121212

	fmt.Printf("Const : %f", pi)
}
