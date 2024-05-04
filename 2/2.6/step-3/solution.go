package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64
	fmt.Scanln(&radius)

	pi := 3.14
	circleArea := pi * math.Pow(radius, 2)

	fmt.Println(circleArea)
}
