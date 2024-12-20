package main

import "fmt"

func main() {
	//s := "((1 + 2) * 3) / ((6 * 3) - ((2 * 3) + 9 + 2))"
	s := "+5.5 - 3.3"
	//s := "(1 + 3 + 5) * 2.5 - 12"

	fmt.Print(Calc(s))

}
