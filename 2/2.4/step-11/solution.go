package main

import "fmt"

func main() {
	var min int

	fmt.Scan(&min)

	fmt.Printf("%d мин - это %d час %d минут.", min, min/60, min%60)
}
