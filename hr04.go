package main

import "fmt"

func main() {

	x := []int32 {4,3,-9,0,4,1}
	plusMinus(x)

}

func plusMinus(arr []int32) {

	var pos, neg, zer float32
	pos, neg, zer  = 0, 0, 0

	lenx := len(arr)

	for i := 0; i < lenx; i++ {
		switch true {
		case arr[i] > 0:
			pos += 1
		case arr[i] < 0:
			neg += 1
		case arr[i] == 0:
			zer += 1
		}
	}

	a := pos/float32(lenx)
	b := neg/float32(lenx)
	c := zer/float32(lenx)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)


}