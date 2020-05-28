package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	variablesInt()
	mixTypeVariables()
	condition()
}

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

func condition() {
	x := 5
	if x > 5 {
		fmt.Println("more than 5")
	} else if x < 2 {
		fmt.Println("less than 2")
	} else {
		fmt.Println("default")
	}
}
