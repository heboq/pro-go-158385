package main

import "fmt"

func main() {
	var base, square int // base - основание степени, square - степень

	fmt.Scan(&base)

	square = base * base

	fmt.Println(square)
}
