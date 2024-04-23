package main

import "fmt"

func main() {
	var num1, num2, num3, product int

	fmt.Scan(&num1, &num2, &num3)

	product = num1 * num2 * num3

	fmt.Println(product)
}
