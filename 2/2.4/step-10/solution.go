package main

import "fmt"

func main() {
	var pirozhokPrice, purchase, rub, kop, n int

	fmt.Scan(&rub, &kop, &n)

	pirozhokPrice = kop + (rub * 100)

	purchase = pirozhokPrice * n

	fmt.Println(purchase/100, purchase%100)
}
