package main

import (
	"fmt"
	"math/big"
)

func main() {

	a := big.NewFloat(0)

	// b := new(big.Float)

	// a.Add(b, big.NewFloat(3))

	// a.Mul(big.NewFloat(5), big.NewFloat(4))

	for _, v := range []int{2, 2, 5, 6} {
		a.Add(a, new(big.Float).Mul(big.NewFloat(float64(v)), big.NewFloat(2)))
	}
	// ---- 30   1+ 2+4 +10 +12
	fmt.Println("----", a)
}
