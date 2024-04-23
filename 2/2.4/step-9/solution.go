package main

import "fmt"

func main() {
	var (
		phonePrice     int
		casePrice      int
		chargePrice    int
		earphonesPrice int
		phoneSetPrice  int
		purchase       int
	)

	fmt.Scan(&phonePrice, &casePrice, &chargePrice, &earphonesPrice)

	phoneSetPrice = phonePrice + casePrice + chargePrice + earphonesPrice

	purchase = phoneSetPrice * 3

	fmt.Println(purchase)
}
