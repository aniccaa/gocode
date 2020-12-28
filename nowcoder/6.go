package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var a int
	_, err := fmt.Scanf("%d", &a)
	if err != nil {
		return
	}
	fmt.Print(getResult(a))

	// fmt.Print(zhishu(180))
}

func getResult(l int) string {

	var a = zhishu(l)
	var b []int
	for i := 0; i < len(a); {
		if l%a[i] == 0 {
			l = l / a[i]
			b = append(b, a[i])
		} else {
			i++
		}

	}
	return toString(b)
}

func toString(a []int) string {
	var s string
	for _, v := range a {

		s = s + strconv.Itoa(v) + " "
	}
	return s
}

func zhishu(a int) []int {
	var b []int

	for i := 2; i <= a; i++ {
		if isZhishu(i) {
			b = append(b, i)
		}
	}
	return b
}

func isZhishu(a int) bool {
	if a <= 3 {
		return a > 1
	}
	for i := 2; i <= int(math.Sqrt(float64(a))); i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}
