package main

import "fmt"

func main() {
	var num, next, prev int

	fmt.Scan(&num)

	next, prev = num+1, num-1

	fmt.Printf("Следующее за числом %d число: %d\n", num, next)
	fmt.Printf("Для числа %d предыдущее число: %d", num, prev)
}
