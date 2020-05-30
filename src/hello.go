package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println("Hello, World!")
	variablesInt()
	mixTypeVariables()
	condition()
	initializeValues()
	shortInitialization()
	appendToSlices()
	vertices()
	iterations()
	otherIteration()
	iterationByRange()
	iterationByRangeInMap()
	sum(2, 3)
	sqrt(0333)
	makePerson()
	formatPrint()
	multipleDeclareVar()
	toStringConversion()
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
	// default
}

func initializeValues() {
	var a [5]int
	a[2] = 7
	fmt.Println(a)
	// [0 0 7 0 0]
}

func shortInitialization() {
	a := [5]int{5, 4, 3, 2, 1}
	fmt.Println(a)
	// [5 4 3 2 1]
}

func appendToSlices() {
	a := []int{1, 2, 3, 4}
	a = append(a, 5)
	fmt.Println(a)
	// [1 2 3 4 5]
}

func vertices() {
	vertices := make(map[string]int)
	vertices["dodecagon"] = 12
	vertices["triangle"] = 3
	vertices["square"] = 5
	delete(vertices, "square")
	fmt.Println(vertices)
	// map[dodecagon:12 triangle:3]
}

func iterations() {
	for i := 1; i < 3; i++ {
		fmt.Println(i)
	}
}

func otherIteration() {
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
}

func iterationByRange() {
	arr := []string{"a", "b", "c"}
	for index, value := range arr {
		fmt.Println("index", index, "value", value)
	}
}

func iterationByRangeInMap() {
	m := make(map[string]string)
	m["a"] = "a"
	m["b"] = "b"
	m["c"] = "c"
	for key, value := range m {
		fmt.Println("key", key, "value", value)
	}
}

func sum(x int, y int) int {
	fmt.Println(x + y)
	return x + y
}

func sqrt(x float64) float64 {
	result := math.Sqrt(x)
	fmt.Println(result)
	return result
}

type person struct {
	name string
	age  int
}

func makePerson() {
	p := person{name: "jack", age: 89}
	fmt.Println(p)
	// {jack 89}
}

func formatPrint() {
	i := 54
	fmt.Printf("%v, %T\n", i, i)
	// 54, int
}

func multipleDeclareVar() {
	var (
		name string = "Elisabeth"
		age  int    = 77
	)
	fmt.Printf("Name: %v Age: %v\n", name, age)
	// Name: Elisabeth Age: 77
}

func toStringConversion() {
	// to avoid * as a conversion from int to str
	var i int = 42
	var j string = strconv.Itoa(i)
	fmt.Printf("%v\n", j)
	// 42
}
