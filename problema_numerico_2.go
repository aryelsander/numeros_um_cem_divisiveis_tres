package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ {
		var temp string
		if i%3 != 0 && i%5 != 0 {
			fmt.Print(" ", i)
			continue
		}
		if i%3 == 0 {
			temp = "Pin"
		}
		if i%5 == 0 {
			temp += "Pan"
		}
		fmt.Print(" ", temp)
	}
	fmt.Println()
}
