package main

import (
	"errors"
	"fmt"
)

func main() {
	// fmt.Println("oie")
	resultado, err := soma(-8, 10)

	if err != nil {
		fmt.Println("erro: " + err.Error())

	} else {

		fmt.Println(resultado)
	}

}

func soma(a int, b int) (int, error) {
	if a < 0 {
		return 0, errors.New("A é menor que 0.")
	}
	return a + b, nil
}
