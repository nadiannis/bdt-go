package main

import (
	"fmt"
	"math/big"
)

func main() {
	num := big.NewInt(30)

	result := factorial(num)
	fmt.Printf("The result of !%v is %v\n", num, result)
}

func factorial(num *big.Int) *big.Int {
	one := big.NewInt(1)

	if num.Cmp(big.NewInt(0)) == 0 {
		return one
	}

	return one.Mul(num, factorial(one.Sub(num, one)))
}
