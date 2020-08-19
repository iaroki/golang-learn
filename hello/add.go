package main

import "fmt"

var a int = 5
var b int = 7
var somestr string = "5+7="

func add(a int, b int) int {
	c := a + b
	return c
}

func main() {

	fmt.Println(somestr, add(a, b))
}
