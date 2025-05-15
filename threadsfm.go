package main

import "fmt"

func divide(a,b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Division by zero gives nothing but 0")
	}
	return a/b, nil
}

func main() {
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Result : ", result)
}

