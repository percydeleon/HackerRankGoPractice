package main

import "fmt"

func main() {
	y := []int32 {1,2,3,4,10,11}
	x := simpleArraySum(y)
	fmt.Println(x)
}

func simpleArraySum(ar []int32) int32 {

	var simpleSum int32
	simpleSum = 0

	for i := 0; i < len(ar); i++ {
		simpleSum += ar[i]
	}
	return simpleSum

}
