package main

import "fmt"

var x string = "string var"
var y int = 99

const immut string = "Const"

func printString(str string) {
	fmt.Println(str)
}

func printInt(num int) {
	fmt.Println(num)
}

func doubleIt() {
	fmt.Print("Enter number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}

func main() {
	printString(x)
	printInt(y)
	printString(immut)
	doubleIt()
}
