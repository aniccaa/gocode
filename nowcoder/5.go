package main

import (
	"fmt"
	"io"
	"strconv"
)

func main() {
	var a []string
	for {
		var t string
		_, err := fmt.Scanf("%s", &t)
		if err != nil || err == io.EOF {
			break
		}
		a = append(a, t)
	}

	for _, v := range a {
		i, e := strconv.ParseUint(v, 0, 32)
		if e != nil {
			return
		}
		fmt.Println(i)
	}

}
