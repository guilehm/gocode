package main

import "fmt"

func Repeat(character string) (repeated string) {
	return character + character + character + character
}

func main() {
	fmt.Println(Repeat("a"))
}
