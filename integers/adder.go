package main

import "fmt"

func Add(x, y int) (sum int) {
	return x + y
}

func main() {
	fmt.Println(Add(2, 2))
}
