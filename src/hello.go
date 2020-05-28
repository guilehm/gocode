package main

import (
	"fmt"
)

func variablesInt() {
	var x int = 5
	var y int = 10
	var sum int = x + y
	fmt.Println(sum)
}

func mixTypeVariables() {
	x := 2
	y := 4
	sum := x + y
	fmt.Println(sum)
}

func main() {
	fmt.Println("Hello, World!")
	variablesInt()
	mixTypeVariables()
}
