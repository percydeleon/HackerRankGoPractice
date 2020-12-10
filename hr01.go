package main

import "fmt"

func main() {
	a := []int32 {1,2,3}
	b := []int32 {1,1,1}
	x := compareTriplets(a,b)

	fmt.Println(x)
}

// Complete the compareTriplets function below.
func compareTriplets(a []int32, b []int32) []int32 {
	points := make([]int32, 2)
	points[0] = 0;
	points[1] = 0;

	for i := 0; i < 3; i++ {
		if a[i] > b[i] {
			points[0] += 1
		} else if a[i] < b[i] {
			points[1] += 1
		}
	}

	return points

}