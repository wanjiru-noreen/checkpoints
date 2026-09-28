package main

import "fmt"

func CheckNumber(arg string) bool {
	for _, a := range arg{ //for each character in arg put it into a
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

//(arg string) receives a string like Hello1
//bool returns either true or false
//range arg = go through arg one rune at a time
