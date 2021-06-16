package main

import "fmt"

func Repeat(character string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}

func main() {
	fmt.Println(Repeat("a", 4))
}
