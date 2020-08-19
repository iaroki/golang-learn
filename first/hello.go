package main

import "fmt"

const englishHelloPrefix = "Hello, "

func HelloFunc(name string) string {
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(HelloFunc("Max"))
}
