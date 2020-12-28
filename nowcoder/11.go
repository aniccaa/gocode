package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int
	_, err := fmt.Scanf("%d", &a)
	if err != nil {
		return
	}
	s := strconv.Itoa(a)
	t := revertByte([]byte(s))
	fmt.Print(string(t))
}

func revertByte(a []byte) []byte {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}
