package main

import "fmt"

//defer waits until the surrounding function returns (func main)
//testing
func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
