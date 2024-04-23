package main

import "fmt"

func main() {
	var num uint

	fmt.Scan(&num)

	fmt.Println((num % 100) / 10)
}
