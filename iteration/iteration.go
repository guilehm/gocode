package main

import "fmt"

const repeatCount = 4

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}

func main() {
	fmt.Println(Repeat("a"))
}
