package main

import (
	"fmt"
)

func main() {
	var a string
	_, err := fmt.Scanf("%s", &a)
	if err != nil {
		return
	}
	t := revertChar([]byte(a))
	fmt.Print(string(t))
}

func revertChar(a []byte) []byte {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}
