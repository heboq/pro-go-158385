package main

import "fmt"

func main() {
	var x, square, power6 int

	fmt.Scan(&x)

	square = x * x
	power6 = square * square * square

	fmt.Println(power6)
}
