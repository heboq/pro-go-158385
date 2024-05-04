package main

import "fmt"

func main() {
	var a, b float64
	fmt.Scanln(&a)
	fmt.Scanln(&b)

	triangleArea := 0.5 * a * b
	fmt.Println(triangleArea)
}
