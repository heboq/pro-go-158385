package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64
	pi := 3.14

	fmt.Scanln(&radius)

	area := pi * math.Pow(radius, 2)

	fmt.Println(area)
}
