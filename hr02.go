package main

import (
	"fmt"
	"math"
)

func main() {

	a := []int64 {1239898798798745,6787867868768765789, 676546542987}
	x := aVeryBigSum(a)
	fmt.Println(x)

	y := math.MaxInt64
	fmt.Println(y)

}

func aVeryBigSum(ar []int64) int64 {
	var bigSum int64
	bigSum = 0
	for i := 0; i < len(ar); i++ {
		bigSum += ar[i]
	}

	return bigSum

}
