package main

import (
	"fmt"
)

func main() {
	for i := range 100 {
		result := ""
		if i%3 == 0 {
			result = "Fizz"
		}
		if i%5 == 0 {
			result += "Buzz"
		}
		if result == "" {
			fmt.Println(i)
			continue
		}
		fmt.Println(result)
	}
}
