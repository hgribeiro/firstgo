package main

import "fmt"

func main() {
	// fmt.Println("oie")
	resultado := soma(170, 10)

	fmt.Println(resultado)

}

func soma(a int, b int) int {
	return a + b
}
