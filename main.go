package main

import (
	"fmt"
	"math/cmplx"
)

type MyType int
type qwe float64

func main() {
	s1 := struct{ x int }{1}
	s2 := struct{ x int }{2}
	s3 := struct{ x int }{3}
	s4 := struct{ x int }{4}
	var arr []struct{ x int }
	arr = append(arr, s1, s2, s3, s4)

	minMax(arr, func(a, b struct{ x int }) bool { return a.x > b.x })
	minMax([][]int{[]int{1, 2, 6}, []int{5, 7}, []int{5, 7, 3, 5}, []int{9}}, func(a, b []int) bool { return len(a) > len(b) })
	minMax([]int{-1, -1}, func(a, b int) bool { return a > b })
	minMax([]int{-1, -1, 5, 2, 7, 4, 8, 9}, func(a, b int) bool { return a > b })
	minMax([]int{-1, -1, 5, 2, 7, 4, 8, 9}, func(a, b int) bool { return a > b })
	minMax([]uint{1, 1, 5, 2, 7, 4, 8, 9}, func(a, b uint) bool { return a > b })
	minMax([]float64{-1.5, -1.95, 5.1482, 2.1, 7.514, 4.5962, -0.58, 9}, func(a, b float64) bool { return a > b })
	minMax([]MyType{1, 2, 3, 4}, func(a, b MyType) bool { return a > b })
	minMax([]qwe{-1.5, -1.95, 5.1482, 2.1, 7.514, 4.5962, -0.58, 9}, func(a, b qwe) bool { return a > b })
	minMax([]complex128{1 + 2i, -1 + 3i, 5 + 1i, 2 + 9i}, func(a, b complex128) bool { return cmplx.Abs(a) > cmplx.Abs(b) })

}

func minMax[T any](arr []T, over func(a, b T) bool) {

	max := arr[0]
	min := arr[0]

	for _, v := range arr {
		if over(v, max) {
			max = v
			continue
		}
		if over(min, v) {
			min = v
			continue
		}
	}
	fmt.Printf("min=%v\tmax=%v\n", min, max)
}
