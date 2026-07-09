package main

import "fmt"

func CheckNumber(arg string) bool {
	for _, a := range arg{
		if a >= '0' && a <= '9' {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(CheckNumber("Hello"))
	fmt.Println(CheckNumber("Hello1"))
}
