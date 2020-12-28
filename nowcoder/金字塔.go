package main

import (
	"fmt"
	"io"
	"strconv"
)

func main() {
	var a []int
	for {
		var b int
		_, err := fmt.Scanf("%d", &b)
		if err != nil || err == io.EOF {
			break
		}
		a = append(a, b)
	}

	for _, v := range a {
		printTower(v)
	}

}

func printTower(n int) {

	for x := 1; x <= n; x++ {
		s := ""
		for i := n - x; i >= 0; i-- {
			s = s + " "
		}
		for j := 1; j <= x; j++ {
			s = s + strconv.Itoa(j)
		}
		for k := x - 1; k > 0; k-- {
			s = s + strconv.Itoa(k)
		}
		fmt.Println(s)
	}

}
