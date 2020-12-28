package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{2, 5, 3, 5, 8, 1, 9, 5, 4, 7, 6}
	data := sort.IntSlice(a[0:])
	sort.Sort(data)
	for _, v := range data {
		fmt.Println(v)
	}
}
