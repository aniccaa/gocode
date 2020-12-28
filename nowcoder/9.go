package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int
	var res []byte
	_, err := fmt.Scanf("%d", &a)
	if err != nil {
		return
	}
	s := strconv.Itoa(a)
	t := revert([]byte(s))
	for _, v := range t {
		if !isExist(res, v) {
			res = append(res, v)
		}
	}
	// 输出结果。
	fmt.Print(string(res))
}

func isExist(a []byte, b byte) bool {
	for _, v := range a {
		if v == b {
			return true
		}
	}
	return false
}

func revert(a []byte) []byte {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}
