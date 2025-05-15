package main

import "fmt"

func main() {
	var days string = "Wednesday"
	switch days {
	case "Monday":
		fmt.Println("Start of week")
	case "Sunday":
		fmt.Println("end of weeks")
	default:
		fmt.Println("mid week")
	}
}

