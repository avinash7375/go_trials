package main

import "fmt"

func main() {
for i := 0; i<5; i++ {
	fmt.Println(i)
}


fmt.Println("-------new----------")

nums := []int{1,2,3,4,5,6}
for i,v := range nums {
	fmt.Printf("index: %d, value: %d\n", i, v*3)
}
}

