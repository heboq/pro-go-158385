package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var str1, str2, str3 string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str1 = scanner.Text()

	scanner.Scan()
	str2 = scanner.Text()

	scanner.Scan()
	str3 = scanner.Text()

	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)
}
