package main

import "fmt"

func main() {
	newFunction()
}

func newFunction() {
	var s string
	fmt.Scan(&s)
	fmt.Println(s)
}
